import * as vscode from "vscode";

export function openProfilePanel(context: vscode.ExtensionContext): void {
  const panel = vscode.window.createWebviewPanel("aospIdeProfile", "AOSP IDE Profile", vscode.ViewColumn.One, { enableScripts: true, retainContextWhenHidden: true });
  const config = vscode.workspace.getConfiguration("aospIde");
  const initial = {
    workspaceViewRoot: config.get<string>("workspaceViewRoot") || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
    metadataRoot: config.get<string>("metadataRoot") ?? "",
    engineExecutable: config.get<string>("engineExecutable") ?? "",
    pathMappings: JSON.stringify(config.get<unknown[]>("pathMappings") ?? [], null, 2),
  };
  panel.webview.html = profileHtml(initial);
  panel.webview.onDidReceiveMessage(async (message) => {
    if (message.type === "browse") {
      const selected = await vscode.window.showOpenDialog({ canSelectFolders: true, canSelectFiles: false, canSelectMany: false, openLabel: "Select directory" });
      if (selected?.[0]) panel.webview.postMessage({ type: "set", field: message.field, value: selected[0].fsPath });
    }
    if (message.type === "save") {
      let mappings: unknown[];
      try { mappings = JSON.parse(message.pathMappings || "[]"); } catch { void vscode.window.showErrorMessage("AOSP IDE pathMappings must be valid JSON."); return; }
      await config.update("workspaceViewRoot", message.workspaceViewRoot, vscode.ConfigurationTarget.Workspace);
      await config.update("metadataRoot", message.metadataRoot, vscode.ConfigurationTarget.Workspace);
      await config.update("engineExecutable", message.engineExecutable, vscode.ConfigurationTarget.Workspace);
      await config.update("pathMappings", mappings, vscode.ConfigurationTarget.Workspace);
      void vscode.window.showInformationMessage("AOSP IDE profile saved.");
    }
    if (message.type === "command" && typeof message.command === "string") {
      await vscode.commands.executeCommand(message.command);
    }
  }, undefined, context.subscriptions);
}

function profileHtml(initial: Record<string, string>): string {
  const nonce = String(Date.now());
  return `<!doctype html><html><head><meta charset="UTF-8"><meta http-equiv="Content-Security-Policy" content="default-src 'none'; style-src 'unsafe-inline'; script-src 'nonce-${nonce}'"><style>body{font:13px sans-serif;padding:20px;max-width:850px}label{display:block;margin-top:14px;font-weight:600}input,textarea{box-sizing:border-box;width:100%;margin-top:5px;padding:7px;background:var(--vscode-input-background);color:var(--vscode-input-foreground);border:1px solid var(--vscode-input-border)}button{margin:14px 8px 0 0;padding:7px 14px}</style></head><body><h1>AOSP IDE Profile</h1><p>Configure the current VS Code workspace. The Workspace View may be a symlink, mount, or selected repository directory.</p><label>Workspace View Root<input id="workspaceViewRoot" value="${escapeHtml(initial.workspaceViewRoot)}"><button data-browse="workspaceViewRoot">Browse</button></label><label>Metadata / out Root<input id="metadataRoot" value="${escapeHtml(initial.metadataRoot)}"><button data-browse="metadataRoot">Browse</button></label><label>Engine executable (optional)<input id="engineExecutable" value="${escapeHtml(initial.engineExecutable)}"></label><label>Path mappings (JSON)<textarea id="pathMappings" rows="8">${escapeHtml(initial.pathMappings)}</textarea></label><button id="save">Save Profile</button><button id="compdb">Refresh clangd Database</button><button id="status">Show Status</button><script nonce="${nonce}">const api=acquireVsCodeApi();document.querySelectorAll('[data-browse]').forEach(b=>b.onclick=()=>api.postMessage({type:'browse',field:b.dataset.browse}));document.getElementById('save').onclick=()=>api.postMessage({type:'save',workspaceViewRoot:document.getElementById('workspaceViewRoot').value,metadataRoot:document.getElementById('metadataRoot').value,engineExecutable:document.getElementById('engineExecutable').value,pathMappings:document.getElementById('pathMappings').value});document.getElementById('compdb').onclick=()=>api.postMessage({type:'command',command:'aospIde.refreshCompdbOverlay'});document.getElementById('status').onclick=()=>api.postMessage({type:'command',command:'aospIde.showStatus'});window.addEventListener('message',e=>{if(e.data.type==='set')document.getElementById(e.data.field).value=e.data.value});</script></body></html>`;
}

function escapeHtml(value: string): string { return value.replaceAll("&", "&amp;").replaceAll("<", "&lt;").replaceAll(">", "&gt;").replaceAll('"', "&quot;"); }
