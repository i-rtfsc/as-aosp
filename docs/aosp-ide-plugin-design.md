# AOSP IDE 插件详细设计

## 1. 设计目标

本设计的目标是：用户安装插件后，通过界面选择 AOSP 源码、可选 out、可选 build target 和路径映射，就能在 Android Studio、IntelliJ IDEA 或 VS Code 中浏览 AOSP，并获得普通源码、AIDL、generated source、C/C++ 和 JNI 的导航能力。

第一原则是：**dev-ide 分支只保留新插件方案，不保留旧 Gradle/CMake 兼容实现。**

旧实现继续保存在 aosp 分支，需要查阅历史行为时从该分支读取。dev-ide 的仓库根目录就是新插件项目根目录，不创建嵌套的 aosp-ide/ 工程，也不迁移到新的 Git 仓库。

## 2. 关键技术决策

| 领域 | 决策 |
|---|---|
| 共享核心 | Go `aosp-index serve --stdio` |
| IDE 通信 | stdin/stdout JSON-RPC 2.0 |
| JetBrains 客户端 | Kotlin + IntelliJ Platform SDK |
| VS Code 客户端 | TypeScript + vscode-languageclient |
| Java/Kotlin 语义 | 优先复用 IDE 原生语言能力 |
| AIDL | 共享 engine 解析和解析引用，客户端适配 PSI/LSP |
| C/C++ | clangd/IDE C++ 能力，engine 负责 compdb 映射 |
| JNI | engine 建立 Java/Kotlin 与 C/C++ 跨语言边 |
| 路径身份 | AOSP 仓库相对路径 RepoPath |
| generated source | 解压到 IDE 系统缓存，不写入源码树 |
| 默认执行模式 | 只读，不自动运行 AOSP 构建 |

### 为什么核心选 Go

Go 不是唯一可选项，但最符合当前需求：

- AOSP ide_query 本身使用 Go，概念和数据模型接近；
- 单文件二进制，方便随插件分发；
- 标准库适合 JSON、zip/srcjar、路径和 daemon；
- Linux/macOS/Windows 交叉编译简单；
- 不要求用户安装 Python、Node 或独立 JVM；
- Kotlin 和 TypeScript 客户端都只需要维护薄适配层。

如果后期 JNI/C++ 需要 libclang AST，可以增加一个可选 native helper，不能让它成为基础索引的硬依赖。

## 3. 总体架构

~~~text
                 AOSP source / mounted workspace
                              |
                 optional out / metadata / srcjar
                              |
                    +---------v---------+
                    |  AOSP Index Engine |
                    | aosp-index serve   |
                    +----+----------+----+
                         | JSON-RPC |
             +-----------+          +-----------+
             |                                  |
   +---------v----------+             +---------v---------+
   | JetBrains Client   |             | VS Code Client    |
   | Kotlin             |             | TypeScript        |
   +----+----------+----+             +----+----------+---+
        |          |                       |          |
 Workspace      PSI/Java                LSP client  clangd
 Model          references              commands   compdb
~~~

共享 engine 只认识 profile、RepoPath、module、symbol、definition 和 diagnostics，不 import 任何 IDE API。

## 4. 工作区 Profile

一个 AOSP 工作区由 profile 描述。

~~~json
{
  "schemaVersion": 1,
  "id": "aosp-main",
  "workspaceViewRoot": "/work/aosp-view",
  "canonicalRepoRoot": "/data/aosp-full",
  "outRoot": "/data/aosp-full/out",
  "buildTarget": {
    "product": "aosp_arm64",
    "release": "trunk_staging",
    "variant": "userdebug"
  },
  "scope": {
    "mode": "selected-repositories",
    "include": [
      "frameworks/base",
      "frameworks/native",
      "system/core",
      "packages/services/Car"
    ],
    "exclude": [
      ".repo",
      "out",
      "external/chromium"
    ]
  },
  "sourceMounts": [
    {
      "name": "framework-base-view",
      "repoPrefix": "frameworks/base",
      "viewPath": "frameworks/base",
      "localPath": "/data/aosp-full/frameworks/base",
      "origin": "symlink",
      "priority": 100,
      "readOnly": true
    },
    {
      "name": "system-core-view",
      "repoPrefix": "system/core",
      "viewPath": "system/core",
      "localPath": "/data/aosp-full/system/core",
      "origin": "symlink",
      "priority": 100,
      "readOnly": true
    }
  ],
  "visibilityPolicy": {
    "mode": "strict-view",
    "allowOpenExternalFile": true,
    "confirmBeforeExpanding": true
  },
  "pathMappings": [
    {
      "name": "source",
      "externalPrefix": "/srv/build/aosp",
      "localPrefix": "/data/aosp-full",
      "kind": "repo"
    },
    {
      "name": "out",
      "externalPrefix": "/srv/build/aosp-out",
      "localPrefix": "/data/aosp-full/out",
      "kind": "out"
    }
  ],
  "buildIntegration": {
    "mode": "read-only"
  }
}
~~~

本机绝对路径默认保存在 IDE 用户配置或 .aosp-ide/profile.local.json，不建议提交到 Git。

可以提交的共享配置只包含：

- include/exclude 的 RepoPath；
- Source Mount 的 repoPrefix 和逻辑名称；
- 模块选择；
- profile 名称；
- build target 建议值；
- 不含用户名和盘符的 mapping key。

## 5. 目录拓扑与路径模型

### 5.1 五类目录角色

目录模型不能只有一个 repoRoot。必须区分：

~~~text
Workspace View Root
  用户在 IDE 中实际打开的轻量目录，例如 /work/aosp-view

