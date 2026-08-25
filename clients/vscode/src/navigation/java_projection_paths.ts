import * as path from "node:path";

export type ProjectionPathRef = { raw?: string; kind: string; repoPath?: string };

export function resolveRepoProjectionPath(
  workspaceRoot: string,
  ref: ProjectionPathRef,
): string | undefined {
  if (ref.kind !== "repo-relative" || !ref.repoPath) return undefined;
  return path.join(workspaceRoot, ...ref.repoPath.split("/"));
}
