# IDE 客户端与跨语言导航设计

> 状态：Partial；JetBrains 工程与 RPC client bootstrap 进入交付，Workspace Model 和语言适配仍为 Target。

## 1. 客户端边界

客户端只负责：

- profile UI 和状态投影；
- engine lifecycle/RPC；
- IDE source root/project view；
- Java/Kotlin 原生语言能力接线；
- VS Code Java projection 只写入经 RepoPath 证明的 workspace source roots；构建机绝对路径
  不直接进入本机 `java.project` 配置，pack 中的本机可解析 libraries 才作为 referenced libraries。
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
- plugin id：`com.rtfsc.aosp`；
- Kotlin package：`com.rtfsc.aosp`。

251 build line 用于同时覆盖同平台线的 IntelliJ IDEA 与 Android Studio。是否兼容具体 Android
Studio patch 版本必须由 plugin verifier 证明，不能仅凭 build number 声称支持。

### 2.2 代码边界

```text
clients/jetbrains/
  build.gradle.kts
  settings.gradle.kts
  gradle.properties
  src/main/kotlin/com/rtfsc/aosp/
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

`EngineProcess`、project-local profile state、profile UI 和 focused root ModuleRootModel adapter
已进入实现；Workspace Model、AIDL/JNI actions 仍在后续 slice。首版不把本机绝对路径写入仓库，
也不内置尚未构建的 engine binary。

### 2.3 Engine process contract

Engine process 的默认启动方式：

```text
configured executable path -> ProcessBuilder([executable, "serve", "--stdio"])
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
focusedSourcePaths
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

项目设置页提供六个字段：profile id、Workspace View Root、Focused Source Paths、Canonical
AOSP Root、Out Root 和 engine executable。Focused Source Paths 使用 Workspace View Root
下的相对路径（逗号分隔），例如 `frameworks/base,system/core,ui`；空值表示只投影
Workspace View Root。Canonical/Out 可以为空；engine executable 是验证和后续启动的唯一
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

### 2.6 Focused source root projection

投影由纯 Kotlin planner 先生成 `WorkspaceRootPlan`，再由 JetBrains adapter 应用：

```text
EngineProfileState
  -> WorkspaceRootPlan
       sourceRoots: explicit focused paths or workspace root
       excludedRoots: .git, .repo, out, .aosp-ide, canonical root when nested
       diagnostics: invalid/escaping/missing paths
  -> ModuleRootModel facade
       non-persistent module + content roots + source folders + excluded folders
  -> IntelliJ Workspace Model legacy bridge persists the projection on current platforms
```

planner 只规范化用户明确填写的相对路径和固定排除目录，不递归列举 Workspace View，
不 follow symlink，不把 Canonical AOSP Root 注册为 content root。投影模块使用
公开 `ModuleRootModel` facade 创建非持久模块，当前 IntelliJ 平台由 Workspace Model legacy
bridge 写入对应实体；已有 Android/IDEA 模块不被重写。

Apply Focused Roots 是显式操作，失败只保留诊断，不修改 profile state。公开 facade 和平台
Workspace Model bridge 使用同一个 `WorkspaceRootPlan`，不会重复扫描或复制 AOSP 事实。
平台 fixture 使用 `BasePlatformTestCase` 验证 content entry、source folder 和固定排除目录。

插件不打包 Kotlin stdlib，运行时使用目标 IDE 自带 Kotlin runtime，避免覆盖平台 API。

## 3. VS Code

### Target

- TypeScript + vscode-languageclient；
- AIDL TextMate grammar + LSP Definition/References/Symbols；
- clangd 消费 path-mapped compdb overlay；
- JNI 使用 command/Definition provider；
- Java/Kotlin 优先使用用户已有 language server；没有 language server 时，VS Code 提供
  focused Workspace View 内的 source-only 类型跳转 fallback，不声称完整语义分析。

### 3.1 首个 VS Code 实现切片

`clients/vscode/` 使用 TypeScript，目录固定为：

