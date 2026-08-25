import * as vscode from "vscode";
import { EngineClient } from "./engine/client";

export function createAospStatusBar(engine: EngineClient, output: vscode.OutputChannel): vscode.StatusBarItem {
  const item = vscode.window.createStatusBarItem(vscode.StatusBarAlignment.Left, 10);
  const refresh = (state: string, error?: string) => {
    item.text = state === "running" ? "AOSP: Ready" : state === "starting" ? "AOSP: Starting…" : state === "error" ? "AOSP: Error" : "AOSP: Stopped";
    item.tooltip = error ?? "AOSP Workspace engine status";
    item.command = "aospIde.showStatus";
    item.show();
  };
  refresh(engine.status.state, engine.status.error);
  const subscription = engine.onStatus(refresh);
  item.dispose = (() => { const original = item.dispose.bind(item); return () => { subscription.dispose(); original(); }; })();
  return item;
}
