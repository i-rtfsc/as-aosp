import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { workspaceViewRoot } from "./paths";
import * as path from "node:path";

export async function refreshCompdbOverlay(engine: EngineClient): Promise<void> {
  const root = workspaceViewRoot();
  const metadataRoot = vscode.workspace.getConfiguration("aospIde").get<string>("metadataRoot")?.trim();
  if (!root || !metadataRoot) {
    void vscode.window.showWarningMessage("Configure AOSP IDE workspaceViewRoot and metadataRoot first.");
    return;
  }
  const result = await engine.compdbOverlay({
    metadataRoot,
    workspaceViewRoot: root,
    pathMappings: vscode.workspace.getConfiguration("aospIde").get<unknown[]>("pathMappings") ?? [],
  });
  const target = vscode.Uri.file(path.join(root, ".aosp-ide", "compdb-overlay", "compile_commands.json"));
  await vscode.workspace.fs.createDirectory(vscode.Uri.file(path.dirname(target.fsPath)));
  await vscode.workspace.fs.writeFile(target, Buffer.from(JSON.stringify(result.entries, null, 2), "utf8"));
  if (result.diagnostics.length > 0) {
    void vscode.window.showWarningMessage(`Compilation database refreshed with ${result.diagnostics.length} diagnostics.`);
  } else {
    void vscode.window.showInformationMessage(`Compilation database refreshed: ${result.entries.length} entries.`);
  }
}
