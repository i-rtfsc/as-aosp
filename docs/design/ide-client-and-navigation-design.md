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

已具备插件 descriptor、251 baseline、Kotlin RPC framing 和
initialize/shutdown client bootstrap。

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

`EngineProcess` 和 project-local profile state 已进入实现；profile UI、Workspace Model、AIDL/JNI
actions 仍在后续 slice。首版不把本机绝对路径写入仓库，也不内置尚未构建的 engine binary。

### 2.3 Engine process contract

Engine process 的默认启动方式：

```text
configured executable path -> ProcessBuilder(argument list) -> aosp-indexd
```

约束：

- 不通过 shell 拼接命令，避免空格、引号和注入问题；
- executable 必须是用户配置或插件 cache 中的明确文件；
- stdout 专用于 JSON-RPC，stderr 由独立 drain thread 读取并写入 IDE log；
- start 后必须先 initialize，initialize 失败则立即 terminate；
- stop 先发送 shutdown，等待有限超时后 destroy，最后才 destroyForcibly；
- EOF、非零退出和 framing 错误都使 session 失效；
- restart 创建新进程和新 RPC client，不复用旧 request id/session；
- Windows 使用 `.exe`，Linux/macOS 使用可执行文件，路径均通过 `java.nio.file.Path` 传递。

首个实现允许注入 process factory 进行单元测试，不要求测试启动真实平台二进制。

### 2.4 Project profile persistence

profile 保存为项目级 IDE workspace state：

```text
profileId
workspaceViewRoot
canonicalRepoRoot
outRoot
engineExecutable
```

这些字段只保存用户本机配置；共享 scope、RepoPath include/exclude 和 mount 规则后续单独
加入 profile editor。加载 state 时缺失字段使用空值，路径不自动 canonicalize，不在持久化
阶段扫描完整 AOSP。启动 engine 前由 RPC `aosp/profile/validate` 检查 profile。

状态服务不启动 engine，也不持有 IDE Project Model；engine lifecycle 由独立 project service
管理，便于测试和后续 VS Code 复用同一 profile schema。

### 2.5 Profile configuration UI

项目设置页提供五个字段：profile id、Workspace View Root、Canonical AOSP Root、Out Root
和 engine executable。Canonical/Out 可以为空；engine executable 是验证和后续启动的唯一
可执行文件，不根据 PATH 猜测。

按钮语义：

- Apply/OK 只保存 workspace state；
- Validate 在 background task 中启动一次临时 engine，调用 `aosp/profile/validate`，显示
  diagnostics 数量后关闭进程；
- Validate 失败不覆盖已保存 state；
- project dispose 或 settings page dispose 不自动启动 engine；
- 路径选择使用 IDE file chooser，保存字符串原样，不跟随软链接或访问远程挂载。

首版 UI 不创建 source root、不读取 Soong metadata、不递归扫描 AOSP，也不把 engine stderr
直接展示为用户诊断；详细诊断后续通过 Tool Window 展示。

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
