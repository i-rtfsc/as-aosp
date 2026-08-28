import assert from "node:assert/strict";
import { EventEmitter } from "node:events";
import { PassThrough } from "node:stream";
import test from "node:test";
import type { ChildProcessWithoutNullStreams } from "node:child_process";
import { EngineClient } from "./client";
import { encodeFrame, FrameDecoder } from "./frame";

class FakeProcess extends EventEmitter {
  readonly stdin = new PassThrough();
  readonly stdout = new PassThrough();
  readonly stderr = new PassThrough();
  killed = false;

  kill(): boolean {
    this.killed = true;
    return true;
  }

  asChild(): ChildProcessWithoutNullStreams {
    return this as unknown as ChildProcessWithoutNullStreams;
  }
}

function respondingProcess(responseFactory: (request: any) => any): FakeProcess {
  const process = new FakeProcess();
  const decoder = new FrameDecoder();
  process.stdin.on("data", (chunk: Buffer) => {
    for (const message of decoder.push(chunk)) {
      const request = JSON.parse(message);
      void Promise.resolve(responseFactory(request)).then((result) => {
        process.stdout.write(encodeFrame(JSON.stringify({
          jsonrpc: "2.0",
          id: request.id,
          result,
        })));
      });
    }
  });
  return process;
}

test("shares one initialize across concurrent starts", async () => {
  const process = respondingProcess(() => ({}));
  let spawns = 0;
  const client = new EngineClient(() => {
    spawns++;
    return process.asChild();
  });
  await Promise.all([client.start("aosp-index"), client.start("aosp-index")]);
  assert.equal(spawns, 1);
  await client.stop();
});

test("cleans a process after initialize timeout and allows restart", async () => {
  const stalled = new FakeProcess();
  const healthy = respondingProcess(() => ({}));
  const processes = [stalled, healthy];
  const client = new EngineClient(() => processes.shift()!.asChild(), 10);
  await assert.rejects(client.start("aosp-index"), /timed out/);
  assert.equal(stalled.killed, true);
  await client.start("aosp-index");
  await client.stop();
});

test("restarts after an unexpected process exit", async () => {
  const first = respondingProcess(() => ({}));
  const second = respondingProcess(() => ({}));
  let spawns = 0;
  const client = new EngineClient(() => {
    spawns++;
    return (spawns === 1 ? first : second).asChild();
  });
  await client.start("aosp-index");
  first.emit("exit", 1);
  await client.start("aosp-index");
  assert.equal(spawns, 2);
  await client.stop();
});

test("waits for an active stop before restarting", async () => {
  let releaseShutdown: (() => void) | undefined;
  const first = respondingProcess((request) => {
    if (request.method === "aosp/shutdown") {
      return new Promise<void>((resolve) => { releaseShutdown = resolve; });
    }
    return {};
  });
  const second = respondingProcess(() => ({}));
  let spawns = 0;
  const client = new EngineClient(() => {
    spawns++;
    return (spawns === 1 ? first : second).asChild();
  });
  await client.start("aosp-index");
  const stopping = client.stop();
  await new Promise((resolve) => setImmediate(resolve));
  const restarting = client.start("aosp-index");
  assert.equal(spawns, 1);
  releaseShutdown?.();
  await stopping;
  await restarting;
  assert.equal(spawns, 2);
  await client.stop();
});
