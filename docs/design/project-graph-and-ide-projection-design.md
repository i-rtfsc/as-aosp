# ProjectGraph 与 IDE Projection 设计

> 状态：Target；客户端只消费 ProjectGraph，不重复拼装 ModuleCatalog。

## 1. Build Facts Source

ProjectGraph 接受两种等价的事实来源：

- `metadataPackRoot`：已复制到本机的正式 Metadata Pack；
- `metadataRoot`：本机 AOSP `out` 或 `out/soong`，由 metadata reader 只读加载。

Pack 存在时优先使用并执行完整校验；未配置 Pack 时直接从本机 out 构建临时
ModuleCatalog 与 ArtifactCatalog。客户端不要求本机 out 用户先导出 Pack，也不自行解析
Soong JSON。两者都未配置时返回明确诊断；不得通过扫描完整源码树猜测 artifact owner。

RPC 输入为：

```text
metadataPackRoot?: string
metadataRoot?: string
targets: TargetSpec[]
```

Artifact Catalog 查询遵循同一来源选择规则，因此图形 selector 在本机 out 和复制 Pack
场景下保持一致。

## 2. JetBrains

- one ProjectGraph target becomes one focused IDE module when feasible;
- multiple targets become multiple modules with stable names;
- shared libraries and framework srcjars become dedicated library/modules;
- module dependency graph becomes `ModuleOrderEntry`;
- duplicate source roots are assigned to one owner;
- current non-persistent focused module remains the initial ProjectGraph implementation.

## 3. VS Code

- current `java.project.sourcePaths/referencedLibraries` is the initial ProjectGraph projection;
- ProjectGraph targets become separate JDT project descriptions/importer inputs;
- generated/source attachments are target-scoped;
- clangd consumes target-scoped compdb overlays;
- AIDL/JNI shared engine links remain cross-target actions.

## 4. Ownership

Source ownership key is normalized RepoPath. Resolution priority is:

```text
more specific target root > framework priority > explicit user priority > lexical stable tie-break
```

Conflicting physical providers produce diagnostics; clients must not silently duplicate the same
source root in multiple Java modules.

## 5. Migration

1. Build ProjectGraph DTO and RPC.
2. Require explicit TargetSpec configuration before projection.
3. Upgrade JetBrains module projection.
4. Add VS Code JDT project/importer integration.
5. Keep source-only only for missing facts; do not add a separate project-selection fallback.
