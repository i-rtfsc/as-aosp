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
  return `<!doctype html><html><head><meta charset="UTF-8"><meta http-equiv="Content-Security-Policy" content="default-src 'none'; style-src 'unsafe-inline'; script-src 'nonce-${nonce}'"><style>
  :root{--purple:#a855f7;--blue:#38bdf8;--teal:#2dd4bf;--card:color-mix(in srgb,var(--vscode-editor-background) 88%,#172554);--muted:var(--vscode-descriptionForeground)}
  *{box-sizing:border-box}body{margin:0;padding:28px;max-width:980px;background:var(--vscode-editor-background);color:var(--vscode-foreground);font:13px -apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif}
  .hero{padding:25px 27px;border-radius:20px;background:linear-gradient(125deg,#581c87,#1d4ed8 56%,#0f766e);box-shadow:0 14px 34px #02061780;margin-bottom:18px}.hero h1{margin:0;font-size:28px;letter-spacing:-.8px;color:white}.hero p{margin:7px 0 0;color:#e0f2fe;font-size:14px}.pill{float:right;margin-top:-23px;padding:6px 12px;border-radius:99px;background:#ffffff24;color:#d1fae5;font-weight:700}.grid{display:grid;grid-template-columns:1.3fr .7fr;gap:14px}.card{background:var(--card);border:1px solid #ffffff14;border-radius:16px;padding:18px;box-shadow:0 8px 22px #02061735}.wide{grid-column:1/-1}.eyebrow{font-size:11px;letter-spacing:1.3px;color:#7dd3fc;font-weight:800;margin-bottom:9px}.value{font-size:15px;font-weight:650;word-break:break-all}.hint{color:var(--muted);margin-top:7px;line-height:1.45}.ok{color:var(--teal);font-weight:700}.dot{display:inline-block;width:8px;height:8px;border-radius:50%;background:var(--teal);margin-right:7px}.field{margin-top:13px}.field label{display:block;color:#c4b5fd;font-weight:700;margin-bottom:6px}.row{display:flex;gap:8px}.row input{flex:1}input,textarea{width:100%;padding:10px 11px;border-radius:9px;border:1px solid #ffffff1f;background:var(--vscode-input-background);color:var(--vscode-input-foreground);outline:none}input:focus,textarea:focus{border-color:var(--blue);box-shadow:0 0 0 2px #38bdf833}button{border:0;border-radius:9px;padding:9px 14px;background:#334155;color:white;font-weight:700;cursor:pointer}button:hover{background:#475569}.primary{background:linear-gradient(100deg,var(--purple),#2563eb)}.actions{display:flex;gap:9px;flex-wrap:wrap;margin-top:17px}.cap{display:grid;grid-template-columns:1fr 1fr;gap:10px;margin-top:10px}.cap div{padding:11px 12px;border-radius:10px;background:#ffffff08}.cap b{display:block;margin-bottom:4px}.footer{margin-top:16px;color:var(--muted);font-size:12px}@media(max-width:700px){.grid{grid-template-columns:1fr}.wide{grid-column:auto}.pill{float:none;display:inline-block;margin:12px 0 0}}
  </style></head><body>
  <section class="hero"><span class="pill">SOURCE-READY</span><h1>AOSP IDE</h1><p>Neon Workspace · focused source navigation for AOSP, MIUI and QCOM</p></section>
  <div class="grid">
    <section class="card"><div class="eyebrow">WORKSPACE VIEW</div><div class="value" id="workspaceValue">${escapeHtml(initial.workspaceViewRoot)}</div><div class="hint"><span class="dot"></span>Only the visible folder is indexed. Symlinks and mounted repositories are supported.</div><div class="field"><label for="workspaceViewRoot">Workspace View Root</label><div class="row"><input id="workspaceViewRoot" value="${escapeHtml(initial.workspaceViewRoot)}"><button data-browse="workspaceViewRoot">Browse</button></div></div></section>
    <section class="card"><div class="eyebrow">OUT / METADATA</div><div class="value" id="outValue">${initial.metadataRoot ? "BUILD-BACKED" : "SOURCE-ONLY"}</div><div class="hint">${initial.metadataRoot ? escapeHtml(initial.metadataRoot) : "No out configured yet. AIDL and JNI source navigation still works."}</div><div class="field"><label for="metadataRoot">Metadata / out Root</label><div class="row"><input id="metadataRoot" value="${escapeHtml(initial.metadataRoot)}"><button data-browse="metadataRoot">Browse</button></div></div></section>
    <section class="card wide"><div class="eyebrow">NAVIGATION CAPABILITIES</div><div class="cap"><div><b><span class="ok">✓</span> AIDL definition</b><span class="hint">Package, import and type resolution</span></div><div><b><span class="ok">✓</span> AIDL references</b><span class="hint">Exact type ranges when available</span></div><div><b><span class="ok">✓</span> Generated peers</b><span class="hint">Java, C++ and NDK candidates</span></div><div><b><span class="ok">✓</span> JNI Java ↔ C++</b><span class="hint">Static and common registration forms</span></div></div></section>
    <section class="card wide"><div class="eyebrow">ADVANCED</div><div class="field"><label for="engineExecutable">Engine executable (optional)</label><input id="engineExecutable" value="${escapeHtml(initial.engineExecutable)}" placeholder="Bundled engine → engines/aosp-index → PATH"></div><div class="field"><label for="pathMappings">Path mappings</label><textarea id="pathMappings" rows="5">${escapeHtml(initial.pathMappings)}</textarea><div class="hint">Use JSON only when build-machine paths differ from this computer.</div></div></section>
  </div><div class="actions"><button class="primary" id="save">Save &amp; Verify Profile</button><button id="compdb">Refresh clangd Database</button><button id="status">Show Status</button></div><div class="footer">AOSP IDE keeps project configuration workspace-scoped. Closing this page is safe — reopen it from the AOSP IDE icon in the Activity Bar.</div>
  <script nonce="${nonce}">const api=acquireVsCodeApi();const $=id=>document.getElementById(id);document.querySelectorAll('[data-browse]').forEach(b=>b.onclick=()=>api.postMessage({type:'browse',field:b.dataset.browse}));$('save').onclick=()=>api.postMessage({type:'save',workspaceViewRoot:$('workspaceViewRoot').value,metadataRoot:$('metadataRoot').value,engineExecutable:$('engineExecutable').value,pathMappings:$('pathMappings').value});$('compdb').onclick=()=>api.postMessage({type:'command',command:'aospIde.refreshCompdbOverlay'});$('status').onclick=()=>api.postMessage({type:'command',command:'aospIde.showStatus'});$('workspaceViewRoot').oninput=e=>$('workspaceValue').textContent=e.target.value;window.addEventListener('message',e=>{if(e.data.type==='set'){ $(e.data.field).value=e.data.value;if(e.data.field==='workspaceViewRoot')$('workspaceValue').textContent=e.data.value}});</script></body></html>`;
}

function escapeHtml(value: string): string { return value.replaceAll("&", "&amp;").replaceAll("<", "&lt;").replaceAll(">", "&gt;").replaceAll('"', "&quot;"); }
