# 源码模块发现与工作区隔离设计

> 状态：Decision；模块选择和 IDE 激活必须以工作区为边界。

## 1. 用户流程

插件打开源码工作区后，先判断是否存在 Android/Soong 标记，再解析可见范围内的
`Android.bp`，列出可用模块并推荐常用模块。用户勾选模块后点击 Apply，客户端内部
生成 TargetSpec 和 ProjectGraph；主界面不要求用户手写 target name。

普通工程即使安装插件，也不能自动启动 engine、扫描全树或显示 AOSP 配置错误。

## 2. 数据来源

```text
Android.bp / Android.mk -> Source Module Catalog
Out Root Build Facts    -> 构建能力和产物增强
两者合并               -> Unified Module Catalog
用户勾选               -> TargetSpec / ProjectGraph
```

### Android.bp

`Android.bp` 是模块发现和无 out fallback 的首要来源。跨平台 Go parser 读取 package、
namespace、module name/type、`srcs`、`path`、`exclude_srcs`、`static_libs`、`libs`、
`required`、`defaults`、generated sources/headers、AIDL、include dirs、资源和 manifest。
解析器不执行 Soong，不依赖 Linux；未知属性只产生诊断，单文件失败不阻塞其他模块。

### Out Root

Out Root 可以是完整 AOSP `out` 或保持 `OUT_DIR` 相对路径的 Portable Out：

| 文件 | 用途 |
|---|---|
| `target/product/<product>/module-info.json` | 构建模块、依赖、installed artifact |
| `soong/module_bp_java_deps.json` | Java 源码、jar、srcjar、Java 依赖 |
| `soong/module_bp_cc_deps.json` | C/C++ 源码、include 和 flags |
| `soong/development/ide/compdb/compile_commands.json` | 单文件真实 native 编译命令 |

Out 不负责唯一的模块发现，而是把源码模块标记为 `build-backed`、补充 generated/jar/
srcjar/APK/CC facts，或提供没有源码的 `build-only` 模块。没有 Out 的模块为 `source-only`。
同名模块必须结合 namespace、Android.bp 路径和 source path 消解，不能按 jar basename 猜 owner。

## 3. Source Module Catalog

建议目录：

```text
core/sourcecatalog/
├── discovery.go
├── android_bp.go
├── android_mk.go
├── module.go
├── namespace.go
├── merge.go
├── recommend.go
└── cache.go
```

每个模块至少包含 name、type、Android.bp 路径、source roots、dependencies、状态、能力
（Java/native/generated/AIDL/JNI/artifact）和 artifact 精确相对路径。扫描遵守 Workspace
View 与 SourceMount 边界，支持软链接、挂载和只读目录，排除 `.git`、`.repo`、out 内部
源码扫描和 IDE cache。

推荐结果分为 presets 和 individual modules。默认 `Platform` scope 包含 Java Foundation、
System Server & WMS、System UI、Native Foundation、Graphics & SurfaceFlinger 和
Input Stack；Automotive 始终显示 Recommended，并按 product 决定是否默认启用。完整列表、
条件和去重规则见 `module-selection-planner-design.md`。

## 4. VS Code 工作区隔离

项目选择必须保存到当前 workspace scope：

```text
workspace: enabled, workspaceViewRoot, outRoot, selectedModules, pathMappings, sourceMounts
global:   engineExecutable  # 仅机器级默认值
```

宿主使用 `ConfigurationTarget.Workspace`。不使用 globalState 保存工程选择；untitled
workspace 使用 workspaceState。这样 `~/code/aosp` 和 `~/code/flyme` 的 profile 完全独立。

激活状态：

```text
DORMANT -> CANDIDATE -> ACTIVE
                   \-> DISABLED
ACTIVE + no out -> SOURCE_ONLY
ACTIVE + out    -> BUILD_BACKED / PARTIAL
```

自动候选只检查根目录或 SourceMount 中的 `Android.bp`、`build/envsetup.sh`、`build/soong`
或 `soong/soong_ui.bash`。不能使用 `workspaceContains:**/*` 这类匹配所有工程的激活条件。
显式执行 Configure 命令仍可启用非标准布局。

侧边栏主界面显示发现数量、Out 状态、推荐模块和模块列表；列表支持按仓库、类型、状态
和能力过滤。Apply 后由 engine 根据选中模块构造 TargetSpec/ProjectGraph。

