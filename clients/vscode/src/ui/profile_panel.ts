import * as path from "node:path";
import * as vscode from "vscode";
import { SelectionCatalogResult, SelectionPlanResult } from "../engine/client";
import { webviewMessages } from "../i18n/messages";
import { normalizeOutRoot } from "../navigation/out_root";

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
  buildBacked: number;
  total: number;
}

type IndexResult = {
  state: string;
  stage: string;
  progress: number;
  modules: number;
  generated: number;
  diagnostics: number;
};

type InitialState = {
  workspaceViewRoot: string;
  workspaceName: string;
  outRoot: string;
  outName: string;
  selectedPresets: string[];
  selectedModules: string[];
  presetSelectionConfigured: boolean;
};

export class ProfileViewProvider implements vscode.WebviewViewProvider {
  private view?: vscode.WebviewView;
  private indexStatus = "Not started";

  constructor(
    private readonly context: vscode.ExtensionContext,
    private readonly discoverModules: (
      workspaceRoot: string, outRoot: string, query?: string, limit?: number,
    ) => Promise<ModuleDiscoveryResult>,
    private readonly loadSelectionCatalog: () => Promise<SelectionCatalogResult>,
    private readonly planSelection: (params: {
      workspaceViewRoot: string;
      outRoot: string;
      cacheRoot?: string;
      presetIds: string[];
      manualModules: string[];
    }) => Promise<SelectionPlanResult>,
    private readonly startIndex: (
      rebuild: boolean, onProgress?: (status: IndexResult) => void,
    ) => Promise<IndexResult>,
  ) {}

  resolveWebviewView(view: vscode.WebviewView): void {
    this.view = view;
    view.webview.options = { enableScripts: true };
    view.webview.html = profileHtml(this.initial(), webviewMessages());
    this.attachMessages(view.webview);
    void view.webview.postMessage({ type: "indexStatus", text: this.indexStatus });
    void this.loadPresets(view.webview);
    void this.refreshModules(view.webview, undefined, undefined, "", 1);
  }

  reveal(): void { this.view?.show?.(true); }

  updateIndexStatus(text: string): void {
    this.indexStatus = text;
    void this.view?.webview.postMessage({ type: "indexStatus", text });
  }

  showCurrentIndexStatus(): void { this.updateIndexStatus(this.indexStatus); }

  private initial(): InitialState {
    const config = vscode.workspace.getConfiguration("aospIde");
    const workspaceRoot = config.get<string>("workspaceViewRoot") ||
      vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "";
    const outRoot = normalizeOutRoot(config.get<string>("outRoot") ?? "");
    const inspected = config.inspect<string[]>("selectedPresets");
    const configured = inspected?.workspaceValue !== undefined;
    return {
      workspaceViewRoot: workspaceRoot,
      workspaceName: path.basename(workspaceRoot) || "AOSP",
      outRoot,
      outName: path.basename(outRoot) || "Out",
      selectedPresets: config.get<string[]>("selectedPresets") ?? [],
      selectedModules: configured ? config.get<string[]>("selectedModules") ?? [] : [],
      presetSelectionConfigured: configured,
    };
  }

  private attachMessages(webview: vscode.Webview): void {
    webview.onDidReceiveMessage(async (message) => {
      switch (message.type) {
      case "browse": await this.browse(webview, message.field); return;
      case "refreshModules":
        await this.refreshModules(webview, message.workspaceViewRoot, message.outRoot, "", 1);
        return;
      case "searchModules":
        await this.refreshModules(
          webview, message.workspaceViewRoot, message.outRoot,
          String(message.query ?? ""), 50,
        );
        return;
      case "applySelection":
        await this.applySelection(message);
        await webview.postMessage({ type: "showWorkbench" });
        await this.runIndex(false);
        return;
      case "rebuildIndex": await this.runIndex(true); return;
      case "openSettings":
        await vscode.commands.executeCommand(
          "workbench.action.openSettings", "@ext:rtfsc.aosp-ide-vscode",
        );
        return;
      }
    }, undefined, this.context.subscriptions);
  }

