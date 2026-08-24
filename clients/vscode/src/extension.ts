import * as vscode from "vscode";
import { EngineClient } from "./engine/client";
import { AidlDefinitionProvider } from "./navigation/aidl";
import { goToJniPeer } from "./navigation/jni";

export function activate(context: vscode.ExtensionContext): void {
  const engine = new EngineClient();
  const executable = vscode.workspace.getConfiguration("aospIde").get<string>("engineExecutable") || "aosp-index";
  const ensureEngine = async (): Promise<void> => engine.start(executable);
  context.subscriptions.push(
    vscode.languages.registerDefinitionProvider({ language: "aidl" }, {
      provideDefinition: async (...args) => {
        await ensureEngine();
        return new AidlDefinitionProvider(engine).provideDefinition(...args);
      },
    }),
    vscode.commands.registerCommand("aospIde.goToJniPeer", async () => {
      await ensureEngine();
      await goToJniPeer(engine);
    }),
    { dispose: () => void engine.stop() },
  );
}

export function deactivate(): void {
  // The extension disposable stops the session.
}
