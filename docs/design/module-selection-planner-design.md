# Module Selection Planner 与去重设计

> 状态：Decision；用户选择不能直接一对一生成 IDE modules。

## 1. 目标

用户可以同时选择多个 presets、手动 modules，甚至选择全部模块。Selection Planner 必须：

- 对相同 module identity 去重；
- 全局只计算一次 dependency closure；
- 同一 source root、generated root、srcjar 和 JAR 只投影一次；
- 为共享源码确定唯一 owner；
- 把依赖关系转换为 Target dependency，而不是复制源码；
- 控制 IDE module 数量，不能因选择 16,000 modules 创建 16,000 IDE modules；
- 保留用户显式选择和高级 pin 意图。

## 2. 输入模型

```text
PresetSelection[]
ManualModuleSelection[]
PinnedModuleSelection[]
SourceRootSelection[]
ArtifactSelection[]
```

选择来源：

```text
explicit-preset
explicit-manual
explicit-pinned
dependency-only
```

Preset、manual 和 dependency 可以同时指向同一个 module，最终只产生一个 canonical entry。

## 3. Module identity

首选 key：

```text
module name + namespace + Android.bp/Android.mk path
```

纯 AOSP 中 module name 通常唯一，但厂商树、namespace 和 overlays 场景不能只按 basename
猜测。Out Facts 与 source catalog 合并时使用 name、namespace、build file path 和 source path
共同消解。

## 4. Presets

专业 preset 名保持英文。

### 4.1 Java Foundation

默认启用：

```text
core-oj
core-libart
ext
framework-minus-apex
framework
```

### 4.2 System Server & WMS

默认启用：

```text
services
services.core
```

WMS 没有独立完整顶层 module；`services -> services.impl -> services.core` 覆盖 AMS、WMS
和核心 system services。Planner 必须去重该依赖链，并允许 `services.core` 作为更具体 owner。

### 4.3 System UI

默认启用：

```text
SystemUI
```

`SystemUI-application`、`SystemUI-core` 和其他子模块通过 closure 自动包含。

### 4.4 Native Foundation

默认启用：

```text
libbase
libutils
liblog
libcutils
libbinder
libbinder_ndk
libnativehelper
```

当前 `module_bp_cc_deps.json` 主要提供源码和编译参数，不提供完整 native dependency graph，
因此这些基础 native modules 必须显式进入 preset，不能完全依赖 closure 推断。

### 4.5 Graphics & SurfaceFlinger

默认启用：

```text
libui
libgui
libnativewindow
surfaceflinger
```

### 4.6 Input Stack

默认启用：

```text
libinput
libinputflinger
libinputdispatcher
libinputreader
libinputreporter
```

当前 Android 17 的核心 InputFlinger module 是 `libinputflinger`，不是 `inputflinger`。

### 4.7 Automotive

始终显示为 Recommended：

```text
android.car
CarService
CarServiceUpdatable
CarSystemUI
```

Automotive product 默认启用；phone product 默认不启用，但保持 Recommended。用户在该
workspace 启用过后必须保存选择，不因 product 检测变化静默关闭。

## 5. Scope presets

```text
Essential
  Java Foundation
  System Server & WMS
  Native Foundation

Platform (default)
  Essential
  System UI
  Graphics & SurfaceFlinger
  Input Stack

Full Platform
  Platform
  Automotive
  WindowManager Shell
  additional system services
```

## 6. 编译算法

```text
1. expand presets
2. union preset and manual selections
3. canonicalize ModuleKey
4. promote explicit selections over dependency-only entries
5. compute one global dependency closure
6. assign modules to bounded Target groups
7. resolve source ownership
8. deduplicate paths, archives and libraries
9. emit ProjectGraph
```

伪代码：

```text
selected = Set(expand(presets) + manual + pinned)
closure = dependencyClosure(selected)

for module in closure:
  target = assignTarget(module, selectionOrigin, preset, kind)
  owner[module] = chooseOwner(target priority, path specificity)

projectGraph = materialize(owner, targetDependencies)
```

