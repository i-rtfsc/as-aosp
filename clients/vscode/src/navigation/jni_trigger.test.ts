import assert from "node:assert/strict";
import test from "node:test";
import { shouldTriggerJniDefinition } from "./jni_trigger";

test("ordinary Java methods do not trigger JNI indexing", () => {
  assert.equal(shouldTriggerJniDefinition("java", "private void resetState() {}"), false);
  assert.equal(shouldTriggerJniDefinition("java", "private native void nativeReset();"), true);
});

test("native registration triggers JNI indexing", () => {
  assert.equal(shouldTriggerJniDefinition("cpp", "static const JNINativeMethod methods[] = {"), true);
  assert.equal(shouldTriggerJniDefinition("cpp", "void resetState() {}"), false);
});