  private async browse(webview: vscode.Webview, field: string): Promise<void> {
    const selected = await vscode.window.showOpenDialog({
      canSelectFolders: true, canSelectFiles: false, canSelectMany: false,
      openLabel: vscode.l10n.t("Select directory"),
    });
    if (!selected?.[0]) return;
    const value = field === "outRoot" ?
      normalizeOutRoot(selected[0].fsPath) : selected[0].fsPath;
    await webview.postMessage({ type: "set", field, value });
  }

  private async loadPresets(webview: vscode.Webview): Promise<void> {
    try {
      const catalog = await this.loadSelectionCatalog();
      const initial = this.initial();
      await webview.postMessage({
        type: "presetCatalog", presets: catalog.presets,
        selected: initial.presetSelectionConfigured ?
          initial.selectedPresets : catalog.defaultPresetIds,
        manualModules: initial.selectedModules,
      });
    } catch (error) {
      await this.postError(webview, error);
    }
  }

  private async refreshModules(
    webview: vscode.Webview,
    rawRoot?: string,
    rawOutRoot?: string,
    query = "",
    limit = 50,
  ): Promise<void> {
    const root = String(rawRoot || this.initial().workspaceViewRoot).trim();
    if (!root) {
      await this.postError(webview, vscode.l10n.t("Open or choose a workspace first."));
      return;
    }
    await webview.postMessage({ type: "moduleLoading", query });
    try {
      const outRoot = normalizeOutRoot(String(rawOutRoot ?? this.initial().outRoot));
      const result = await this.discoverModules(root, outRoot, query, limit);
      await webview.postMessage({
        type: "moduleCatalog", query, modules: result.modules,
        total: result.total, buildBacked: result.buildBacked,
        filesScanned: result.filesScanned, cacheHit: result.cacheHit,
        diagnostics: result.diagnostics ?? 0,
      });
    } catch (error) {
      await this.postError(webview, error);
    }
  }

  private async applySelection(message: any): Promise<void> {
    const presetIds = strings(message.selectedPresets);
    const manualModules = strings(message.selectedModules);
    const workspaceViewRoot = String(message.workspaceViewRoot ?? "").trim();
    const outRoot = normalizeOutRoot(String(message.outRoot ?? ""));
    const plan = await this.planSelection({
      workspaceViewRoot, outRoot, cacheRoot: this.context.storageUri?.fsPath,
      presetIds, manualModules,
    });
    const config = vscode.workspace.getConfiguration("aospIde");
    await config.update("workspaceViewRoot", workspaceViewRoot, vscode.ConfigurationTarget.Workspace);
    await config.update("outRoot", outRoot, vscode.ConfigurationTarget.Workspace);
    await config.update("selectedPresets", plan.presetIds, vscode.ConfigurationTarget.Workspace);
    await config.update("selectedModules", manualModules, vscode.ConfigurationTarget.Workspace);
    await config.update("targets", plan.targets, vscode.ConfigurationTarget.Workspace);
  }

  private async runIndex(rebuild: boolean): Promise<void> {
    this.updateIndexStatus(vscode.l10n.t(
      rebuild ? "Rebuilding workspace index..." : "Applying changes...",
    ));
    try {
      const status = await this.startIndex(rebuild, (progress) => {
        this.updateIndexStatus(formatIndexStatus(progress));
      });
      this.updateIndexStatus(formatIndexStatus(status));
    } catch (error) {
      this.updateIndexStatus(vscode.l10n.t(
        "Index unavailable: {0}", error instanceof Error ? error.message : String(error),
      ));
    }
  }

  private async postError(webview: vscode.Webview, error: unknown): Promise<void> {
    await webview.postMessage({
      type: "moduleError",
      text: error instanceof Error ? error.message : String(error),
    });
  }
}

function strings(value: unknown): string[] {
  return Array.isArray(value) ? value.filter(
    (item): item is string => typeof item === "string" && item.trim() !== "",
  ) : [];
}

function formatIndexStatus(status: IndexResult): string {
  return `${status.state} · ${status.stage} · ${status.progress}% · ` +
    `${status.modules} modules · ${status.generated} generated · ` +
    `${status.diagnostics} diagnostics`;
}