## 7. 相同 Java 目录只索引一次

例如：

```text
Framework
  frameworks/base/core/java
  frameworks/base/graphics/java

System Server & WMS
  frameworks/base/core/java
  frameworks/base/services/core
```

最终：

```text
Framework
  owns frameworks/base/core/java
  owns frameworks/base/graphics/java

System Server & WMS
  depends on Framework
  owns frameworks/base/services/core
```

`frameworks/base/core/java` 只交给 IDE 一次。

## 8. 显式选择与依赖重叠

用户同时选择：

```text
services
services.core
framework
```

即使 `services` 依赖 `services.core` 和 `framework`，canonical selection 中仍只有三个
modules。显式选择提升 ownership 优先级，依赖方引用 owner target，不复制 source roots。

## 9. Source ownership

优先级：

```text
pinned target
> more specific source root
> explicit manual selection
> explicit preset selection
> dependency-only shared target
> stable lexical tie-break
```

Nested roots：

```text
frameworks/base
frameworks/base/services
frameworks/base/services/core
```

更具体 target 获得 child root；父 target 生成 `excludedSourceRoots`。WorkspaceTopology 先把
软链接、挂载和 canonical source 解析为一个 active provider，再执行 ownership。

## 10. 资源去重

### Source root

key 为 normalized RepoPath。相同 RepoPath 只投影一次。

### Generated root

key 为 archive out-relative path、archive hash 和 extracted root。相同 srcjar 只解压一次。

### JAR/AAR

key 首选精确 Out 相对路径，不能按 basename 去重。内容 hash 可以作为相同内容的二级合并
证据，但不同路径和 owner 的 artifact 不能仅因 basename 相同被合并。

### Native source

key 为 normalized RepoPath；同一文件只进入一个 clangd scope。variant compile command
冲突必须选择 primary variant 并保留 diagnostics。

## 11. Target 数量控制

Preset 生成稳定 grouped targets：

```text
Java Foundation
System Server & WMS
System UI
Native Foundation
Graphics & SurfaceFlinger
Input Stack
Automotive
```

Manual modules 按 kind 进入 bounded buckets：

```text
Custom Java
Custom Android
Custom Native
Custom Rust
Custom Generated
```

只有 `Pin as separate target` 的高级选择才生成独立 target。

选择 16,000 modules 时仍只能生成有限数量的 IDE targets，不允许一模块一 target。

## 12. Select All

去重不能消除范围本身的成本。Apply 前必须显示影响预估：

```text
16,745 selected modules
126,565 modules in closure
8,420 source roots
5,284 generated archives
1,627 libraries

Large workspace scope
Estimated memory: High
```

超过阈值时要求确认，但允许继续。Select All 仍执行全局 closure、ownership 和资源去重。

## 13. Current Gap

当前 engine 已有：

- dependency closure visited set；
- 单 target path 去重；
- generated archive/hash cache；
- ProjectGraph ownership；
- nested root exclusion；
- exact artifact path resolution。

当前 Go Selection Planner 和 VS Code 已实现 preset grouped targets、manual kind buckets 和
module identity 去重。下一阶段仍需实现 pinned targets、Select All 影响预估、product-aware
Automotive 默认策略，并让 JetBrains 消费相同 Workbench/Selection Plan。

## 14. Acceptance

- 两个 presets 引用同一 Java 目录时 IDE 只索引一次；
- preset 与 manual 同选一个 module 时 canonical entry 只有一个；
- 显式选择依赖 module 时它成为独立 owner，依赖方引用该 target；
- 相同 srcjar 只解压和投影一次；
- 同 basename 不同路径的 JAR 不被错误合并；
- Select All 不创建与 module 数量线性增长的 IDE modules；
- VS Code 与 JetBrains 消费同一个 Selection Plan 和 ProjectGraph。
