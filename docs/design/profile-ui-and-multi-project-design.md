# Profile UI 与多项目体验设计

> 状态：Target；Neon Workspace 首个 VS Code panel 已进入实现，JetBrains Tool Window
> 与统一视觉主题随后接入。

## 1. 产品原则

- Settings JSON 是兼容存储，不是主要用户入口；
- 首次打开未配置 Workspace 时自动提示一次，关闭后不重复打扰；
- 永远提供稳定入口：VS Code Activity Bar、状态栏和 Command Palette；JetBrains
  Tool Window、Settings 和项目菜单；
- 色彩鲜艳、青春、现代，但信息密度和错误状态必须适合开发工具；
- 自动发现 Workspace Root 和 `out`，用户可覆盖为外部挂载路径或 source-only；
- 只配置一个 Out Root；完整 out 与稀疏 Portable Out 使用完全相同的目录合同；
- 提供从 Android.bp 发现的 Modules 列表和 artifact/source 能力；不要求用户手写 target name；
- 用户勾选模块后由 engine 生成内部 Targets；Settings JSON/Targets JSON 不是用户入口；
- profile 是项目级状态，不把 AOSP/MIUI/QCOM 不同源码树静默合并。
- 模块发现结果是派生缓存，不写入 `.vscode/settings.json` 或 JetBrains workspace 配置；
  配置只保存 enabled、Out Root、模块选择和路径映射。

## 2. Neon Workspace 视觉

深色背景上使用紫色→电蓝→青绿色渐变 Hero，橙色表示 warning，红色表示 error。组件使用
大圆角卡片、状态 pill、清晰的主按钮和次按钮。页面结构：

```text
┌──────────────────────────────────────────────┐
│ AOSP Workspace                  ● Ready      │
│ Focused source navigation workspace          │
├──────────────────────────────────────────────┤
│ Workspace View card      Out Detected card   │
│ Engine status card       Capabilities card   │
│                                              │
│ [Save & Verify] [Refresh clangd] [Show Log]  │
└──────────────────────────────────────────────┘
```

状态模式：`Source-only`、`Build-backed`、`Starting`、`Ready`、`Error`、`Stale`。

## 3. VS Code

注册 `viewsContainers.activitybar` 的 AOSP Workspace 图标和唯一的 `Workspace Setup`
Webview sidebar view（使用新 id，避免旧版本 view 状态残留）。
主视图直接显示可点击的 profile 表单、Workspace View、out/metadata、engine 和 capability
状态；不使用文章式 TreeView。第二个 Profile view 是同一配置页的稳定别名。Command
Palette 和状态栏继续作为备用入口。

Modules 区域显示 Android.bp 发现的模块，每行包含名称、kind、Android.bp 路径、状态、
Java/native/generated 能力和 artifact 摘要。`framework-minus-apex`、`services`、`SystemUI`
等匹配推荐规则的模块默认勾选；用户可以搜索、过滤、全选或取消。点击 Apply 后由 engine
根据选中模块生成内部 TargetSpec；高级入口才允许手动配置 artifact/source selector。

编辑器同时提供 `Resolve Target` 预览。该操作不会保存配置，也不会启动完整索引；它将当前
草稿直接提交给 `aosp/projects/graph`，显示 exact、unresolved 或 error，以及已解析的
modules、artifacts、source roots、generated roots、libraries 和 diagnostics 数量。这样用户
可以在保存或启动索引前确认 `framework.jar`、`services.jar`、`SystemUI.apk` 或自定义
source/module 组合是否能被当前 Out Root 解析。

配置 Out Root 后，artifact selector 通过 `aosp/artifacts/catalog` 加载真实 jar、
srcjar、apk、owner、path、size 和 digest；用户可选 catalog artifact，也可输入自定义
selector。APK 只作为 Android Target anchor，不会被加入 Java classpath。

主 view 声明 `type: webview` 和 `onView:<view-id>` activation event，确保用户点击 Activity Bar 时先激活
扩展并注册 WebviewViewProvider，不依赖某个语言文件或其他 command 先触发 activation。