export function profileHtml(initial: InitialState, messages: Record<string, string>): string {
  const nonce = String(Date.now());
  return `<!doctype html><html><head><meta charset="UTF-8"><meta http-equiv="Content-Security-Policy" content="default-src 'none'; style-src 'unsafe-inline'; script-src 'nonce-${nonce}'"><style>
  :root{--cyan:#22d3ee;--violet:#8b5cf6;--pink:#ec4899;--green:#34d399;--orange:#f59e0b;--red:#f87171;--muted:var(--vscode-descriptionForeground);--border:color-mix(in srgb,var(--vscode-foreground) 16%,transparent);--soft:color-mix(in srgb,var(--vscode-editor-background) 74%,var(--vscode-sideBar-background))}*{box-sizing:border-box}body{margin:0;background:var(--vscode-sideBar-background);color:var(--vscode-foreground);font:12px -apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif;letter-spacing:0}.accent{height:3px;background:linear-gradient(90deg,var(--cyan),var(--violet),var(--pink),var(--green))}.shell{padding:10px}.header{display:flex;align-items:center;justify-content:space-between;padding:2px 0 9px}.brand{display:flex;align-items:center;gap:7px;min-width:0}.mark{width:9px;height:18px;border-radius:3px;background:linear-gradient(180deg,var(--cyan),var(--violet),var(--pink))}.brand b{font-size:14px}.state-dot{width:7px;height:7px;border-radius:50%;background:var(--orange)}.icon-actions{display:flex;gap:3px}.icon{width:27px;height:27px;padding:0;border:0;border-radius:6px;background:transparent;color:var(--vscode-foreground);font-size:16px;cursor:pointer}.icon:hover{background:var(--vscode-toolbar-hoverBackground)}.section{padding:10px 0;border-top:1px solid var(--border)}.section:first-of-type{border-top:0}.section-title{display:flex;justify-content:space-between;align-items:center;margin-bottom:7px;color:var(--muted);font-size:10px;font-weight:700;text-transform:uppercase}.metric{display:flex;justify-content:space-between;gap:8px;align-items:flex-start;margin:7px 0}.metric-main{min-width:0}.metric-name{font-weight:700;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}.metric-path{margin-top:2px;color:var(--muted);font-size:10px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}.status{display:inline-flex;padding:3px 6px;border-radius:5px;background:color-mix(in srgb,var(--green) 14%,transparent);color:var(--green);font-size:10px;white-space:nowrap}.status.source{background:color-mix(in srgb,var(--cyan) 12%,transparent);color:var(--cyan)}.steps{display:grid;grid-template-columns:repeat(6,1fr);gap:3px;margin:9px 0}.step{height:4px;border-radius:2px;background:var(--border)}.step.active{background:linear-gradient(90deg,var(--cyan),var(--violet))}.index-text{line-height:1.45;color:var(--muted);word-break:break-word}.preset-list{display:flex;flex-direction:column;gap:3px}.preset-row{display:flex;justify-content:space-between;gap:8px;padding:5px 0}.preset-row span:first-child{white-space:nowrap;overflow:hidden;text-overflow:ellipsis}.count,.summary{color:var(--muted);font-size:10px}.button{border:1px solid var(--border);border-radius:6px;padding:7px 9px;background:var(--vscode-button-secondaryBackground);color:var(--vscode-button-secondaryForeground);font-weight:700;cursor:pointer}.button:hover{background:var(--vscode-button-secondaryHoverBackground)}.button.primary{border:0;background:linear-gradient(100deg,var(--violet),var(--pink));color:white}.button.full{width:100%;margin-top:7px}.actions,.manage-footer{display:grid;grid-template-columns:1fr 1fr;gap:6px;padding-top:9px;border-top:1px solid var(--border)}.path-editor{display:none;padding:8px;margin:3px 0 9px;background:var(--soft);border-radius:7px}.path-editor.visible{display:block}.field{margin-bottom:7px}.field label{display:block;color:var(--muted);font-size:10px;margin-bottom:3px}.field-row{display:flex;gap:4px}.field-row input,.search{min-width:0;width:100%;height:30px;border:1px solid var(--border);border-radius:6px;padding:5px 7px;background:var(--vscode-input-background);color:var(--vscode-input-foreground)}.manage{display:none}.manage.visible{display:block}.workbench.hidden{display:none}.manage-head{display:flex;align-items:center;justify-content:space-between;padding-bottom:9px}.manage-title{display:flex;align-items:center;gap:6px;font-weight:700;font-size:14px}.tabs{display:flex;gap:4px;overflow-x:auto;padding:7px 0}.tab{border:0;border-radius:5px;padding:5px 7px;background:var(--soft);color:var(--muted);white-space:nowrap;cursor:pointer}.tab.active{background:color-mix(in srgb,var(--violet) 22%,transparent);color:var(--vscode-foreground)}.preset-option,.module-option{display:grid;grid-template-columns:18px 1fr auto;gap:6px;align-items:start;padding:7px 0;border-top:1px solid var(--border)}.preset-option:first-child,.module-option:first-child{border-top:0}.preset-option input,.module-option input{margin:2px 0 0}.option-name{font-weight:700}.option-meta{color:var(--muted);font-size:10px;line-height:1.4;margin-top:2px;word-break:break-word}.badges{display:flex;gap:3px;flex-wrap:wrap;margin-top:4px}.badge{padding:2px 4px;border-radius:4px;background:color-mix(in srgb,var(--cyan) 12%,transparent);color:var(--cyan);font-size:9px}.recommended{color:var(--pink);font-size:9px}.results{max-height:290px;overflow:auto}.empty{padding:12px 0;color:var(--muted);text-align:center;line-height:1.5}.error{color:var(--red)}</style></head><body><div class="accent"></div><div class="shell">
  <header class="header"><div class="brand"><span class="mark"></span><b>AOSP Workspace</b><span class="state-dot" id="stateDot"></span></div><div class="icon-actions"><button class="icon" id="refresh" title="${escapeHtml(messages.refreshCatalog)}">↻</button><button class="icon" id="settings" title="${escapeHtml(messages.settings)}">⚙</button></div></header>
  <main class="workbench" id="workbench"><div class="path-editor" id="pathEditor"><div class="field"><label>${escapeHtml(messages.workspace)}</label><div class="field-row"><input id="workspaceRoot" value="${escapeHtml(initial.workspaceViewRoot)}"><button class="button" data-browse="workspaceRoot">${escapeHtml(messages.browse)}</button></div></div><div class="field"><label>${escapeHtml(messages.outRoot)}</label><div class="field-row"><input id="outRoot" value="${escapeHtml(initial.outRoot)}"><button class="button" data-browse="outRoot">${escapeHtml(messages.browse)}</button></div></div></div><section class="section"><div class="section-title"><span>${escapeHtml(messages.workspace)}</span><button class="icon" id="paths" title="${escapeHtml(messages.configurePaths)}">⋯</button></div><div class="metric"><div class="metric-main"><div class="metric-name" id="workspaceName">${escapeHtml(initial.workspaceName)}</div><div class="metric-path" id="workspacePath" title="${escapeHtml(initial.workspaceViewRoot)}">${escapeHtml(initial.workspaceViewRoot)}</div></div></div><div class="metric"><div class="metric-main"><div class="metric-name" id="outName">${escapeHtml(initial.outName)}</div><div class="metric-path" id="outPath" title="${escapeHtml(initial.outRoot)}">${escapeHtml(initial.outRoot)}</div></div><span class="status source" id="outStatus">${escapeHtml(initial.outRoot ? messages.buildBacked : messages.sourceOnly)}</span></div><div class="summary" id="catalogSummary">${escapeHtml(messages.discovering)}</div></section><section class="section"><div class="section-title"><span>${escapeHtml(messages.index)}</span><span id="indexPercent">0%</span></div><div class="steps">${Array.from({length:6},()=>'<span class="step"></span>').join('')}</div><div class="index-text" id="indexStatus">${escapeHtml(messages.notStarted)}</div></section><section class="section"><div class="section-title"><span>${escapeHtml(messages.enabledPresets)}</span><span id="presetCount">0</span></div><div class="preset-list" id="enabledPresets"></div><button class="button full" id="manage">${escapeHtml(messages.manageModules)}</button></section><div class="actions"><button class="button primary" id="applyChanges">${escapeHtml(messages.applyChanges)}</button><button class="button" id="rebuild">${escapeHtml(messages.rebuild)}</button></div></main>
  <main class="manage" id="manageView"><div class="manage-head"><div class="manage-title"><button class="icon" id="back">←</button><span>${escapeHtml(messages.modules)}</span></div><span class="count" id="selectionCount"></span></div><input class="search" id="query" placeholder="${escapeHtml(messages.searchModules)}"><div class="tabs"><button class="tab active" data-filter="all">${escapeHtml(messages.all)}</button><button class="tab" data-filter="recommended">${escapeHtml(messages.recommended)}</button><button class="tab" data-filter="java">Java</button><button class="tab" data-filter="android">Android</button><button class="tab" data-filter="native">Native</button><button class="tab" data-filter="rust">Rust</button></div><section class="section"><div class="section-title"><span>${escapeHtml(messages.presets)}</span></div><div id="presetOptions"></div></section><section class="section"><div class="section-title"><span>${escapeHtml(messages.matchingModules)}</span></div><div class="summary" id="searchSummary"></div><div class="results" id="moduleResults"><div class="empty">${escapeHtml(messages.searchHint)}</div></div></section><div class="manage-footer"><button class="button" id="cancel">${escapeHtml(messages.cancel)}</button><button class="button primary" id="applySelection">${escapeHtml(messages.apply)}</button></div></main>
  </div><script nonce="${nonce}">${webviewScript(initial, messages)}</script></body></html>`;
}

