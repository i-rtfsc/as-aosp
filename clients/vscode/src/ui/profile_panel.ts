import * as vscode from "vscode";

export class ProfileViewProvider implements vscode.WebviewViewProvider {
  private view?: vscode.WebviewView;
  constructor(
    private readonly context: vscode.ExtensionContext,
    private readonly inspectPack?: (root: string) => Promise<{ modules: number; generated: number; diagnostics: number }>,
    private readonly startIndex?: (rebuild: boolean) => Promise<{ state: string; stage: string; progress: number; modules: number; generated: number; diagnostics: number }>,
  ) {}
  resolveWebviewView(view: vscode.WebviewView): void {
    this.view = view;
    view.webview.options = { enableScripts: true };
    view.webview.html = profileHtml(this.initial());
    this.attachMessages(view.webview);
  }
  reveal(): void { this.view?.show?.(true); }
  private initial(): Record<string, string> {
    const config = vscode.workspace.getConfiguration("aospIde");
    return {
      workspaceViewRoot: config.get<string>("workspaceViewRoot") || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
      metadataRoot: config.get<string>("metadataRoot") ?? "",
      metadataPackRoot: config.get<string>("metadataPackRoot") ?? "",
      engineExecutable: config.get<string>("engineExecutable") ?? "",
      pathMappings: JSON.stringify(config.get<unknown[]>("pathMappings") ?? [], null, 2),
    };
  }
  private attachMessages(webview: vscode.Webview): void {
    webview.onDidReceiveMessage(async (message) => {
      if (message.type === "browse") {
        const selected = await vscode.window.showOpenDialog({ canSelectFolders: true, canSelectFiles: false, canSelectMany: false, openLabel: "Select directory" });
        if (selected?.[0]) webview.postMessage({ type: "set", field: message.field, value: selected[0].fsPath });
      }
      if (message.type === "save") {
        let mappings: unknown[];
        try { mappings = JSON.parse(message.pathMappings || "[]"); } catch { void vscode.window.showErrorMessage("AOSP IDE pathMappings must be valid JSON."); return; }
        const config = vscode.workspace.getConfiguration("aospIde");
        await config.update("workspaceViewRoot", message.workspaceViewRoot, vscode.ConfigurationTarget.Workspace);
        await config.update("metadataRoot", message.metadataRoot, vscode.ConfigurationTarget.Workspace);
        await config.update("metadataPackRoot", message.metadataPackRoot, vscode.ConfigurationTarget.Workspace);
        await config.update("engineExecutable", message.engineExecutable, vscode.ConfigurationTarget.Workspace);
        await config.update("pathMappings", mappings, vscode.ConfigurationTarget.Workspace);
        void vscode.window.showInformationMessage("AOSP IDE profile saved.");
      }
      if (message.type === "inspectPack") {
        if (!this.inspectPack || !message.metadataPackRoot?.trim()) {
          webview.postMessage({ type: "catalogStatus", text: "Choose a Metadata Pack first." });
        } else {
          try {
            const summary = await this.inspectPack(message.metadataPackRoot);
            webview.postMessage({ type: "catalogStatus", text: `${summary.modules} modules · ${summary.generated} generated entries · ${summary.diagnostics} diagnostics` });
          } catch (error) {
            webview.postMessage({ type: "catalogStatus", text: `Catalog unavailable: ${error instanceof Error ? error.message : String(error)}` });
          }
        }
      }
      if (message.type === "startIndex") {
        if (!this.startIndex) return;
        try {
          const status = await this.startIndex(Boolean(message.rebuild));
          webview.postMessage({ type: "indexStatus", text: `${status.state} · ${status.stage} · ${status.progress}% · ${status.modules} modules · ${status.generated} generated · ${status.diagnostics} diagnostics` });
        } catch (error) {
          webview.postMessage({ type: "indexStatus", text: `Index unavailable: ${error instanceof Error ? error.message : String(error)}` });
        }
      }
      if (message.type === "command" && typeof message.command === "string") await vscode.commands.executeCommand(message.command);
    }, undefined, this.context.subscriptions);
  }
}

