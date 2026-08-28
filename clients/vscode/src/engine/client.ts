import { ChildProcessWithoutNullStreams, spawn } from "node:child_process";
import { encodeFrame, FrameDecoder } from "./frame";

type SpawnProcess = (executable: string, arguments_: string[]) => ChildProcessWithoutNullStreams;
type PendingRequest = {
  resolve: (value: any) => void;
  reject: (error: Error) => void;
  timeout: NodeJS.Timeout;
};

export interface AidlResolveResult {
  resolution: {
    state: string;
    symbols?: Array<{ path: string; range: { start: { offset: number } } }>;
  };
  diagnostics: unknown[];
}

export interface JniPeersResult {
  peers: Array<{
    java: { path: string; range: { start: number } };
    native: { path: string; range: { start: number } };
  }>;
  diagnostics: unknown[];
}


export interface AidlPeersResult {
  peers: Array<{
    path: { raw: string; repoPath?: string };
    localPath?: string;
    confidence: string;
    exists: boolean;
  }>;
  diagnostics: unknown[];
}

export interface AidlReferencesResult {
  references: Array<{ path: string; range: { start: { offset: number }; end: { offset: number } } }>;
  diagnostics: unknown[];
}


export interface CompdbOverlayResult {
  entries: Array<{ directory: string; arguments: string[]; file: string; output?: string }>;
  responseFiles?: Record<string, string>;
  diagnostics: unknown[];
}

export interface ModulesCatalogResult {
  modules: Array<{ name: string; sourceRoots: number; generated: number; dependencies: number }>;
  diagnostics: unknown[];
}

export interface ModuleFact {
  name: string;
  sourceRoots?: Array<{ raw: string; kind: string; repoPath?: string }>;
  sources?: Array<{ raw: string; kind: string; repoPath?: string }>;
  generated?: Array<{ raw: string; kind: string; repoPath?: string }>;
  jars?: Array<{ raw: string; kind: string; repoPath?: string }>;
  srcjars?: Array<{ raw: string; kind: string; repoPath?: string }>;
}

export interface ModulesClosureResult {
  modules: ModuleFact[];
  diagnostics: unknown[];
}

export interface ProjectGraphResult {
  graph: {
    targets: Array<{
      spec: { id: string; kind: string };
      resolvedModules: string[];
      resolvedArtifacts: Array<{ id: string; kind: string; name: string; path: { raw: string; repoPath?: string } }>;
      sourceRoots: Array<{ raw: string; kind: string; repoPath?: string }>;
      excludedSourceRoots: Array<{ raw: string; kind: string; repoPath?: string }>;
      generatedRoots: Array<{ raw: string; kind: string; repoPath?: string }>;
      libraries: Array<{ raw: string; kind: string; repoPath?: string }>;
      dependencies: string[];
      targetDependencies: string[];
      diagnostics: unknown[];
    }>;
    ownership: Array<{ repoPath: string; targetId: string }>;
    duplicateSourceDecisions: string[];
    diagnostics: unknown[];
  };
  diagnostics: unknown[];
}

export interface ArtifactsCatalogResult {
  artifacts: Array<{
    id: string;
    kind: string;
    name: string;
    path: { raw: string; repoPath?: string };
    ownerModules?: string[];
    size?: number;
    sha256?: string;
  }>;
  diagnostics: unknown[];
}

export interface SourceModuleResult {
  modules: Array<{
    name: string;
    type: string;
    blueprint: string;
    directory: string;
    sourceRoots?: string[];
    sources?: string[];
    dependencies?: string[];
    status: string;
    capabilities?: string[];
    recommended: boolean;
  }>;
  diagnostics: unknown[] | null;
  filesScanned: number;
  cacheHit: boolean;
  buildBacked: number;
  total: number;
}

export interface RustProjectResult {
  projectPath?: string;
  crates: number;
  diagnostics: unknown[];
}

