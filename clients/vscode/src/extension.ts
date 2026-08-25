import * as vscode from "vscode";
import { EngineClient } from "./engine/client";
import { AidlDefinitionProvider } from "./navigation/aidl";
import { goToAidlGeneratedPeer } from "./navigation/aidl_peers";
import { AidlReferenceProvider } from "./navigation/aidl_references";
import { goToJniPeer, JniDefinitionProvider } from "./navigation/jni";
import { refreshCompdbOverlay } from "./navigation/compdb";
import { createAospStatusBar } from "./status";
import { resolveEngineExecutable } from "./engine/resolve";
import { openProfilePanel, ProfileViewProvider } from "./ui/profile_panel";
import { detectOut } from "./ui/overview";
import { JavaSourceDefinitionProvider } from "./navigation/java";
import { applyJavaProjection } from "./navigation/java_projection";

export function activate(context: vscode.ExtensionContext): void {
  const engine = new EngineClient();
  const output = vscode.window.createOutputChannel("AOSP IDE");
  output.appendLine("AOSP Workspace extension activated");
  context.subscriptions.push(output);
  const ensureEngine = async (): Promise<void> => {
    const configured = vscode.workspace.getConfiguration("aospIde").get<string>("engineExecutable");
    const executable = resolveEngineExecutable({ configured, extensionPath: context.extensionPath });
    await engine.start(executable);
  };
  const profileView = new ProfileViewProvider(context, async (metadataPackRoot) => {
    await ensureEngine();
    const result = await engine.modulesCatalog(metadataPackRoot);
    return {
      modules: result.modules.length,
      generated: result.modules.reduce((total, module) => total + module.generated, 0),
      diagnostics: result.diagnostics.length,
    };
  }, async (rebuild) => {
    const config = vscode.workspace.getConfiguration("aospIde");
    const profile = {
      schemaVersion: 1,
      id: vscode.workspace.name || "vscode-workspace",
      workspaceViewRoot: config.get<string>("workspaceViewRoot") || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
      outRoot: config.get<string>("metadataRoot") || "",
      metadataPackRoot: config.get<string>("metadataPackRoot") || "",
      pathMappings: config.get<unknown[]>("pathMappings") || [],
      buildIntegration: { mode: "read-only" },
    };
    const status = await engine.startIndex(profile, rebuild);
    if (status.state === "running") {
      for (let attempt = 0; attempt < 40; attempt++) {
        await new Promise((resolve) => setTimeout(resolve, 50));
        const current = await engine.indexStatus();
        if (current.state !== "running") {
          if (current.state === "ready" || current.state === "partial") {
            const packRoot = profile.metadataPackRoot;
            if (packRoot && profile.workspaceViewRoot) {
              await applyJavaProjection(engine, profile.workspaceViewRoot, packRoot, context.globalStorageUri.fsPath, profile.outRoot);
            }
          }
          return {
            state: current.state,
            stage: current.stage,
            progress: current.progress,
            modules: current.moduleCount,
            generated: current.generatedCount,
            diagnostics: current.diagnostics.length,
          };
        }
      }
    }
    return {
      state: status.state,
      stage: status.stage,
      progress: status.progress,
      modules: status.moduleCount,
      generated: status.generatedCount,
      diagnostics: status.diagnostics.length,
    };
  });
  context.subscriptions.push(vscode.window.registerWebviewViewProvider("aospWorkspace.setup", profileView));
  context.subscriptions.push(createAospStatusBar(engine, output));
  context.subscriptions.push(
    vscode.languages.registerDefinitionProvider({ language: "aidl" }, {
      provideDefinition: async (...args) => {
        await ensureEngine();
        return new AidlDefinitionProvider(engine).provideDefinition(...args);
      },
    }),
    vscode.languages.registerReferenceProvider({ language: "aidl" }, {
      provideReferences: async (...args) => {
        await ensureEngine();
        return new AidlReferenceProvider(engine).provideReferences(...args);
      },
    }),
    vscode.languages.registerDefinitionProvider(
      ["java", "kotlin", "c", "cpp", "objective-c", "objective-cpp"].map((language) => ({ language })), {
      provideDefinition: async (document, position, token) => {
        await ensureEngine();
        return new JniDefinitionProvider(engine).provideDefinition(document, position, token);
      },
      }),
    vscode.languages.registerDefinitionProvider(
      [{ language: "java" }, { language: "kotlin" }],
      new JavaSourceDefinitionProvider(),
    ),
    vscode.commands.registerCommand("aospIde.goToJniPeer", async () => {
      await ensureEngine();
      await goToJniPeer(engine);
    }),
    vscode.commands.registerCommand("aospIde.goToAidlGeneratedPeer", async () => {
      await ensureEngine();
      await goToAidlGeneratedPeer(engine);
    }),
    vscode.commands.registerCommand("aospIde.refreshCompdbOverlay", async () => {
      await ensureEngine();
      output.show(true);
      await refreshCompdbOverlay(engine, output);
    }),
    vscode.commands.registerCommand("aospIde.showStatus", () => {
      const config = vscode.workspace.getConfiguration("aospIde");
      output.appendLine(`engine: ${config.get<string>("engineExecutable") || "aosp-index"}`);
      output.appendLine(`workspaceViewRoot: ${config.get<string>("workspaceViewRoot") || "<workspace>"}`);
      output.appendLine(`metadataRoot: ${config.get<string>("metadataRoot") || "<unset>"}`);
      output.appendLine(`metadataPackRoot: ${config.get<string>("metadataPackRoot") || "<unset>"}`);
      output.appendLine(`status: ${engine.status.state}${engine.status.error ? ` (${engine.status.error})` : ""}`);
      output.show(true);
    }),
    vscode.commands.registerCommand("aospIde.openProfile", async () => profileView.reveal()),
    vscode.workspace.onDidChangeConfiguration((event) => {
      if (event.affectsConfiguration("aospIde.engineExecutable") ||
        event.affectsConfiguration("aospIde.workspaceViewRoot") ||
        event.affectsConfiguration("aospIde.metadataRoot") ||
        event.affectsConfiguration("aospIde.metadataPackRoot") ||
        event.affectsConfiguration("aospIde.pathMappings")) {
        void engine.stop();
      }
    }),
    { dispose: () => void engine.stop() },
  );
  const configuredRoot = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot")?.trim();
  if (vscode.workspace.workspaceFolders?.length && !configuredRoot) {
    const root = vscode.workspace.workspaceFolders[0].uri.fsPath;
    void detectOut(root).then(async (detection) => {
      if (detection.root && !vscode.workspace.getConfiguration("aospIde").get<string>("metadataRoot")?.trim()) {
        await vscode.workspace.getConfiguration("aospIde").update("metadataRoot", detection.root, vscode.ConfigurationTarget.Workspace);
      }
      setTimeout(() => profileView.reveal(), 300);
    });
  }
}

export function deactivate(): void {
  // The extension disposable stops the session.
}