Canonical AOSP Root
  可选的完整 AOSP，提供自动推导、metadata 回退和缺失源码查询
  默认永远不作为 IDE content root

Out Root
  当前或远程构建产物所在目录
  默认不递归加入 IDE，只精确选取 generated files

Source Mount
  RepoPath 前缀到 View/Local 目录的映射

External Build Root
  Soong metadata、compdb 或远程执行服务中记录的绝对路径
~~~

例如本机完整 AOSP + 轻量目录 A：

~~~text
/data/aosp-full/

/work/aosp-view/
  frameworks/base   -> /data/aosp-full/frameworks/base
  system/core       -> /data/aosp-full/system/core
  ui                -> /data/aosp-full/frameworks/base/packages/SystemUI
~~~

对应 topology：

~~~text
RepoPrefix frameworks/base
  ViewPath frameworks/base
  LocalPath /data/aosp-full/frameworks/base

RepoPrefix system/core
  ViewPath system/core
  LocalPath /data/aosp-full/system/core

RepoPrefix frameworks/base/packages/SystemUI
  ViewPath ui
  LocalPath /data/aosp-full/frameworks/base/packages/SystemUI
~~~

前两个可以按 view 相对路径自动推导。第三个使用任意别名 ui，需要完整根目录辅助反推或用户显式确认。

### 5.2 RepoPath

所有源码身份使用 AOSP 根目录相对路径：

~~~go
type RepoPath string

type FileIdentity struct {
    RepoPath RepoPath
    Origin   FileOrigin
}
~~~

例如：

~~~text
frameworks/base/core/java/android/os/Bundle.java
frameworks/base/core/jni/android_os_Binder.cpp
system/tools/aidl/aidl.cpp
~~~

RepoPath 约束：

- 永远是相对路径；
- 永远使用 /；
- 不包含盘符；
- 不允许越过根目录的 ..；
- 保留源码真实大小写；
- 作为索引 key、RPC 字段和缓存 key。

### 5.3 ViewPath、LocalPath 与 ExternalPath

~~~go
type ViewPath string
type LocalPath string
type ExternalPath string
~~~

- ViewPath 是 Workspace View Root 下用户看到的路径；
- LocalPath 是当前 IDE 所在机器可访问的路径；
- ExternalPath 是 metadata 或 compdb 中记录的构建机路径；
- PathMapper 在二者和 RepoPath 之间转换。

### 5.4 WorkspaceTopology 与 SourceMount

~~~go
type SourceMount struct {
    Name       string
    RepoPrefix RepoPath
    ViewPath   ViewPath
    LocalPath  LocalPath
    Origin     MountOrigin
    Priority   int
    ReadOnly   bool
}

type WorkspaceTopology struct {
    WorkspaceViewRoot  LocalPath
    CanonicalRepoRoot  *LocalPath
    OutRoot            *LocalPath
    SourceMounts       []SourceMount
    ExternalMappings   []PathMapping
}
~~~

MountOrigin 至少支持：

~~~text
directory
symlink
junction
bind-mount
network-mount
copied-tree
overlay
canonical-fallback
~~~

### 5.5 Source Provider 与优先级

同一个 RepoPath 可能有多个来源：

~~~text
1. Workspace View 中的显式 Source Mount
2. 用户配置的 overlay/vendor modified tree
3. Canonical AOSP Root
4. metadata-only missing provider
~~~

解析规则：

1. 最长 RepoPrefix 优先；
2. 同样前缀时 priority 高者优先；
3. 同 priority 且内容不同，返回 conflict；
4. 非 active provider 标记 shadowed；
5. 同一 physical file 的别名只建立一个 RepoPath symbol identity；
6. 一个 physical file 被映射到两个不同 RepoPath 时必须诊断，不能自动合并。

### 5.6 可见性策略

~~~text
strict-view
  只索引 Workspace View 中 active 且被 scope 选中的 mount

open-external-on-demand
  允许从 Canonical AOSP Root 临时打开一个精确文件
  不把其父目录加入 content/source root

auto-expand
  根据模块依赖向 ProjectView 增加 Source Mount
  每次扩展前显示预计新增文件数
~~~

默认是 strict-view。Canonical AOSP Root 的存在不能改变默认索引范围。

依赖源码未暴露时：

1. 有 out：优先使用 jar、srcjar、generated header 和 compdb；
2. 无 out：保留 unresolved dependency，并提供 Add Repository to View；
3. 用户导航到隐藏源码：按 visibility policy 决定提示、临时打开或扩展；
4. 临时打开的外部文件不进入全局 source index。

### 5.7 PathMapper API

~~~text
LocalToRepo(localPath) -> RepoPath
ViewToRepo(viewPath) -> RepoPath
RepoToLocal(repoPath) -> LocalPath or missing
RepoToView(repoPath) -> ViewPath or hidden
ExternalToLocal(externalPath) -> LocalPath or unmapped
ExternalToRepo(externalPath) -> RepoPath or unmapped
Explain(path) -> resolution trace
~~~

Explain 必须能输出：

~~~text
input: /srv/build/aosp/frameworks/base/core/java/android/os/Bundle.java
matched mapping: source
repo path: frameworks/base/core/java/android/os/Bundle.java
local path: Z:\aosp\frameworks\base\core\java\android\os\Bundle.java
active provider: framework-base-view
shadowed provider: canonical-aosp-root
status: resolved
~~~

### 5.8 自动发现目录 A

扫描 Workspace View Root 时使用 lstat 语义，先识别目录项本身，再决定是否解析 target。

自动发现流程：

