const assert = require("node:assert/strict");
const vscode = require("vscode");

suite("AOSP IDE extension host", () => {
  test("activates and registers AIDL language", async () => {
    const extension = vscode.extensions.getExtension("rtfsc.aosp-ide-vscode");
    assert.ok(extension);
    await extension.activate();
    const commands = await vscode.commands.getCommands(true);
    assert.ok(commands.includes("aospIde.goToJniPeer"));
    assert.ok(commands.includes("aospIde.goToAidlGeneratedPeer"));
  });
});
