import * as vscode from "vscode";
import * as fs from "node:fs";
import * as path from "node:path";
import { EngineClient } from "../engine/client";
import { normalizeWorkspacePath, resolveRepoProjectionPath } from "./java_projection_paths";
import { array, normalizeProjectionInput } from "./java_projection_model";

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
  outRoot: string,
  cacheRoot: string,
  targets: unknown[] = [],
): Promise<JavaProjectionSummary> {
  const graphResult = await engine.projectGraph(outRoot, targets);
  const projection = normalizeProjectionInput(graphResult);
  const generated = await engine.materializeGenerated(
    outRoot, cacheRoot, projection.archiveRefs,
  );
  const graphTargets = projection.targets;
  const sourcePaths = new Set<string>();
  const referencedLibraries = new Set<string>();
  let skipped = projection.diagnostics + array(generated.diagnostics).length;
  for (const sourceRoot of array(generated.sourceRoots)) {
    const exposed = exposeGeneratedSourceRoot(workspaceRoot, sourceRoot.root, sourceRoot.hash);
    if (exposed) sourcePaths.add(exposed);
    else skipped++;
  }
  const graphModules = graphTargets.flatMap((target) => ({
    sourceRoots: target.sourceRoots,
    sources: [],
    jars: target.libraries,
    srcjars: [],
  }));
  for (const module of graphModules) {
    for (const ref of [...(module.sourceRoots ?? []), ...(module.sources ?? [])]) {
      const pathName = resolveRepoPath(workspaceRoot, ref);
      if (pathName && ref.repoPath) sourcePaths.add(normalizeWorkspacePath(ref.repoPath));
      else skipped++;
    }
    for (const ref of [...(module.jars ?? []), ...(module.srcjars ?? [])]) {
      const pathName = resolveLocalPath(outRoot, ref);
      if (pathName) referencedLibraries.add(pathName);
      else skipped++;
    }
  }
  const config = vscode.workspace.getConfiguration("java.project");
  await config.update("sourcePaths", [...sourcePaths].sort(), vscode.ConfigurationTarget.Workspace);
  await config.update("referencedLibraries", [...referencedLibraries].sort(), vscode.ConfigurationTarget.Workspace);
  // AOSP contains unrelated Gradle projects whose build scripts are not a
  // complete standalone build. Let JDT use this projection instead of trying
  // to import those projects and failing workspace initialization.
  const javaImport = vscode.workspace.getConfiguration("java.import");
  await javaImport.update("gradle.enabled", false, vscode.ConfigurationTarget.Workspace);
  await javaImport.update("maven.enabled", false, vscode.ConfigurationTarget.Workspace);
  await config.update("importOnFirstTimeStartup", "disabled", vscode.ConfigurationTarget.Workspace);
  const gradle = vscode.workspace.getConfiguration("gradle");
  await gradle.update("autoDetect", "off", vscode.ConfigurationTarget.Workspace);
  await gradle.update("nestedProjects", false, vscode.ConfigurationTarget.Workspace);
  return {
    sourcePaths: [...sourcePaths].sort(),
    referencedLibraries: [...referencedLibraries].sort(),
    generatedRoots: array(generated.sourceRoots).length,
    skipped,
  };
}

function resolveRepoPath(workspaceRoot: string, ref: PathRef): string | undefined {
  return resolveRepoProjectionPath(workspaceRoot, ref);
}

function exposeGeneratedSourceRoot(
  workspaceRoot: string,
  generatedRoot: string,
  hash: string,
): string | undefined {
  if (!generatedRoot || !hash || !fs.existsSync(path.join(generatedRoot, "entries"))) return undefined;
  const relative = path.join(".aosp-ide", "java", "srcjar", hash);
  const destination = path.join(workspaceRoot, relative);
  try {
    fs.mkdirSync(path.dirname(destination), { recursive: true });
    const current = fs.lstatSync(destination);
    if (current.isSymbolicLink() || current.isDirectory()) {
      let sameRoot = false;
      try { sameRoot = fs.realpathSync(destination) === fs.realpathSync(generatedRoot); } catch { /* stale link */ }
      if (!sameRoot) fs.rmSync(destination, { recursive: true, force: true });
    } else {
      fs.rmSync(destination, { recursive: true, force: true });
    }
  } catch (error) {
    if ((error as NodeJS.ErrnoException).code !== "ENOENT") {
      try { fs.rmSync(destination, { recursive: true, force: true }); } catch { return undefined; }
    }
  }
  if (!fs.existsSync(destination)) {
    try {
      fs.symlinkSync(generatedRoot, destination, process.platform === "win32" ? "junction" : "dir");
    } catch {
      try { fs.cpSync(generatedRoot, destination, { recursive: true }); } catch { return undefined; }
    }
  }
  return `${normalizeWorkspacePath(relative)}/entries`;
}

function resolveLocalPath(outRoot: string, ref: PathRef): string | undefined {
  let candidate: string | undefined;
  if (ref.kind === "external-absolute" && ref.raw) candidate = ref.raw;
  if (candidate && fs.existsSync(candidate)) return candidate;
  if (outRoot && ref.raw && !ref.raw.includes("..")) {
    const normalized = ref.raw.replaceAll("\\", "/");
    if (!path.isAbsolute(normalized)) {
      const suffix = normalized.startsWith("out/") ? normalized.slice(4) : normalized;
      const localOut = path.join(outRoot, ...suffix.split("/"));
      if (fs.existsSync(localOut)) return localOut;
    }
  }
  return undefined;
}