1. 读取目录 A 的直接和受控深度子项；
2. 识别普通目录、symlink、junction、mount point；
3. 如果 view 相对路径像合法 AOSP RepoPath，生成候选 repoPrefix；
4. 如果 target 位于 Canonical AOSP Root 下，用 target 相对路径生成候选；
5. 两种候选一致时自动接受；
6. 不一致或使用任意别名时要求用户确认；
7. 计算预计可见文件数和 mount 覆盖范围；
8. 生成 topology snapshot，不修改用户目录。

不应无界递归跟随链接。每个 mount 有 traversal budget、最大深度和 cycle detection。

### 5.9 Windows、UNC、junction 和 symlink

Windows 处理规则：

- 接受 C:\、D:\ 和 \\server\share；
- 内部先转换为 platform-neutral path segments；
- 盘符比较大小写不敏感；
- RepoPath 比较保持大小写；
- 检测同一个目录中仅大小写不同的冲突文件；
- 识别 Windows junction 与目录 symlink；
- 不要求 Windows 能创建或跟随 Linux symlink；
- Linux symlink 在 SMB/NFS 客户端表现为普通文件或不可用时，允许用户将 viewPath 显式映射到另一个 localPath。

symlink 记录：

~~~text
logical link path
local visible path
resolved target, if available
target visibility
file identity, if available
~~~

不能一开始就 realpath 所有文件。否则 Linux 软链接工作区在 macOS/Windows 挂载后会失去逻辑 AOSP 目录结构。

### 5.10 异常目录拓扑

必须检测：

- link 指向自身；
- A -> B -> A 的循环；
- dangling link；
- nested mount 覆盖父 mount；
- 一个 RepoPath 有多个不同 active candidates；
- 同一 physical target 通过多个 view path 暴露；
- file-level symlink forest；
- case-only collision；
- mount 暂时不可访问；
- 目录 A 混入不同 AOSP release 的仓库；
- out 被错误映射为 source mount；
- source mount 指向 Workspace View Root 的祖先目录。

循环检测优先使用文件系统 identity；拿不到 identity 时使用归一化路径链和最大跳转次数。

### 5.11 混合版本工作区

一个 profile 对应一个 canonical AOSP namespace、一个 metadata snapshot 和一个 build target。

用户可以在目录 A 中放入厂商修改仓库或 cherry-pick 后的目录，这属于 overlay；但如果 mount 明显来自不同 Android release：

- 普通文件、AIDL 语法和 JNI 文本索引继续工作；
- build-backed 状态降级为 mixed/source-only；
- 不合并两个 release 的 module graph；
- 不使用一个 release 的 compdb 解释另一个 release 的 C++ 文件；
- 建议为不同 release 创建独立 profile；
- diagnostics 显示 release fingerprint 冲突。

release fingerprint 可以由 build/make 版本文件、manifest revision、Soong metadata fingerprint 和用户配置共同判断，不能只看目录名。

## 6. 索引工作模式

### source-only

只有源码，没有 out：

- 只扫描 WorkspaceTopology 中 active 且在 scope 内的 Source Mount；
- 扫描可见的 Android.bp/Android.mk；
- 扫描 Java/Kotlin package、class 和 native declaration；
- 解析 AIDL；
- 扫描 JNI 注册表和导出符号；
- 使用目录规则推断 source root 和 include root；
- 所有非确定结果标记 inferred。

### metadata-partial

部分 out metadata 可用：

- Java metadata 可用但 compdb 不存在；
- compdb 可用但部分 generated file 未挂载；
- metadata 引用的部分仓库不可见；
- build target 信息缺失或可能过期。

插件继续工作，同时在 diagnostics 中列出缺失覆盖率。

### build-backed

以下主要数据匹配：

- repo/out mapping 有效；
- Workspace View、Source Mount 和 overlay 冲突已经解析；
- build target 已选择；
- module_bp_java_deps.json 可解析；
- module-info.json 或等价模块信息可用；
- compdb 可解析；
- 目标模块所需 generated file 可访问。

build-backed 不表示 AOSP 可以完整编译，只表示当前索引数据来自当前构建产物。AOSP compdb 可能对同一源码的多个 variant 去重，只保留一条记录，因此单条 C/C++ 编译参数仍可能存在 variant 歧义。

### stale

metadata 时间、hash 或 build target 与 profile 不一致。stale 结果仍可用于浏览，但 UI 必须提示重新生成或重新扫描。

## 7. 数据发现与优先级

### 一级：Soong/AOSP 构建产物

优先读取：

~~~text
out/soong/module_bp_java_deps.json
out/soong/module_bp_cc_deps.json
out/target/product/<product>/module-info.json
out/soong/module-info-<product>.json
out/soong/development/ide/compdb/compile_commands.json
out/soong/.intermediates 中 metadata 指向的 srcjar/generated files
~~~

每个 reader 独立容错。Java metadata 解析失败不能阻止 AIDL 源码索引；compdb 缺失不能让 Java 导航失效。

### 二级：源码元数据

- repo manifest 和 project path；
- Android.bp；
- Android.mk；
- Java/Kotlin package；
- AIDL package/import；
- JNI 注册与导出模式。

### 三级：目录约定

- java、src、src/main/java；
- aidl、aidl_api；
- jni、cpp、native、include；
- res、assets；
- Framework、SystemUI、Car 等已知 AOSP 目录。

三级数据只用于 fallback，不覆盖一级数据。

## 8. 内部模块图与 IDE 投影

### 8.1 内部 ModuleGraph

engine 保存真实 Soong 模块：

