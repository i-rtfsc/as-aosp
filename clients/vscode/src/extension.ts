import * as vscode from "vscode";
import { EngineClient } from "./engine/client";
import { AidlDefinitionProvider } from "./navigation/aidl";
import { goToAidlGeneratedPeer } from "./navigation/aidl_peers";
import { AidlReferenceProvider } from "./navigation/aidl_references";
import { goToJniPeer, JniDefinitionProvider } from "./navigation/jni";
import { refreshCompdbOverlay } from "./navigation/compdb";
import { createAospStatusBar } from "./status";
import { resolveEngineExecutable } from "./engine/resolve";
import { openProfilePanel } from "./ui/profile_panel";

export function activate(context: vscode.ExtensionContext): void {
  const engine = new EngineClient();
  const output = vscode.window.createOutputChannel("AOSP IDE");
  context.subscriptions.push(output);
  context.subscriptions.push(createAospStatusBar(engine, output));
  const ensureEngine = async (): Promise<void> => {
    const configured = vscode.workspace.getConfiguration("aospIde").get<string>("engineExecutable");
    const executable = resolveEngineExecutable({ configured, extensionPath: context.extensionPath });
    await engine.start(executable);
  };
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
      output.appendLine(`status: ${engine.status.state}${engine.status.error ? ` (${engine.status.error})` : ""}`);
      output.show(true);
    }),
    vscode.commands.registerCommand("aospIde.openProfile", () => openProfilePanel(context)),
    vscode.workspace.onDidChangeConfiguration((event) => {
      if (event.affectsConfiguration("aospIde.engineExecutable") ||
        event.affectsConfiguration("aospIde.workspaceViewRoot") ||
        event.affectsConfiguration("aospIde.metadataRoot") ||
        event.affectsConfiguration("aospIde.pathMappings")) {
        void engine.stop();
      }
    }),
    { dispose: () => void engine.stop() },
  );
  const configuredRoot = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot")?.trim();
  if (vscode.workspace.workspaceFolders?.length && !configuredRoot) {
    setTimeout(() => openProfilePanel(context), 300);
  }
}

export function deactivate(): void {
  // The extension disposable stops the session.
}
