import assert from "node:assert/strict";
import test from "node:test";
import { sourceRootForPackage } from "./java_source_roots";

test("derives an AOSP Java root from the package", () => {
  assert.equal(
    sourceRootForPackage(
      "frameworks/base/core/java/android/os/Binder.java",
      "package android.os; public class Binder {}",
    ),
    "frameworks/base/core/java",
  );
});

test("derives a Kotlin root without a semicolon", () => {
  assert.equal(
    sourceRootForPackage(
      "frameworks/base/packages/SystemUI/src/com/android/systemui/Foo.kt",
      "package com.android.systemui\nclass Foo",
    ),
    "frameworks/base/packages/SystemUI/src",
  );
});