~~~text
Module
  id
  name
  module type
  source paths
  generated paths
  jars
  srcjars
  dependencies
  resource dirs
  AIDL include dirs
  compiler flags
  build variants
  origin
  confidence
~~~

### 8.2 不直接创建数千个 IDE module

内部保留真实 Soong graph，不代表 JetBrains 必须创建同样数量的 module。大量 module 会增加 Workspace Model、索引和 UI 成本。

IDE 投影提供三种模式：

1. **Focused**：用户选择一个或多个 root module，只投影其源码和依赖闭包；
2. **Profile groups**：Framework、Services、SystemUI、Car 等用户分组；
3. **Full graph**：高级选项，只建议强机器和完整挂载使用。

默认使用 Focused。

无论选择哪种模式，Canonical AOSP Root 都不自动成为 IDE module/content root。它只是 engine 的备用 provider。只有实际进入 ProjectView 的 RepoPath 才会被投影到 IDE。

### 8.3 ProjectView

~~~text
ProjectView
  selected modules
  dependency closure
  visible source roots
  generated source roots
  external jars
  excluded roots
  IDE module projection
~~~

JetBrains 客户端可以将多个只读依赖源聚合到少量 IDE module，同时保留 engine 中的真实 owner 信息。这样既能跳转，也避免当前工程“所有模块互相 compileOnly”的模型。

## 9. Generated Source Store

srcjar 和 generated file 写入 IDE 系统缓存：

~~~text
<ide-system>/aosp-ide/
  generated/<profile-id>/<content-hash>/
  metadata/<profile-id>/
  symbols/<profile-id>/
~~~

记录：

- source jar RepoPath/ExternalPath；
- local path；
- entry name；
- size、mtime、hash；
- owning module；
- generated language；
- origin AIDL/proto/R/sysprop/logtags；
- stale 状态。

解压策略：

- 只解压 module closure 需要的 srcjar；
- 内容 hash 去重；
- 原子写入临时目录后 rename；
- 不在源码树创建 build/aosp_srcjars；
- 失效缓存延迟清理，避免索引过程中目录消失。

## 10. AIDL 设计

### 10.1 Parser 范围

第一版解析跳转需要的语法：

- package；
- import；
- interface；
- parcelable；
- enum；
- union；
- method、参数、返回类型；
- constant；
- annotation；
- source range。

第一版不复刻完整 AIDL compiler 的所有稳定性和兼容性检查。

### 10.2 AIDL Symbol

~~~text
AidlSymbol
  package
  name
  kind
  methods/fields
  repo path
  source range
  module owner
  include root
  version/stability, if known
~~~

### 10.3 Resolver 顺序

1. 当前文件；
2. 显式 import；
3. 当前 module 的 AIDL include dirs；
4. 依赖 module 的 include dirs；
5. profile 额外 include dirs；
6. aidl_api/current 或版本目录；
7. package + filename fallback。

如果多个 include root 命中，返回 ambiguous candidates，不随机选择第一个。

### 10.4 Java/AIDL 映射

~~~text
android.foo.IFoo
android.foo.IFoo.Stub
android.foo.IFoo.Proxy
android.foo.FooParcelable
          |
          v
android/foo/IFoo.aidl
android/foo/FooParcelable.aidl
~~~

精确来源：

- Soong Aidl metadata；
- srcjar entry；
- generated source path；
- depfile；
- module owner。

无 out 时使用 package/type 规则，confidence 为 inferred。

### 10.5 AIDL 到 C++/NDK

保存以下映射：

~~~text
IFoo.aidl
  -> IFoo.h
  -> BpFoo.h
  -> BnFoo.h
  -> aidl/.../IFoo.h
  -> generated cpp source
~~~

有 out 时由 generated file 和 module graph 确认；无 out 时只提供命名候选。

## 11. JNI 设计

### 11.1 Java/Kotlin native declaration

索引：

~~~text
Java: native returnType method(parameters)
Kotlin: external fun method(parameters)
Annotations: FastNative, CriticalNative
~~~

唯一 key：

~~~text
java FQCN + method name + JNI descriptor
~~~

### 11.2 C/C++ 注册模式

支持：

- Java_pkg_Class_method；
- Java_pkg_Class_method__signature；
- JNINativeMethod 数组；
- MAKE_JNI_NATIVE_METHOD 等常见宏；
- jniRegisterNativeMethods；
- RegisterMethodsOrDie；
- AndroidRuntime::registerNativeMethods；
- JNIEnv::RegisterNatives；
- JNI_OnLoad；
- FindClass/GetMethodID/GetStaticMethodID/GetFieldID。

### 11.3 JNI Descriptor

单独实现 descriptor package，处理：

- primitive；
- object；
- array；
- inner class；
- overloaded method；
- static 和 instance 的 native receiver 差异。

### 11.4 JNI Link

~~~text
JniLink
  java symbol
  native function
  registration site
  registration class
  descriptor
  module owner
  origin
  confidence
  diagnostics
~~~

导航优先级：

1. 精确动态注册；
2. 精确静态导出；
3. 带签名的命名规则；
4. 不带签名的唯一候选；
5. 多候选选择列表；
6. unresolved 诊断。

### 11.5 JNI 不能保证的情况

- class name 由运行时复杂逻辑生成；
- 注册表由宏和多个编译单元拼接；
- 函数指针经多层变量传递；
- dlopen 后运行时注册；
- ART 内部非标准绑定。

这些情况仍可以用文本候选辅助，但不能标 exact。

## 12. C/C++ 与 clangd

### 12.1 Compdb 映射

AOSP compdb 中包含构建机路径。engine 生成本机 overlay compdb：

