# AOSP IDE 当前状态

> 复核日期：2026-08-23。

## 1. Current

- `dev-ide` 分支已切换为新方案仓库边界；旧 Gradle/CMake 实现在 `aosp` 分支。
- 产品可行性、详细设计、拓扑和关键决策文档已存在。
- Phase 0 已有 Go model、topology/pathmap、metadata readers、ModuleGraph、AIDL source index 和 CLI；JetBrains plugin 或 VS Code extension 尚未实现。
- AOSP 源码研究已覆盖 Soong IdeInfo、Java/CC metadata、compdb、AIDL 和 JNI 模式。

## 2. Partial

- 根目录、Makefile 和 Phase 0 工程入口已完成；IDE 和远程执行入口尚未实现。
- WorkspaceTopology、PathMapper、SourceMount 和四类 metadata reader 已有初始实现，但真实 AOSP schema 覆盖、路径映射接线和完整 topology 发现仍未闭合。
- AIDL source parser/index、owning module、srcjar entry 和 Java/C++/NDK build-backed peers 已接入 ModuleGraph；JNI 和 IDE adapters 尚未闭合。

## 3. Target

- Phase 0：Go core、profile、topology、path mapping、doctor/scan CLI 和 fixture tests。
- Phase 1：JetBrains MVP；随后 AIDL、C/C++/JNI 和 VS Code。

## 4. Negative capability

- 当前不能打开 AOSP 并自动提供导航；
- 当前没有发布版跨平台 engine binary；
- 当前没有任何远程执行服务；
- 当前不声称 Android Studio/IDEA/VS Code 已支持；
- 当前不声称真实 Android Studio/VS Code 已可跳转，也不声称 AIDL compiler 的 stable API/hash 或全部厂商 backend 已兼容。

## 5. Evidence

- 设计入口：`docs/aosp-ide-plugin-design.md`；
- 规范化设计：`docs/design/`；
- Phase 0 计划：`docs/roadmap/delivery-plan.md`；
- 代码入口：`cmd/aosp-index`、`core/model`、`core/topology`、`core/pathmap`、`core/metadata`、`core/graph`、`core/generated`、`core/aidl`、`core/index`；
- 当前提交通过 `make check` 和 `go test -race ./core/...`；真实 Framework 探测为 Partial，不证明真实 IDE、Windows 或完整 AOSP 兼容。
