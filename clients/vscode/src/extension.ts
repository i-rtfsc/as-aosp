import * as vscode from "vscode";
import { EngineClient } from "./engine/client";
import { AidlDefinitionProvider } from "./navigation/aidl";
import { goToAidlGeneratedPeer } from "./navigation/aidl_peers";
import { AidlReferenceProvider } from "./navigation/aidl_references";
import { goToJniPeer, JniDefinitionProvider } from "./navigation/jni";
import { refreshCompdbOverlay } from "./navigation/compdb";
import { createAospStatusBar } from "./status";
import { resolveEngineExecutable } from "./engine/resolve";
import { ProfileViewProvider } from "./ui/profile_panel";
import { detectOut } from "./ui/overview";
import { JavaSourceDefinitionProvider } from "./navigation/java";
import { applyJavaProjection } from "./navigation/java_projection";

export function activate(context: vscode.ExtensionContext): void {
  const engine = new EngineClient();
  const output = vscode.window.createOutputChannel("AOSP Workspace");
  output.appendLine("AOSP Workspace extension activated");
  context.subscriptions.push(output);
  const ensureEngine = async (): Promise<void> => {
    const configured = vscode.workspace.getConfiguration("aospIde").get<string>("engineExecutable");
    const executable = resolveEngineExecutable({ configured, extensionPath: context.extensionPath });
    await engine.start(executable);
  };
  const startIndex = async (rebuild: boolean) => {
    await ensureEngine();
    const config = vscode.workspace.getConfiguration("aospIde");
    const profile = {
      schemaVersion: 1,
      id: vscode.workspace.name || "vscode-workspace",
      workspaceViewRoot: config.get<string>("workspaceViewRoot") || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
      outRoot: config.get<string>("outRoot") || "",
      targets: config.get<unknown[]>("targets") || [],
      pathMappings: config.get<unknown[]>("pathMappings") || [],
      buildIntegration: { mode: "read-only" },
    };
    if (profile.targets.length === 0 && profile.workspaceViewRoot) {
      const discovered = await engine.sourceModules(profile.workspaceViewRoot, context.storageUri?.fsPath);
      const recommended = discovered.modules.filter((module) => module.recommended);
      if (recommended.length > 0) {
        profile.targets = recommended.map((module) => ({
          id: module.name,
          displayName: module.name,
          kind: module.type.toLowerCase().includes("cc_") ? "native" :
            module.type.toLowerCase().includes("android_") ? "android" : "java",
          selectors: { modules: [module.name] },
          includeDependencies: true,
          includeGeneratedSources: true,
          includeResources: module.type.toLowerCase().includes("android_"),
          includeManifest: module.type.toLowerCase().includes("android_"),
        }));
        await config.update("targets", profile.targets, vscode.ConfigurationTarget.Workspace);
        await config.update("selectedModules", recommended.map((module) => module.name), vscode.ConfigurationTarget.Workspace);
      }
    }
    if (profile.targets.length === 0) throw new Error("No Android.bp modules selected for indexing.");
    const status = await engine.startIndex(profile, rebuild);
    let current = status;
    for (let attempt = 0; attempt < 40 && current.state === "running"; attempt++) {
      await new Promise((resolve) => setTimeout(resolve, 50));
      current = await engine.indexStatus();
    }
    if (current.state === "ready" || current.state === "partial") {
      if (profile.outRoot && profile.workspaceViewRoot) {
        const cacheRoot = context.storageUri?.fsPath || context.globalStorageUri.fsPath;
        const projection = await applyJavaProjection(engine, profile.workspaceViewRoot, profile.outRoot, cacheRoot, profile.targets);
        current = await engine.reportProjection({
          profileId: profile.id,
          sourceRoots: projection.sourcePaths.length,
          generatedRoots: projection.generatedRoots,
          libraries: projection.referencedLibraries.length,
          skipped: projection.skipped,
          diagnostics: [],
        });
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
  };
  const profileView = new ProfileViewProvider(context, async (workspaceRoot, query) => {
    await ensureEngine();
    const result = await engine.sourceModules(workspaceRoot, context.storageUri?.fsPath, query);
    return {
      modules: result.modules,
      diagnostics: result.diagnostics.length,
      filesScanned: result.filesScanned,
      cacheHit: result.cacheHit,
    };
  }, startIndex);
  context.subscriptions.push(vscode.commands.registerCommand("aospIde.refreshModules", async () => {
    await ensureEngine();
    const root = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot")?.trim()
      || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "";
    const result = await engine.sourceModules(root, context.storageUri?.fsPath);
    output.appendLine(`source modules: ${result.modules.length}; files: ${result.filesScanned}; cache: ${result.cacheHit}`);
    output.show(true);
  }));
  let autoGeneration = 0;
  let autoTimer: NodeJS.Timeout | undefined;
  const runAutomaticIndex = async (rebuild = false): Promise<void> => {
    const generation = ++autoGeneration;
    const root = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot")?.trim()
      || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath;
    if (!root) return;
    profileView.updateIndexStatus("Starting automatically…");
    try {
      const status = await startIndex(rebuild);
      if (generation !== autoGeneration) return;
      profileView.updateIndexStatus(`${status.state} · ${status.stage} · ${status.progress}% · ${status.modules} modules · ${status.generated} generated · ${status.diagnostics} diagnostics`);
    } catch (error) {
      if (generation !== autoGeneration) return;
      profileView.updateIndexStatus(`Index unavailable: ${error instanceof Error ? error.message : String(error)}`);
    }
  };
  const scheduleAutomaticIndex = (rebuild = false): void => {
    if (autoTimer) clearTimeout(autoTimer);
    autoTimer = setTimeout(() => void runAutomaticIndex(rebuild), 300);
  };
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
      await goToAidlGeneratedPeer(engine, context.storageUri?.fsPath || context.globalStorageUri.fsPath);
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
      output.appendLine(`outRoot: ${config.get<string>("outRoot") || "<unset>"}`);
      output.appendLine(`status: ${engine.status.state}${engine.status.error ? ` (${engine.status.error})` : ""}`);
      output.show(true);
    }),
    vscode.commands.registerCommand("aospIde.openProfile", async () => profileView.reveal()),
    vscode.workspace.onDidChangeConfiguration((event) => {
      if (event.affectsConfiguration("aospIde.engineExecutable") ||
        event.affectsConfiguration("aospIde.workspaceViewRoot") ||
        event.affectsConfiguration("aospIde.outRoot") ||
        event.affectsConfiguration("aospIde.pathMappings")) {
        autoGeneration++;
        void engine.stop().finally(() => scheduleAutomaticIndex(true));
      }
    }),
    { dispose: () => { if (autoTimer) clearTimeout(autoTimer); void engine.stop(); } },
  );
  const configuredRoot = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot")?.trim();
  if (vscode.workspace.workspaceFolders?.length && !configuredRoot) {
    const root = vscode.workspace.workspaceFolders[0].uri.fsPath;
    void detectOut(root).then(async (detection) => {
      if (detection.root && !vscode.workspace.getConfiguration("aospIde").get<string>("outRoot")?.trim()) {
        await vscode.workspace.getConfiguration("aospIde").update("outRoot", detection.root, vscode.ConfigurationTarget.Workspace);
      }
      setTimeout(() => profileView.reveal(), 300);
      scheduleAutomaticIndex(false);
    });
  } else if (configuredRoot) {
    scheduleAutomaticIndex(false);
  }
}

export function deactivate(): void {
  // The extension disposable stops the session.
}