1. 映射 directory；
2. 映射 file；
3. 映射 -I、-isystem、--sysroot、-include、response file；
4. 处理 out 路径；
5. 保留 target 和宏；
6. 对不可映射参数记录 diagnostics；
7. 写入缓存目录，不改原始 compdb。

### 12.2 Compiler 路径

Linux 构建机 compiler 在 macOS/Windows 不可执行。处理策略：

- clangd 解析 flags，不直接用该 compiler 构建；
- 可配置本机 clangd；
- 必要时将 argv[0] 替换为本机 clang；
- 保留 --target、sysroot、include 和宏；
- 如果 resource dir 无法映射，标记 partial。

### 12.3 无 compdb fallback

- 当前文件目录；
- 当前 module source/include roots；
- AOSP 常用公共 include；
- 用户额外 flags；
- 不猜测高风险 product 宏；
- 显示 heuristic 状态。

## 13. JetBrains 客户端

### 13.1 兼容范围

同一 Kotlin 源码支持 Android Studio 和 IntelliJ IDEA，但 JetBrains Platform API 有版本范围。发布时需要：

- 选定最低 IDE baseline；
- 使用 verifier 测试 Android Studio 和 IDEA；
- 对 Android-only API 使用 optional dependency；
- 必要时为不同 IDE build range 发布不同版本，但保持同一源码分支。

操作系统可以使用同一个插件包；IDE 大版本兼容不应与操作系统兼容混为一谈。

### 13.2 Workspace Model

插件创建：

- content roots；
- source roots；
- generated source roots；
- libraries/jars；
- selected module dependency；
- excluded roots。

默认不创建 Android Gradle module，不依赖 AGP，不修改 android.jar。

对于目录 A：

- Workspace View Root 是项目入口；
- 只把 ProjectView 选中的 Source Mount 或其子目录登记为 source/content root；
- Canonical AOSP Root 不登记为 content root；
- 高层 symlink 暴露了大量未选择目录时，尽早登记 exclude roots；
- 同一文件通过 view path 和 physical path 可见时，只向 IDE 暴露 view path；
- 如果 IDE 平台强制解析到 physical target，客户端仍通过 RepoPath 去重并阻止创建第二份 module/root。

需要承认一个平台边界：用户若在目录 A 中软链接整个 AOSP 顶层，IDE 在插件完成配置前可能已经枚举部分目录。插件可以通过项目模板、首次打开向导和早期 exclude 降低成本，但无法控制所有 IDE 在插件加载前的文件系统行为。

### 13.3 AIDL PSI

JetBrains 侧需要：

- AIDL Language/FileType；
- lexer/parser definition；
- syntax highlighter；
- PSI element/source range；
- reference contributor；
- find usages；
- go to declaration；
- structure view；
- engine RPC adapter。

parser 可以由 engine 与客户端共享 grammar 设计，但 PSI 节点必须由 JetBrains 客户端构造。

### 13.4 JNI UI

提供：

- Java native 方法 gutter 图标；
- C/C++ registration entry gutter 图标；
- Go to JNI peer；
- Find JNI peers；
- 多候选 popup；
- unresolved inspection。

### 13.5 工具窗口

页面：

1. Overview；
2. Modules；
3. Generated Sources；
4. AIDL；
5. JNI；
6. Path Mappings；
7. Diagnostics。

## 14. VS Code 客户端

### 14.1 组成

- TypeScript extension；
- engine process manager；
- JSON-RPC client；
- AIDL language client；
- JNI commands；
- clangd/compdb integration；
- status bar 和 tree view。

### 14.2 AIDL

提供：

- TextMate syntax；
- Definition；
- References；
- DocumentSymbol；
- WorkspaceSymbol；
- Hover；
- Diagnostics；
- Rename 可放到第二阶段。

### 14.3 Java/Kotlin

VS Code 客户端不重写完整 Java/Kotlin language server。它做：

- generated source 暴露；
- external jars/profile 配置；
- AIDL 和 JNI cross-language definition provider；
- 与用户已有 Java/Kotlin 扩展协作。

因此 VS Code 的普通 Android resource 和 Kotlin 体验不会自动等同于 Android Studio。

## 15. UI 详细设计

### 15.1 Welcome Wizard

步骤一：识别用户打开的目录

~~~text
Workspace View Root [current project] /work/aosp-view
Detected:
  Directories          3
  Symlinks             7
  Junctions/Mounts     0
  Broken links         1
~~~

步骤二：可选完整 AOSP 与 out

~~~text
Canonical AOSP Root [Browse] /data/aosp-full  (不会加入 IDE 索引)
Out Root            [Browse] /data/aosp-full/out

[ ] 将完整 AOSP 作为 fallback provider
[x] 仅索引 Workspace View 中选中的内容
~~~

步骤三：目录拓扑

~~~text
Repo prefix                         View path             Target/status
frameworks/base                     frameworks/base       symlink, ready
system/core                         system/core           symlink, ready
frameworks/base/packages/SystemUI   ui                    inferred, confirm
packages/modules/Wifi               wifi                  broken

[Scan links] [Add mapping] [Resolve conflicts] [Estimate size]
~~~

步骤四：检测 out

~~~text
Out root [Browse] [Skip]
Detected:
  Java module metadata  Ready
  CC module metadata    Ready
  Compdb                Missing
  Generated files       Partial
~~~

步骤五：外部构建路径映射

~~~text
External prefix             Local prefix
/srv/build/aosp             Z:\aosp
/srv/build/aosp-out         Z:\aosp-out

[Auto detect] [Test all mappings]
~~~