function webviewScript(initial: InitialState, messages: Record<string, string>): string {
  return `const api=acquireVsCodeApi(),S=${jsonForScript(initial)},T=${jsonForScript(messages)},$=id=>document.getElementById(id);let presets=[],selectedPresets=new Set(S.selectedPresets),selectedModules=new Set(S.selectedModules),modules=[],filter='all',timer;const esc=v=>String(v??'').replaceAll('&','&amp;').replaceAll('<','&lt;').replaceAll('>','&gt;').replaceAll('"','&quot;');const count=()=>{const n=selectedPresets.size+selectedModules.size;$('selectionCount').textContent=T.selected.replace('{0}',n);$('presetCount').textContent=selectedPresets.size;$('applySelection').textContent=T.apply+' '+n};const renderEnabled=()=>{$('enabledPresets').innerHTML=presets.filter(p=>selectedPresets.has(p.id)).map(p=>'<div class="preset-row"><span>'+esc(p.name)+'</span><span class="count">'+p.modules.length+'</span></div>').join('')||'<div class="empty">'+esc(T.notStarted)+'</div>';count()};const renderPresets=()=>{$('presetOptions').innerHTML=presets.map(p=>'<label class="preset-option"><input type="checkbox" data-preset="'+esc(p.id)+'" '+(selectedPresets.has(p.id)?'checked':'')+'><div><div class="option-name">'+esc(p.name)+(p.recommended?' <span class="recommended">'+esc(T.recommended)+'</span>':'')+'</div><div class="option-meta">'+p.modules.map(esc).join(' · ')+'</div></div><span class="count">'+p.modules.length+'</span></label>').join('');document.querySelectorAll('[data-preset]').forEach(i=>i.onchange=()=>{i.checked?selectedPresets.add(i.dataset.preset):selectedPresets.delete(i.dataset.preset);renderEnabled()})};const match=m=>filter==='all'||filter==='recommended'&&m.recommended||filter==='java'&&m.type.toLowerCase().includes('java')||filter==='android'&&m.type.toLowerCase().includes('android')||filter==='native'&&(m.type.toLowerCase().includes('cc_')||m.capabilities?.includes('native'))||filter==='rust'&&m.capabilities?.includes('rust');const renderModules=()=>{const v=modules.filter(match);$('moduleResults').innerHTML=v.length?v.map(m=>'<label class="module-option"><input type="checkbox" data-module="'+esc(m.name)+'" '+(selectedModules.has(m.name)?'checked':'')+'><div><div class="option-name">'+esc(m.name)+(m.recommended?' <span class="recommended">'+esc(T.recommended)+'</span>':'')+'</div><div class="option-meta">'+esc(m.type)+' · '+esc(m.status)+'<br>'+esc(m.blueprint)+'</div><div class="badges">'+(m.capabilities||[]).map(x=>'<span class="badge">'+esc(x)+'</span>').join('')+'</div></div></label>').join(''):'<div class="empty">'+esc(T.noResults)+'</div>';document.querySelectorAll('[data-module]').forEach(i=>i.onchange=()=>{i.checked?selectedModules.add(i.dataset.module):selectedModules.delete(i.dataset.module);count()})};const showManage=()=>{$('workbench').classList.add('hidden');$('manageView').classList.add('visible');renderPresets();renderModules();count()};const showWorkbench=()=>{$('manageView').classList.remove('visible');$('workbench').classList.remove('hidden');renderEnabled()};const search=()=>{const q=$('query').value.trim();if(q.length<2){modules=[];$('searchSummary').textContent='';$('moduleResults').innerHTML='<div class="empty">'+esc(T.searchHint)+'</div>';return}api.postMessage({type:'searchModules',workspaceViewRoot:$('workspaceRoot').value,outRoot:$('outRoot').value,query:q})};document.querySelectorAll('[data-browse]').forEach(b=>b.onclick=()=>api.postMessage({type:'browse',field:b.dataset.browse}));$('refresh').onclick=()=>api.postMessage({type:'refreshModules',workspaceViewRoot:$('workspaceRoot').value,outRoot:$('outRoot').value});$('settings').onclick=()=>api.postMessage({type:'openSettings'});$('paths').onclick=()=>$('pathEditor').classList.toggle('visible');$('manage').onclick=showManage;$('back').onclick=showWorkbench;$('cancel').onclick=showWorkbench;const apply=()=>api.postMessage({type:'applySelection',workspaceViewRoot:$('workspaceRoot').value,outRoot:$('outRoot').value,selectedPresets:[...selectedPresets],selectedModules:[...selectedModules]});$('applyChanges').onclick=apply;$('applySelection').onclick=apply;$('rebuild').onclick=()=>api.postMessage({type:'rebuildIndex'});$('query').oninput=()=>{clearTimeout(timer);timer=setTimeout(search,350)};document.querySelectorAll('[data-filter]').forEach(b=>b.onclick=()=>{filter=b.dataset.filter;document.querySelectorAll('[data-filter]').forEach(x=>x.classList.toggle('active',x===b));renderModules()});window.addEventListener('message',e=>{const d=e.data;if(d.type==='set'){ $(d.field).value=d.value;const n=d.value.split(/[\\/]/).filter(Boolean).pop()||d.value;if(d.field==='workspaceRoot'){$('workspaceName').textContent=n;$('workspacePath').textContent=d.value;$('workspacePath').title=d.value}else{$('outName').textContent=n;$('outPath').textContent=d.value;$('outPath').title=d.value;$('outStatus').textContent=d.value?T.buildBacked:T.sourceOnly}}if(d.type==='presetCatalog'){presets=d.presets;selectedPresets=new Set(d.selected);selectedModules=new Set(d.manualModules);renderPresets();renderEnabled()}if(d.type==='moduleLoading'){$('catalogSummary').textContent=T.discovering;if(d.query)$('searchSummary').textContent=T.discovering}if(d.type==='moduleCatalog'){if(d.query){modules=d.modules;$('searchSummary').textContent=T.matching.replace('{0}',d.total)+' · '+T.showing.replace('{0}',d.modules.length);renderModules()}else $('catalogSummary').textContent=T.matching.replace('{0}',d.total)+' · '+T.buildBackedCount.replace('{0}',d.buildBacked)}if(d.type==='moduleError'){$('catalogSummary').innerHTML='<span class="error">'+esc(d.text)+'</span>';$('searchSummary').innerHTML='<span class="error">'+esc(d.text)+'</span>'}if(d.type==='indexStatus'){const p=Number(/(\\d+)%/.exec(d.text)?.[1]||0);$('indexPercent').textContent=p+'%';$('indexStatus').textContent=d.text;document.querySelectorAll('.step').forEach((s,i)=>s.classList.toggle('active',p>=(i+1)*16));$('stateDot').style.background=p===100?'var(--green)':'var(--orange)'}if(d.type==='showWorkbench')showWorkbench()});renderEnabled();`;
}

function jsonForScript(value: unknown): string {
  return JSON.stringify(value).replaceAll("<", "\\u003c");
}

function escapeHtml(value: string): string {
  return value.replaceAll("&", "&amp;").replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;").replaceAll('"', "&quot;");
}
