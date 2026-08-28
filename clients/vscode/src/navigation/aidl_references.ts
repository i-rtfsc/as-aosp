import * as vscode from "vscode";
import { EngineClient } from "../engine/client";
import { repoPath, targetUri, workspaceViewRoot } from "./paths";

export class AidlReferenceProvider implements vscode.ReferenceProvider {
  constructor(private readonly engine: EngineClient) {}

  async provideReferences(document: vscode.TextDocument, position: vscode.Position, _context?: vscode.ReferenceContext, _token?: vscode.CancellationToken): Promise<vscode.Location[]> {
    const root = workspaceViewRoot();
    const currentPath = root && repoPath(root, document.uri);
    const word = document.getWordRangeAtPosition(position, /[A-Za-z_]\w*/);
    if (!root || !currentPath || !word) return [];
    const name = document.getText(word);
    const result = await this.engine.aidlReferences({ workspaceViewRoot: root, currentPath, name });
    const locations = await Promise.all(result.references.map(async (reference) => {
      const target = targetUri(root, reference.path);
      if (!target) return undefined;
      const targetDocument = document.uri.toString() === target.toString()
        ? document : await vscode.workspace.openTextDocument(target);
      const start = targetDocument.positionAt(reference.range.start.offset);
      const end = targetDocument.positionAt(reference.range.end.offset);
      return new vscode.Location(target, new vscode.Range(start, end));
    }));
    return locations.filter((location): location is vscode.Location => location !== undefined);
  }
}
