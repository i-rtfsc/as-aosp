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
import { applyRustProjection } from "./navigation/rust_projection";
import { normalizeOutRoot } from "./navigation/out_root";

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
  const startIndex = async (
    rebuild: boolean,
    onProgress?: (status: { state: string; stage: string; progress: number; modules: number; generated: number; diagnostics: number }) => void,
  ) => {
    await ensureEngine();
    const config = vscode.workspace.getConfiguration("aospIde");
    const profile = {
      schemaVersion: 1,
      id: vscode.workspace.name || "vscode-workspace",
      workspaceViewRoot: config.get<string>("workspaceViewRoot") || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
      outRoot: normalizeOutRoot(config.get<string>("outRoot") || ""),
      targets: config.get<unknown[]>("targets") || [],
      pathMappings: config.get<unknown[]>("pathMappings") || [],
      buildIntegration: { mode: "read-only" },
    };
    const presetInspection = config.inspect<string[]>("selectedPresets");
    const selectionConfigured = presetInspection?.workspaceValue !== undefined;
    if ((!selectionConfigured || profile.targets.length === 0) &&
      profile.workspaceViewRoot) {
      const catalog = await engine.selectionCatalog();
      const presetIds = !selectionConfigured ?
        catalog.defaultPresetIds : config.get<string[]>("selectedPresets") ?? [];
      const manualModules = !selectionConfigured ? [] :
        config.get<string[]>("selectedModules") ?? [];
      const plan = await engine.selectionPlan({
        workspaceViewRoot: profile.workspaceViewRoot,
        outRoot: profile.outRoot,
        cacheRoot: context.storageUri?.fsPath,
        presetIds,
        manualModules,
      });
      profile.targets = plan.targets;
      await config.update("selectedPresets", plan.presetIds, vscode.ConfigurationTarget.Workspace);
      await config.update("selectedModules", manualModules, vscode.ConfigurationTarget.Workspace);
      await config.update("targets", plan.targets, vscode.ConfigurationTarget.Workspace);
    }
    if (profile.targets.length === 0) throw new Error("No Android.bp modules selected for indexing.");
    const status = await engine.startIndex(profile, rebuild);
    let current = status;
    const describe = () => ({
      state: current.state,
      stage: current.stage,
      progress: current.progress,
      modules: current.moduleCount,
      generated: current.generatedCount,
      diagnostics: current.diagnostics?.length ?? 0,
    });
    onProgress?.(describe());
    for (let attempt = 0; attempt < 120 && current.state === "running"; attempt++) {
      await new Promise((resolve) => setTimeout(resolve, 500));
      current = await engine.indexStatus();
      if (current.state === "running") onProgress?.(describe());
    }
    if (current.state === "running") {
      throw new Error("AOSP build facts did not finish loading within 60 seconds.");
    }
    if (current.state === "ready" || current.state === "partial") {
      if (profile.outRoot && profile.workspaceViewRoot) {
        onProgress?.({
          ...describe(), state: "running", stage: "projection", progress: 90,
        });
        const cacheRoot = context.storageUri?.fsPath || context.globalStorageUri.fsPath;
        const projection = await applyJavaProjection(engine, profile.workspaceViewRoot, profile.outRoot, cacheRoot, profile.targets);
        const rust = await applyRustProjection(
          engine, profile.workspaceViewRoot, profile.outRoot, cacheRoot,
          profile.pathMappings,
        );
        current = await engine.reportProjection({
          profileId: profile.id,
          sourceRoots: projection.sourcePaths.length,
          generatedRoots: projection.generatedRoots,
          libraries: projection.referencedLibraries.length,
          skipped: projection.skipped,
          diagnostics: rust.diagnostics,
        });
      }
    }
    return describe();
  };
  const profileView = new ProfileViewProvider(context, async (workspaceRoot, outRoot, query, limit) => {
    await ensureEngine();
    const result = await engine.sourceModules(
      workspaceRoot, outRoot, context.storageUri?.fsPath, query, limit,
    );
    return {
      modules: result.modules,
      diagnostics: result.diagnostics?.length ?? 0,
      filesScanned: result.filesScanned,
      cacheHit: result.cacheHit,
      buildBacked: result.buildBacked,
      total: result.total,
    };
  }, async () => {
    await ensureEngine();
    return engine.selectionCatalog();
  }, async (params) => {
    await ensureEngine();
    return engine.selectionPlan(params);
  }, startIndex);
  context.subscriptions.push(vscode.commands.registerCommand("aospIde.refreshModules", async () => {
    await ensureEngine();
    const root = vscode.workspace.getConfiguration("aospIde").get<string>("workspaceViewRoot")?.trim()
      || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "";
    const outRoot = normalizeOutRoot(
      vscode.workspace.getConfiguration("aospIde").get<string>("outRoot") || "",
    );
    const result = await engine.sourceModules(root, outRoot, context.storageUri?.fsPath);
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
      const status = await startIndex(rebuild, (progress) => {
        if (generation !== autoGeneration) return;
        profileView.updateIndexStatus(
          `${progress.state} · ${progress.stage} · ${progress.progress}% · ` +
          `${progress.modules} modules · ${progress.generated} generated · ` +
          `${progress.diagnostics} diagnostics`,
        );
      });
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
      output.appendLine(`outRoot: ${normalizeOutRoot(config.get<string>("outRoot") || "") || "<unset>"}`);
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
  const configuredOut = vscode.workspace.getConfiguration("aospIde").get<string>("outRoot") || "";
  const normalizedOut = normalizeOutRoot(configuredOut);
  if (configuredOut && configuredOut !== normalizedOut) {
    void vscode.workspace.getConfiguration("aospIde").update(
      "outRoot", normalizedOut, vscode.ConfigurationTarget.Workspace,
    );
  }
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