打开 workspace 后：

1. 自动使用当前 workspace folder 作为候选 Workspace View Root；
2. 检查 `<root>/out`、`<root>/out/soong` 的 metadata/compdb；
3. 没有已保存 profile 时自动打开一次 panel；
4. 已保存 profile 不重复弹窗；
5. 保存后写 workspace-scoped `aospIde.*`，并触发 engine session 重启。
6. 已保存 Workspace View 的窗口激活后自动启动 coordinator；不要求用户点击
   `Start Index`。profile 相关配置变化合并为一次 cancel/restart，手动按钮仍用于
   显式 rebuild 和故障恢复。
7. 侧边栏 provider 保存最后一个 coordinator snapshot，view 关闭并重新打开后立即
   恢复 state、stage、progress、module/generated/diagnostic counts。
8. 首次打开或 fingerprint 变化时扫描 Android.bp/Android.mk；未变化文件从工作区缓存
   复用，不在每次启动时重新解析全部源码。

## 4. JetBrains

注册 `AOSP Workspace` Tool Window，提供与 VS Code 相同的 Overview/Profile/Sources/Build & Out/
Diagnostics 卡片。Settings Configurable 保留为高级编辑入口，默认项目操作从 Tool Window
完成。打开项目时以项目根为 Workspace View 候选并自动检测 `out`；用户可以通过 Browse
覆盖为外部 out 或选择 Source-only。

Tool Window 直接嵌入 Modules 图形列表，支持加载 Android.bp 发现结果、显示 Out 能力、
勾选推荐模块并 Apply。Settings 与 Tool Window 复用同一模块选择和 TargetSpec 生成逻辑，
不允许复制一份字段映射或 JSON 拼装逻辑。

JetBrains startup activity 在已有 profile 时自动执行 Build-Facts projection；Tool Window
提供状态、模块/generated 计数、diagnostics 和 `Apply/Rebuild` 操作。状态由项目级
coordinator service 保存，窗口关闭后重新打开仍恢复最后快照。

模块列表、解析 AST、Build Facts 合并结果、generated 和 compdb overlay 写入 IDE system
cache，并以 Project location hash 隔离；workspace.xml 只保存用户选择和路径配置。

JetBrains Targets 编辑器与 VS Code 使用相同的 `Resolve Target` 预览语义和 ProjectGraph
RPC；解析在后台任务中执行，结果显示在编辑器内，不阻塞 Swing UI。

## 5. 自动 out 检测

优先级：已保存 Out Root > `<workspace>/out` > `<workspace>/out/soong` > 未配置。检测
字段包括 `module-info.json`、Java/CC metadata 和 `compile_commands.json`。检测只读，不扫描
canonical AOSP 全树，不覆盖用户手动选择。

## 6. 多项目模型

### VS Code

- 多窗口：每个窗口一个 Workspace View/profile，推荐用于 AOSP、MIUI、QCOM；
- multi-root workspace：可以同时打开多个目录，但首版 profile 绑定 workspace 而不是每个
  root，混合不同 release 时显示 warning，不静默合并索引；
- 每个窗口/extension host 自己启动 engine session。

### JetBrains

- 多窗口：每个 Project 一个 profile，推荐用于 AOSP、MIUI、QCOM；
- 单项目多模块：可以把多个只读目录投影到同一 Project，但 engine 仍只允许一个 active
  profile；不同 release 必须分项目，否则 AIDL/JNI/compdb 结果标记 mixed/source-only；
- 每个 Project 自己启动 engine session，不共享 daemon 状态。

## 7. Acceptance

- 关闭配置 panel 后可从 Activity Bar/Tool Window/命令重新打开；
- 新 workspace 首次自动出现配置 panel；
- workspace 有 `out` 时默认显示 detected，不存在时明确显示 Source-only；
- 手动 out 不被自动检测覆盖；
- AOSP、MIUI、QCOM 多窗口 profile 互不污染；
- VS Code 与 JetBrains 的卡片状态和术语一致。
