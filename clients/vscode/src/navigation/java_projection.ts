import * as vscode from "vscode";
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
): Promise<JavaProjectionSummary> {
  const catalog = await engine.modulesCatalog(metadataPackRoot);
  const closure = await engine.modulesClosure(
    metadataPackRoot,
    catalog.modules.map((module) => module.name),
  );
  const sourcePaths = new Set<string>();
  const referencedLibraries = new Set<string>();
  let skipped = closure.diagnostics.length;
  for (const module of closure.modules) {
    for (const ref of [...(module.sourceRoots ?? []), ...(module.sources ?? [])]) {
      const pathName = resolveRepoPath(workspaceRoot, ref);
      if (pathName) sourcePaths.add(pathName);
      else skipped++;
    }
    for (const ref of [...(module.jars ?? []), ...(module.srcjars ?? [])]) {
      const pathName = resolveLocalPath(metadataPackRoot, ref);
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

function resolveLocalPath(packRoot: string, ref: PathRef): string | undefined {
  if (ref.kind !== "repo-relative" && ref.kind !== "generated") return undefined;
  if (!ref.raw || ref.raw.includes("..")) return undefined;
  const candidate = vscode.Uri.joinPath(vscode.Uri.file(packRoot), ...ref.raw.replaceAll("\\", "/").split("/")).fsPath;
  return candidate;
}