### 4.5 配置与缓存分离

发现结果不是用户配置，不能把完整模块列表、解析后的 Android.bp AST 或诊断复制到
`.vscode/settings.json`。两者职责如下：

| 内容 | VS Code 存放位置 | 是否跨窗口共享 |
|---|---|---|
| enabled、Out Root、选中模块、path mapping | Workspace Configuration | 否 |
| 当前索引状态和进度 | 内存 + workspaceState | 否 |
| Source Module Catalog、解析诊断、fingerprint | `ExtensionContext.storageUri` 下的工作区缓存 | 否 |
| srcjar 解压、compdb overlay | 同一工作区缓存的子目录 | 否 |
| engineExecutable 默认值 | User Configuration | 是 |

有 workspace 时优先使用 `storageUri`，不能使用未分区的 `globalStorageUri` 保存模块
catalog。没有 workspace 的临时窗口才允许使用带 workspace identity 的 fallback 目录。

缓存键至少包含：workspace URI、可见 SourceMount 集合、parser schema version、每个
`Android.bp`/`Android.mk` 的规范化路径和 fingerprint，以及 Out Root 标准 Build Facts
文件的 fingerprint。启动时只重新枚举候选文件并比较 fingerprint；未变化的文件复用已
解析结果，变化或新增文件才重新解析。软链接目标、挂载变化和显式 Refresh 会使对应
目录失效，但不会把整个用户配置清空。

选中模块是用户意图，必须持久化；已发现模块列表是派生数据，只写缓存。缓存损坏时删掉
当前 workspace 的派生缓存并重新扫描，不修改 `settings.json`。

## 5. 多项目和 JetBrains

VS Code 每个窗口拥有独立 profile、engine session 和索引状态；multi-root 首版以整个
workspace 为 profile，但为每个 folder 记录 repo prefix，混合 release 必须提示 warning。
JetBrains 以 Project 为边界独立发现 Android.bp、配置 Out Root 和选择模块；Project 内的
多个目录只能通过显式 SourceMount 合并。

### 5.1 JetBrains 存储边界

JetBrains 的用户配置继续由 Project 级 `PersistentStateComponent` 保存到 workspace storage
（通常体现为 `.idea/workspace.xml` 或对应 Project workspace 文件），包括 enabled、Out Root、
选中模块和 path mapping。它不能保存完整 Source Module Catalog。

Source Module Catalog、Build Facts merge、srcjar 解压和 compdb overlay 应写入 IntelliJ
系统缓存目录，例如：

```text
<IDE system path>/aosp-workspace/<project location hash>/
├── source-catalog/
├── build-facts/
├── generated/
└── compdb/
```

缓存目录由 Project service 计算并按 project location hash 隔离。不同 Project、不同 AOSP/
Flyme 目录和不同 Out Root 不共享派生索引；项目关闭或插件卸载不会向源码树写入索引文件。
当前实现中仍有少量 `.aosp-ide` 项目目录缓存，迁移完成前不能把它视为最终合同。

### 5.2 启动扫描策略

首次打开工作区或缓存失效时执行一次完整的可见范围发现。后续启动先读取 fingerprint
manifest，只增量解析新增或变化的 Android.bp/Android.mk；用户点击 Refresh、SourceMount
变化、Out Root 变化或 parser schema 升级时重新计算对应缓存。扫描不会把结果写入用户
settings，也不会因为普通工程启动插件而运行。

## 6. Out 自动发现和降级

发现顺序为已保存 Out Root、`<workspace>/out`、可见 `out/soong`、用户 Browse 的外部目录，
否则保持 source-only。缺少 module-info 或 Java deps 不阻塞源码模块；缺少 CC deps/compdb
只降低 native 能力；普通非 Android 工程保持 DORMANT。

## 7. 验收

- 打开 AOSP 后自动列出 Android.bp 模块，不要求输入 target name；
- 无 out 可选择 source-only，有 out 显示 build-backed/generated/artifact 能力；
- AOSP、Flyme、QCOM 工作区配置互不覆盖；
- 普通工程安装插件但不启用时不启动 engine、不扫描、不弹错误；
- 软链接、挂载和自定义 Workspace View 遵守可见范围；
- VS Code 与 JetBrains 使用同一模块发现、推荐和 ProjectGraph 语义。