export interface IndexSnapshot {
  profileId?: string;
  state: string;
  stage: string;
  completedStages: string[];
  totalStages: number;
  progress: number;
  moduleCount: number;
  generatedCount: number;
  sourceRootCount: number;
  generatedRootCount: number;
  libraryCount: number;
  skippedCount: number;
  diagnostics: unknown[];
}

export interface GeneratedMaterializeResult {
  sourceRoots: Array<{ archive: string; root: string; hash: string; entries: number }>;
  diagnostics: unknown[];
}

export class EngineClient {
  private process?: ChildProcessWithoutNullStreams;
  private startPromise?: Promise<void>;
  private stopPromise?: Promise<void>;
  private decoder = new FrameDecoder();
  private nextId = 1;
  private pending = new Map<number, PendingRequest>();
  private queue = Promise.resolve();
  private state: "stopped" | "starting" | "running" | "stopping" | "error" = "stopped";
  private lastError?: string;
  private listeners = new Set<(state: string, error?: string) => void>();

  constructor(
    private readonly spawnProcess: SpawnProcess = (executable, arguments_) =>
      spawn(executable, arguments_, { stdio: "pipe" }),
    private readonly timeoutMilliseconds = 60_000,
  ) {}

  get status(): { state: string; error?: string } { return { state: this.state, error: this.lastError }; }
  onStatus(listener: (state: string, error?: string) => void): { dispose: () => void } {
    this.listeners.add(listener);
    return { dispose: () => this.listeners.delete(listener) };
  }

  async start(executable: string): Promise<void> {
    if (this.stopPromise) {
      await this.stopPromise;
    }
    if (this.startPromise) return this.startPromise;
    if (this.process && this.state === "running") return;
    this.setStatus("starting");
    const operation = this.launch(executable);
    this.startPromise = operation;
    try {
      await operation;
    } finally {
      if (this.startPromise === operation) this.startPromise = undefined;
    }
  }

  private async launch(executable: string): Promise<void> {
    const process = this.spawnProcess(executable, ["serve", "--stdio"]);
    this.process = process;
    process.stdout.on("data", (chunk: Buffer) => this.onData(chunk));
    process.on("error", (error) => this.handleExit(process, error));
    process.on("exit", () => this.handleExit(process, new Error("aosp-index exited")));
    try {
      await this.request("aosp/initialize", { protocolVersion: "1.0" });
      this.setStatus("running");
    } catch (error) {
      process.kill();
      this.handleExit(process, error instanceof Error ? error : new Error(String(error)));
      throw error;
    }
  }

  async stop(): Promise<void> {
    if (this.stopPromise) return this.stopPromise;
    const operation = this.performStop();
    this.stopPromise = operation;
    try {
      await operation;
    } finally {
      if (this.stopPromise === operation) this.stopPromise = undefined;
    }
  }

  private async performStop(): Promise<void> {
    this.setStatus("stopping");
    try {
      await this.startPromise;
    } catch {
      return;
    }
    const process = this.process;
    if (!process) return;
    try {
      await this.request("aosp/shutdown", null);
    } catch {
      // The process may have already closed its protocol streams.
    }
    process.kill();
    this.handleExit(process, new Error("aosp-index stopped"));
    this.setStatus("stopped");
  }

  async aidlResolve(params: { workspaceViewRoot: string; currentPath: string; name: string }): Promise<AidlResolveResult> {
    return this.request("aosp/aidl/resolve", params);
  }


  async jniPeers(params: { workspaceViewRoot: string; currentPath: string }): Promise<JniPeersResult> {
    return this.request("aosp/jni/peers", params);
  }


  async aidlPeers(params: { workspaceViewRoot: string; currentPath: string; name: string }): Promise<AidlPeersResult> {
    return this.request("aosp/aidl/peers", params);
  }

  async aidlReferences(params: { workspaceViewRoot: string; currentPath: string; name: string }): Promise<AidlReferencesResult> {
    return this.request("aosp/aidl/references", params);
  }

