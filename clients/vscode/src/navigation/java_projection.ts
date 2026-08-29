import * as vscode from "vscode";
import * as fs from "node:fs";
import * as path from "node:path";
import { EngineClient } from "../engine/client";
import { normalizeWorkspacePath, resolveRepoProjectionPath } from "./java_projection_paths";
import { array, normalizeProjectionInput, preferredLibraries } from "./java_projection_model";
import { deriveJavaSourceRoots } from "./java_source_roots";

type PathRef = { raw?: string; kind: string; repoPath?: string };

export interface JavaProjectionSummary {
  sourcePaths: string[];
  referencedLibraries: string[];
  generatedRoots: number;
  generatedEntries: number;
  moduleCount: number;
  projectCount: number;
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
  const libraryCandidates: string[] = [];
  const claimedRoots = new Set<string>();
  let skipped = projection.diagnostics + array(generated.diagnostics).length;
  const generatedRoots = array(generated.sourceRoots);
  const generatedOverlay = exposeGeneratedSourceRoots(workspaceRoot, generatedRoots);
  if (!generatedOverlay) skipped += generatedRoots.length;
  else sourcePaths.add(generatedOverlay);
  for (const target of graphTargets) {
    const derivedRoots = deriveJavaSourceRoots(workspaceRoot, target.sourceFiles);
    const targetRoots = new Set<string>();
    if (derivedRoots.length > 0) {
      for (const root of derivedRoots) targetRoots.add(root);
    } else {
      for (const ref of target.sourceRoots) {
        const pathName = resolveRepoPath(workspaceRoot, ref);
        if (pathName && ref.repoPath) targetRoots.add(normalizeWorkspacePath(ref.repoPath));
        else skipped++;
      }
    }
    const ownedRoots = [...targetRoots].filter((root) => {
      if (claimedRoots.has(root)) return false;
      claimedRoots.add(root);
      sourcePaths.add(root);
      return true;
    }).map((root) => path.join(workspaceRoot, ...root.split("/")));
    for (const ref of target.libraries) {
      const libraryPath = resolveLocalPath(outRoot, ref);
	  if (libraryPath) libraryCandidates.push(libraryPath);
      else skipped++;
    }
  }
	const referencedLibraries = preferredLibraries(libraryCandidates);
	fs.rmSync(path.join(workspaceRoot, ".aosp-ide", "java", "projects"), {
	  recursive: true, force: true,
	});
  const config = vscode.workspace.getConfiguration("java.project");
	await config.update(
	  "sourcePaths", [...sourcePaths].sort(), vscode.ConfigurationTarget.Workspace,
	);
	await config.update(
	  "referencedLibraries", referencedLibraries, vscode.ConfigurationTarget.Workspace,
	);
	await config.update(
	  "outputPath", ".aosp-ide/java/classes", vscode.ConfigurationTarget.Workspace,
	);
  await config.update("resourceFilters", [
    "node_modules", "\\.git", "\\.repo", "out", "prebuilts",
    "tests?", "benchmarks?", "testdata",
  ], vscode.ConfigurationTarget.Workspace);
  // AOSP contains unrelated Gradle projects whose build scripts are not a
  // complete standalone build. Let JDT use this projection instead of trying
  // to import those projects and failing workspace initialization.
  const javaImport = vscode.workspace.getConfiguration("java.import");
  await javaImport.update("gradle.enabled", false, vscode.ConfigurationTarget.Workspace);
  await javaImport.update("maven.enabled", false, vscode.ConfigurationTarget.Workspace);
  const gradle = vscode.workspace.getConfiguration("gradle");
  await gradle.update("autoDetect", "off", vscode.ConfigurationTarget.Workspace);
  await gradle.update("nestedProjects", false, vscode.ConfigurationTarget.Workspace);
  await vscode.workspace.getConfiguration("java.autobuild").update(
    "enabled", false, vscode.ConfigurationTarget.Workspace,
  );
  await vscode.workspace.getConfiguration("java.server").update(
    "launchMode", "Standard", vscode.ConfigurationTarget.Workspace,
  );
  const jdt = vscode.workspace.getConfiguration("java.jdt.ls");
  if (jdt.inspect<string>("vmargs")?.workspaceValue === undefined) {
    const current = jdt.get<string>("vmargs") ?? "";
    const vmargs = /-Xmx\S+/.test(current) ? current.replace(/-Xmx\S+/, "-Xmx6G") : `${current} -Xmx6G`;
    await jdt.update("vmargs", vmargs.trim(), vscode.ConfigurationTarget.Workspace);
  }
  if (vscode.extensions.getExtension("redhat.java")?.isActive) {
    setTimeout(() => {
      void vscode.commands.executeCommand("java.projectConfiguration.update");
    }, 0);
  }
  return {
    sourcePaths: [...sourcePaths].sort(),
	  referencedLibraries,
    generatedRoots: array(generated.sourceRoots).length,
    generatedEntries: generatedRoots.reduce((total, root) => total + root.entries, 0),
    moduleCount: projection.moduleCount,
	  projectCount: sourcePaths.size > 0 ? 1 : 0,
    skipped,
  };
}

