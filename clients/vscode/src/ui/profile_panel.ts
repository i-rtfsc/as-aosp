import * as vscode from "vscode";

export interface WorkspaceModule {
  name: string;
  type: string;
  blueprint: string;
  directory: string;
  status: string;
  capabilities?: string[];
  recommended: boolean;
}

export interface ModuleDiscoveryResult {
  modules: WorkspaceModule[];
  diagnostics: number;
  filesScanned: number;
  cacheHit: boolean;
}

type IndexResult = {
  state: string;
  stage: string;
  progress: number;
  modules: number;
  generated: number;
  diagnostics: number;
};

export class ProfileViewProvider implements vscode.WebviewViewProvider {
  private view?: vscode.WebviewView;
  private indexStatus = "Not started";
  private discovered = new Map<string, WorkspaceModule>();

  constructor(
    private readonly context: vscode.ExtensionContext,
    private readonly discoverModules: (workspaceRoot: string, query?: string) => Promise<ModuleDiscoveryResult>,
    private readonly startIndex: (rebuild: boolean) => Promise<IndexResult>,
  ) {}

  resolveWebviewView(view: vscode.WebviewView): void {
    this.view = view;
    view.webview.options = { enableScripts: true };
    view.webview.html = profileHtml(this.initial());
    this.attachMessages(view.webview);
    void view.webview.postMessage({ type: "indexStatus", text: this.indexStatus });
    void this.refreshModules(view.webview);
  }

  reveal(): void {
    this.view?.show?.(true);
  }

  updateIndexStatus(text: string): void {
    this.indexStatus = text;
    void this.view?.webview.postMessage({ type: "indexStatus", text });
  }

  showCurrentIndexStatus(): void {
    this.updateIndexStatus(this.indexStatus);
  }

  private initial(): Record<string, string> {
    const config = vscode.workspace.getConfiguration("aospIde");
    return {
      workspaceViewRoot: config.get<string>("workspaceViewRoot") ||
        vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "",
      outRoot: config.get<string>("outRoot") ?? "",
      selectedModules: JSON.stringify(config.get<string[]>("selectedModules") ?? []),
    };
  }

  private attachMessages(webview: vscode.Webview): void {
    webview.onDidReceiveMessage(async (message) => {
      if (message.type === "browse") {
        const selected = await vscode.window.showOpenDialog({
          canSelectFolders: true,
          canSelectFiles: false,
          canSelectMany: false,
          openLabel: "Select directory",
        });
        if (selected?.[0]) {
          await webview.postMessage({ type: "set", field: message.field, value: selected[0].fsPath });
        }
        return;
      }
      if (message.type === "refreshModules") {
        await this.refreshModules(webview, message.workspaceViewRoot, message.query);
        return;
      }
      if (message.type === "applyModules") {
        await this.applyModules(message);
        await this.runIndex(false);
        return;
      }
      if (message.type === "rebuildIndex") {
        await this.applyModules(message);
        await this.runIndex(true);
        return;
      }
      if (message.type === "command" && typeof message.command === "string") {
        await vscode.commands.executeCommand(message.command);
      }
    }, undefined, this.context.subscriptions);
  }

  private async refreshModules(webview: vscode.Webview, rawRoot?: string, query?: string): Promise<void> {
    const root = String(rawRoot || this.initial().workspaceViewRoot).trim();
    if (!root) {
      await webview.postMessage({ type: "moduleError", text: "Open or choose a workspace first." });
      return;
    }
    await webview.postMessage({ type: "moduleLoading" });
    try {
      const result = await this.discoverModules(root, query);
      this.discovered = new Map(result.modules.map((module) => [module.name, module]));
      await webview.postMessage({
        type: "moduleCatalog",
        modules: result.modules,
        selected: vscode.workspace.getConfiguration("aospIde").get<string[]>("selectedModules") ?? [],
        summary: `${result.modules.length} modules from ${result.filesScanned} files` +
          ` · ${result.cacheHit ? "cache reused" : "scanned"}` +
          ` · ${result.diagnostics} diagnostics`,
      });
    } catch (error) {
      await webview.postMessage({
        type: "moduleError",
        text: error instanceof Error ? error.message : String(error),
      });
    }
  }