```text
src/extension.ts          activation、commands 和 DefinitionProvider
src/engine/frame.ts       Content-Length framing
src/engine/client.ts      aosp-index 生命周期和 typed requests
src/navigation/paths.ts   Workspace View 与 RepoPath 映射
src/navigation/aidl.ts    AIDL definition 请求和目标打开
src/navigation/jni.ts     JNI peer command 和目标打开
src/ui/profile_panel.ts   Workspace View/profile 图形配置页
```

extension 不实现 AIDL/JNI parser。AIDL DefinitionProvider 和 JNI command 都只发送
当前 `workspaceViewRoot + currentPath (+ name)`，由统一 engine 扫描 view 并返回
`RepoPath + range`。engine 进程由 workspace extension host 管理，退出或配置改变时
销毁，不把绝对路径写入共享协议之外的 fixture。engine client 启动失败必须清理
半启动进程；异常退出后下一次命令允许重新启动。

VS Code 路径 adapter 必须先验证 engine 返回的 RepoPath：拒绝绝对路径、Windows
盘符、UNC、空路径、`.`、`..` 和任何 escaping segment，再按当前平台路径 API 映射
到 Workspace View。测试同时使用 `path.posix` 与 `path.win32`，不能只在开发机平台
上验证斜杠。固定排除目录不能成为导航输入或目标。

VS Code profile 配置使用 Webview panel，不要求用户编辑 Settings JSON。面板提供
Workspace View、metadata/out、engine executable 和 path mappings 字段，并提供目录
选择、保存、验证、刷新 compdb、打开状态等按钮。保存写入 workspace-scoped
`aospIde.*` 配置；JSON 仍作为兼容底层存储，不作为产品主入口。
扩展在打开 workspace 且尚未配置 `workspaceViewRoot` 时自动打开该面板，并把当前
workspace 文件夹预填为 Workspace View Root；已配置的 workspace 不重复弹出。

## 4. AIDL navigation

engine 返回 `RepoPath + range + origin + confidence`。有 out 时依据 AIDL metadata/generated path；无 out 时依据 package/import/type 规则并标记 `inferred`。

### 4.1 JetBrains 首个实现切片

首个 JetBrains 切片先注册 `.aidl` `LanguageFileType`，并使用
`GotoDeclarationHandler` 调用项目内 source-only resolver。它不把 AIDL compiler
或完整 PSI parser 作为插件启动依赖：handler 在当前文件读取 package、import 和
声明名，再通过 `FilenameIndex` 找候选文件，最后返回目标 AIDL 文件。由于首版
使用 Plain Text PSI，目标是文件级导航；自定义 AIDL parser/PSI 接入后再缩小到
声明 token，并保持同一 resolver 和候选排序。

这个边界先保证无 `out`、软链接 view 和 focused source roots 都能工作；候选按
完整 import、当前 package、同名声明排序，多候选时不猜测。后续再把同一 resolver
接到 AIDL PSI/reference、Find Usages 和 engine build-backed generated peers，
不改变 `RepoPath` 与 confidence 语义。

`aosp/aidl/resolve` 已启用，但 JetBrains 的同步 `GotoDeclarationHandler` 不能阻塞 EDT
等待进程启动或 Workspace View 扫描，因此 handler 保留本地文件级 fallback。精确
engine range 由异步 `Go to AIDL Definition` action 承载；后台任务完成后再回到 UI
线程打开目标。Plain Text PSI 仍只能返回文件元素，完整 declaration PSI 留到 parser
definition。

`aosp/aidl/peers` 负责 AIDL 到 generated Java/C++/NDK 的异步跳转。首版 UI 显示并打开
唯一可见 peer；`inferred` 候选如果当前 view 中不存在则只显示诊断，不自动创建文件。

`aosp/aidl/references` 首先接入 VS Code `ReferenceProvider` 和 JetBrains 异步
`Find AIDL Usages` action。首版 range 是所属 import/member 声明范围，客户端显示
这些 source locations；不会声称已经具备 token-level PSI。

VS Code JNI 首个 DefinitionProvider 只消费 `aosp/jni/peers` 的唯一 link。Java/Kotlin
或 C/C++ 文件位于 link 的一侧时，provider 返回另一侧的 `RepoPath + byte range`；
多候选、`inferred`/`unresolved` 或 Workspace View 外目标返回空，让 Java/Kotlin
language server 和 clangd 继续拥有本语言定义权威。

