import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { workspaceViewRoot } from "./paths";
import * as path from "node:path";
import { overlayClangdConfig } from "./clangd_config";

export async function refreshCompdbOverlay(engine: EngineClient, output?: vscode.OutputChannel): Promise<void> {
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
  output?.appendLine(`compdb overlay: ${result.entries.length} entries`);
  for (const diagnostic of result.diagnostics) output?.appendLine(`diagnostic: ${JSON.stringify(diagnostic)}`);
  const target = vscode.Uri.file(path.join(root, ".aosp-ide", "compdb-overlay", "compile_commands.json"));
  const directory = vscode.Uri.file(path.dirname(target.fsPath));
  await vscode.workspace.fs.createDirectory(directory);
  await vscode.workspace.fs.writeFile(target, Buffer.from(JSON.stringify(result.entries, null, 2), "utf8"));
  if (result.responseFiles) {
    for (const [name, content] of Object.entries(result.responseFiles)) {
      const responseTarget = vscode.Uri.file(path.join(path.dirname(target.fsPath), path.basename(name)));
      await vscode.workspace.fs.writeFile(responseTarget, Buffer.from(content, "utf8"));
    }
  }
  const clangdConfig = vscode.Uri.file(path.join(path.dirname(target.fsPath), ".clangd"));
  await vscode.workspace.fs.writeFile(clangdConfig, Buffer.from(overlayClangdConfig, "utf8"));
  if (result.diagnostics.length > 0) {
    void vscode.window.showWarningMessage(`Compilation database refreshed with ${result.diagnostics.length} diagnostics.`);
  } else {
    void vscode.window.showInformationMessage(`Compilation database refreshed: ${result.entries.length} entries.`);
  }
}
