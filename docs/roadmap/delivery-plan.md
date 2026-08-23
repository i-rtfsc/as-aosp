# AOSP IDE 当前交付计划

> 只保留未完成任务。完成项从本表移除，必要证据进入历史记录。

## P0：索引核心前置

### AOSP-IDE-CORE-01

- 目标：建立 Go engine、profile、RepoPath、WorkspaceTopology、SourceMount、PathMapper。
- 技术前置：已接受仓库、Go、目录和路径决策。
- 下一验收：fixture 覆盖普通目录、软链接、别名、overlay、断链、循环、冲突和 Windows 路径。
- 负能力：不读取真实 AOSP 全树，不接 IDE，不执行构建。

### AOSP-IDE-CLI-01

- 目标：提供 `aosp-index doctor/scan` 和仓库 Make 入口。
- 技术前置：CORE-01。
- 下一验收：profile 校验、topology snapshot、结构化 diagnostics 输出。
- 负能力：不启动 IDE，不自动扩大 scope。

## P1：语言和 IDE

### AOSP-IDE-AIDL-GENERATED-01

- 目标：接入 Soong AIDL include roots、srcjar/generated files，以及 Java/C++/NDK generated peer 映射。
- 技术前置：AIDL source parser/index 和 metadata reader 基线已具备。
- 下一验收：同一 AIDL fixture 在 source-only 与 build-backed 模式返回不同来源和置信度，且 generated file 缺失时只降级不失败。
- 负能力：不执行 AIDL compiler，不复刻 stable API/hash 检查。

### AOSP-IDE-JNI-01

- 目标：Java native、静态导出、JNINativeMethod 和 JNI descriptor。
- 技术前置：CORE-01、METADATA-01。
- 下一验收：精确/推断/未解析三类结果 fixture。

### AOSP-IDE-JB-01

- 目标：JetBrains 配置页、Workspace Model、Java/generated source 和 AIDL PSI。
- 技术前置：AIDL source baseline 和 CLI baseline 已具备；generated peer 接线可以按能力渐进启用。
- 下一验收：Android Studio/IDEA baseline plugin verification。

### AOSP-IDE-VSCODE-01

- 目标：VS Code extension、AIDL LSP、clangd compdb overlay 和 JNI commands。
- 技术前置：AIDL-01、JNI-01、metadata reader。
- 下一验收：VS Code extension host fixture。

## P2：外部环境

### AOSP-IDE-REMOTE-01

- 目标：Linux remote-execution bridge。
- 技术前置：CLI/RPC、路径映射和安全边界。
- 下一验收：受控 Linux fixture 刷新 metadata 并返回 RepoPath。
- 状态：External。