VS Code 提供 `AOSP: Refresh clangd Compilation Database` command。command 调用
`aosp/compdb/overlay`，把返回 entries 写到 Workspace View 的
`.aosp-ide/compdb-overlay/compile_commands.json`，并写入同目录 `.clangd` 的
`CompilationDatabase: .`，不覆盖 AOSP out 或用户根目录 `.clangd`。用户可将 clangd
工作目录指向该 overlay；command 通过 `AOSP IDE` OutputChannel 记录 entry 数、overlay
路径和每条 diagnostic，便于判断 compdb 是否只覆盖了部分 variant。
`.clangd` 内容由独立纯函数生成并固定为 overlay-local `CompilationDatabase: .`，
不会引用 workspace 外路径或覆盖用户配置。
response-file 输出只允许 basename，客户端拒绝 `..`、盘符和 UNC 目标，确保 engine
返回的 overlay artifact 不能写出 `.aosp-ide/compdb-overlay`。

compdb 参数映射覆盖拼接形式和分离形式（例如 `-I/path`、`-I /path`、
`--sysroot=/path`、`--sysroot /path`），并处理 `@response-file` 的本机路径映射。
同一字段有多个 mapping 时按最长 external prefix 选择；overlay 不改写原始 response
file，无法读取或映射时保留原参数并返回稳定 diagnostic。

Extension Host 验收使用 `@vscode/test-electron` 的显式 `host-test` 脚本，验证真实
extension activation、AIDL language registration、definition/reference provider 和
commands。默认 `make test-vscode` 不下载 VS Code runtime；离线 CI 仍由 TypeScript/
framing/navigation fixtures 提供确定性 gate。

VS Code 状态栏只消费 `EngineClient` 的显式状态：`stopped`、`starting`、`running`、
`error`。状态栏点击执行 `AOSP: Show Status`，输出当前 engine executable、Workspace
View、metadata root 和最近错误；UI 不通过子进程对象猜测状态。

## 5. JNI navigation

支持：

- Java `native`/Kotlin `external`；
- `Java_pkg_Class_method`；
- `JNINativeMethod`；
- `jniRegisterNativeMethods`、`RegisterMethodsOrDie`、`RegisterNatives`；
- JNI descriptor。

结果等级：`exact`、`inferred`、`unresolved`。复杂宏、运行时反射和动态加载不升级为 exact。

### 5.1 JetBrains 首个实现切片

JetBrains 提供异步 `Go to JNI Peer` action。action 只发送 profile 中的
`workspaceViewRoot` 和当前 `RepoPath`，通过 `aosp/jni/peers` 让共享 engine 扫描
focused view，再用返回的 `RepoPath + byte range` 选择当前文件对应的唯一 peer 并
打开目标。没有 engine、没有唯一 peer 或只有 `unresolved` 时显示诊断，不用 Kotlin
正则猜测目标。显式 source map 仅保留给远程桥接和协议 fixture。

engine 在同一 daemon 会话内缓存 Workspace View 的 JNI snapshot。JetBrains action
不负责缓存失效，也不根据 canonical real path 合并 symlink alias；所有目标仍以
用户打开的 view-relative RepoPath 返回。

所有 engine 启动、Workspace View 扫描和 RPC 等待必须在 background task 中执行；
只有 warning dialog 和 editor navigation 回到 UI 线程。客户端不能在 EDT 上等待
daemon 或读取大目录。

首版不依赖 Java 或 C++ PSI 的私有实现，也不要求 `out`。后续可在 Java native、
Kotlin external、`Java_` export 和 `JNINativeMethod` 上增加 gutter，但 gutter 与
VS Code command 必须复用同一 RPC 结果。

## 6. Acceptance

- JetBrains bootstrap 可执行 `test`、`buildPlugin` 和 plugin descriptor 校验；
- Kotlin RPC framing 与 Go daemon framing fixture 语义一致；
- JetBrains fixture 能从 profile 创建 focused view；
- AIDL 文件和 Java generated peer 可互跳；
- JNI 常见注册模式双向可跳；
- VS Code LSP 结果与 engine CLI 返回的 RepoPath 一致；
- 客户端重启 engine 后可以重新 initialize。
