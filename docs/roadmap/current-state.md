# AOSP IDE 当前状态

> 复核日期：2026-08-23。

## 1. Current

- `dev-ide` 分支已切换为新方案仓库边界；旧 Gradle/CMake 实现在 `aosp` 分支。
- 产品可行性、详细设计、拓扑和关键决策文档已存在。
- Phase 0 已有 Go model、topology/pathmap、metadata readers、ModuleGraph、AIDL/JNI source index 和 CLI；JetBrains plugin bootstrap 已可构建，VS Code extension 尚未实现。
- AOSP 源码研究已覆盖 Soong IdeInfo、Java/CC metadata、compdb、AIDL 和 JNI 模式。

## 2. Partial

- 根目录、Makefile、Phase 0 CLI、首版 stdio JSON-RPC daemon 和 JetBrains 251 bootstrap 已完成；Workspace Model、语言适配和远程执行入口尚未实现，ModuleRootModel focused adapter 已具备。
- WorkspaceTopology、PathMapper、SourceMount 和四类 metadata reader 已有初始实现，但真实 AOSP schema 覆盖、路径映射接线和完整 topology 发现仍未闭合。
- AIDL source parser/index、owning module、srcjar entry 和 Java/C++/NDK build-backed peers 已接入 ModuleGraph；JNI source-only index 与 profile/topology RPC 已接入 CLI/daemon，IDE adapters 尚未闭合。

## 3. Target

- Phase 0：Go core、profile、topology、path mapping、doctor/scan CLI 和 fixture tests。
- Phase 1：JetBrains MVP；随后 AIDL、C/C++/JNI 和 VS Code。

## 4. Negative capability

- 当前不能打开 AOSP 并自动提供导航；
- 当前 JetBrains 插件已有 descriptor、Kotlin RPC framing/lifecycle client、EngineProcess、project-local profile state、profile UI 和 focused root planner/非持久投影模块；仍没有 AIDL PSI 或 JNI action；
- 当前 daemon 只启用 initialize、shutdown、profile validate、topology discover/validate；索引、AIDL、JNI RPC capability 尚未启用；
- 当前 JNI 只覆盖 Java/Kotlin declaration、静态导出、常见 `JNINativeMethod` 注册表和有限注册调用；复杂宏、运行时 class 解析和完整 C++ 语义仍未覆盖；
- 当前没有发布版跨平台 engine binary；
- 当前没有任何远程执行服务；
- 当前不声称 Android Studio/IDEA/VS Code 已支持；
- 当前不声称真实 Android Studio/VS Code 已可跳转，也不声称 AIDL compiler 的 stable API/hash 或全部厂商 backend 已兼容。

## 5. Evidence

- 设计入口：`docs/aosp-ide-plugin-design.md`；
- 规范化设计：`docs/design/`；
- Phase 0 计划：`docs/roadmap/delivery-plan.md`；
- 代码入口：`cmd/aosp-index`、`core/model`、`core/topology`、`core/pathmap`、`core/metadata`、`core/graph`、`core/generated`、`core/aidl`、`core/index`；
- 代码入口：`cmd/aosp-index`、`core/model`、`core/topology`、`core/pathmap`、`core/metadata`、`core/graph`、`core/generated`、`core/aidl`、`core/jni`、`core/index`；
- JNI fixture 通过 `make check`；只读 `frameworks/base/core/java` 与 `core/jni` 探测得到 1126 个 Java/Kotlin native 符号、2686 个 native 符号、988 个注册项、673 条链接，其中 561 条 exact、112 条 inferred。该结果不证明真实 IDE、Windows 或完整 AOSP 兼容。
- JetBrains bootstrap、EngineProcess、profile state 和 focused adapter 已使用 JDK 21 与本机 IntelliJ IDEA 2025.2.4 SDK 通过 `test` 与 `buildPlugin`，包含真实 `BasePlatformTestCase` fixture；生成 `aosp-ide-jetbrains-0.1.0-dev.zip`，且 zip 不携带 Kotlin stdlib。251 baseline 的下载式构建和 Android Studio verifier 尚未闭合。
