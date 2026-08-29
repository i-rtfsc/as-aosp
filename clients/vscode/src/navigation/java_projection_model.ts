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
  resolvedModules: string[];
};

export type ProjectionInput = {
  targets: ProjectionTarget[];
  archiveRefs: string[];
  diagnostics: number;
  moduleCount: number;
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
    resolvedModules: array<string>(target?.resolvedModules),
  }));
  const javaTargets = targets.filter((target) => !target.kind ||
    target.kind === "java" || target.kind === "android");
  const archiveRefs = javaTargets.flatMap((target) => target.generatedRoots
    .filter((ref) => ref?.raw?.toLowerCase().endsWith(".srcjar"))
    .map((ref) => ref.raw as string));
  const moduleCount = new Set(javaTargets.flatMap((target) => target.resolvedModules)).size;
  return { targets: javaTargets, archiveRefs: [...new Set(archiveRefs)], diagnostics, moduleCount };
}

export function array<T>(value: T[] | null | undefined): T[] {
  return Array.isArray(value) ? value : [];
}

export function preferredLibraries(values: string[]): string[] {
  const selected = new Map<string, string>();
  for (const value of [...new Set(values)].sort()) {
    const key = basename(value).toLowerCase();
    const current = selected.get(key);
    if (!current || libraryScore(value) > libraryScore(current)) {
      selected.set(key, value);
    }
  }
  return [...selected.values()].sort();
}

function basename(value: string): string {
  return value.replaceAll("\\", "/").split("/").pop() ?? value;
}

function libraryScore(value: string): number {
  const normalized = value.replaceAll("\\", "/");
  let score = 0;
  if (normalized.includes("/android_common/")) score += 100;
  if (normalized.includes("/turbine-combined/")) score += 30;
  if (normalized.includes("/javac-header/")) score += 20;
  if (normalized.includes("/combined/")) score += 10;
  if (normalized.includes("_apex")) score -= 20;
  if (normalized.includes("/linux_")) score -= 40;
  return score;
}
