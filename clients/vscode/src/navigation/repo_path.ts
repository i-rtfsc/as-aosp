import type { PlatformPath } from "node:path";

const excludedSegments = new Set([".git", ".repo", "out", ".aosp-ide"]);

export function normalizeRepoPath(repo: string): string | undefined {
  const normalized = repo.replaceAll("\\", "/");
  if (!normalized || normalized.startsWith("/") ||
    /^[A-Za-z]:\//.test(normalized) || normalized.startsWith("//")) {
    return undefined;
  }
  const parts = normalized.split("/");
  if (parts.some((part) => !part || part === "." || part === "..")) return undefined;
  if (parts.some((part) => excludedSegments.has(part))) return undefined;
  return parts.join("/");
}

export function toRepoPath(paths: PlatformPath, root: string, file: string): string | undefined {
  const relative = paths.relative(root, file);
  if (!relative || relative.startsWith("..") || paths.isAbsolute(relative)) return undefined;
  return normalizeRepoPath(relative.split(paths.sep).join("/"));
}

export function resolveRepoTarget(paths: PlatformPath, root: string, repo: string): string | undefined {
  const normalized = normalizeRepoPath(repo);
  if (!normalized) return undefined;
  const target = paths.resolve(root, ...normalized.split("/"));
  return toRepoPath(paths, root, target) === normalized ? target : undefined;
}
