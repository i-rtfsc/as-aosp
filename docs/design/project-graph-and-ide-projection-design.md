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

- VS Code uses one focused unmanaged JDT project backed by
  `java.project.sourcePaths/referencedLibraries`. The paths are generated from
  the selected ProjectGraph, not from a whole-tree scan.
- Java projection consumes only `java` and `android` ProjectGraph targets. Native targets are
  projected to clangd and Rust targets to rust-analyzer; they must never become JDT source roots.
- Java source roots are derived from explicit `.java`/`.kt` source paths and package declarations.
  A broad module path such as `frameworks/base` is only a fallback when no explicit Java sources
  are available; it must not be preferred over `frameworks/base/core/java`.
- Aggregator modules follow source-composition edges such as `static_libs` to find their owned
  sources (`services -> services.impl`, `SystemUI -> SystemUI-application -> SystemUI-core`).
  Ordinary classpath dependencies remain libraries and do not transfer source ownership.
- AOSP workspace projection uses a dedicated `.aosp-ide/java/classes` output directory, disables
  JDT auto-build diagnostics, and requests a 6 GiB JDT heap unless the workspace already has an
  explicit VM setting. Navigation must not wait for full AOSP diagnostics publication.
- Target-scoped srcjar archives are materialized into one deduplicated
  `.aosp-ide/java/generated` overlay source root. A stable archive-hash manifest reuses the
  overlay; JDT must not receive hundreds of individual srcjar source roots.
- JDT workspace resource filters exclude `.repo`, Out, prebuilts, test/testdata, and benchmark
  directories. The generated Java overlay remains visible and is not filtered.
- VS Code Java projection disables the Maven and Gradle importers for the workspace. AOSP
  contains embedded Gradle projects that are not standalone builds; allowing JDT LS to import
  them can fail workspace initialization and disable Java navigation even for local symbols.
- The VS Code Gradle extension is also configured with `gradle.autoDetect=off` and
  `gradle.nestedProjects=false`; its independent Gradle server must not scan AOSP sample and
  vendor projects. This does not disable AOSP builds from a terminal.
- ProjectGraph Java targets are merged into one focused VS Code build path;
- generated/source attachments are target-scoped;
- clangd consumes target-scoped compdb overlays;
- AIDL/JNI shared engine links remain cross-target actions.

Generated layout:

```text
.aosp-ide/java/generated/
.aosp-ide/java/classes/
```

The projection writes repository-relative Java roots directly to
`java.project.sourcePaths`. Therefore a file opened at its original AOSP URI is
on the default project's build path. Generated sources use one overlay root.
Dependency jars are reduced to one preferred Android/header variant per jar
basename before they are written to `java.project.referencedLibraries`.

Nested Eclipse projects with linked resources were rejected after host
verification. Red Hat Java imported the default `aosp_<hash>` project instead
of the descriptors below `.aosp-ide`, and original AOSP files produced
`is not on its project's build path`. Linked descriptors are not a supported
runtime path for VS Code unless a dedicated JDT importer extension is added.

The Java server uses `Standard` mode for this projection.

Projection readiness is independent from JDT background indexing. Writing the
focused build path marks the AOSP projection ready;
`java.projectConfiguration.update` runs asynchronously. The sidebar reports
selected closure modules and generated entries. It must not display the full
Out catalog counts as if all modules were selected.

Host acceptance requires the original
`frameworks/base/services/core/java/...` URI to be on the default project's
build path and cross-file definition requests to resolve without linked paths.

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
4. Generate the focused VS Code unmanaged-project build path.
5. Keep source-only only for missing facts; do not add a separate project-selection fallback.
