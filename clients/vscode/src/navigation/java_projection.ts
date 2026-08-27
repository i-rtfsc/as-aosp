import * as vscode from "vscode";
import * as fs from "node:fs";
import * as path from "node:path";
import { EngineClient } from "../engine/client";
import { resolveRepoProjectionPath } from "./java_projection_paths";

type PathRef = { raw?: string; kind: string; repoPath?: string };

export interface JavaProjectionSummary {
  sourcePaths: string[];
  referencedLibraries: string[];
  generatedRoots: number;
  skipped: number;
}

export async function applyJavaProjection(
  engine: EngineClient,
  workspaceRoot: string,
  metadataPackRoot: string,
  cacheRoot: string,
  outRoot = "",
  targets: unknown[] = [],
): Promise<JavaProjectionSummary> {
  const graphResult = await engine.projectGraph(metadataPackRoot, targets, outRoot);
  const archiveRefs = graphResult.graph.targets.flatMap((target) =>
    target.libraries.filter((ref) => ref.raw?.toLowerCase().endsWith(".srcjar")).map((ref) => ref.raw),
  );
  const generated = metadataPackRoot
    ? await engine.materializeGenerated(metadataPackRoot, cacheRoot)
    : await engine.materializeGenerated("", cacheRoot, outRoot, archiveRefs);
  const packInfo: { outRoot?: string; diagnostics: unknown[] } = metadataPackRoot
    ? await engine.packInfo(metadataPackRoot)
    : { diagnostics: [] };
  const graphTargets = graphResult.diagnostics.length === 0 ? graphResult.graph.targets : [];
  const sourcePaths = new Set<string>();
  const referencedLibraries = new Set<string>();
  let skipped = graphResult.diagnostics.length + generated.diagnostics.length + packInfo.diagnostics.length;
  for (const sourceRoot of generated.sourceRoots) sourcePaths.add(sourceRoot.root);
  const graphModules = graphTargets.flatMap((target) => ({
    sourceRoots: target.sourceRoots,
    sources: [],
    jars: target.libraries,
    srcjars: [],
  }));
  for (const module of graphModules) {
    for (const ref of [...(module.sourceRoots ?? []), ...(module.sources ?? [])]) {
      const pathName = resolveRepoPath(workspaceRoot, ref);
      if (pathName) sourcePaths.add(pathName);
      else skipped++;
    }
    for (const ref of [...(module.jars ?? []), ...(module.srcjars ?? [])]) {
      const pathName = resolveLocalPath(metadataPackRoot, outRoot, packInfo.outRoot, ref);
      if (pathName) referencedLibraries.add(pathName);
      else skipped++;
    }
  }
  const config = vscode.workspace.getConfiguration("java.project");
  await config.update("sourcePaths", [...sourcePaths].sort(), vscode.ConfigurationTarget.Workspace);
  await config.update("referencedLibraries", [...referencedLibraries].sort(), vscode.ConfigurationTarget.Workspace);
  return {
    sourcePaths: [...sourcePaths].sort(),
    referencedLibraries: [...referencedLibraries].sort(),
    generatedRoots: generated.sourceRoots.length,
    skipped,
  };
}

function resolveRepoPath(workspaceRoot: string, ref: PathRef): string | undefined {
  return resolveRepoProjectionPath(workspaceRoot, ref);
}

function resolveLocalPath(packRoot: string, outRoot: string, packOutRoot: string | undefined, ref: PathRef): string | undefined {
  let candidate: string | undefined;
  if (ref.kind === "external-absolute" && ref.raw) candidate = ref.raw;
  if (packRoot && (ref.kind === "repo-relative" || ref.kind === "generated") && ref.raw && !ref.raw.includes("..")) {
    candidate = path.join(packRoot, ...ref.raw.replaceAll("\\", "/").split("/"));
  }
  if (candidate && fs.existsSync(candidate)) return candidate;
  if (packOutRoot && ref.raw) {
    const normalized = ref.raw.replaceAll("\\", "/");
    const sourceRoot = packOutRoot.replaceAll("\\", "/");
    if (normalized.startsWith(sourceRoot + "/")) {
      const suffix = normalized.slice(sourceRoot.length + 1);
      const packed = path.join(packRoot, "out", ...suffix.split("/"));
      if (fs.existsSync(packed)) return packed;
    }
  }
  if (outRoot && ref.raw && !ref.raw.includes("..")) {
    const normalized = ref.raw.replaceAll("\\", "/");
    if (packOutRoot && normalized.startsWith(packOutRoot.replaceAll("\\", "/") + "/")) {
      const suffix = normalized.slice(packOutRoot.replaceAll("\\", "/").length + 1);
      const fallback = path.join(outRoot, ...suffix.split("/"));
      if (fs.existsSync(fallback)) return fallback;
    }
    if (!path.isAbsolute(normalized)) {
      const suffix = normalized.startsWith("out/") ? normalized.slice(4) : normalized;
      const localOut = path.join(outRoot, ...suffix.split("/"));
      if (fs.existsSync(localOut)) return localOut;
    }
  }
  return undefined;
}
