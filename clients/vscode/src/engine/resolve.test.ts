import assert from "node:assert/strict";
import test from "node:test";
import { platformTarget, resolveEngineExecutable } from "./resolve";

test("normalizes platform and architecture names", () => {
  assert.equal(platformTarget("win32", "x64"), "windows-amd64");
  assert.equal(platformTarget("darwin", "arm64"), "darwin-arm64");
});

test("prefers configured, then bundled, then PATH engine", () => {
  assert.equal(resolveEngineExecutable({ configured: "/custom/aosp-index", extensionPath: "/ext", exists: () => true }), "/custom/aosp-index");
  assert.equal(resolveEngineExecutable({ extensionPath: "/ext", exists: (value) => value.endsWith("linux-amd64/aosp-index"), platform: "linux", architecture: "x64" }), "/ext/engine/linux-amd64/aosp-index");
  assert.equal(resolveEngineExecutable({ extensionPath: "/ext", exists: () => false, platform: "linux", architecture: "x64" }), "aosp-index");
});