步骤六：范围

~~~text
(*) Focused modules
( ) Repository selection
( ) Full visible tree

Search modules: [SystemUI____________]
Selected: SystemUI-core, services
~~~

步骤七：确认

显示：

- active/shadowed/missing Source Mount；
- IDE 预计可见文件数；
- Canonical AOSP Root 是否保持在索引外；
- generated 数；
- 不可见路径；
- source-only/partial/build-backed 模式。

### 15.2 Diagnostics

每条诊断包含：

- severity；
- data source；
- external path；
- RepoPath；
- local mapping；
- suggested action；
- 是否影响 Java/AIDL/C++/JNI。

## 16. RPC 协议

建议第一版方法：

~~~text
aosp/initialize
aosp/shutdown
aosp/profile/validate
aosp/profile/detect
aosp/topology/discover
aosp/topology/validate
aosp/topology/resolveConflict
aosp/topology/estimateVisibility
aosp/index/status
aosp/index/rebuild
aosp/index/refreshFiles
aosp/projectView/resolve
aosp/document/symbols
aosp/definition
aosp/references
aosp/workspace/symbols
aosp/aidl/resolve
aosp/aidl/generatedPeers
aosp/jni/peers
aosp/path/explain
aosp/diagnostics/list
~~~

文件位置使用 RepoPath：

~~~json
{
  "profileId": "aosp-main",
  "repoPath": "frameworks/base/core/java/android/os/Bundle.java",
  "line": 120,
  "column": 17
}
~~~

客户端负责：

~~~text
IDE VirtualFile/URI -> RepoPath -> engine
engine RepoPath -> LocalPath -> IDE navigation
~~~

## 17. 缓存与增量更新

### 17.1 缓存内容

~~~text
profiles/
metadata/
module-graph/
generated/
symbols/
compdb-overlay/
diagnostics/
~~~

### 17.2 失效条件

- Workspace View Root、Canonical AOSP Root、Out Root 或 external path mapping 改变；
- Workspace View Root、Canonical AOSP Root 或 Source Mount 改变；
- symlink/junction target、overlay priority 或 active provider 改变；
- 原先 missing 的 mount 恢复，或原先可见的 mount 断开；
- product/release/variant 改变；
- metadata mtime、size、hash 改变；
- Android.bp/Android.mk 改变；
- AIDL 或 JNI source 改变；
- srcjar hash 改变；
- mapping 目标从 available 变成 missing；
- engine schema 版本改变。

### 17.3 网络挂载

网络文件系统不能只依赖 watcher：

- metadata 定时轻量轮询；
- 用户手动 refresh；
- 打开文件时按需验证；
- 大规模重扫需要 debounce；
- out 正在同步时等待文件大小稳定后解析。

## 18. 可选远程执行服务

第一版只要求挂载文件可读。后续可以增加运行在 Linux AOSP 构建机上的远程执行服务：

~~~text
IDE client
   |
SSH or HTTPS
   |
aosp-build-bridge on Linux
   |
soong_ui / ide_query / metadata refresh / build
~~~

远程执行服务的职责：

- 执行 Linux-only ide_query；
- 生成或刷新 compdb；
- 构建选定 generated source；
- 返回 RepoPath、module、generated content 和 diagnostics；
- 不向客户端暴露不可解释的服务器绝对路径。

远程执行必须是用户显式开启的模式，默认只读。

## 19. 推荐代码目录结构

下面的结构将共享核心、IDE 客户端、协议、打包和测试严格分开。

dev-ide 分支的当前 Git 仓库根目录就是插件项目根目录，不创建 aosp-ide/ 子目录，也不迁移到另一个 Git 仓库：

~~~text
/Users/solo/code/aosp/                    真实 AOSP，不放插件源码

/Users/solo/code/github/as-aosp/          dev-ide 新插件项目根目录
  core/
  clients/
  protocol/
  remote-execution/
  packaging/
  docs/
~~~

旧 Gradle/CMake 方案只存在于 aosp 分支，不进入 dev-ide 的目录结构。

