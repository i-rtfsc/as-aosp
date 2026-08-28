# Profile UI、i18n 与多项目体验设计

> 状态：Decision；日常入口采用紧凑工作台和独立模块选择模式。

## 1. 产品原则

- UI 是侧边栏工作台，不是全屏落地页；
- 不使用占据大量纵向空间的 Hero；
- 日常界面只展示工作区、Out、索引状态和已启用 presets；
- 完整模块目录只在 Manage Modules 模式出现；
- Refresh 只刷新模块目录，Apply 才生成 ProjectGraph 并启动 IDE projection；
- Settings JSON、Targets JSON 和原始模块名不是主要编辑入口；
- 专业名称不翻译，例如 `Graphics & SurfaceFlinger`、`Input Stack`、
  `System Server & WMS`、`CarService`、`libbinder`；
- 按钮、状态、说明和错误信息支持中英文；
- 色彩鲜艳、青春、现代，但信息密度必须适合开发工具；
- profile、模块选择、索引状态和 cache 都以 workspace/Project 隔离。

## 2. 信息架构

侧边栏存在两个互斥模式：

```text
Workbench       日常状态、启用范围和索引操作
Manage Modules  presets、搜索、筛选和自定义模块选择
```

模块目录不能常驻 Workbench，也不能一次渲染数百或上万行。

## 3. Workbench

```text
┌──────────────────────────────────┐
│ ▰ AOSP Workspace        ● Ready  │
│                         ↻   ⚙    │
├──────────────────────────────────┤
│ WORKSPACE                        │
│ aosp                             │
│ /workspace/aosp                  │
│                                  │
│ OUT                              │
│ ● Build-backed        Android 17 │
│ 14,803 modules · 5,284 generated │
├──────────────────────────────────┤
│ INDEX                            │
│ ● Discover ━● Graph ━● Sources ━●│
│ Ready                            │
│ 1,842 roots · 1,463 generated    │
│ 3,451 libraries                  │
├──────────────────────────────────┤
│ ENABLED PRESETS              7   │
│ ● Java Foundation               │
│ ● System Server & WMS           │
│ ● Graphics & SurfaceFlinger     │
│ ● Input Stack                   │
│                                  │
│ [ Manage Modules ]               │
├──────────────────────────────────┤
│ [ Apply Changes ]   [ Rebuild ]  │
└──────────────────────────────────┘
```

### 3.1 顶部

- 使用 3px 多色强调线，不使用大面积渐变背景；
- 显示产品名、状态点、Refresh 图标和 Settings 图标；
- Refresh tooltip 明确为 `Refresh module catalog`；
- Settings 进入高级路径、SourceMount 和 engine 设置。

### 3.2 Workspace 与 Out

默认显示目录 basename，完整路径作为次级文本和 tooltip。Out Root 必须指向 `OUT_DIR`，
用户选择 `out/soong` 时自动规范化为父级 `out`。

状态：

```text
Source-only
Build-backed
Partial
Out missing
Out stale
```

### 3.3 Index

必须显示真实阶段，不能只显示一行模糊文本：

```text
Discover -> Catalog -> ProjectGraph -> Generated -> Projection -> IDE Index
```

Refresh 完成不代表可跳转。只有 Apply/Rebuild 完成 `Projection` 后，插件才已把 source roots、
libraries、srcjar 和 Rust project 交给 IDE。Java Language Server 或 IDE 自身索引继续运行时，
状态显示 `Waiting for Java index`，不能提前显示 Ready。

### 3.4 Enabled Presets

Workbench 只显示已启用的 preset 名称和总数，不显示其全部底层 modules。点击一行可查看
module 数、能力和 pending changes；点击 Manage Modules 进入选择模式。

## 4. Manage Modules

```text
┌──────────────────────────────────┐
│ ← Modules                  8 ✓   │
├──────────────────────────────────┤
│ 🔍 Search modules               │
├──────────────────────────────────┤
│ Recommended  Java  Native  Apps  │
├──────────────────────────────────┤
│ PRESETS                          │
│ ☑ Java Foundation        5       │
│ ☑ System Server & WMS    2       │
│ ☑ Graphics & SF          4       │
│ ☑ Input Stack            5       │
│ ☐ Automotive             4       │
├──────────────────────────────────┤
│ MATCHING MODULES                 │
│ ☑ framework-minus-apex           │
│   Java · Build-backed            │
│   frameworks/base                │
│                                  │
│ ☐ CarService                     │
│   Android · Build-backed         │
│   packages/services/Car/...      │
├──────────────────────────────────┤
│ 16,745 matching · showing 50     │
├──────────────────────────────────┤
│ [ Cancel ]      [ Apply 8 ]      │
└──────────────────────────────────┘
```

