import * as path from "node:path";
import * as vscode from "vscode";
import { resolveRepoTarget, toRepoPath } from "./repo_path";

export function workspaceViewRoot(): string | undefined {
  const configured = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot");
  return configured?.trim() || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
}

export function repoPath(root: string, file: vscode.Uri): string | undefined {
  return toRepoPath(path, root, file.fsPath);
}

export function targetUri(root: string, repo: string): vscode.Uri | undefined {
  const target = resolveRepoTarget(path, root, repo);
  return target ? vscode.Uri.file(target) : undefined;
}
