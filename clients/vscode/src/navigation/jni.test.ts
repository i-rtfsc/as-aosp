import assert from "node:assert/strict";
import test from "node:test";
import { selectJniDefinitionTarget } from "./jni_selection";

const link = {
  java: { path: "demo/Bridge.java", range: { start: 42 } },
  native: { path: "demo/bridge.cpp", range: { start: 17 } },
  confidence: "exact",
};

test("selects the native side for a Java file", () => {
  assert.deepEqual(selectJniDefinitionTarget({ peers: [link] }, "demo/Bridge.java"), {
    path: "demo/bridge.cpp", offset: 17,
  });
});

test("selects the Java side for a native file", () => {
  assert.deepEqual(selectJniDefinitionTarget({ peers: [link] }, "demo/bridge.cpp"), {
    path: "demo/Bridge.java", offset: 42,
  });
});

test("rejects inferred or ambiguous links", () => {
  assert.equal(selectJniDefinitionTarget({ peers: [{ ...link, confidence: "inferred" }] }, "demo/Bridge.java"), undefined);
  assert.equal(selectJniDefinitionTarget({ peers: [link, link] }, "demo/Bridge.java"), undefined);
});
