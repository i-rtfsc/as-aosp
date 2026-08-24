import assert from "node:assert/strict";
import test from "node:test";
import { overlayResponseFilePath } from "./compdb_paths";

test("keeps response files inside the overlay directory", () => {
  assert.equal(overlayResponseFilePath("/view/.aosp-ide/compdb-overlay", "args.rsp"), "/view/.aosp-ide/compdb-overlay/args.rsp");
  assert.equal(overlayResponseFilePath("/view/.aosp-ide/compdb-overlay", "../outside.rsp"), undefined);
  assert.equal(overlayResponseFilePath("/view/.aosp-ide/compdb-overlay", "C:/outside.rsp"), undefined);
});
