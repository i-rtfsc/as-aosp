# AOSP 多 IDE 导航方案重新审计

> 复核日期：2026-08-26。证据包含 AOSP 官方 AIDEGen、Soong compdb、
> `build/make/tools/ide_query`、Eclipse JDT LS、VS Code Java、GitHub AOSP compdb
> 项目，以及本仓当前实现。

主要可复核来源：

- AOSP `tools/asuite/aidegen`：https://android.googlesource.com/platform/tools/asuite/+/main/aidegen/
- AOSP Soong compdb：`/Users/solo/code/aosp/build/soong/cc/compdb.go` 和
  `build/soong/docs/compdb.md`；
- AOSP ide_query：`/Users/solo/code/aosp/build/make/tools/ide_query/ide_query.go`；
- AIDEGen Git mirror：https://android.googlesource.com/platform/tools/asuite/+/main/aidegen/README.md
- Eclipse JDT LS：https://github.com/eclipse-jdtls/eclipse.jdt.ls；
- VS Code Java：https://github.com/redhat-developer/vscode-java；
- AOSP VS Code compdb generator：GitHub `amezin/aosp-vscode` 的
  `generate_compdb.py`（按 Ninja rule/build 生成 compile_commands）。

## 1. 结论

当前方案的基础方向合理，不需要推翻：

```text
Soong/out facts or portable Metadata Pack
  -> ModuleCatalog / generated / classpath / compdb
  -> IDE-native Java project model and clangd
  -> shared AIDL/JNI cross-language edges
```

需要调整的是客户端 project projection，而不是 Pack 或共享 engine。当前 JetBrains
将 selected closure 压进一个 non-persistent module，VS Code 将全部 roots/libraries
写入 unmanaged Java project settings。这适合作为首个 MVP，但不是最终模型；大型 AOSP
会遇到重复 source root、同包不同实现、依赖顺序、bootclasspath 和 variant 隔离问题。

## 2. 官方方案的关键事实

### 2.1 AIDEGen

AIDEGen 的目标是为 Android Studio、IntelliJ、Eclipse、CLion 和 VS Code 配置项目，
而不是提供统一语言服务器。它支持 module/path 选择、dependency depth、whole-tree、
skip-build 和不同 IDE target。

IntelliJ 生成器的关键行为：

- 按 project/module 生成独立 `.iml`；
- source、test、jar、srcjar 和 module dependency 分开写入；
- 为重复 `.iml` 名生成稳定唯一名；
- framework source 优先于普通 source，解决 duplicate package；
- `ProjectSplitter` 删除跨 module 重复 source roots；
- 单独生成 `framework_srcjars.iml` 和 dependencies module；
- 多 project 场景显式建立 module dependencies。

因此“把全部 source roots 放进一个 module”不能完整复现官方行为。

### 2.2 ide_query

`ide_query` 的行为证明 Build Facts 是正确事实源：

- 运行 Soong `nothing` 并显式设置 `SOONG_GEN_COMPDB=1`；
- Java 读取 `out/soong/module_bp_java_deps.json`；
- Java module 提供 path、deps、srcs、jars、srcjars；
- C/C++ 由 compdb 和 cc analyzer 计算 active file inputs；
- 根据 active file 找 owning module，并按需构建 target/generated files；
- 输出的 analysis 是 file/build-unit scoped，不是全树静态真理。

当前 Pack 读取相同 facts 的方向正确，但后续应增加 owner/build-unit 查询，而不只返回
全 catalog closure。

### 2.3 Soong compdb

Soong 仅在 `SOONG_GEN_COMPDB=1` 时生成：

```text
out/soong/development/ide/compdb/compile_commands.json
```

官方建议使用 `make nothing` 生成全图；通过 mm/mmm 等有限构建生成的数据库只覆盖
被访问模块。`SOONG_GEN_COMPDB_MODULES` 还能进一步过滤。当前 UI 必须显示 compdb
coverage/freshness，不能只显示 entry count 后宣称 Native Ready。

### 2.4 VS Code Java

VS Code Java 使用 Eclipse JDT LS。其一等 project import 是 Maven/Gradle；unmanaged
folder 的 `java.project.sourcePaths` 和 `referencedLibraries` 是可用 fallback，但它没有
AOSP/Soong module dependency、bootclasspath 和 generated ownership 语义。

所以当前 VS Code settings projection 可用于 MVP 和 framework baseline，最终需要生成
workspace-scoped JDT project descriptions，或者实现 JDT LS extension/importer，将
ModuleCatalog 直接转换为多个 Java projects。

