export function javaCandidatePaths(packageName: string, imports: string[], name: string): string[] {
  const candidates = new Set<string>();
  for (const imported of imports) if (imported.endsWith(`.${name}`)) candidates.add(`${imported.replaceAll(".", "/")}.java`);
  if (packageName) candidates.add(`${packageName.replaceAll(".", "/")}/${name}.java`);
  candidates.add(`${name}.java`);
  return [...candidates];
}
