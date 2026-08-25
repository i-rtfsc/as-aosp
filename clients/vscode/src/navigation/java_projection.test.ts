import test from "node:test";
import assert from "node:assert/strict";
import { resolveRepoProjectionPath } from "./java_projection_paths";

test("maps repo paths into the focused workspace root", () => {
  assert.equal(
    resolveRepoProjectionPath("/workspace/view", { kind: "repo-relative", repoPath: "frameworks/base" }),
    "/workspace/view/frameworks/base",
  );
});

test("does not map build-host absolute paths", () => {
  assert.equal(
    resolveRepoProjectionPath("/workspace/view", { kind: "external-absolute", raw: "/srv/aosp/frameworks/base" }),
    undefined,
  );
});
