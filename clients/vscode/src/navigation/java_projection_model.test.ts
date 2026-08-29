import assert from "node:assert/strict";
import test from "node:test";
import { normalizeProjectionInput } from "./java_projection_model";

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
