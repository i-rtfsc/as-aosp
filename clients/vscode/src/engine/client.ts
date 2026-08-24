import { ChildProcessWithoutNullStreams, spawn } from "node:child_process";
import { encodeFrame, FrameDecoder } from "./frame";

export interface AidlResolveResult {
  resolution: {
    state: string;
    symbols?: Array<{ path: string; range: { start: { offset: number } } }>;
  };
  diagnostics: unknown[];
}

export interface JniPeersResult {
  peers: Array<{
    java: { path: string };
    native: { path: string; range: { start: number } };
  }>;
  diagnostics: unknown[];
}

export class EngineClient {
  private process?: ChildProcessWithoutNullStreams;
  private decoder = new FrameDecoder();
  private nextId = 1;
  private pending = new Map<number, { resolve: (value: any) => void; reject: (error: Error) => void }>();
  private queue = Promise.resolve();

  async start(executable: string): Promise<void> {
    if (this.process) return;
    this.process = spawn(executable, ["serve", "--stdio"], { stdio: "pipe" });
    this.process.stdout.on("data", (chunk: Buffer) => this.onData(chunk));
    this.process.on("error", (error) => this.failPending(error));
    this.process.on("exit", () => this.failPending(new Error("aosp-index exited")));
    await this.request("aosp/initialize", { protocolVersion: "1.0" });
  }

  async stop(): Promise<void> {
    if (!this.process) return;
    try {
      await this.request("aosp/shutdown", null);
    } catch {
      // The process may have already closed its protocol streams.
    }
    this.process.kill();
    this.process = undefined;
  }

  async aidlResolve(params: { workspaceViewRoot: string; currentPath: string; name: string }): Promise<AidlResolveResult> {
    return this.request("aosp/aidl/resolve", params);
  }

  async jniPeers(params: { workspaceViewRoot: string; currentPath: string }): Promise<JniPeersResult> {
    return this.request("aosp/jni/peers", params);
  }

  private request(method: string, params: unknown): Promise<any> {
    const operation = this.queue.then(() => new Promise((resolve, reject) => {
      if (!this.process) {
        reject(new Error("aosp-index is not running"));
        return;
      }
      const id = this.nextId++;
      this.pending.set(id, { resolve, reject });
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
      if (response.error) pending.reject(new Error(response.error.message ?? "aosp-index request failed"));
      else pending.resolve(response.result);
    }
  }

  private failPending(error: Error): void {
    for (const pending of this.pending.values()) pending.reject(error);
    this.pending.clear();
  }
}
