import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { repoPath, targetUri, workspaceViewRoot } from "./paths";
import { selectJniDefinitionTarget } from "./jni_selection";
export { selectJniDefinitionTarget } from "./jni_selection";

export class JniDefinitionProvider implements vscode.DefinitionProvider {
  constructor(private readonly engine: EngineClient) {}

  async provideDefinition(document: vscode.TextDocument, _position?: vscode.Position, _token?: vscode.CancellationToken): Promise<vscode.Definition | undefined> {
    const root = workspaceViewRoot();
    const currentPath = root && repoPath(root, document.uri);
    if (!root || !currentPath) return undefined;
    const result = await this.engine.jniPeers({ workspaceViewRoot: root, currentPath });
    const selected = selectJniDefinitionTarget(result, currentPath);
    if (!selected) return undefined;
    const target = targetUri(root, selected.path);
    if (!target) return undefined;
    const targetDocument = await vscode.workspace.openTextDocument(target);
    const position = targetDocument.positionAt(Math.min(selected.offset, targetDocument.getText().length));
    return new vscode.Location(target, new vscode.Range(position, position));
  }
}

export async function goToJniPeer(engine: EngineClient): Promise<void> {
  const editor = vscode.window.activeTextEditor;
  const root = workspaceViewRoot();
  const currentPath = editor && root && repoPath(root, editor.document.uri);
  if (!editor || !root || !currentPath) return;
  const result = await engine.jniPeers({ workspaceViewRoot: root, currentPath });
  const selected = selectJniDefinitionTarget(result, currentPath);
  if (!selected) {
    void vscode.window.showWarningMessage("No unique JNI peer was found for this file.");
    return;
  }
  const target = targetUri(root, selected.path);
  if (!target) {
    void vscode.window.showWarningMessage("JNI peer path is outside the Workspace View.");
    return;
  }
  const document = await vscode.workspace.openTextDocument(target);
  const position = document.positionAt(Math.min(selected.offset, document.getText().length));
  await vscode.window.showTextDocument(document, { selection: new vscode.Range(position, position) });
}
