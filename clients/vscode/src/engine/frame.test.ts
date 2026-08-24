import assert from "node:assert/strict";
import test from "node:test";
import { encodeFrame, FrameDecoder } from "./frame";

test("decodes split Content-Length frames", () => {
  const decoder = new FrameDecoder();
  const frame = encodeFrame(JSON.stringify({ method: "aosp/initialize" }));
  assert.deepEqual(decoder.push(frame.subarray(0, 9)), []);
  assert.deepEqual(decoder.push(frame.subarray(9)), [JSON.stringify({ method: "aosp/initialize" })]);
});

test("decodes multiple frames in one chunk", () => {
  const decoder = new FrameDecoder();
  const first = JSON.stringify({ id: 1 });
  const second = JSON.stringify({ id: 2 });
  assert.deepEqual(decoder.push(Buffer.concat([encodeFrame(first), encodeFrame(second)])), [first, second]);
});
