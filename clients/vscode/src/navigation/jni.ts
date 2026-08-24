import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { repoPath, targetUri, workspaceViewRoot } from "./paths";

export async function goToJniPeer(engine: EngineClient): Promise<void> {
  const editor = vscode.window.activeTextEditor;
  const root = workspaceViewRoot();
  const currentPath = editor && root && repoPath(root, editor.document.uri);
  if (!editor || !root || !currentPath) return;
  const result = await engine.jniPeers({ workspaceViewRoot: root, currentPath });
  const targets = result.peers.flatMap((peer) => {
    if (peer.java.path === currentPath) return [{ path: peer.native.path, offset: peer.native.range.start }];
    if (peer.native.path === currentPath) return [{ path: peer.java.path, offset: peer.java.range.start }];
    return [];
  });
  if (targets.length !== 1) {
    void vscode.window.showWarningMessage("No unique JNI peer was found for this file.");
    return;
  }
  const target = targetUri(root, targets[0].path);
  if (!target) {
    void vscode.window.showWarningMessage("JNI peer path is outside the Workspace View.");
    return;
  }
  const document = await vscode.workspace.openTextDocument(target);
  const position = document.positionAt(Math.min(targets[0].offset, document.getText().length));
  await vscode.window.showTextDocument(document, { selection: new vscode.Range(position, position) });
}
