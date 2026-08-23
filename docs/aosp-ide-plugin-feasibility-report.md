# AOSP IDE 插件可行性报告

## 结论摘要

这个需求可以实现。dev-ide 分支将直接替换历史 Gradle/CMake 方案，成为“索引引擎 + IDE 客户端”的新产品实现。

需要区分三个目标：

1. **跨操作系统**：同一套源码支持 Linux、macOS、Windows。
2. **跨 IDE**：Android Studio/IntelliJ IDEA 与 VS Code 的插件形态不同，但可以复用同一个索引引擎。
3. **跳转准确性**：没有 out 时可以提供源码级和启发式跳转；有匹配当前构建变体的 out 时，才能提供完整的生成代码、编译参数和模块依赖信息。

推荐的最终形态：

~~~text
AOSP 源码 + 可选 out + 可选构建变体
                    |
             AOSP Index Engine
       路径、Soong、生成文件、AIDL、JNI、C/C++
          /                         \
 JetBrains 客户端                 VS Code 客户端
 Android Studio / IDEA             VS Code + clangd
~~~

## 研究依据

本报告参考了：

- 当前工程 as-aosp 的 Gradle 配置和脚本；
- ~/code/github/gs_work/backup/as-aosp 中的 Gradle、Android.bp、srcjar 和 CMake 生成脚本；
- BAS025-Android-Studio配置.md；
- BAS026-Compdb编译数据库.md；
- BAS029-AIDL编译器.md；
- AOSP 的 build/make/tools/ide_query；
- AOSP Soong 的 IdeInfo、Java dependency、C/C++ compdb 和 AIDL provider；
- AOSP Framework JNI 的静态导出和 JNINativeMethod 动态注册代码。

当前 AOSP 的 Soong IDE 数据已经包含 srcs、dependencies、jars、srcjars、AIDL 源文件、AIDL include 路径和 Java/Kotlin 编译参数。C/C++ 还可以通过 compile_commands.json 获得编译参数。这是插件化的基础。

### AOSP 源码证据

| 结论 | AOSP 位置 |
|---|---|
| Java IDE metadata 包含依赖、srcjar、AIDL、资源和编译参数 | build/soong/android/module.go 的 IdeInfo |
| Java module metadata 默认生成 | build/soong/java/jdeps.go |
| C/C++ module metadata 默认生成 | build/soong/cc/ccdeps.go |
| compdb 由 SOONG_GEN_COMPDB 控制 | build/soong/cc/compdb.go |
| ide_query 消费 Java metadata 和 compdb | build/make/tools/ide_query/ide_query.go |
| ide_query shell 入口只支持 Linux | build/make/tools/ide_query/ide_query.sh |
| AIDL provider 提供 srcs 和 include dirs | build/soong/aidl_library/aidl_library.go |
| AIDL 后端生成 Java/C++/NDK/Rust | system/tools/aidl/aidl.cpp |
| Framework JNI 大量使用动态注册表 | frameworks/base/core/jni 和 services/core/jni |

需要注意：AOSP compdb 对同一源码的多个 variant 通常只保留一条记录，因此“有 compdb”仍不等于保存了 host/device、架构、vendor/recovery 的全部组合。插件必须显示选中记录及其来源，不能把它宣传成所有 C++ variant 的绝对真相。

## 当前工程真正做了什么

当前工程不是在构建完整 AOSP，而是在人工构造 IDE 模型：

