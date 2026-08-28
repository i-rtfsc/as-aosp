import * as fs from "node:fs";
import * as path from "node:path";
import * as vscode from "vscode";
import { EngineClient } from "../engine/client";

export interface RustProjectionSummary {
  projectPath?: string;
  crates: number;
  diagnostics: unknown[];
}

export async function applyRustProjection(
  engine: EngineClient,
  workspaceViewRoot: string,
  outRoot: string,
  cacheRoot: string,
  pathMappings: unknown[],
): Promise<RustProjectionSummary> {
  const source = path.join(outRoot, "soong", "rust-project.json");
  if (!fs.existsSync(source)) {
    return { crates: 0, diagnostics: [] };
  }
  const result = await engine.rustProject({
    workspaceViewRoot, outRoot, cacheRoot, pathMappings,
  });
  if (!result.projectPath || result.diagnostics.length > 0) {
    return result;
  }
  const configuration = vscode.workspace.getConfiguration("rust-analyzer");
  const current = configuration.get<unknown[]>("linkedProjects") ?? [];
  const linked = current.some((value) => value === result.projectPath) ?
    current : [...current, result.projectPath];
  await configuration.update(
    "linkedProjects", linked, vscode.ConfigurationTarget.Workspace,
  );
  return result;
}
