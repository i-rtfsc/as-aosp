# IDE 客户端与跨语言导航设计

> 状态：Partial；JetBrains 工程与 RPC client bootstrap 进入交付，Workspace Model 和语言适配仍为 Target。

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

### 2.1 首版平台 baseline

- IntelliJ Platform Gradle Plugin：`2.18.1`；
- compile/test platform：IntelliJ IDEA Community `2025.1.7.2`；
- compatible build line：`sinceBuild = 251`，`untilBuild = 251.*`；
- Kotlin/JVM toolchain：21；
- plugin id：`io.github.irtfsc.aospide`；
- Kotlin package：`io.github.irtfsc.aospide`。

251 build line 用于同时覆盖同平台线的 IntelliJ IDEA 与 Android Studio。是否兼容具体 Android
Studio patch 版本必须由 plugin verifier 证明，不能仅凭 build number 声称支持。

### 2.2 代码边界

```text
clients/jetbrains/
  build.gradle.kts
  settings.gradle.kts
  gradle.properties
  src/main/kotlin/io/github/irtfsc/aospide/
    engine/
      JsonRpcFraming.kt
      EngineRpcClient.kt
      EngineProcess.kt
    model/
    settings/
    ui/
    project/
    aidl/
    jni/
  src/main/resources/META-INF/plugin.xml
  src/test/kotlin/
```

首个 bootstrap slice 只交付：

- 可构建的插件包和 `plugin.xml`；
- Kotlin `Content-Length` framing；
- initialize/shutdown RPC client；
- 不启动外部进程的 framing/client unit tests。

`EngineProcess`、profile UI、Workspace Model、AIDL/JNI actions 在后续 slice 接入。首版不把本机
绝对路径写入仓库，也不内置尚未构建的 engine binary。

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

- JetBrains bootstrap 可执行 `test`、`buildPlugin` 和 plugin descriptor 校验；
- Kotlin RPC framing 与 Go daemon framing fixture 语义一致；
- JetBrains fixture 能从 profile 创建 focused view；
- AIDL 文件和 Java generated peer 可互跳；
- JNI 常见注册模式双向可跳；
- VS Code LSP 结果与 engine CLI 返回的 RepoPath 一致；
- 客户端重启 engine 后可以重新 initialize。