规则：

- 默认只显示 presets、已选择模块和少量推荐模块；
- 输入至少两个字符后向 engine 查询普通模块；
- 每页最多 50，允许 Load More 到 100，不渲染完整 catalog；
- 支持 Recommended、Java、Android、Native、Rust 和 Source-only 筛选；
- Apply 按钮显示最终 preset/custom 数量和 pending changes；
- 模块名、Soong type 和路径保持原文；
- 选中一个 module 不等于创建一个 IDE module，具体编译规则见
  `module-selection-planner-design.md`。

## 5. 色彩与组件

| 色彩 | 用途 |
|---|---|
| Cyan | source discovery、Workspace |
| Violet | ProjectGraph、selection |
| Pink | generated source、srcjar |
| Green | Ready、Build-backed |
| Orange | Partial、Waiting、Stale |
| Red | Error、invalid path |

要求：

- 使用 VS Code/JetBrains theme foreground/background 作为主体；
- 鲜艳颜色只用于强调线、状态、能力 badge 和 primary action；
- card radius 不超过 8px；
- 不嵌套 cards；
- icon button 使用平台图标或 Lucide 对应语义；
- compact panel 内不使用 hero-scale 字号；
- 文本在窄侧边栏中允许换行，路径使用 ellipsis 和 tooltip。

## 6. i18n

### VS Code

```text
clients/vscode/package.nls.json
clients/vscode/package.nls.zh-cn.json
clients/vscode/src/i18n/messages.ts
```

contribution 文本通过 package NLS，runtime 文本通过 `vscode.l10n`。Webview 只接收已经
本地化的 dictionary，不在 HTML 中硬编码中英文。

### JetBrains

```text
messages/AospWorkspaceBundle.properties
messages/AospWorkspaceBundle_zh_CN.properties
```

使用 IntelliJ ResourceBundle。语言默认跟随 IDE，不单独维护插件语言开关。

| English | 中文 |
|---|---|
| Manage Modules | 管理模块 |
| Apply Changes | 应用更改 |
| Build-backed | 构建信息可用 |
| Source-only | 仅源码 |
| Discovering Modules | 正在发现模块 |
| Building ProjectGraph | 正在构建 ProjectGraph |
| Preparing Generated Sources | 正在准备生成源码 |
| Waiting for Java Index | 等待 Java 索引 |
| Ready | 就绪 |

Preset 名称和模块专业名词不翻译。

## 7. VS Code 行为

1. 普通工程保持 Dormant；
2. AOSP marker 或用户显式打开入口后激活；
3. 自动检测 Workspace View 和 Out Root；
4. Refresh 只更新 Source Module Catalog 和 Build-backed 状态；
5. Apply 保存 selection，编译 Selection Plan，启动 ProjectGraph 和 projection；
6. Rebuild 清理本 workspace 的派生 cache 后重新执行；
7. view 关闭后重新打开恢复最后状态和 pending changes；
8. 所有项目配置写 workspace scope；
9. 生成源码、Build Facts 和 Rust project 写 `storageUri`。

## 8. JetBrains 行为

JetBrains Tool Window 使用相同 Workbench/Manage Modules 信息架构。Project 级
PersistentStateComponent 只保存用户意图，派生数据写 IDE system cache。Settings 保留为
高级入口，不复制一份模块选择算法。

## 9. 多项目

- VS Code 每个窗口一个 profile 和 engine session；
- multi-root 首版一个 profile，但 folder provider 必须保留所属根；
- JetBrains 每个 Project 一个 profile 和 engine session；
- 不同 AOSP/MIUI/QCOM release 不静默合并；
- 普通工程不启动 engine、不扫描、不显示 AOSP 错误。

## 10. Acceptance

- Workbench 首屏不显示完整模块 catalog；
- Manage Modules 不一次渲染超过 100 个普通模块；
- Refresh 和 Apply 的行为、文案和进度明确不同；
- 中文 IDE 显示中文说明，preset/module 专业名保持英文；
- Index 阶段能区分 engine projection 与 IDE language index；
- 侧边栏窄宽度下无文本重叠和横向溢出；
- VS Code 与 JetBrains 使用一致的信息层级、术语和状态语义。
