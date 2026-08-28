export type ProjectionPathRef = {
  raw?: string;
  kind: string;
  repoPath?: string;
};

export type ProjectionTarget = {
  sourceRoots: ProjectionPathRef[];
  libraries: ProjectionPathRef[];
};

export type ProjectionInput = {
  targets: ProjectionTarget[];
  archiveRefs: string[];
  diagnostics: number;
};

export function normalizeProjectionInput(result: any): ProjectionInput {
  const diagnostics = array(result?.diagnostics).length +
    array(result?.graph?.diagnostics).length;
  const targets = array(result?.graph?.targets).map((target: any) => ({
    sourceRoots: array<ProjectionPathRef>(target?.sourceRoots),
    libraries: array<ProjectionPathRef>(target?.libraries),
  }));
  const archiveRefs = targets.flatMap((target) => target.libraries
    .filter((ref) => ref?.raw?.toLowerCase().endsWith(".srcjar"))
    .map((ref) => ref.raw as string));
  return { targets, archiveRefs: [...new Set(archiveRefs)], diagnostics };
}

export function array<T>(value: T[] | null | undefined): T[] {
  return Array.isArray(value) ? value : [];
}
