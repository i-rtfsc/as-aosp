# AOSP IDE 当前状态

> 复核日期：2026-08-23。

## 1. Current

- `dev-ide` 分支已切换为新方案仓库边界；旧 Gradle/CMake 实现在 `aosp` 分支。
- 产品可行性、详细设计、拓扑和关键决策文档已存在。
- 当前没有 Go engine、JetBrains plugin 或 VS Code extension 实现。
- AOSP 源码研究已覆盖 Soong IdeInfo、Java/CC metadata、compdb、AIDL 和 JNI 模式。

## 2. Partial

- 根目录和文档骨架已完成，但工程构建入口尚未实现。
- WorkspaceTopology、PathMapper、SourceMount 已完成设计，未有代码实现。
- AIDL/JNI 目标算法已定义，未有 parser/index 实现。

## 3. Target

- Phase 0：Go core、profile、topology、path mapping、doctor/scan CLI 和 fixture tests。
- Phase 1：JetBrains MVP；随后 AIDL、C/C++/JNI 和 VS Code。

## 4. Negative capability

- 当前不能打开 AOSP 并自动提供导航；
- 当前没有跨平台 engine binary；
- 当前没有任何远程执行服务；
- 当前不声称 Android Studio/IDEA/VS Code 已支持；
- 当前不声称 AIDL/JNI 已可跳转。

## 5. Evidence

- 设计入口：`docs/aosp-ide-plugin-design.md`；
- 规范化设计：`docs/design/`；
- Phase 0 计划：`docs/roadmap/delivery-plan.md`；
- 代码和测试结果以本分支后续工程入口为准。
