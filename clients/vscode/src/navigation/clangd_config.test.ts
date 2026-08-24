import assert from "node:assert/strict";
import test from "node:test";
import { overlayClangdConfig } from "./clangd_config";

test("points clangd at the local overlay database", () => {
  assert.equal(overlayClangdConfig, "CompileFlags:\n  CompilationDatabase: .\n");
  assert.equal(overlayClangdConfig.includes("../"), false);
});
