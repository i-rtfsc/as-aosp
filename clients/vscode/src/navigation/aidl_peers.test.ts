import assert from "node:assert/strict";
import test from "node:test";
import { selectAidlPeerPath } from "./aidl_peer_selection";

test("selects a visible exact repo peer and ignores inferred candidates", () => {
  assert.equal(selectAidlPeerPath({
    peers: [
      { path: { raw: "demo/IFoo.java", repoPath: "demo/IFoo.java" }, confidence: "inferred", exists: false },
      { path: { raw: "out/generated/demo/IFoo.java", repoPath: "demo/IFoo.java" }, confidence: "exact", exists: true },
    ],
  }), "demo/IFoo.java");
});

test("rejects absolute exact peers", () => {
  assert.equal(selectAidlPeerPath({
    peers: [{ path: { raw: "/srv/out/IFoo.java" }, confidence: "exact", exists: true }],
  }), undefined);
});