  async compdbOverlay(params: { outRoot?: string; workspaceViewRoot: string; pathMappings: unknown[] }): Promise<CompdbOverlayResult> {
    return this.request("aosp/compdb/overlay", params);
  }

  async modulesCatalog(outRoot: string): Promise<ModulesCatalogResult> {
    return this.request("aosp/modules/catalog", { outRoot });
  }

  async projectGraph(outRoot: string, targets: unknown[]): Promise<ProjectGraphResult> {
    return this.request("aosp/projects/graph", { outRoot, targets });
  }

  async artifactsCatalog(outRoot: string, kind?: string, query?: string): Promise<ArtifactsCatalogResult> {
    return this.request("aosp/artifacts/catalog", { outRoot, kind, query });
  }

  async sourceModules(workspaceViewRoot: string, outRoot?: string, cacheRoot?: string, query?: string, limit = 500): Promise<SourceModuleResult> {
    return this.request("aosp/source/modules", { workspaceViewRoot, outRoot, cacheRoot, query, limit });
  }

  async rustProject(params: { workspaceViewRoot: string; outRoot: string; cacheRoot: string; pathMappings: unknown[] }): Promise<RustProjectResult> {
    return this.request("aosp/rust/project", params);
  }

  async startIndex(profile: unknown, rebuild = false): Promise<IndexSnapshot> {
    return this.request(rebuild ? "aosp/index/rebuild" : "aosp/index/start", { profile });
  }

  async indexStatus(): Promise<IndexSnapshot> {
    return this.request("aosp/index/status", null);
  }

  async reportProjection(params: { profileId: string; sourceRoots: number; generatedRoots: number; libraries: number; skipped: number; diagnostics: unknown[] }): Promise<IndexSnapshot> {
    return this.request("aosp/index/projection", params);
  }

  async materializeGenerated(outRoot: string, cacheRoot: string, archives?: string[]): Promise<GeneratedMaterializeResult> {
    return this.request("aosp/generated/materialize", { outRoot, archives, cacheRoot });
  }


  private request(method: string, params: unknown): Promise<any> {
    const operation = this.queue.then(() => new Promise((resolve, reject) => {
      if (!this.process) {
        reject(new Error("aosp-index is not running"));
        return;
      }
      const id = this.nextId++;
      const timeout = setTimeout(() => {
        if (!this.pending.delete(id)) return;
        reject(new Error(`aosp-index request timed out: ${method}`));
      }, this.timeoutMilliseconds);
      this.pending.set(id, { resolve, reject, timeout });
      this.process.stdin.write(encodeFrame(JSON.stringify({ jsonrpc: "2.0", id, method, params })));
    }));
    this.queue = operation.then(() => undefined, () => undefined);
    return operation;
  }

  private onData(chunk: Buffer): void {
    for (const message of this.decoder.push(chunk)) {
      const response = JSON.parse(message) as { id: number; result?: unknown; error?: { message?: string } };
      const pending = this.pending.get(response.id);
      if (!pending) continue;
      this.pending.delete(response.id);
      clearTimeout(pending.timeout);
      if (response.error) pending.reject(new Error(response.error.message ?? "aosp-index request failed"));
      else pending.resolve(response.result);
    }
  }

  private failPending(error: Error): void {
    for (const pending of this.pending.values()) {
      clearTimeout(pending.timeout);
      pending.reject(error);
    }
    this.pending.clear();
  }

  private handleExit(process: ChildProcessWithoutNullStreams, error: Error): void {
    if (this.process !== process) return;
    this.process = undefined;
    this.failPending(error);
    this.setStatus("error", error.message);
  }

  private setStatus(state: "stopped" | "starting" | "running" | "stopping" | "error", error?: string): void {
    this.state = state;
    this.lastError = error;
    for (const listener of this.listeners) listener(state, error);
  }
}
