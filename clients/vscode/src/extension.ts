import * as vscode from "vscode";
import { EngineClient } from "./engine/client";
import { AidlDefinitionProvider } from "./navigation/aidl";
import { goToAidlGeneratedPeer } from "./navigation/aidl_peers";
import { AidlReferenceProvider } from "./navigation/aidl_references";
import { goToJniPeer } from "./navigation/jni";

export function activate(context: vscode.ExtensionContext): void {
  const engine = new EngineClient();
  const ensureEngine = async (): Promise<void> => {
    const executable = vscode.workspace.getConfiguration("aospIde").get<string>("engineExecutable") || "aosp-index";
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
    vscode.commands.registerCommand("aospIde.goToJniPeer", async () => {
      await ensureEngine();
      await goToJniPeer(engine);
    }),
    vscode.commands.registerCommand("aospIde.goToAidlGeneratedPeer", async () => {
      await ensureEngine();
      await goToAidlGeneratedPeer(engine);
    }),
    vscode.workspace.onDidChangeConfiguration((event) => {
      if (event.affectsConfiguration("aospIde.engineExecutable") ||
        event.affectsConfiguration("aospIde.workspaceViewRoot")) {
        void engine.stop();
      }
    }),
    { dispose: () => void engine.stop() },
  );
}

export function deactivate(): void {
  // The extension disposable stops the session.
}
