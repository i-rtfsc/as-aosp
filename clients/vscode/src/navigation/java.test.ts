import assert from "node:assert/strict";
import test from "node:test";
import { javaCandidatePaths } from "./java_paths";

test("builds package and import Java candidates", () => {
  assert.deepEqual(javaCandidatePaths("demo.client", ["demo.api.IFoo", "java.util.List"], "IFoo"), ["demo/api/IFoo.java", "demo/client/IFoo.java", "IFoo.java"]);
});
