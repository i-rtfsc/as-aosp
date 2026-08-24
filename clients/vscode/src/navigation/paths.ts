import * as path from "node:path";
import * as vscode from "vscode";

export function workspaceViewRoot(): string | undefined {
  const configured = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot");
  return configured?.trim() || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
}

export function repoPath(root: string, file: vscode.Uri): string | undefined {
  const relative = path.relative(root, file.fsPath);
  if (!relative || relative.startsWith("..") || path.isAbsolute(relative)) return undefined;
  return relative.split(path.sep).join("/");
}

export function targetUri(root: string, repo: string): vscode.Uri {
  return vscode.Uri.file(path.resolve(root, ...repo.split("/")));
}
