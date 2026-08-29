export type ProjectionPathRef = {
  raw?: string;
  kind: string;
  repoPath?: string;
};

export type ProjectionTarget = {
  kind?: string;
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
    kind: typeof target?.spec?.kind === "string" ? target.spec.kind : undefined,
    sourceRoots: array<ProjectionPathRef>(target?.sourceRoots),
    libraries: array<ProjectionPathRef>(target?.libraries),
  }));
  const javaTargets = targets.filter((target) => !target.kind ||
    target.kind === "java" || target.kind === "android");
  const archiveRefs = javaTargets.flatMap((target) => target.libraries
    .filter((ref) => ref?.raw?.toLowerCase().endsWith(".srcjar"))
    .map((ref) => ref.raw as string));
  return { targets: javaTargets, archiveRefs: [...new Set(archiveRefs)], diagnostics };
}

export function array<T>(value: T[] | null | undefined): T[] {
  return Array.isArray(value) ? value : [];
}
