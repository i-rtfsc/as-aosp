import * as fs from "node:fs";
import * as path from "node:path";

export function normalizeOutRoot(value: string): string {
  const trimmed = value.trim();
  if (!trimmed) return "";
  const clean = path.resolve(trimmed);
  if (path.basename(clean) !== "soong") return clean;
  const parent = path.dirname(clean);
  if (hasModuleFacts(clean) && fs.existsSync(path.join(parent, "soong"))) {
    return parent;
  }
  return clean;
}

export function detectOutRoot(workspaceRoot: string): { root?: string; files: string[] } {
  for (const candidate of [path.join(workspaceRoot, "out"), workspaceRoot]) {
    const root = normalizeOutRoot(candidate);
    const files = moduleFactFiles(root);
    if (files.length > 0) return { root, files };
  }
  return { files: [] };
}

function hasModuleFacts(root: string): boolean {
  return moduleFactFiles(path.dirname(root)).length > 0 ||
    fs.existsSync(path.join(root, "module_bp_java_deps.json"));
}

function moduleFactFiles(root: string): string[] {
  const candidates = [
    "soong/module_bp_java_deps.json",
    "soong/module_bp_cc_deps.json",
    "soong/development/ide/compdb/compile_commands.json",
    "module_bp_java_deps.json",
    "module_bp_cc_deps.json",
    "development/ide/compdb/compile_commands.json",
  ];
  return candidates.filter((relative) => fs.existsSync(path.join(root, relative)));
}
