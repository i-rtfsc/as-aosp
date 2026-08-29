import * as fs from "node:fs";
import * as path from "node:path";

export type SourceRef = { kind: string; repoPath?: string };

export function deriveJavaSourceRoots(
  workspaceRoot: string,
  refs: SourceRef[],
): string[] {
  const roots = new Set<string>();
  for (const ref of refs) {
    const repoPath = ref.kind === "repo-relative" ? ref.repoPath : undefined;
    if (!repoPath || !/\.(?:java|kt)$/i.test(repoPath)) continue;
    const local = path.join(workspaceRoot, ...repoPath.split("/"));
    let source: string;
    try {
      const handle = fs.openSync(local, "r");
      try {
        const buffer = Buffer.alloc(64 * 1024);
        const length = fs.readSync(handle, buffer, 0, buffer.length, 0);
        source = buffer.subarray(0, length).toString("utf8");
      } finally {
        fs.closeSync(handle);
      }
    } catch {
      continue;
    }
    const root = sourceRootForPackage(repoPath, source);
    if (root) roots.add(root);
  }
  return [...roots].sort();
}

export function sourceRootForPackage(repoPath: string, source: string): string | undefined {
  const normalized = repoPath.replaceAll("\\", "/");
  const directory = normalized.slice(0, normalized.lastIndexOf("/"));
  const packageName = /\bpackage\s+([A-Za-z_$][\w$]*(?:\.[A-Za-z_$][\w$]*)*)\s*;?/.exec(source)?.[1];
  if (packageName) {
    const packagePath = packageName.replaceAll(".", "/");
    if (directory === packagePath) return ".";
    if (directory.endsWith("/" + packagePath)) {
      return directory.slice(0, -packagePath.length - 1);
    }
  }
  const parts = directory.split("/");
  for (let index = parts.length - 1; index >= 0; index--) {
    if (parts[index] === "java" || parts[index] === "kotlin") {
      return parts.slice(0, index + 1).join("/");
    }
  }
  return undefined;
}