### 2.5 VS Code / CLion Native

AIDEGen 的 VS Code native generator 和 GitHub `amezin/aosp-vscode` 都把 compdb 作为
核心入口。AIDEGen CLion 路径则从 CC module facts 生成 CMakeLists。当前 shared compdb
overlay 是正确方向；自定义 C++ parser 不应成为主路径。

## 3. 当前方案逐项判断

| 当前能力 | 判断 | 调整 |
|---|---|---|
| Go engine + stdio RPC | 保留 | 继续作为共享 facts/query owner |
| Metadata Pack | 保留并加强 | 增加 identity、coverage、build-unit facts |
| ModuleCatalog closure | 保留 | 增加 owner/file/build-unit 查询 |
| generated srcjar cache | 正确 | 增加 module/source attachment 关系 |
| AIDL/JNI shared edges | 正确 | 保持辅助导航，不替代 Java/C++ 语义引擎 |
| JetBrains 单 focused module | 过渡实现 | 升级为多 module graph projection |
| VS Code sourcePaths settings | 过渡实现 | 升级为 JDT 多 project/importer |
| compdb overlay | 正确 | 增加 coverage、response files、compiler/toolchain relocation |
| source-only file scan | 仅 fallback | 不再扩展为主架构 |

## 4. 必须立即调整的设计

### 4.1 ProjectGraph DTO

在 ModuleCatalog 与客户端之间增加稳定 projection DTO：

```text
ProjectGraph
  projects[]
    id / module names / source roots / test roots
    generated roots / jars / source attachments
    project dependencies / compiler level / bootclasspath role
  shared libraries[]
  duplicate source ownership decisions[]
  diagnostics[]
```

这层负责 AIDEGen `ProjectSplitter` 等价规则，客户端不再各自把 catalog 扁平化。

### 4.2 JetBrains 多 module projection

按 selected top-level modules 生成稳定 module，依赖 closure 可合并到 dependencies module
以控制数量；framework srcjars 独立 module/library。必须保证一个 physical source root
只有一个 owner，并为 module dependency 建 `ModuleOrderEntry`。

### 4.3 VS Code JDT project integration

短期继续支持 `sourcePaths/referencedLibraries`。中期增加 JDT LS importer 或生成隔离的
project descriptors，每个 project 使用独立 source/classpath/dependency，而不是把所有
root 作为一个 unmanaged folder。不能恢复 Gradle 兼容方案作为产品主路径。

### 4.4 Pack identity 与 coverage

manifest 应增加：

- Android release、manifest digest、product/release/variant/arch；
- source revision digest；
- selected module closure；
- Java facts coverage；
- compdb generation mode、module filter、entry count；
- generated/archive owner coverage；
- exporter/engine schema version。

匹配等级继续按域判断，不能用 product/variant 阻塞 framework Java baseline。

## 5. 调整后的实施顺序

1. 等真实服务器 Pack，验证真实 schema/path/artifact coverage。
2. 实现 `ProjectGraph` 与 duplicate source ownership。
3. JetBrains 多 module/dependencies/framework-srcjars projection。
4. VS Code JDT 多 project/importer，保留当前 settings fallback。
5. Native compdb coverage/freshness 与 toolchain relocation。
6. file -> owning module/build-unit 查询和增量 invalidation。
7. Android Studio、VS Code、CLion/clangd、Windows/macOS/Linux 验证矩阵。

## 6. 不应继续扩展的方向

- 不开发自有 Java compiler/type checker；
- 不开发自有完整 C++ semantic engine；
- 不恢复 Gradle 兼容工程作为新方案主线；
- 不用 basename 或全树扫描补偿缺失 build facts；
- 不把一个兼容 AOSP Pack 宣称为 vendor tree 精确构建图；
- 不在真实 IDE 验证前继续增加大量 UI 状态或协议方法。

## 7. 最终判断

当前方案不是旧方案 workaround，而是与 AOSP 官方工具相同的 Build-Facts 路线。
Metadata Pack 是对 AIDEGen/ide_query “必须在 Linux 构建树运行”限制的合理跨机器扩展。

当前最大的结构性差距是 IDE project model fidelity：我们已经能把 facts 带到 Mac，
但还没有像 AIDEGen 一样保留 module 边界、dependency order、duplicate source ownership
和 framework srcjar 分层。下一阶段应集中修正这一层。
