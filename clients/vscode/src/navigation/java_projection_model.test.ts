import assert from "node:assert/strict";
import test from "node:test";
import { normalizeProjectionInput, preferredLibraries } from "./java_projection_model";

test("normalizes omitted project graph collections", () => {
  const input = normalizeProjectionInput({
    graph: {
      targets: [
        { sourceRoots: [{ raw: "frameworks/base", kind: "repo-relative" }] },
        { libraries: null },
      ],
    },
    diagnostics: null,
  });
  assert.equal(input.targets.length, 2);
  assert.deepEqual(input.targets[0].libraries, []);
  assert.deepEqual(input.targets[1].sourceRoots, []);
  assert.deepEqual(input.archiveRefs, []);
  assert.equal(input.diagnostics, 0);
  assert.equal(input.moduleCount, 0);
});

test("collects srcjar archives from available libraries", () => {
  const input = normalizeProjectionInput({
    graph: {
      targets: [{ ownedGeneratedRoots: [
        { raw: "out/a.srcjar", kind: "generated" },
      ], libraries: [
        { raw: "out/a.jar", kind: "generated" },
      ] }],
      diagnostics: [{}],
    },
    diagnostics: [{}],
  });
  assert.deepEqual(input.archiveRefs, ["out/a.srcjar"]);
  assert.equal(input.diagnostics, 2);
});

test("excludes native targets from Java projection", () => {
  const input = normalizeProjectionInput({
    graph: { targets: [
      { spec: { kind: "native" }, sourceRoots: [{ raw: "system/core", kind: "repo-relative" }] },
      { spec: { kind: "java" }, sourceRoots: [{ raw: "frameworks/base", kind: "repo-relative" }] },
    ] },
  });
  assert.deepEqual(input.targets.map((target) => target.kind), ["java"]);
  assert.deepEqual(input.targets[0].sourceRoots, [{ raw: "frameworks/base", kind: "repo-relative" }]);
});

test("counts selected Java closure modules", () => {
  const input = normalizeProjectionInput({ graph: { targets: [
    { spec: { id: "java", kind: "java" }, resolvedModules: ["framework", "core"] },
    { spec: { id: "services", kind: "java" }, resolvedModules: ["framework", "services"] },
    { spec: { id: "native", kind: "native" }, resolvedModules: ["libbinder"] },
  ] } });
  assert.equal(input.moduleCount, 3);
});

test("keeps one preferred Android jar for each basename", () => {
  const libraries = preferredLibraries([
    "/out/module/android_common_apex31/combined/example.jar",
    "/out/module/linux_glibc_common/javac-header/example.jar",
    "/out/module/android_common/turbine-combined/example.jar",
    "/out/other/android_common/turbine-combined/other.jar",
  ]);
  assert.deepEqual(libraries, [
    "/out/module/android_common/turbine-combined/example.jar",
    "/out/other/android_common/turbine-combined/other.jar",
  ]);
});
