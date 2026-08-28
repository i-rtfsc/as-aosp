import assert from "node:assert/strict";
import * as fs from "node:fs";
import * as os from "node:os";
import * as path from "node:path";
import test from "node:test";
import { detectOutRoot, normalizeOutRoot } from "./out_root";

test("normalizes an out/soong selection to OUT_DIR", () => {
  const root = fs.mkdtempSync(path.join(os.tmpdir(), "aosp-out-"));
  const soong = path.join(root, "out", "soong");
  fs.mkdirSync(soong, { recursive: true });
  fs.writeFileSync(path.join(soong, "module_bp_java_deps.json"), "{}");
  assert.equal(normalizeOutRoot(soong), path.join(root, "out"));
});

test("detects OUT_DIR instead of the nested soong directory", () => {
  const root = fs.mkdtempSync(path.join(os.tmpdir(), "aosp-workspace-"));
  const soong = path.join(root, "out", "soong");
  fs.mkdirSync(soong, { recursive: true });
  fs.writeFileSync(path.join(soong, "module_bp_java_deps.json"), "{}");
  assert.equal(detectOutRoot(root).root, path.join(root, "out"));
});
