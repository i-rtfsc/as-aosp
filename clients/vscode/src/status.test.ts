import assert from "node:assert/strict";
import test from "node:test";

function label(state: string): string {
  return state === "running" ? "AOSP: Ready" : state === "starting" ? "AOSP: Starting…" : state === "error" ? "AOSP: Error" : "AOSP: Stopped";
}

test("maps engine states to stable status labels", () => {
  assert.equal(label("running"), "AOSP: Ready");
  assert.equal(label("starting"), "AOSP: Starting…");
  assert.equal(label("error"), "AOSP: Error");
  assert.equal(label("stopped"), "AOSP: Stopped");
});
