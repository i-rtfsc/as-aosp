# ProjectGraph 与 IDE Projection 设计

> 状态：Target；客户端只消费 ProjectGraph，不重复拼装 ModuleCatalog。

## 1. Build Facts Source

ProjectGraph 只接受一个 Build Facts Root：`outRoot`。该目录可以是本机完整 AOSP out，
也可以是从另一台机器复制来的稀疏 Portable Out。两者必须保持完全相同的 OUT_DIR
相对路径，客户端和 engine 不区分来源，也不解析另一套 Pack schema。

ModuleCatalog、ArtifactCatalog、AIDL/JNI index 等派生数据由本机 engine 在 IDE cache 中
生成，不写入源码树、out 或 Portable Out。两者都未配置时返回明确诊断；不得通过扫描
完整源码树猜测 artifact owner。

RPC 输入为：

```text
outRoot: string
targets: TargetSpec[]
```

用户的 presets、manual modules 和 Select All 不直接生成该数组。Selection Planner 先执行
全局 module/closure/path 去重、ownership 和 bounded target grouping，再把稳定 TargetSpec
集合提交给 ProjectGraph。详细规则见 `module-selection-planner-design.md`。

Artifact Catalog 查询使用同一 out reader，因此图形 selector 在本机完整 out 和复制来的
稀疏 out 场景下保持一致。

Generated source materialization 必须是 target-scoped：客户端先
取得 ProjectGraph，再把图中明确引用的 `.srcjar` archives 提交给 materialize RPC。服务端
只接受位于 Out Root 内的绝对路径，或可安全解析到 Out Root 的 `out/...`、`soong/...`
相对路径；不得遍历并解压 out 中所有 srcjar。

## 2. JetBrains

- one ProjectGraph target becomes one focused IDE module when feasible;
- multiple targets become multiple modules with stable names;
- shared libraries and framework srcjars become dedicated library/modules;
- module dependency graph becomes `ModuleOrderEntry`;
- target-to-target dependency edges are computed by the engine from resolved module ownership;
- duplicate source roots are assigned to one owner;
- nested roots produce `excludedSourceRoots` on the parent target so a more specific target owns
  Services/SystemUI without duplicate indexing inside Framework;
- current non-persistent focused module remains the initial ProjectGraph implementation.

## 3. VS Code

- current `java.project.sourcePaths/referencedLibraries` is the initial ProjectGraph projection;
- VS Code Java projection disables the Maven and Gradle importers for the workspace. AOSP
  contains embedded Gradle projects that are not standalone builds; allowing JDT LS to import
  them can fail workspace initialization and disable Java navigation even for local symbols.
- The VS Code Gradle extension is also configured with `gradle.autoDetect=off` and
  `gradle.nestedProjects=false`; its independent Gradle server must not scan AOSP sample and
  vendor projects. This does not disable AOSP builds from a terminal.
- ProjectGraph targets become separate JDT project descriptions/importer inputs;
- generated/source attachments are target-scoped;
- clangd consumes target-scoped compdb overlays;
- AIDL/JNI shared engine links remain cross-target actions.

## 4. Ownership

Source ownership key is normalized RepoPath. Resolution priority is:

```text
pinned > more specific root > manual > preset > dependency-only > stable tie-break
```

该 priority 由 Selection Planner 编译。多个 targets 引用同一 source root 时只能有一个
owner，其他 targets 通过 dependency edge 使用它。

Conflicting physical providers produce diagnostics; clients must not silently duplicate the same
source root in multiple Java modules.

## 5. Migration

1. Build ProjectGraph DTO and RPC.
2. Require explicit TargetSpec configuration before projection.
3. Upgrade JetBrains module projection.
4. Add VS Code JDT project/importer integration.
5. Keep source-only only for missing facts; do not add a separate project-selection fallback.