function resolveRepoPath(workspaceRoot: string, ref: PathRef): string | undefined {
  return resolveRepoProjectionPath(workspaceRoot, ref);
}

function exposeGeneratedSourceRoots(
  workspaceRoot: string,
  roots: Array<{ root: string; hash: string }>,
): string | undefined {
  const valid = roots.filter((root) => root.root && root.hash &&
    fs.existsSync(path.join(root.root, "entries"))).sort((a, b) => a.hash.localeCompare(b.hash));
  if (valid.length === 0) return undefined;
  const relative = path.join(".aosp-ide", "java", "generated");
  const destination = path.join(workspaceRoot, relative);
  const manifestPath = path.join(workspaceRoot, ".aosp-ide", "java", "generated.json");
  const fingerprint = valid.map((root) => root.hash).join("\n");
  try {
    if (fs.existsSync(destination) && fs.readFileSync(manifestPath, "utf8") === fingerprint) {
      return normalizeWorkspacePath(relative);
    }
  } catch { /* rebuild an incomplete overlay */ }
  const temporary = destination + `.tmp-${process.pid}`;
  try {
    fs.rmSync(temporary, { recursive: true, force: true });
    fs.mkdirSync(temporary, { recursive: true });
    for (const root of valid) {
      linkGeneratedEntries(path.join(root.root, "entries"), temporary);
    }
    fs.rmSync(destination, { recursive: true, force: true });
    fs.renameSync(temporary, destination);
    fs.writeFileSync(manifestPath, fingerprint, "utf8");
    return normalizeWorkspacePath(relative);
  } catch {
    try { fs.rmSync(temporary, { recursive: true, force: true }); } catch { /* best effort */ }
    return undefined;
  }
}

function linkGeneratedEntries(sourceRoot: string, destinationRoot: string): void {
  const pending: Array<{ source: string; relative: string }> = [{ source: sourceRoot, relative: "" }];
  while (pending.length > 0) {
    const current = pending.pop()!;
    for (const entry of fs.readdirSync(current.source, { withFileTypes: true })) {
      const relative = path.join(current.relative, entry.name);
      const source = path.join(current.source, entry.name);
      if (entry.isDirectory()) {
        pending.push({ source, relative });
        continue;
      }
      if (!entry.isFile() || !/\.(?:java|kt)$/i.test(entry.name)) continue;
      const destination = path.join(destinationRoot, relative);
      if (fs.existsSync(destination)) continue;
      fs.mkdirSync(path.dirname(destination), { recursive: true });
      try {
        fs.symlinkSync(source, destination, "file");
      } catch {
        fs.copyFileSync(source, destination, fs.constants.COPYFILE_EXCL);
      }
    }
  }
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
