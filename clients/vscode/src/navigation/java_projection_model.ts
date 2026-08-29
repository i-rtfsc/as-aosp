export type ProjectionPathRef = {
  raw?: string;
  kind: string;
  repoPath?: string;
};

export type ProjectionTarget = {
  id: string;
  kind?: string;
  sourceRoots: ProjectionPathRef[];
  sourceFiles: ProjectionPathRef[];
  generatedRoots: ProjectionPathRef[];
  libraries: ProjectionPathRef[];
  targetDependencies: string[];
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
    id: typeof target?.spec?.id === "string" ? target.spec.id : "java-target",
    kind: typeof target?.spec?.kind === "string" ? target.spec.kind : undefined,
    sourceRoots: array<ProjectionPathRef>(target?.ownedSourceRoots ?? target?.sourceRoots),
    sourceFiles: array<ProjectionPathRef>(target?.ownedSources),
    generatedRoots: array<ProjectionPathRef>(target?.ownedGeneratedRoots ?? target?.generatedRoots),
    libraries: array<ProjectionPathRef>(target?.libraries).filter(
      (ref: ProjectionPathRef) => !ref?.raw?.toLowerCase().endsWith(".srcjar"),
    ),
    targetDependencies: array<string>(target?.targetDependencies),
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
