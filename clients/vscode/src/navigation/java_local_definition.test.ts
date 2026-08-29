import assert from "node:assert/strict";
import test from "node:test";
import { findLocalJavaConstructors, findLocalJavaDeclarations } from "./java_local_definition";

test("finds a private method declaration", () => {
  const source = `class Demo {
  void run() { resetState(); }
  private static void resetState() {}
}`;
  const offsets = findLocalJavaDeclarations(source, "resetState");
  assert.deepEqual(offsets, [source.indexOf("resetState() {}")]);
});

test("finds a package-private multiline constructor", () => {
  const source = `private final class AppDeathRecipient {
  AppDeathRecipient(ProcessRecord app, int pid,
      IApplicationThread thread) {}
}`;
  assert.deepEqual(
    findLocalJavaConstructors(source, "AppDeathRecipient"),
    [source.indexOf("AppDeathRecipient(ProcessRecord")],
  );
});

test("finds fields and nested classes", () => {
  const source = `class Demo {
  private final Object lock = new Object();
  static class State {}
}`;
  assert.deepEqual(findLocalJavaDeclarations(source, "lock"), [source.indexOf("lock =")]);
  assert.deepEqual(findLocalJavaDeclarations(source, "State"), [source.indexOf("State {}")]);
});
