import assert from "node:assert/strict";
import path from "node:path";
import test from "node:test";
import { normalizeRepoPath, resolveRepoTarget, toRepoPath } from "./repo_path";

test("accepts portable repository paths", () => {
  assert.equal(normalizeRepoPath("frameworks/base/core/java/Foo.java"), "frameworks/base/core/java/Foo.java");
  assert.equal(normalizeRepoPath("frameworks\\base\\Foo.java"), "frameworks/base/Foo.java");
});

test("rejects escaping and absolute repository paths", () => {
  for (const value of ["../outside.java", "/tmp/outside.java", "C:/outside.java", "\\\\server\\share\\x.java", "a/../b.java"]) {
    assert.equal(normalizeRepoPath(value), undefined, value);
  }
});

test("rejects fixed excluded directories", () => {
  assert.equal(normalizeRepoPath("out/soong/generated/Foo.java"), undefined);
  assert.equal(normalizeRepoPath("frameworks/base/.git/config"), undefined);
});

test("maps POSIX paths relative to the workspace view", () => {
  assert.equal(toRepoPath(path.posix, "/work/aosp-view", "/work/aosp-view/frameworks/base/Foo.java"),
    "frameworks/base/Foo.java");
  assert.equal(resolveRepoTarget(path.posix, "/work/aosp-view", "frameworks/base/Foo.java"),
    "/work/aosp-view/frameworks/base/Foo.java");
});

test("maps Windows paths without leaking drive separators", () => {
  assert.equal(toRepoPath(path.win32, "C:\\aosp-view", "C:\\aosp-view\\frameworks\\base\\Foo.java"),
    "frameworks/base/Foo.java");
  assert.equal(resolveRepoTarget(path.win32, "C:\\aosp-view", "frameworks/base/Foo.java"),
    "C:\\aosp-view\\frameworks\\base\\Foo.java");
  assert.equal(toRepoPath(path.win32, "C:\\aosp-view", "D:\\other\\Foo.java"), undefined);
});
