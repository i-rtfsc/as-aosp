# IDE 客户端与跨语言导航设计

> 状态：Target；JetBrains 先于 VS Code 开发。

## 1. 客户端边界

客户端只负责：

- profile UI 和状态投影；
- engine lifecycle/RPC；
- IDE source root/project view；
- Java/Kotlin 原生语言能力接线；
- AIDL PSI 或 LSP；
- clangd 配置；
- JNI 跨语言导航命令。

客户端不负责读取 Soong JSON、选择 Source Provider 或复制一套 index。

## 2. JetBrains

### Current

无插件代码。

### Target

- Kotlin + IntelliJ Platform SDK；
- Workspace Model 创建 focused ProjectView；
- Canonical AOSP Root 不登记为 content root；
- generated source 使用 cache root；
- AIDL FileType、PSI、ReferenceContributor、Find Usages；
- Java/AIDL、Java/JNI 自定义 Definition；
- Tool Window 显示 mode、provider、missing、stale 和 diagnostics。

## 3. VS Code

### Target

- TypeScript + vscode-languageclient；
- AIDL TextMate grammar + LSP Definition/References/Symbols；
- clangd 消费 path-mapped compdb overlay；
- JNI 使用 command/Definition provider；
- Java/Kotlin 依赖用户已有 language server。

## 4. AIDL navigation

engine 返回 `RepoPath + range + origin + confidence`。有 out 时依据 AIDL metadata/generated path；无 out 时依据 package/import/type 规则并标记 `inferred`。

## 5. JNI navigation

支持：

- Java `native`/Kotlin `external`；
- `Java_pkg_Class_method`；
- `JNINativeMethod`；
- `jniRegisterNativeMethods`、`RegisterMethodsOrDie`、`RegisterNatives`；
- JNI descriptor。

结果等级：`exact`、`inferred`、`unresolved`。复杂宏、运行时反射和动态加载不升级为 exact。

## 6. Acceptance

- JetBrains fixture 能从 profile 创建 focused view；
- AIDL 文件和 Java generated peer 可互跳；
- JNI 常见注册模式双向可跳；
- VS Code LSP 结果与 engine CLI 返回的 RepoPath 一致；
- 客户端重启 engine 后可以重新 initialize。
