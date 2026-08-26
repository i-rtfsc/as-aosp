# AOSP IDE 当前状态

> 复核日期：2026-08-26。

## 1. Current

- `dev-ide` 分支已切换为新方案仓库边界；旧 Gradle/CMake 实现在 `aosp` 分支。
- 产品可行性、详细设计、拓扑和关键决策文档已存在。
- Phase 0 已有 Go model、topology/pathmap、metadata readers、ModuleGraph、AIDL/JNI source index 和 CLI；JetBrains plugin bootstrap 已可构建，VS Code extension 首个 TypeScript client slice 已可编译，并覆盖 framing、并发启动、超时清理、异常退出重启和 POSIX/Windows 路径 fixture。
- AOSP 源码研究已覆盖 Soong IdeInfo、Java/CC metadata、compdb、AIDL 和 JNI 模式。

## 2. Partial

- 根目录、Makefile、Phase 0 CLI、`aosp-index serve --stdio`、JetBrains 251 bootstrap 和 VS Code TypeScript bootstrap 已完成；ModuleRootModel/Workspace Model bridge、首个 clangd compdb overlay 和 VS Code engine status UI 已具备，完整 AIDL PSI 和远程执行入口尚未实现。
- WorkspaceTopology、PathMapper、SourceMount 和四类 metadata reader 已有初始实现，但真实 AOSP schema 覆盖、路径映射接线和完整 topology 发现仍未闭合。
- AIDL source parser/index、exact type ranges、references、generated peers 和 JNI source-only links 已接入 daemon；JetBrains/VS Code adapters 已有首个导航切片，但真实 IDE host、Android Studio verifier 和完整语言 UI 尚未闭合。
- Metadata Pack 首个 export slice 已实现：默认 Core JSON + Java/AIDL generated srcjar，`--native` 附加 CC/compdb；VS Code 已有首个 source-root projection，完整 pack import 和 IDE project projection 尚未闭合。
- `aosp-index` CLI 使用指南和 `pack validate` 已加入，用户可以从 Linux out 导出并校验可复制 Metadata Pack。
- `core/modules` 首个 ModuleCatalog slice 已实现，可合并 Java/CC/module-info facts 并计算依赖闭包；尚未接入 IDE project projection。
- daemon 已提供 `aosp/modules/catalog|get|closure`，可从显式 Metadata Pack
  返回稳定排序的模块摘要、单模块 facts 和依赖闭包；VS Code 侧边栏已能
  主动检查 pack 并显示模块、generated 和 diagnostics 数量。

## 3. Target

- Phase 0：Go core、profile、topology、path mapping、doctor/scan CLI 和 fixture tests。
- Phase 1：JetBrains MVP；随后 AIDL、C/C++/JNI 和 VS Code。

## 4. Negative capability

- 当前不能在未配置 engine/profile 时打开完整 AOSP 并自动提供导航；
- 当前 JetBrains 插件已有 descriptor、Kotlin RPC framing/lifecycle client、EngineProcess、project-local profile state、profile UI、focused root planner/非持久投影模块、source-only AIDL fallback、异步 AIDL definition/generated peer/references 和 JNI peer actions，并查找公共 `engines/aosp-index`；VS Code 已有 Java/Kotlin source-only type fallback、AIDL providers、AIDL generated peer/JNI commands 和 JNI DefinitionProvider；仍没有 AIDL PSI 或 JNI gutter；
- 当前 `aosp-index serve --stdio` 已启用 initialize、shutdown、profile validate、topology discover/validate、ModuleCatalog 查询、`aosp/jni/peers`、`aosp/aidl/resolve`、`aosp/aidl/peers` 和 `aosp/aidl/references`；长期增量索引尚未启用；
- IndexCoordinator 首个状态 slice 已启用 `aosp/index/start|rebuild|status|cancel`；它能从 Metadata Pack 产生阶段、进度、模块和 generated 计数，尚未投影 Java source roots 或 IDE 原生索引。
- VS Code 在 coordinator 完成后已把可证明的 RepoPath source roots 投影到
  `java.project.sourcePaths`，并通过 `aosp/generated/materialize` 将 pack 内 srcjar
  安全解压到 engine cache；构建机绝对路径不会写入本机配置，完整 Java LS classpath
  仍未闭合。
- JetBrains 配置页已增加 `Apply Build Facts`，会把 engine 返回的可见 RepoPath
  source roots、generated cache roots 和本机存在的 jar libraries 投影到非持久
  focused module；完整 Workspace Model module closure 仍未闭合。
- Metadata Pack identity RPC 已提供原始 `outRoot`，VS Code/JetBrains 仅在前缀可证明
  时将构建机 jar 重定位到本机 `outRoot`，不做 basename 猜测。
- pack export 已包含 `aidl/symbols.json` 与 `aidl/modules.json`；VS Code AIDL definition
  provider 配置 pack 后优先查询这些 facts，未配置 pack 时保留 workspace scan fallback。
- pack export 现在同时包含 `aidl/index.json`；VS Code AIDL definition、references 和
  generated peer actions 已优先查询 pack index，pack diagnostics 或缺失结果会回退
  workspace engine 查询。
- JetBrains AIDL 右键 definition、references、generated peer actions 已 pack-first；
  同步 PSI declaration handler 保持本地解析，避免阻塞编辑器线程。
- AIDL pack generated peer RPC 已支持 cacheRoot，能够把匹配的 srcjar Java entry
  materialize 为可打开的 exact localPath。
- `--native` Metadata Pack 已包含 `jni/index.json`；VS Code 和 JetBrains JNI peer action
  优先查询 native pack links，pack 缺失或 peer 不唯一时回退 source-only workspace index。
