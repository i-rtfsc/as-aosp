# ProjectGraph 与 IDE Projection 设计

> 状态：Target；客户端只消费 ProjectGraph，不重复拼装 ModuleCatalog。

## 1. JetBrains

- one ProjectGraph target becomes one focused IDE module when feasible;
- multiple targets become multiple modules with stable names;
- shared libraries and framework srcjars become dedicated library/modules;
- module dependency graph becomes `ModuleOrderEntry`;
- duplicate source roots are assigned to one owner;
- current non-persistent focused module remains fallback during migration.

## 2. VS Code

- current `java.project.sourcePaths/referencedLibraries` remains unmanaged fallback;
- ProjectGraph targets become separate JDT project descriptions/importer inputs;
- generated/source attachments are target-scoped;
- clangd consumes target-scoped compdb overlays;
- AIDL/JNI shared engine links remain cross-target actions.

## 3. Ownership

Source ownership key is normalized RepoPath. Resolution priority is:

```text
more specific target root > framework priority > explicit user priority > lexical stable tie-break
```

Conflicting physical providers produce diagnostics; clients must not silently duplicate the same
source root in multiple Java modules.

## 4. Migration

1. Build ProjectGraph DTO and RPC.
2. Translate legacy `selectedModules` to implicit targets.
3. Upgrade JetBrains module projection.
4. Add VS Code JDT project/importer integration.
5. Keep source-only and existing settings projection as degraded fallback.