- settings.gradle 固定声明一组模块；
- scripts/configs/*.gradle 固定列出源码、AIDL、srcjar 和 out/soong/.intermediates 路径；
- scripts/task.gradle 把模块互相添加为 compileOnly 依赖；
- scripts/copy_aosp_scrjar.py 在 Gradle 配置阶段解压生成源码；
- 备份脚本根据 Android.bp 和 Soong JSON 生成新的 Gradle/CMake 文本。

这套方案的问题是：

1. 模块清单和路径必须跟着 AOSP 版本维护；
2. 生成目录会随 product、架构和 variant 变化；
3. Gradle 配置阶段执行 Python，环境依赖会阻塞导入；
4. Java、AIDL、C/C++ 的真实 owner 是 Soong、AIDL 编译器和 Clang，而不是 Gradle。

本地执行 ./gradlew projects 时，当前配置会因为直接执行 python 命令而失败。这说明“安装插件后自动工作”不能继续依赖现有脚本链路。

## 1. Linux、macOS、Windows 通用性

### 结论

**同一套源码可以跨平台；一个未打包的二进制不能直接在三个系统运行。**

推荐拆分如下：

| 层 | 推荐实现 | 跨平台方式 |
|---|---|---|
| 共享索引引擎 | Go | 编译 Linux/macOS/Windows 多架构二进制 |
| Android Studio/IDEA | Kotlin/JVM | 同一 JetBrains 插件源码 |
| VS Code | TypeScript | 同一 VS Code 扩展源码 |
| C/C++ | clangd/Clang | 使用本机或随包工具 |
| AIDL/JNI | 共享 engine | 客户端只做 API 适配 |

“跨平台”有三个含义：

- 核心源码跨平台：可以；
- JetBrains 插件包覆盖三种操作系统：可以，随包携带多个 engine；
- 一个完全相同的安装包同时安装到 JetBrains 和 VS Code：不可以，两个 IDE 的扩展 API 不同。

当前 AOSP 的 ide_query.sh 明确只支持 Linux，因此插件不能把它当作跨平台唯一依赖。Linux 上可以调用它或可选远程执行服务；macOS/Windows 应优先读取 JSON、srcjar 和 compdb。

### 为什么不继续使用 Python

Python 适合一次性生成脚本，不适合作为 IDE 常驻索引服务的必需运行时。当前工程已经出现 python 命令不存在导致 Gradle 配置失败的问题。engine 应该是插件自带的可执行文件，用户不需要安装 Python。

## 2. 插件源码使用什么

### 共享核心：Go

原因：

- AOSP ide_query 本身使用 Go；
- JSON、protobuf、路径处理和常驻 daemon 适合 Go；
- Go 可以输出单文件跨平台二进制；
- Kotlin 和 TypeScript 只通过 JSON-RPC 调用同一 engine；
- 不需要在 JVM 和 Node 中各实现一套 Soong、AIDL、JNI 解析逻辑。

不能直接 import AOSP 内部 Go 包。应复用 AOSP 的协议和数据模型，在独立 module 中实现稳定的 normalized model。

### JetBrains 客户端：Kotlin

- IntelliJ Platform SDK；
- Workspace Model 或 Project Model；
- 设置页、工具窗口、Actions；
- Java/Kotlin 使用 IDE 原生类型系统；
- 插件补充 generated source、AIDL 和 JNI cross-language reference。

### VS Code 客户端：TypeScript

- vscode-languageclient；
- AIDL 使用共享 engine 的 LSP 结果；
- C/C++ 使用 clangd；
- JNI 使用扩展命令和共享 engine 的跨语言边；
- Java/Kotlin 依赖用户安装的 language server，不能假设拥有 Android Studio 的完整 Android 模型。

## 3. 插件是否需要界面

需要。界面负责配置、状态和诊断，不暴露 Soong 的全部复杂性。

首次配置：

~~~text
AOSP Workspace

Source root       [选择目录]
Out root          [选择目录]  (可选)
Build target      [自动检测]  product / release / variant

Path mappings     [管理]
Index scope       [全仓库 / 选定模块 / 当前文件依赖]
Build integration [只读 / 本机 / 远程执行]

[验证配置] [开始索引]
~~~

设置页建议分组：

1. Workspace：repo root、out root、缓存目录；
2. Build variant：product、release、variant；
3. Path mappings：服务器路径到本机路径；
4. Index scope：全仓、模块、当前文件依赖；
5. Generated sources：srcjar 和 generated file 策略；
6. C/C++：clangd、compdb、fallback flags；
7. Build integration：只读、本机执行、SSH/远程执行服务。

工具窗口显示状态和来源：

~~~text
Source       Ready     18,421 files
Out          Partial   6,032 generated files
Java graph   Ready     1,842 modules
C/C++ graph  Ready     4,931 compile entries
AIDL         Ready     612 interfaces
JNI          Partial   1,204 links, 87 unresolved

[重新扫描] [查看诊断] [打开配置]
~~~

跳转结果必须显示 Soong metadata、compile_commands、source heuristic 或 JNI naming convention 等来源。

## 4. 目录拓扑、轻量工作区、挂载和软链接

### 结论

**可以支持，而且“IDE 打开的目录”和“完整 AOSP 所在目录”必须从第一天就是两个独立概念。**

一个非常重要的本机场景是：

~~~text
/data/aosp-full/                     完整 AOSP，不直接用 IDE 打开

/work/aosp-view/                     IDE 只打开这个轻量目录
  frameworks/base       -> /data/aosp-full/frameworks/base
  frameworks/native     -> /data/aosp-full/frameworks/native
  system/core           -> /data/aosp-full/system/core
  packages/SystemUI     -> /data/aosp-full/frameworks/base/packages/SystemUI
  out                   -> /data/aosp-full/out
~~~

这个场景不属于远程挂载的特例，而应该是第一等工作模式。插件内部至少区分：

| 概念 | 含义 | 是否自动加入 IDE 索引 |
|---|---|---:|
| Workspace View Root | 用户实际打开的目录 A | 是，但只限选定范围 |
| Canonical AOSP Root | 完整 AOSP 的真实位置 | 否 |
| Out Root | 构建产物位置 | 否，只精确挂载需要的生成文件 |
| Source Mount | 某个 RepoPath 前缀到本机目录的映射 | 按 scope |
| External Build Root | metadata/compdb 中记录的构建机路径 | 否，只用于路径转换 |

因此用户可以拥有完整 AOSP，同时让 IDE 只看到目录 A 中选中的部分。插件读取完整 AOSP 根目录或 out 的 metadata，不代表把它们设成 content root。

唯一身份使用 AOSP 根目录下的仓库相对路径：

~~~text
frameworks/base/core/java/android/os/Bundle.java
~~~

磁盘访问由 Source Mount 和 Path Mapping 完成：

~~~text
RepoPath:       frameworks/base/core/java/android/os/Bundle.java
View path:      /work/aosp-view/frameworks/base/core/java/android/os/Bundle.java
Physical path:  /data/aosp-full/frameworks/base/core/java/android/os/Bundle.java
Build path:     /srv/aosp/frameworks/base/core/java/android/os/Bundle.java
~~~

路径模型必须：

- 内部统一使用 /；
- Windows 盘符、UNC、macOS volume、Linux mount 只存在于 filesystem 层；
- 仓库相对路径保持大小写；
- 同时保存 view path、link path、physical target 和 external build path；
- 不能把 realpath 后的完整 AOSP 目录误设为 IDE 根目录；
- 不能依赖本机 realpath 才能建立索引；
- 同一个 RepoPath 只建立一个符号身份，避免软链接路径和真实路径被重复索引；
- 一个物理目录映射到多个 RepoPath，或一个 RepoPath 存在多个候选目录时，必须报告冲突或使用显式优先级。

可支持：

1. 本机有完整 AOSP，只打开软链接拼装的目录 A；
2. 目录 A 保持 AOSP 原始相对结构，例如 A/frameworks/base；
3. 目录 A 使用任意别名，例如 A/framework 指向 frameworks/base，此时由用户显式填写 RepoPath 前缀；
4. 目录、软链接、复制文件和独立 vendor 仓库混合存在；
5. Linux bind mount、macOS volume、Windows junction/UNC/share；
6. Linux AOSP 通过 NFS、SMB、SSHFS 等挂载到 macOS/Windows；
7. Linux 先将多个 project 软链接到工作目录，再整体挂载；
8. 只挂载 Framework、SystemUI、packages/modules 等需要的仓库；
9. 源码和 out 分别挂载到不同本机目录；
10. out 的 JSON、srcjar 和 generated file 通过独立映射提供；
11. 用另一个本地仓库覆盖完整 AOSP 中的某个 RepoPath，例如用厂商修改版替换 frameworks/base 的局部目录。

### 自动发现与显式配置

如果目录 A 保持 AOSP 相对结构，插件可以自动推导：

~~~text
A/frameworks/base -> RepoPath prefix frameworks/base
A/system/core     -> RepoPath prefix system/core
~~~

如果目录名称被改写，不能只根据 target 猜测。例如：

~~~text
A/framework -> /data/aosp-full/frameworks/base
~~~

插件应该建议映射，但必须由用户确认：

~~~text
RepoPath prefix: frameworks/base
View path:       /work/aosp-view/framework
Physical target: /data/aosp-full/frameworks/base
~~~

完整 AOSP 根目录可以帮助自动反推出 target 相对于 AOSP 的 RepoPath，但仍不能自动把完整根目录加入索引。

### 覆盖和优先级

用户可能将某个源码目录替换成修改版：

~~~text
A/frameworks/base -> /work/vendor-modified/frameworks/base
Canonical root    -> /data/aosp-full
~~~

此时 Source Mount 的优先级应是：

1. Workspace View 中的显式 mount；
2. 用户配置的 overlay；
3. Canonical AOSP Root 回退；
4. metadata 中只有记录但本机不可见的 missing entry。

同一个 RepoPath 如果同时出现在多个可见位置，插件不能静默索引两份。应选定一个 active provider，并在诊断中列出 shadowed providers。

### 完整 AOSP 存在，不等于允许自动扩大工作区

默认策略应是 strict-view：

- 只索引目录 A 明确暴露并被 scope 选中的 RepoPath；
- Canonical AOSP Root 只用于识别链接目标、读取 metadata 和判断缺失源码位置；
- Java/C++ 依赖优先使用 out 中的 jar、srcjar、generated header 或编译参数；
- 未链接进目录 A 的源码不自动成为 source root；
- 用户跳转到未暴露源码时，可以选择临时打开单个文件，或者把对应仓库/模块加入 Workspace View；
- 插件不能因为发现完整 AOSP 就在后台索引整棵树。

可以提供三个策略：

| 策略 | 行为 |
|---|---|
| strict-view | 默认，只索引目录 A 中选定内容 |
| open-external-on-demand | 允许临时打开完整 AOSP 中的单个目标文件，不递归加入 source root |
| auto-expand | 高级选项，按模块依赖自动扩大 ProjectView |

auto-expand 必须显示预计增加的文件数并由用户确认。

必须诊断：

- 软链接循环、自引用或嵌套循环；
- dangling symlink；
- 同一物理文件通过多个链接重复暴露；
- 同一 RepoPath 指向多个不同内容；
- 一个 source mount 覆盖另一个 mount 的子目录；
- 任意别名无法自动推断 RepoPath；
- Windows junction、UNC 或大小写冲突；
- file-level symlink 数量过多导致扫描成本上升；
- 目录 A 中混入另一个 AOSP 版本；
- JSON 中引用的 source 在客户端不可见；
- out 仍保留服务器绝对路径；
- symlink target 在客户端不存在；
- Windows 无法表达 Linux symlink；
- 网络文件系统 mtime、文件通知和大小写语义不同；
- 服务器生成文件还没有同步到客户端。

一个 profile 默认只对应一个 canonical AOSP namespace 和一个 build target。如果目录 A 故意混入两个不同 Android release 的仓库，普通源码和文本跳转仍可工作，但 Soong module graph、AIDL generated mapping 和 compdb 不能再假设来自同一构建图。插件应降级为 mixed/source-only，除非用户提供分别隔离的 profile。

用户只选择部分仓库时采用三层索引：

1. workspace topology：记录目录 A、Source Mount、overlay 和优先级；
2. 模块图摘要：只加载这些路径所属模块及递归依赖；
3. 按需文件索引：打开文件时再读具体源码和 generated file。

### 性能边界

如果用户创建：

~~~text
A/frameworks -> /data/aosp-full/frameworks
~~~

那么从文件系统角度，A 仍然暴露整个 frameworks。插件可以通过 IDE exclude、Focused ProjectView 和按需索引减少 IDE 工作量，但不能消除网络文件系统或 IDE 自身对高层链接的目录枚举成本。

为了获得最好的加载速度，推荐链接 project 或更细粒度源码根，而不是直接链接 AOSP 顶层大目录。插件的向导应根据预计可见文件数给出提示，但不能禁止用户的目录布局。

不要把 Canonical AOSP Root 或整个 out/soong/.intermediates 加入 IDE 普通源码索引。

## 5. AIDL 跳转

### 有 out

可以提供：

- Java/Kotlin 接口类到原始 .aidl；
- Stub、Proxy、parcelable、enum 到原始声明；
- AIDL import 到定义文件；
- AIDL 到 Java/C++/NDK/Rust generated file；
- 读取 srcjars、AIDL include dirs、generated output 和 depfile；
- 按 build variant 区分同名接口和不同 API 版本。

### 没有 out

仍可提供：

- AIDL 文件内 package、import、类型和方法跳转；
- com.example.IFoo 到 com/example/IFoo.aidl；
- IFoo.Stub、IFoo.Proxy 到 IFoo.aidl；
- 基于 package 和目录约定的源码级跳转。

不能保证：

- 当前 product 是否真的使用该接口；
- 实际 aidl_interface variant；
- jarjar、稳定性和版本规则；
- generated file 真实存在。

无 out 的结果必须标记为 source heuristic。

## 6. JNI 双向跳转

### 结论

**可以做到 Java 与 C/C++ 的 JNI 互跳。常见静态导出和动态注册可以精确处理，复杂宏和运行时绑定需要降级。**

支持的 Java 到 C/C++ 模式：

1. 静态导出命名：

~~~cpp
Java_android_os_Foo_nativeBar(...)
~~~

2. 动态注册：

~~~cpp
static const JNINativeMethod methods[] = {
    {"nativeBar", "(I)V", (void*)nativeBar},
};
jniRegisterNativeMethods(env, "android/os/Foo", methods, NELEM(methods));
~~~

3. RegisterMethodsOrDie、RegisterNatives、jniRegisterNativeMethods；
4. JNI_OnLoad；
5. FindClass、GetMethodID、GetFieldID。

支持的 C/C++ 到 Java 模式：

- 从注册类字符串定位 Java 类；
- 从 JNINativeMethod.name 定位 Java native 方法；
- 从 FindClass、GetMethodID、GetFieldID 关联 Java 符号；
- 从导出函数名反解 Java 包名、类名、方法名；
- 解析 JNI descriptor。

不能完全保证：

- 宏生成并跨文件拼接的注册表；
- 复杂函数计算 class name；
- 运行时反射、动态加载；
- ART 内部特殊绑定；
- 条件编译和模板导致的多版本实现。

结果需要有 exact、inferred、unresolved 三个等级。C++ 类型系统和宏语义交给 clangd/IDE，插件只维护 JNI 跨语言边。

## 7. 无 out 与有 out 的能力矩阵

| 能力 | 无 out | 有匹配 out |
|---|---:|---:|
| 普通 Java/Kotlin 源码跳转 | 高 | 高 |
| AIDL 文件内跳转 | 中高 | 高 |
| Java 到 AIDL | 启发式 | 高 |
| Java 到 R/生成类 | 低 | 高 |
| Java 模块依赖 | 低 | 高 |
| C/C++ 普通跳转 | 中 | 高 |
| JNI 双向跳转 | 中 | 高 |
| generated C/C++ header | 低 | 高 |
| 当前 product/variant 过滤 | 低 | 高 |
| 构建目标反查 | 低 | 高 |

无 out 时进入 source-only 模式：

- 扫描 Android.bp/Android.mk 的源码声明；
- 扫描 Java/Kotlin package 和 import；
- 解析 AIDL；
- 扫描 JNI 静态符号、注册表和 Java native 方法；
- 使用目录约定推断 include roots；
- 对不确定结果标记来源和置信度。

## 8. 最终判断

| 要求 | 判断 |
|---|---|
| 不再运行现有脚本，安装插件后配置目录即可 | 可行 |
| Android Studio / IDEA | 可行，共享 JetBrains 客户端源码 |
| VS Code | 可行，独立 TypeScript 客户端 |
| Linux/macOS/Windows | 可行，engine 按平台打包 |
| 本机完整 AOSP，但只打开软链接目录 A | 可行；完整根目录仅作 fallback，不加入 IDE 索引 |
| 只挂载部分 AOSP 仓库和 out | 可行，必须使用仓库相对路径和 Path Mapping |
| Linux 软链接后整体挂载 | 可行，但要诊断 target 是否可见 |
| AIDL 跳转 | 可行，有 out 精确，无 out 降级 |
| Java/C JNI 互跳 | 可行，常见模式精确，复杂模式降级 |
| 一个相同安装包覆盖 AS、IDEA、VS Code | 不可行；共享 engine 和协议即可 |
| 无 out 时拥有完整 variant 语义 | 不可行，只能源码级或启发式 |

## 开发顺序

1. Go engine 的 profile、路径模型、metadata reader 和缓存；
2. JetBrains 配置界面、Java/generated source 导航；
3. AIDL parser 和 Java/AIDL 双向跳转；
4. compdb、clangd 和 JNI cross-language index；
5. VS Code 客户端；
6. 可选的 Linux ide_query 或远程构建桥接服务。