  private async applyModules(message: any): Promise<void> {
    const selected = Array.isArray(message.selectedModules) ?
      message.selectedModules.filter((value: unknown): value is string => typeof value === "string") : [];
    const targets = selected.map((name: string) => moduleTarget(this.discovered.get(name), name));
    const config = vscode.workspace.getConfiguration("aospIde");
    await config.update("workspaceViewRoot", String(message.workspaceViewRoot ?? ""), vscode.ConfigurationTarget.Workspace);
    await config.update("outRoot", String(message.outRoot ?? ""), vscode.ConfigurationTarget.Workspace);
    await config.update("selectedModules", selected, vscode.ConfigurationTarget.Workspace);
    await config.update("targets", targets, vscode.ConfigurationTarget.Workspace);
  }

  private async runIndex(rebuild: boolean): Promise<void> {
    try {
      const status = await this.startIndex(rebuild);
      this.updateIndexStatus(
        `${status.state} · ${status.stage} · ${status.progress}% · ` +
        `${status.modules} modules · ${status.generated} generated · ` +
        `${status.diagnostics} diagnostics`,
      );
    } catch (error) {
      this.updateIndexStatus(`Index unavailable: ${error instanceof Error ? error.message : String(error)}`);
    }
  }
}

function moduleTarget(module: WorkspaceModule | undefined, name: string): unknown {
  const type = module?.type.toLowerCase() ?? "";
  const kind = type.includes("cc_") || type.includes("rust_") ? "native" :
    type.includes("android_") ? "android" : "java";
  return {
    id: name,
    displayName: name,
    kind,
    selectors: { modules: [name] },
    includeDependencies: true,
    includeGeneratedSources: true,
    includeResources: kind === "android",
    includeManifest: kind === "android",
  };
}

