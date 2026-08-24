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

export class EngineClient {
  private process?: ChildProcessWithoutNullStreams;
  private startPromise?: Promise<void>;
  private decoder = new FrameDecoder();
  private nextId = 1;
  private pending = new Map<number, PendingRequest>();
  private queue = Promise.resolve();

  constructor(
    private readonly spawnProcess: SpawnProcess = (executable, arguments_) =>
      spawn(executable, arguments_, { stdio: "pipe" }),
    private readonly timeoutMilliseconds = 5_000,
  ) {}

  async start(executable: string): Promise<void> {
    if (this.startPromise) return this.startPromise;
    if (this.process) return;
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
    } catch (error) {
      process.kill();
      this.handleExit(process, error instanceof Error ? error : new Error(String(error)));
      throw error;
    }
  }

  async stop(): Promise<void> {
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

  async compdbOverlay(params: { metadataRoot: string; workspaceViewRoot: string; pathMappings: unknown[] }): Promise<CompdbOverlayResult> {
    return this.request("aosp/compdb/overlay", params);
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
  }
}
