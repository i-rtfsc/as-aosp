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
  const name = editor.document.getText(word);
  const packRoot = vscode.workspace.getConfiguration("aospIde").get<string>("metadataPackRoot")?.trim();
  const packed = packRoot ? await engine.aidlPackPeers(packRoot, name) : undefined;
  const result = packed && packed.diagnostics.length === 0 && packed.peers.length > 0
    ? packed
    : await engine.aidlPeers({ workspaceViewRoot: root, currentPath, name });
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
