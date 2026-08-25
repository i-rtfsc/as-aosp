import * as vscode from "vscode";
import { repoPath, targetUri, workspaceViewRoot } from "./paths";
import { javaCandidatePaths } from "./java_paths";
export { javaCandidatePaths } from "./java_paths";

export class JavaSourceDefinitionProvider implements vscode.DefinitionProvider {
  async provideDefinition(document: vscode.TextDocument, position: vscode.Position): Promise<vscode.Definition | undefined> {
    const root = workspaceViewRoot();
    if (!root) return undefined;
    const word = document.getWordRangeAtPosition(position, /[A-Za-z_$][\w$]*/);
    if (!word) return undefined;
    const name = document.getText(word);
    const source = document.getText();
    const packageName = /\bpackage\s+([\w.]+)\s*;/.exec(source)?.[1] ?? "";
    const imports = [...source.matchAll(/\bimport\s+([\w.]+)\s*;/g)].map((match) => match[1]);
    for (const relative of javaCandidatePaths(packageName, imports, name)) {
      const target = targetUri(root, relative);
      if (!target) continue;
      try {
        const targetDocument = await vscode.workspace.openTextDocument(target);
        const declaration = new RegExp(`\\b(?:class|interface|enum|record)\\s+${escapeRegExp(name)}\\b`).exec(targetDocument.getText());
        const offset = declaration?.index ?? 0;
        if (repoPath(root, target)) return new vscode.Location(target, targetDocument.positionAt(offset));
      } catch { /* candidate is not visible */ }
    }
    return undefined;
  }
}


function escapeRegExp(value: string): string { return value.replace(/[.*+?^${}()|[\]\\]/g, "\\$&"); }
