import * as fs from "node:fs";
import * as path from "node:path";
import * as vscode from "vscode";

type Dependency = {
  id: string;
  name: string;
  capability: string;
  required: boolean;
};

const dependencies: Dependency[] = [
  {
    id: "redhat.java",
    name: "Language Support for Java (Red Hat)",
    capability: "Java/Kotlin navigation",
    required: true,
  },
  {
    id: "llvm-vs-code-extensions.vscode-clangd",
    name: "clangd",
    capability: "C/C++ and JNI navigation",
    required: true,
  },
  {
    id: "rust-lang.rust-analyzer",
    name: "rust-analyzer",
    capability: "Rust navigation",
    required: false,
  },
];

export function missingDependencies(): Dependency[] {
  return dependencies.filter((dependency) => !vscode.extensions.getExtension(dependency.id));
}

export function isAospWorkspace(): boolean {
  const configured = vscode.workspace.getConfiguration("aospIde")
    .get<string>("workspaceViewRoot")?.trim();
  const root = configured || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
  if (!root) return false;
  return fs.existsSync(path.join(root, "Android.bp")) ||
    fs.existsSync(path.join(root, "build", "envsetup.sh")) ||
    fs.existsSync(path.join(root, "soong", "soong_ui.bash"));
}

export async function checkDependencies(output: vscode.OutputChannel): Promise<void> {
  const missing = missingDependencies();
  if (missing.length === 0) {
    output.appendLine("language dependencies: all available");
    return;
  }
  const details = missing.map((dependency) =>
    `${dependency.name} (${dependency.capability}${dependency.required ? ", required" : ", optional"})`)
    .join("; ");
  output.appendLine(`missing language dependencies: ${details}`);
  const install = vscode.l10n.t("Install missing extensions");
  const action = await vscode.window.showWarningMessage(
    vscode.l10n.t("AOSP navigation extensions are missing: {0}", details), install,
  );
  if (action !== install) return;
  for (const dependency of missing) {
    try {
      await vscode.commands.executeCommand("workbench.extensions.installExtension", dependency.id);
    } catch (error) {
      output.appendLine(`extension install failed (${dependency.id}): ${String(error)}`);
    }
  }
}
