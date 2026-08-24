import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { repoPath, targetUri, workspaceViewRoot } from "./paths";

export class AidlDefinitionProvider implements vscode.DefinitionProvider {
  constructor(private readonly engine: EngineClient) {}

  async provideDefinition(
    document: vscode.TextDocument,
    position: vscode.Position,
    _token?: vscode.CancellationToken,
  ): Promise<vscode.Definition | undefined> {
    const root = workspaceViewRoot();
    const currentPath = root && repoPath(root, document.uri);
    if (!root || !currentPath) return undefined;
    const word = document.getWordRangeAtPosition(position, /[A-Za-z_]\w*/);
    if (!word) return undefined;
    const result = await this.engine.aidlResolve({ workspaceViewRoot: root, currentPath, name: document.getText(word) });
    if (result.resolution.state !== "found" || result.resolution.symbols?.length !== 1) return undefined;
    const symbol = result.resolution.symbols[0];
    const target = targetUri(root, symbol.path);
    if (!target) return undefined;
    const targetDocument = await vscode.workspace.openTextDocument(target);
    const offset = Math.min(symbol.range.start.offset, targetDocument.getText().length);
    return new vscode.Location(target, targetDocument.positionAt(offset));
  }
}