export function openProfilePanel(context: vscode.ExtensionContext): void {
  const panel = vscode.window.createWebviewPanel("aospWorkspaceProfile", "AOSP Workspace", vscode.ViewColumn.One, { enableScripts: true, retainContextWhenHidden: true });
  const config = vscode.workspace.getConfiguration("aospIde");
  const initial = {
    workspaceViewRoot: config.get<string>("workspaceViewRoot") || vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
    metadataRoot: config.get<string>("metadataRoot") ?? "",
    metadataPackRoot: config.get<string>("metadataPackRoot") ?? "",
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
      await config.update("metadataPackRoot", message.metadataPackRoot, vscode.ConfigurationTarget.Workspace);
      await config.update("engineExecutable", message.engineExecutable, vscode.ConfigurationTarget.Workspace);
      await config.update("pathMappings", mappings, vscode.ConfigurationTarget.Workspace);
      void vscode.window.showInformationMessage("AOSP IDE profile saved.");
    }
    if (message.type === "command" && typeof message.command === "string") {
      await vscode.commands.executeCommand(message.command);
    }
  }, undefined, context.subscriptions);
}

export function profileHtml(initial: Record<string, string>): string {
  const nonce = String(Date.now());
  return `<!doctype html><html><head><meta charset="UTF-8"><meta http-equiv="Content-Security-Policy" content="default-src 'none'; style-src 'unsafe-inline'; script-src 'nonce-${nonce}'"><style>
  :root{--purple:#a855f7;--blue:#38bdf8;--teal:#2dd4bf;--card:color-mix(in srgb,var(--vscode-editor-background) 88%,#172554);--muted:var(--vscode-descriptionForeground)}
  *{box-sizing:border-box}body{margin:0;padding:14px;background:var(--vscode-sideBar-background);color:var(--vscode-foreground);font:12px -apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif}
  .hero{padding:15px;border-radius:13px;background:linear-gradient(125deg,#581c87,#1d4ed8 56%,#0f766e);margin-bottom:10px}.hero h1{margin:0;font-size:19px;color:white}.hero p{margin:4px 0 0;color:#e0f2fe;font-size:11px}.pill{display:inline-block;margin-bottom:7px;padding:4px 8px;border-radius:99px;background:#ffffff24;color:#d1fae5;font-size:10px;font-weight:700}.grid{display:flex;flex-direction:column;gap:8px}.card{background:var(--card);border:1px solid #ffffff14;border-radius:11px;padding:12px}.eyebrow{font-size:10px;letter-spacing:1.1px;color:#7dd3fc;font-weight:800;margin-bottom:6px}.value{font-size:12px;font-weight:650;word-break:break-all}.hint{color:var(--muted);margin-top:5px;line-height:1.35;font-size:11px}.ok{color:var(--teal);font-weight:700}.dot{display:inline-block;width:7px;height:7px;border-radius:50%;background:var(--teal);margin-right:5px}.field{margin-top:9px}.field label{display:block;color:#c4b5fd;font-weight:700;margin-bottom:4px}.row{display:flex;gap:5px}.row input{flex:1}input,textarea{width:100%;padding:7px 8px;border-radius:7px;border:1px solid #ffffff1f;background:var(--vscode-input-background);color:var(--vscode-input-foreground);outline:none}input:focus,textarea:focus{border-color:var(--blue)}button{border:0;border-radius:7px;padding:7px 9px;background:#334155;color:white;font-weight:700;cursor:pointer;font-size:11px}button:hover{background:#475569}.primary{background:linear-gradient(100deg,var(--purple),#2563eb)}.actions{display:flex;flex-direction:column;gap:6px;margin-top:10px}.cap{display:flex;flex-direction:column;gap:5px;margin-top:6px}.cap div{padding:7px 8px;border-radius:7px;background:#ffffff08}.cap b{display:block;margin-bottom:2px}.footer{margin-top:10px;color:var(--muted);font-size:10px;line-height:1.35}
  </style></head><body>
  <section class="hero"><span class="pill">SOURCE-READY</span><h1>AOSP Workspace</h1><p>Focused source navigation for AOSP, MIUI and QCOM</p></section>
  <div class="grid">
    <section class="card"><div class="eyebrow">WORKSPACE VIEW</div><div class="value" id="workspaceValue">${escapeHtml(initial.workspaceViewRoot)}</div><div class="hint"><span class="dot"></span>Only the visible folder is indexed. Symlinks and mounted repositories are supported.</div><div class="field"><label for="workspaceViewRoot">Workspace View Root</label><div class="row"><input id="workspaceViewRoot" value="${escapeHtml(initial.workspaceViewRoot)}"><button data-browse="workspaceViewRoot">Browse</button></div></div></section>
    <section class="card"><div class="eyebrow">OUT / METADATA</div><div class="value" id="outValue">${initial.metadataRoot ? "BUILD-BACKED" : "SOURCE-ONLY"}</div><div class="hint">${initial.metadataRoot ? escapeHtml(initial.metadataRoot) : "No out configured yet. AIDL and JNI source navigation still works."}</div><div class="field"><label for="metadataRoot">Out / compdb Root</label><div class="row"><input id="metadataRoot" value="${escapeHtml(initial.metadataRoot)}"><button data-browse="metadataRoot">Browse</button></div></div><div class="field"><label for="metadataPackRoot">Metadata Pack Root</label><div class="row"><input id="metadataPackRoot" value="${escapeHtml(initial.metadataPackRoot)}"><button data-browse="metadataPackRoot">Browse</button></div><div class="hint">Use a pack exported on another machine for module and generated-source facts.</div><div class="hint" id="catalogStatus">Not inspected</div></div></section>
    <section class="card wide"><div class="eyebrow">INDEX STATUS</div><div class="hint" id="indexStatus">Not started</div></section>
    <section class="card wide"><div class="eyebrow">NAVIGATION CAPABILITIES</div><div class="cap"><div><b><span class="ok">✓</span> AIDL definition</b><span class="hint">Package, import and type resolution</span></div><div><b><span class="ok">✓</span> AIDL references</b><span class="hint">Exact type ranges when available</span></div><div><b><span class="ok">✓</span> Generated peers</b><span class="hint">Java, C++ and NDK candidates</span></div><div><b><span class="ok">✓</span> JNI Java ↔ C++</b><span class="hint">Static and common registration forms</span></div></div></section>
    <section class="card wide"><div class="eyebrow">ADVANCED</div><div class="field"><label for="engineExecutable">Engine executable (optional)</label><input id="engineExecutable" value="${escapeHtml(initial.engineExecutable)}" placeholder="Bundled engine → engines/aosp-index → PATH"></div><div class="field"><label for="pathMappings">Path mappings</label><textarea id="pathMappings" rows="5">${escapeHtml(initial.pathMappings)}</textarea><div class="hint">Use JSON only when build-machine paths differ from this computer.</div></div></section>
  </div><div class="actions"><button class="primary" id="save">Save &amp; Verify Profile</button><button id="startIndex">Start Index</button><button id="rebuildIndex">Rebuild Index</button><button id="inspectPack">Inspect Module Pack</button><button id="compdb">Refresh clangd Database</button><button id="status">Show Status</button></div><div class="footer">AOSP Workspace keeps configuration workspace-scoped. Reopen it from the AOSP Workspace icon in the Activity Bar.</div>
  <script nonce="${nonce}">const api=acquireVsCodeApi();const $=id=>document.getElementById(id);document.querySelectorAll('[data-browse]').forEach(b=>b.onclick=()=>api.postMessage({type:'browse',field:b.dataset.browse}));$('save').onclick=()=>api.postMessage({type:'save',workspaceViewRoot:$('workspaceViewRoot').value,metadataRoot:$('metadataRoot').value,metadataPackRoot:$('metadataPackRoot').value,engineExecutable:$('engineExecutable').value,pathMappings:$('pathMappings').value});$('startIndex').onclick=()=>api.postMessage({type:'startIndex',rebuild:false});$('rebuildIndex').onclick=()=>api.postMessage({type:'startIndex',rebuild:true});$('inspectPack').onclick=()=>api.postMessage({type:'inspectPack',metadataPackRoot:$('metadataPackRoot').value});$('compdb').onclick=()=>api.postMessage({type:'command',command:'aospIde.refreshCompdbOverlay'});$('status').onclick=()=>api.postMessage({type:'command',command:'aospIde.showStatus'});$('workspaceViewRoot').oninput=e=>$('workspaceValue').textContent=e.target.value;window.addEventListener('message',e=>{if(e.data.type==='set'){ $(e.data.field).value=e.data.value;if(e.data.field==='workspaceViewRoot')$('workspaceValue').textContent=e.data.value}if(e.data.type==='catalogStatus')$('catalogStatus').textContent=e.data.text;if(e.data.type==='indexStatus')$('indexStatus').textContent=e.data.text});</script></body></html>`;
}

function escapeHtml(value: string): string { return value.replaceAll("&", "&amp;").replaceAll("<", "&lt;").replaceAll(">", "&gt;").replaceAll('"', "&quot;"); }
