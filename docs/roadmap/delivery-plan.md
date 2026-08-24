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

### AOSP-IDE-RPC-01

- 目标：由 `aosp-index serve --stdio` 提供版本化 JSON-RPC，作为 JetBrains 与 VS Code 的共享边界。
- 技术前置：CORE-01、CLI-01。
- 下一验收：initialize、profile/topology validation、topology discovery 和 shutdown smoke。
- 负能力：首版不持有长期索引；AIDL/JNI capability 按独立 slice 启用。

### AOSP-IDE-JNI-01

- 目标：Java native、静态导出、JNINativeMethod 和 JNI descriptor。
- 技术前置：CORE-01、METADATA-01。
- 状态：Partial；Go core 与 `aosp-index jni` source-only fixture 已交付。
- 已验收：Java/Kotlin declaration、静态导出、常见注册表、描述符校验、exact/inferred/unresolved 结果和 Java/C++ 双向查询。
- 下一验收：接入 VS Code adapter，并补充 AIDL generated peers、复杂宏、`RegisterNatives` 变量流和真实多平台路径测试。

### AOSP-IDE-JB-01

- 目标：JetBrains 配置页、Workspace Model、Java/generated source 和 AIDL PSI。
- 技术前置：AIDL source baseline 和 CLI baseline 已具备；generated peer 接线可以按能力渐进启用。
- 当前切片：251 baseline、插件工程、Kotlin RPC framing、EngineProcess、profile persistence、profile UI 和 focused root adapters 已完成。
- 下一验收：AIDL declaration PSI、Android Studio/IDEA baseline plugin build、descriptor verification 和真实 IDE root projection。

### AOSP-IDE-VSCODE-01

- 目标：VS Code extension、AIDL LSP、clangd compdb overlay 和 JNI commands。
- 技术前置：AIDL-01、JNI-01、metadata reader。
- 当前切片：TypeScript extension、统一 engine framing、AIDL DefinitionProvider 和 JNI peer command。
- 下一验收：VS Code extension host fixture、AIDL references、generated peers 和 clangd compdb overlay。

## P2：外部环境

### AOSP-IDE-REMOTE-01

- 目标：Linux remote-execution bridge。
- 技术前置：CLI/RPC、路径映射和安全边界。
- 下一验收：受控 Linux fixture 刷新 metadata 并返回 RepoPath。
- 状态：External。
