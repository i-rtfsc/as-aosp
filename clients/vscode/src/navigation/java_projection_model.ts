export type ProjectionPathRef = {
  raw?: string;
  kind: string;
  repoPath?: string;
};

export type ProjectionTarget = {
  kind?: string;
  sourceRoots: ProjectionPathRef[];
  generatedRoots: ProjectionPathRef[];
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
    sourceRoots: array<ProjectionPathRef>(target?.ownedSourceRoots ?? target?.sourceRoots),
    generatedRoots: array<ProjectionPathRef>(target?.ownedGeneratedRoots ?? target?.generatedRoots),
    libraries: array<ProjectionPathRef>(target?.libraries).filter(
      (ref: ProjectionPathRef) => !ref?.raw?.toLowerCase().endsWith(".srcjar"),
    ),
  }));
  const javaTargets = targets.filter((target) => !target.kind ||
    target.kind === "java" || target.kind === "android");
  const archiveRefs = javaTargets.flatMap((target) => target.generatedRoots
    .filter((ref) => ref?.raw?.toLowerCase().endsWith(".srcjar"))
    .map((ref) => ref.raw as string));
  return { targets: javaTargets, archiveRefs: [...new Set(archiveRefs)], diagnostics };
}

export function array<T>(value: T[] | null | undefined): T[] {
  return Array.isArray(value) ? value : [];
}