~~~text
<repository-root>/
├── README.md
├── LICENSE
├── go.mod
├── go.sum
├── Makefile
├── Taskfile.yml
│
├── docs/
│   ├── aosp-ide-plugin-feasibility-report.md
│   ├── aosp-ide-plugin-design.md
│   ├── protocol.md
│   ├── path-mapping.md
│   ├── index-modes.md
│   └── compatibility.md
│
├── protocol/
│   ├── schema/
│   │   ├── profile.schema.json
│   │   ├── rpc.schema.json
│   │   └── diagnostics.schema.json
│   └── examples/
│
├── core/
│   ├── model/
│   │   ├── profile.go
│   │   ├── path.go
│   │   ├── module.go
│   │   ├── symbol.go
│   │   ├── location.go
│   │   └── diagnostics.go
│   │
│   ├── pathmap/
│   │   ├── mapper.go
│   │   ├── normalize.go
│   │   ├── windows.go
│   │   ├── symlink.go
│   │   └── explain.go
│   │
│   ├── topology/
│   │   ├── workspace_view.go
│   │   ├── source_mount.go
│   │   ├── discover.go
│   │   ├── overlay.go
│   │   ├── priority.go
│   │   ├── cycle_detection.go
│   │   ├── conflicts.go
│   │   ├── file_identity.go
│   │   └── topology_test.go
│   │
│   ├── discover/
│   │   ├── repo.go
│   │   ├── out.go
│   │   ├── manifest.go
│   │   ├── build_target.go
│   │   └── doctor.go
│   │
│   ├── metadata/
│   │   ├── java_deps.go
│   │   ├── cc_deps.go
│   │   ├── module_info.go
│   │   ├── compdb.go
│   │   ├── ide_query.go
│   │   ├── fingerprint.go
│   │   └── compatibility.go
│   │
│   ├── buildfiles/
│   │   ├── blueprint_lexer.go
│   │   ├── blueprint_parser.go
│   │   ├── android_bp.go
│   │   ├── android_mk.go
│   │   └── fallback_graph.go
│   │
│   ├── generated/
│   │   ├── store.go
│   │   ├── srcjar.go
│   │   ├── depfile.go
│   │   ├── mapping.go
│   │   └── cleanup.go
│   │
│   ├── java/
│   │   ├── source_scan.go
│   │   ├── package_index.go
│   │   ├── classpath.go
│   │   └── resolver.go
│   │
│   ├── aidl/
│   │   ├── lexer.go
│   │   ├── parser.go
│   │   ├── ast.go
│   │   ├── symbols.go
│   │   ├── resolver.go
│   │   ├── include_roots.go
│   │   └── generated_peers.go
│   │
│   ├── cpp/
│   │   ├── compdb_mapper.go
│   │   ├── arguments.go
│   │   ├── include_roots.go
│   │   ├── clangd_config.go
│   │   └── fallback.go
│   │
│   ├── jni/
│   │   ├── java_native.go
│   │   ├── exported_name.go
│   │   ├── registration_table.go
│   │   ├── descriptor.go
│   │   ├── resolver.go
│   │   └── diagnostics.go
│   │
│   ├── graph/
│   │   ├── module_graph.go
│   │   ├── dependency_closure.go
│   │   ├── project_view.go
│   │   └── scope.go
│   │
│   ├── index/
│   │   ├── symbol_index.go
│   │   ├── cross_language.go
│   │   ├── snapshot.go
│   │   ├── incremental.go
│   │   └── storage.go
│   │
│   ├── rpc/
│   │   ├── server.go
│   │   ├── session.go
│   │   ├── methods.go
│   │   ├── types.go
│   │   └── errors.go
│   │
│   └── app/
│       ├── engine.go
│       ├── lifecycle.go
│       └── logging.go
│
├── cmd/
│   ├── aosp-index/
│   │   └── main.go
│   └── aosp-doctor/
│       └── main.go
│
├── clients/
│   ├── jetbrains/
│   │   ├── build.gradle.kts
│   │   ├── settings.gradle.kts
│   │   ├── gradle.properties
│   │   └── src/main/
│   │       ├── kotlin/com/example/aospide/
│   │       │   ├── engine/
│   │       │   ├── model/
│   │       │   ├── settings/
│   │       │   ├── ui/
│   │       │   ├── project/
│   │       │   ├── aidl/
│   │       │   ├── jni/
│   │       │   └── actions/
│   │       └── resources/
│   │           ├── META-INF/plugin.xml
│   │           └── icons/
│   │
│   └── vscode/
│       ├── package.json
│       ├── tsconfig.json
│       ├── src/
│       │   ├── extension.ts
│       │   ├── engineProcess.ts
│       │   ├── engineRpcClient.ts
│       │   ├── profile.ts
│       │   ├── configuration.ts
│       │   ├── aidlLanguageClient.ts
│       │   ├── jniCommands.ts
│       │   ├── clangd.ts
│       │   ├── diagnostics.ts
│       │   └── statusView.ts
│       ├── syntaxes/aidl.tmLanguage.json
│       └── language-configuration.json
│
├── remote-execution/
│   ├── bridge/
│   ├── transport/
│   └── build/
│
├── packaging/
│   ├── engine/
│   │   ├── darwin-arm64.json
│   │   ├── darwin-amd64.json
│   │   ├── linux-arm64.json
│   │   ├── linux-amd64.json
│   │   ├── windows-arm64.json
│   │   └── windows-amd64.json
│   ├── jetbrains/
│   └── vscode/
│
├── testdata/
│   ├── minimal-aosp/
│   ├── source-only/
│   ├── metadata/
│   ├── generated-srcjar/
│   ├── aidl/
│   ├── jni/
│   ├── compdb/
│   ├── partial-mount/
│   ├── local-full-aosp-light-view/
│   ├── alias-view/
│   ├── nested-overlays/
│   ├── duplicate-links/
│   ├── dangling-links/
│   ├── link-cycles/
│   ├── junction-view/
│   ├── mixed-copied-and-linked/
│   ├── symlink-view/
│   └── windows-paths/
│
└── .github/workflows/
    ├── core.yml
    ├── jetbrains.yml
    ├── vscode.yml
    └── release.yml
~~~

### 目录依赖规则

- core/model 不依赖任何其他业务 package；
- core/pathmap 只能依赖 model；
- core/topology 负责 Workspace View、Source Mount、overlay、优先级和循环/冲突检测；
- core/pathmap 不决定哪个 provider 生效；
- core/topology 不把任何完整 AOSP 根目录自动转换成 IDE content root；
- metadata reader 只读数据，不创建 IDE module；
- aidl/jni 只输出 normalized symbols 和 links；
- graph 决定 module closure 和 ProjectView；
- index 负责快照、查询和增量；
- clients 不直接解析 Soong JSON；
- packaging 不包含业务逻辑；
- testdata 必须覆盖 Linux build path 到 macOS/Windows local path。

## 20. 打包与发布

### Engine 构建目标

第一版至少：

~~~text
linux-amd64
linux-arm64
darwin-amd64
darwin-arm64
windows-amd64
~~~

