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
      outRoot: config.get<string>("metadataRoot") || "",
      metadataPackRoot: config.get<string>("metadataPackRoot") || "",
      targets: config.get<unknown[]>("targets") || [],
      pathMappings: config.get<unknown[]>("pathMappings") || [],
      buildIntegration: { mode: "read-only" },
    };
    if (profile.targets.length === 0) throw new Error("Configure at least one TargetSpec before indexing.");
    if (profile.metadataPackRoot) {
      const validation = await engine.validatePack(profile.metadataPackRoot);
      if (!validation.valid) throw new Error(validation.diagnostics.map((item) => `${item.code}: ${item.message}`).join("; "));
    }
    const status = await engine.startIndex(profile, rebuild);
    let current = status;
    for (let attempt = 0; attempt < 40 && current.state === "running"; attempt++) {
      await new Promise((resolve) => setTimeout(resolve, 50));
      current = await engine.indexStatus();
    }
    if (current.state === "ready" || current.state === "partial") {
      const packRoot = profile.metadataPackRoot;
      if ((packRoot || profile.outRoot) && profile.workspaceViewRoot) {
        const projection = await applyJavaProjection(engine, profile.workspaceViewRoot, packRoot, context.globalStorageUri.fsPath, profile.outRoot, profile.targets);
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
  const profileView = new ProfileViewProvider(context, async (metadataPackRoot, metadataRoot) => {
    await ensureEngine();
    metadataRoot = metadataRoot || vscode.workspace.getConfiguration("aospIde").get<string>("metadataRoot") || "";
    if (metadataPackRoot) {
      const validation = await engine.validatePack(metadataPackRoot);
      if (!validation.valid) {
        throw new Error(validation.diagnostics.map((item) => `${item.code}: ${item.message}`).join("; "));
      }
    }
    const modules = await engine.modulesCatalog(metadataPackRoot, metadataRoot);
    const artifacts = await engine.artifactsCatalog(metadataPackRoot, undefined, undefined, metadataRoot);
    return {
      modules: modules.modules.length,
      generated: modules.modules.reduce((total, module) => total + module.generated, 0),
      diagnostics: modules.diagnostics.length + artifacts.diagnostics.length,
      artifacts: artifacts.artifacts,
    };
  }, startIndex, async (metadataPackRoot, metadataRoot, target) => {
    await ensureEngine();
    metadataRoot = metadataRoot || vscode.workspace.getConfiguration("aospIde").get<string>("metadataRoot") || "";
    if (!metadataPackRoot?.trim() && !metadataRoot.trim()) {
      throw new Error("Metadata Pack Root or Out Root is required to resolve a target.");
    }
    const result = await engine.projectGraph(metadataPackRoot, [target], metadataRoot);
    const project = result.graph.targets[0];
    const diagnostics = [
      ...result.diagnostics,
      ...(project?.diagnostics ?? []),
    ];
    return {
      state: diagnostics.length ? "error" : project ? "exact" : "unresolved",
      modules: project?.resolvedModules.length ?? 0,
      artifacts: project?.resolvedArtifacts.length ?? 0,
      sources: project?.sourceRoots.length ?? 0,
      generated: project?.generatedRoots.length ?? 0,
      libraries: project?.libraries.length ?? 0,
      diagnostics: diagnostics.length,
    };
  });
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
      await goToAidlGeneratedPeer(engine, context.globalStorageUri.fsPath);
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
      if (detection.root && !vscode.workspace.getConfiguration("aospIde").get<string>("metadataRoot")?.trim()) {
        await vscode.workspace.getConfiguration("aospIde").update("metadataRoot", detection.root, vscode.ConfigurationTarget.Workspace);
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
