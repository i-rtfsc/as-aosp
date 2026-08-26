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
- 独立配置可复制的 Metadata Pack Root，不把 pack 误当成完整 `out`；
- 提供 Targets 列表和 artifact/module/source selector；不提供 Selected Modules 入口；
- Targets 必须使用图形化列表和编辑器；Settings JSON/Targets JSON 不是用户入口；
- profile 是项目级状态，不把 AOSP/MIUI/QCOM 不同源码树静默合并。

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

Targets 区域显示重复项目行，每行包含名称、kind、artifact/module/source 摘要、状态、
编辑和删除操作。`Add Target` 打开内联 editor：Target name、kind、artifact kind/name、
Soong modules、source roots 和 dependencies/generated/resources/manifest 开关。保存后生成
TargetSpec；用户不编辑 JSON。

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

## 4. JetBrains

注册 `AOSP Workspace` Tool Window，提供与 VS Code 相同的 Overview/Profile/Sources/Build & Out/
Diagnostics 卡片。Settings Configurable 保留为高级编辑入口，默认项目操作从 Tool Window
完成。打开项目时以项目根为 Workspace View 候选并自动检测 `out`；用户可以通过 Browse
覆盖为外部 out 或选择 Source-only。

JetBrains startup activity 在已有 profile 时自动执行 Build-Facts projection；Tool Window
提供状态、模块/generated 计数、diagnostics 和 `Apply/Rebuild` 操作。状态由项目级
coordinator service 保存，窗口关闭后重新打开仍恢复最后快照。

## 5. 自动 out 检测

优先级：已保存 metadata/out > `<workspace>/out` > `<workspace>/out/soong` > 未配置。检测
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
