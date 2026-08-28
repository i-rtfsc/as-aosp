import * as vscode from "vscode";
import { detectOutRoot } from "../navigation/out_root";

export type OutDetection = { root?: string; mode: "detected" | "source-only"; files: string[] };

export async function detectOut(root: string): Promise<OutDetection> {
  const detection = detectOutRoot(root);
  if (detection.root) return { root: detection.root, mode: "detected", files: detection.files };
  return { mode: "source-only", files: [] };
}

export class OverviewProvider implements vscode.TreeDataProvider<OverviewItem> {
  private readonly changeEmitter = new vscode.EventEmitter<void>();
  readonly onDidChangeTreeData = this.changeEmitter.event;
  constructor(private readonly context: vscode.ExtensionContext) {}
  refresh(): void { this.changeEmitter.fire(); }
  getTreeItem(element: OverviewItem): vscode.TreeItem { return element; }
  async getChildren(): Promise<OverviewItem[]> {
    const root = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
    if (!root) return [new OverviewItem("Open a workspace folder", "aospIde.openProfile")];
    const config = vscode.workspace.getConfiguration("aospIde");
    const detection = await detectOut(root);
    const saved = config.get<string>("outRoot")?.trim();
    return [
      new OverviewItem(`Workspace: ${root}`, "aospIde.openProfile"),
      new OverviewItem(`Profile: ${saved ? "Build-backed" : detection.mode === "detected" ? "Out detected" : "Source-only"}`, "aospIde.openProfile"),
      new OverviewItem(`Engine: ${config.get<string>("engineExecutable") || "bundled / PATH"}`, "aospIde.showStatus"),
      new OverviewItem(`Build facts: ${saved || detection.root || "not detected"}`, "aospIde.openProfile"),
      new OverviewItem("Configure Profile", "aospIde.openProfile"),
      new OverviewItem("Refresh clangd Database", "aospIde.refreshCompdbOverlay"),
    ];
  }
}

class OverviewItem extends vscode.TreeItem {
  constructor(label: string, command: string) {
    super(label, vscode.TreeItemCollapsibleState.None);
    this.command = { command, title: label };
    this.contextValue = "aospIdeOverview";
  }
}