- `pack export --module` 已按 ModuleCatalog dependency closure 限制 AIDL/JNI source
  facts，并按 closure 的 `SrcJars/Generated` 完整相对路径过滤 generated archives；
  未知模块 fail closed，不按 basename 猜 owner。
- CLI、daemon、VS Code 和 JetBrains 现在复用 `core/pack` 校验；侧边栏检查和索引启动
  在消费 pack 前先 fail closed，避免损坏数据进入项目模型。
- daemon 的 ModuleCatalog、generated、AIDL、JNI 和 pack identity 查询也会在服务端
  校验正式 pack；缓存指纹覆盖所有文件的 path/size/mtime 与 manifest/checksum，
  pack 原子替换或原地文件变化后自动失效。
- daemon pack 校验采用 per-pack singleflight；同一 fingerprint 的并发查询共享一次
  checksum 工作，不同 pack 可以并行，校验期间发生变化的结果不会写入旧缓存。
- VS Code 已在已有 profile 的 workspace 激活时自动启动 coordinator；配置变化经
  300ms debounce 后取消旧 session 并 rebuild，侧边栏重新打开会恢复最后状态。
- coordinator 已支持 IDE projection 回报；VS Code 和 JetBrains 会上报 source/generated/
  library counts 与 diagnostics，snapshot 的最终状态覆盖 facts 与本机投影两段链路。
- VS Code 和 JetBrains 已切换到 Targets；TargetSpec/ArtifactResolver/ProjectGraph
  尚未完整实现；`core/model` 已加入 TargetSpec/ArtifactSelector，`core/artifacts` 已加入
  基于 path/name/digest 的 Catalog/Resolver，
  `core/projects` 已能从 target/module/artifact facts 构建首个 ProjectGraph 和 source
  ownership decisions；`aosp/projects/graph` RPC 与 pack artifact catalog 已接入，
  VS Code/JetBrains 投影只使用 ProjectGraph，selectedModules 路径已移除。
- pack export 已复制 ModuleCatalog 明确引用的 out 内 Java jars 到 `pack/out/**`；
  两端可在没有本机 out 时使用 pack classpath，inspect 会报告 jar/srcjar 统计。
- Native Pack 的 `compile_commands.json` 已可在 VS Code compdb overlay 中直接消费；
  `metadataRoot` 为空时 daemon 使用已校验的 `metadataPackRoot`，JetBrains compdb UI
  仍待接入。
- pack export 已规范化相对/绝对 OUT_DIR 输入，并提供 `scripts/export-aosp-pack.sh` /
  `make export-pack` 服务器入口，自动执行 export、validate、inspect。
- 服务器入口新增 `pack preflight`，导出前报告 module-info、Java/CC deps、compdb、
  srcjar 和 jar 可用性，并在 compdb 缺失时提示 `SOONG_GEN_COMPDB=1 m nothing`。
- JetBrains 已有 `AOSP Workspace` Tool Window、startup auto-apply 和 Settings 保存后
  自动 rebuild；窗口显示 source/generated/library/skipped/diagnostic counts，完整
  persistent Workspace Model module closure 仍未闭合。
- JetBrains profile state 已可观察；Tool Window 路径标签随 Settings 更新，coordinator
  以 300ms debounce 自动重建，所有 Swing 标签更新切回 EDT。
- `core/pack` 已对 generated manifest 与 srcjar entry 做内容一致性校验，拒绝 entry
  漂移和 zip-slip；不要求部分 Workspace View 的源码实体位于 pack 内。
- 当前 JNI 只覆盖 Java/Kotlin declaration、静态导出、常见 `JNINativeMethod` 注册表和有限注册调用；复杂宏、运行时 class 解析和完整 C++ 语义仍未覆盖；
- 当前已有本地六平台 engine/JetBrains ZIP/VS Code VSIX 打包入口，以及 VS Code 显式/bundled/PATH engine 选择逻辑；尚未完成把 bundled engine 嵌入 VSIX、签名、公证、发布版下载和 marketplace verifier；
- 当前没有任何远程执行服务；
- 当前不声称 Android Studio/IDEA/VS Code 已支持；
- 当前不声称真实 Android Studio/VS Code 已可跳转，也不声称 AIDL compiler 的 stable API/hash 或全部厂商 backend 已兼容。
- `npm run host-test` 已具备真实 Extension Host fixture；本次运行进入 `@vscode/test-electron` 后因下载 VS Code runtime 请求超时，未完成外部 runtime 验证。

## 5. Evidence

- 设计入口：`docs/README.md`；
- 规范化设计：`docs/architecture/`、`docs/design/`、`docs/decisions/`；
- Phase 0 计划：`docs/roadmap/delivery-plan.md`；
- 代码入口：`cmd/aosp-index`、`core/model`、`core/topology`、`core/pathmap`、`core/metadata`、`core/graph`、`core/generated`、`core/aidl`、`core/jni`、`core/index`；
- JNI fixture 通过 `make check`；只读 `frameworks/base/core/java` 与 `core/jni` 探测得到 1126 个 Java/Kotlin native 符号、2686 个 native 符号、988 个注册项、673 条链接，其中 561 条 exact、112 条 inferred。该结果不证明真实 IDE、Windows 或完整 AOSP 兼容。
- JetBrains bootstrap、EngineProcess、profile state 和 focused adapter 已使用 JDK 21 与本机 IntelliJ IDEA 2025.2.4 SDK 通过 `test` 与 `buildPlugin`，包含真实 `BasePlatformTestCase` fixture；生成 `aosp-ide-jetbrains-0.1.0-dev.zip`，且 zip 不携带 Kotlin stdlib。251 baseline 的下载式构建和 Android Studio verifier 尚未闭合。