Windows arm64 可在有用户需求后加入。

### JetBrains 插件包

两种方式：

1. 一个较大的通用插件包，包含全部 engine；
2. Marketplace 按平台发布变体。

第一版推荐通用包，减少用户选择成本。启动时根据 os/arch 选择资源并复制到 IDE cache，校验 hash 后执行。

macOS 需要考虑可执行文件签名和 quarantine；Windows 需要考虑 Defender；Linux 需要保留 executable permission。发布流程必须对 engine 做 checksum。

### VS Code

优先使用 VS Code 平台专用 VSIX，避免一个包携带所有二进制。源码和版本保持一致。

## 21. 测试策略

### Core 单元测试

- path normalization；
- Windows/UNC；
- path mapping；
- Workspace View 与 Canonical Root 分离；
- symlink/junction visible/missing；
- symlink cycle；
- duplicate physical target；
- RepoPath provider conflict；
- nested overlay priority；
- metadata schema 容错；
- srcjar 解压和缓存；
- AIDL import；
- JNI descriptor；
- JNINativeMethod 注册；
- compdb argument mapping；
- module closure。

### Fixture 集成测试

1. source-only；
2. 完整 metadata；
3. 只有 Java out；
4. 只有 compdb；
5. 部分仓库挂载；
6. 本机完整 AOSP + 轻量软链接目录 A；
7. 任意别名目录到 RepoPath 的显式映射；
8. 同一物理目录被多个链接暴露；
9. overlay 覆盖 Canonical AOSP Root；
10. Linux 外部路径映射到 macOS；
11. Linux 外部路径映射到 Windows；
12. symlink target 不可见；
13. symlink 循环；
14. AIDL 同名 import 歧义；
15. JNI 多候选。

### JetBrains 测试

- plugin verifier；
- Android Studio 与 IDEA baseline；
- source root 安装；
- generated root 更新；
- AIDL PSI/reference；
- JNI action；
- profile persistence。

### VS Code 测试

- extension host；
- engine lifecycle；
- AIDL LSP；
- JNI command；
- clangd overlay compdb；
- Windows URI。

## 22. 分阶段实施

### Phase 0：技术验证

- RepoPath 和 PathMapper；
- WorkspaceTopology、SourceMount 和 provider priority；
- 本机完整 AOSP + 软链接目录 A；
- broken link、循环、别名和重复 provider 诊断；
- 读取 module_bp_java_deps.json；
- srcjar 缓存；
- CLI 输出 module/source/generated 信息；
- Linux 到 macOS/Windows 路径映射测试。

### Phase 1：JetBrains MVP

- 设置界面；
- profile validation；
- Focused ProjectView；
- source/generated roots；
- Java 普通跳转；
- engine 状态窗口。

### Phase 2：AIDL

- AIDL parser；
- import/type Definition；
- Java/AIDL 双向跳转；
- generated peers；
- source-only 和 build-backed 两种结果。

### Phase 3：C/C++ 与 JNI

- compdb overlay；
- clangd integration；
- Java/C++ JNI 双向跳转；
- 多候选和 diagnostics。

### Phase 4：VS Code

- TypeScript client；
- AIDL LSP；
- JNI commands；
- clangd config；
- profile UI。

### Phase 5：远程执行

- Linux metadata refresh；
- ide_query；
- 按需构建 generated source；
- SSH/HTTPS；
- 客户端 partial mount。

## 23. 第一版验收标准

1. Linux、macOS、Windows 对同一 fixture 产生相同 RepoPath 索引；
2. 配置 source root 后，无 out 也能跳普通 Java、AIDL 和常见 JNI；
3. 配置 out 后能挂载 srcjar/generated Java；
4. Linux server path 映射到 macOS/Windows 后仍能导航；
5. 完整 AOSP 在本机存在但只打开软链接目录 A 时，不会把完整 AOSP 加入 IDE content root；
6. 只挂载 frameworks/base 等部分仓库时不会强制扫描全树；
7. compdb 路径经过转换后 clangd 能打开一个 AOSP C++ 文件；
8. Java native 与静态 JNI、动态注册表能够双向跳；
9. 缺失 mapping、broken link、shadowed provider 和 generated file 有清晰 diagnostics；
10. 不修改 AOSP checkout；
11. 不要求用户安装 Python。

## 24. 第一版明确不做

- 完整复刻 Soong 所有 mutator 和 product variable；
- macOS/Windows 本机运行 Linux-only ide_query.sh；
- 自己实现完整 C++ 类型系统；
- 把整个 out/soong/.intermediates 加入源码根；
- 修改 Android SDK android.jar；
- 自动修改用户挂载或 symlink；
- 保证运行时动态 JNI 百分之百可解析；
- 一个安装包同时安装到 JetBrains 和 VS Code。

## 最终建议

dev-ide 分支的仓库根目录直接承载新方案，后续继续使用当前 Git 仓库和已有 Star、Issue、Release 与用户入口，不创建新仓库。

旧 Gradle/CMake 方案只保留在 aosp 分支。需要研究历史实现时跨分支读取，不在 dev-ide 中复制兼容目录、旧脚本或旧模块。

第一条实现主线应固定为：

~~~text
配置 repo/out
 -> WorkspaceTopology + PathMapper
 -> metadata reader 或 source-only scan
 -> ModuleGraph
 -> AIDL/JNI cross-language index
 -> ProjectView
 -> JetBrains/VS Code navigation
~~~

只要这条主线保持 IDE 中立、路径中立和构建系统中立，后续才能稳定扩展到更多 AOSP 版本、厂商树、远程构建机和部分仓库挂载。