export function profileHtml(initial: Record<string, string>): string {
  const nonce = String(Date.now());
  return `<!doctype html><html><head><meta charset="UTF-8">
  <meta http-equiv="Content-Security-Policy" content="default-src 'none'; style-src 'unsafe-inline'; script-src 'nonce-${nonce}'">
  <style>
  :root{--cyan:#22d3ee;--violet:#8b5cf6;--pink:#ec4899;--green:#34d399;--muted:var(--vscode-descriptionForeground);--panel:color-mix(in srgb,var(--vscode-sideBar-background) 86%,#172554)}
  *{box-sizing:border-box}body{margin:0;padding:12px;background:var(--vscode-sideBar-background);color:var(--vscode-foreground);font:12px -apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif}.hero{padding:12px;border-radius:8px;background:linear-gradient(120deg,#7c3aed,#2563eb 55%,#0891b2);color:white}.hero h1{margin:0;font-size:17px}.hero p{margin:4px 0 0;color:#e0f2fe}.section{margin-top:10px;padding:10px;border:1px solid #ffffff18;border-radius:8px;background:var(--panel)}.title{font-size:10px;font-weight:800;color:var(--cyan);letter-spacing:0;margin-bottom:7px}.field{margin-top:7px}label{display:block;margin-bottom:4px;color:var(--muted)}input{width:100%;padding:6px;border:1px solid #ffffff25;border-radius:6px;background:var(--vscode-input-background);color:var(--vscode-input-foreground)}.row{display:flex;gap:5px}.row input{flex:1}button{padding:6px 8px;border:0;border-radius:6px;background:#334155;color:white;font-weight:700;cursor:pointer}.primary{background:linear-gradient(100deg,var(--violet),var(--pink))}.toolbar{display:flex;gap:5px;margin:7px 0}.summary,.status{color:var(--muted);line-height:1.4}.modules{max-height:390px;overflow:auto;border-top:1px solid #ffffff12}.module{display:grid;grid-template-columns:18px 1fr;gap:6px;padding:8px 2px;border-bottom:1px solid #ffffff0d}.module input{width:14px}.name{font-weight:700}.meta{font-size:10px;color:var(--muted);word-break:break-all}.badges{display:flex;gap:4px;flex-wrap:wrap;margin-top:3px}.badge{padding:2px 4px;border-radius:4px;background:#ffffff12;color:#bae6fd}.recommended{color:#f9a8d4}.actions{display:flex;gap:6px;margin-top:10px}.actions button{flex:1}
  </style></head><body>
  <section class="hero"><h1>AOSP Workspace</h1><p>Discover modules, choose scope, apply navigation.</p></section>
  <section class="section"><div class="title">WORKSPACE</div>
    <div class="field"><label>Workspace View Root</label><div class="row"><input id="workspaceRoot" value="${escapeHtml(initial.workspaceViewRoot)}"><button data-browse="workspaceRoot">Browse</button></div></div>
    <div class="field"><label>Out Root</label><div class="row"><input id="outRoot" value="${escapeHtml(initial.outRoot)}"><button data-browse="outRoot">Browse</button></div></div>
  </section>
  <section class="section"><div class="title">MODULES</div>
    <input id="query" placeholder="Search module, path or type">
    <div class="toolbar"><button id="refresh">Refresh</button><button id="recommended">Recommended</button><button id="clear">Clear</button></div>
    <div id="moduleSummary" class="summary">Discovering Android.bp modules...</div>
    <div id="modules" class="modules"></div>
  </section>
  <section class="section"><div class="title">INDEX STATUS</div><div id="indexStatus" class="status">Not started</div></section>
  <div class="actions"><button class="primary" id="apply">Apply Modules</button><button id="rebuild">Rebuild</button></div>
  <script nonce="${nonce}">
  const api=acquireVsCodeApi(), byId=id=>document.getElementById(id);
  let modules=[], selected=new Set(JSON.parse(${JSON.stringify(initial.selectedModules || "[]")}));
  const render=()=>{const query=byId('query').value.trim().toLowerCase();const visible=modules.filter(m=>!query||m.name.toLowerCase().includes(query)||m.blueprint.toLowerCase().includes(query)||m.type.toLowerCase().includes(query));byId('modules').innerHTML=visible.map(m=>'<label class="module"><input type="checkbox" data-module="'+escapeText(m.name)+'" '+(selected.has(m.name)?'checked':'')+'><div><div class="name">'+escapeText(m.name)+(m.recommended?' <span class="recommended">recommended</span>':'')+'</div><div class="meta">'+escapeText(m.type)+' · '+escapeText(m.blueprint)+' · '+escapeText(m.status)+'</div><div class="badges">'+(m.capabilities||[]).map(v=>'<span class="badge">'+escapeText(v)+'</span>').join('')+'</div></div></label>').join('');document.querySelectorAll('[data-module]').forEach(input=>input.onchange=()=>{input.checked?selected.add(input.dataset.module):selected.delete(input.dataset.module)})};
  const escapeText=value=>String(value||'').replaceAll('&','&amp;').replaceAll('<','&lt;').replaceAll('>','&gt;').replaceAll('"','&quot;');
  document.querySelectorAll('[data-browse]').forEach(button=>button.onclick=()=>api.postMessage({type:'browse',field:button.dataset.browse}));
  byId('refresh').onclick=()=>api.postMessage({type:'refreshModules',workspaceViewRoot:byId('workspaceRoot').value,query:byId('query').value});
  byId('recommended').onclick=()=>{selected=new Set(modules.filter(m=>m.recommended).map(m=>m.name));render()};
  byId('clear').onclick=()=>{selected.clear();render()};byId('query').oninput=render;
  const apply=type=>api.postMessage({type,workspaceViewRoot:byId('workspaceRoot').value,outRoot:byId('outRoot').value,selectedModules:[...selected]});
  byId('apply').onclick=()=>apply('applyModules');byId('rebuild').onclick=()=>apply('rebuildIndex');
  window.addEventListener('message',event=>{const data=event.data;if(data.type==='set')byId(data.field).value=data.value;if(data.type==='moduleLoading')byId('moduleSummary').textContent='Scanning module files...';if(data.type==='moduleCatalog'){modules=data.modules;selected=new Set(data.selected?.length?data.selected:modules.filter(m=>m.recommended).map(m=>m.name));byId('moduleSummary').textContent=data.summary;render()}if(data.type==='moduleError')byId('moduleSummary').textContent='Module discovery failed: '+data.text;if(data.type==='indexStatus')byId('indexStatus').textContent=data.text});
  </script></body></html>`;
}

function escapeHtml(value: string): string {
  return value.replaceAll("&", "&amp;").replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;").replaceAll('"', "&quot;");
}
