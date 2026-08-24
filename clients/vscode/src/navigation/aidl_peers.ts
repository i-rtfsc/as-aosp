import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { repoPath, targetUri, workspaceViewRoot } from "./paths";
import { selectAidlPeerPath } from "./aidl_peer_selection";

export { selectAidlPeerPath } from "./aidl_peer_selection";

export async function goToAidlGeneratedPeer(engine: EngineClient): Promise<void> {
  const editor = vscode.window.activeTextEditor;
  const root = workspaceViewRoot();
  const currentPath = editor && root && repoPath(root, editor.document.uri);
  if (!editor || !root || !currentPath) return;
  const word = editor.document.getWordRangeAtPosition(editor.selection.active, /[A-Za-z_]\w*/);
  if (!word) {
    void vscode.window.showWarningMessage("Place the caret on an AIDL declaration.");
    return;
  }
  const result = await engine.aidlPeers({
    workspaceViewRoot: root,
    currentPath,
    name: editor.document.getText(word),
  });
  const selected = selectAidlPeerPath(result);
  if (!selected || !targetUri(root, selected)) {
    void vscode.window.showWarningMessage("No unique visible AIDL generated peer was found.");
    return;
  }
  const target = targetUri(root, selected);
  if (!target) return;
  const document = await vscode.workspace.openTextDocument(target);
  const position = new vscode.Position(0, 0);
  await vscode.window.showTextDocument(document, { selection: new vscode.Range(position, position) });
}
