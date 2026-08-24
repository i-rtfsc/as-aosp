import * as path from "node:path";

export function overlayResponseFilePath(directory: string, name: string): string | undefined {
  const normalized = name.replaceAll("\\", "/");
  const base = path.basename(normalized);
  if (!base || base === "." || base === ".." || normalized !== base) return undefined;
  return path.join(directory, base);
}
