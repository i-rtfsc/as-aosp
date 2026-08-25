import * as vscode from "vscode";
import * as fs from "node:fs";
import * as path from "node:path";
import { EngineClient } from "../engine/client";
import { resolveRepoProjectionPath } from "./java_projection_paths";

type PathRef = { raw?: string; kind: string; repoPath?: string };

export interface JavaProjectionSummary {
  sourcePaths: string[];
  referencedLibraries: string[];
  skipped: number;
}

export async function applyJavaProjection(
  engine: EngineClient,
  workspaceRoot: string,
  metadataPackRoot: string,
  cacheRoot: string,
  outRoot = "",
): Promise<JavaProjectionSummary> {
  const generated = await engine.materializeGenerated(metadataPackRoot, cacheRoot);
  const catalog = await engine.modulesCatalog(metadataPackRoot);
  const closure = await engine.modulesClosure(
    metadataPackRoot,
    catalog.modules.map((module) => module.name),
  );
  const sourcePaths = new Set<string>();
  const referencedLibraries = new Set<string>();
  let skipped = closure.diagnostics.length + generated.diagnostics.length;
  for (const sourceRoot of generated.sourceRoots) sourcePaths.add(sourceRoot.root);
  for (const module of closure.modules) {
    for (const ref of [...(module.sourceRoots ?? []), ...(module.sources ?? [])]) {
      const pathName = resolveRepoPath(workspaceRoot, ref);
      if (pathName) sourcePaths.add(pathName);
      else skipped++;
    }
    for (const ref of [...(module.jars ?? []), ...(module.srcjars ?? [])]) {
      const pathName = resolveLocalPath(metadataPackRoot, outRoot, ref);
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
    skipped,
  };
}

function resolveRepoPath(workspaceRoot: string, ref: PathRef): string | undefined {
  return resolveRepoProjectionPath(workspaceRoot, ref);
}

function resolveLocalPath(packRoot: string, outRoot: string, ref: PathRef): string | undefined {
  let candidate: string | undefined;
  if (ref.kind === "external-absolute" && ref.raw) candidate = ref.raw;
  if ((ref.kind === "repo-relative" || ref.kind === "generated") && ref.raw && !ref.raw.includes("..")) {
    candidate = path.join(packRoot, ...ref.raw.replaceAll("\\", "/").split("/"));
  }
  if (candidate && fs.existsSync(candidate)) return candidate;
  if (outRoot && ref.raw && !ref.raw.includes("..")) {
    const fallback = path.join(outRoot, ...ref.raw.replaceAll("\\", "/").split("/"));
    if (fs.existsSync(fallback)) return fallback;
  }
  return undefined;
}
