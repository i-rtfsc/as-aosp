# AOSP Workspace 跳转方案调研与架构调整

> 复核日期：2026-08-25。调研对象：AOSP 官方 AIDEGen、Soong compdb/CLion generator、
> `ide_query`、本地旧 Gradle/CMake 方案，以及 GitHub 上的 AOSP VS Code 配置项目。

## 1. 结论先行

当前插件“能启动、能按请求扫描、能做部分 AIDL/JNI 跳转”，但还没有达到用户旧 Gradle
工程“打开后 Java 立即可跳”的能力。根因不是单纯索引速度，而是没有把 AOSP 构建事实
投影给 IDE 的原生语言索引器：

```text
当前实现：文件扫描/猜路径 -> 自定义少量跳转
正确方向：Android.bp/Soong metadata/AIDEGen facts
        -> module + source roots + dependencies + jars/srcjars
        -> IDE/Java index；compdb -> clangd；AIDL/JNI -> shared edges
```

因此后续必须从“请求时扫描优先”调整为“构建事实驱动的 Project Model/Workspace Index
优先，请求时扫描仅作为 source-only fallback”。

## 2. AOSP 官方方案证据

### 2.1 AIDEGen：生成 IDE project，而不是让 IDE 猜

AOSP 官方 AIDEGen README 明确将自身定义为：自动为 Android Studio、IntelliJ、Eclipse、
CLion 和 VS Code 配置项目依赖。典型使用方式是：

```text
source build/envsetup.sh && lunch <TARGET>
aidegen Settings framework
aidegen <module> -i c
aidegen -a
```

关键点：

- Java 项目按 module/target 选择并生成 IntelliJ 项目文件；
- `-d/--depth` 控制依赖深度；
- `-a/--android-tree` 表示整个 Android tree；
- `-s/--skip-build` 明确允许跳过构建，但官方提示可能缺少 jar、R/AIDL/logtags generated
  文件，导致 IDE 红线；
- AIDEGen 读取 module-info/Java dependency 等构建结果，生成 `.idea/modules.xml`、多个
  `.iml` 和 dependencies module，而不是把所有目录平铺成一个普通源码目录。

官方 AIDEGen 的 IntelliJ generator 代码显示：

- 每个 ProjectInfo 可以生成独立 `.iml`；
- `.iml` 中写入真实 source folders、generated source、srcjar 和依赖 module；
- `ProjectSplitter` 会处理多个项目共享 source root 的重复问题；
- 主项目的 `.idea/modules.xml` 会引用依赖 `.iml`；
- 这是旧 Gradle 工程能快速跳转的关键机制之一。

### 2.2 Soong compdb：C/C++ 依赖编译数据库

AOSP Soong 官方文档说明：

```text
export SOONG_GEN_COMPDB=1
make nothing
```

生成：

```text
out/soong/development/ide/compdb/compile_commands.json
```

该数据库供 YouCompleteMe、libclang、clangd、CLion 等消费。Soong 还可以生成 CLion
项目：

```text
SOONG_GEN_CMAKEFILES=1
make <target>
```

输出位于：

```text
out/development/ide/clion/<module>/<target-arch-os>/CMakeLists.txt
```

官方文档还说明可以通过 super project 聚合多个 CMake 子项目。这说明 C/C++ 的正确
入口是 compdb/CMake project，而不是扫描 C++ 文件再猜 include path。

### 2.3 `ide_query`：构建产物查询层

AOSP `build/make/tools/ide_query` 负责分析 build artifacts，使用：

- `module_bp_java_deps.json`；
- `compile_commands.json`；
- `module-info.json`；
- C/C++ analyzer；
- build target 和生成文件信息。

代码中还明确记录 compdb 可能过期、只包含当前构建目标，以及对多个 build unit 的选择
逻辑。因此 plugin 应显示 metadata freshness/variant ambiguity，不能把一条 compdb 记录
宣传成全部 AOSP C++ 真相。

## 3. 本地旧方案为什么能立刻跳

旧 Gradle/CMake 脚本虽然不是理想的长期方案，但它做对了几件关键事情：

1. 把 AOSP Java source roots 显式写进 Android Studio/IntelliJ project；
2. 把 `out` 下 jars、srcjars、generated Java 接入 classpath/source roots；
3. 把模块依赖关系写成 IDE 能理解的 project/module 关系；
4. C/C++ 侧生成/复制 CMakeLists 或 compdb，让 CLion/clangd 获得编译参数；
5. 打开的是“生成后的 IDE project”，不是只有一个完整 AOSP 普通目录。

所以它快不是因为 Gradle 本身更懂 AOSP，而是因为脚本已经提前完成了 AOSP → IDE model
转换。当前插件缺少的正是这个转换层。

## 4. GitHub 方案对比

### 4.1 `amezin/aosp-vscode`

该项目的 README 采用非常直接的方案：

```text
1. build AOSP
2. clone 到 AOSP 根的 .vscode
3. python .vscode/generate_compdb.py out/build-<target>.ninja
4. 用 VS Code 打开 AOSP 根
```

其 `c_cpp_properties.json` 指向：

```json
"compileCommands": "${workspaceFolder}/compile_commands.json"
```

其 `generate_compdb.py` 从 Ninja rule/build 行提取 clang 命令，生成 compile_commands。
这是纯 C/C++ 方案，不能解决 Java/AIDL，但验证了 VS Code/clangd 的最小可靠入口就是
compdb，而不是自定义 C++ parser。

### 4.2 AOSP 官方 CLion/VS Code 方向

AIDEGen README 将 CLion 和 VS Code 都作为目标 IDE，但底层仍按语言分别生成：

- Java/IntelliJ：`.idea`、`.iml`、dependencies、srcjars；
- C/C++：CLion CMake 或 VS Code compdb；
- AOSP 构建：Soong/`ide_query` 产出事实；
- IDE：消费标准 project model/LSP/clangd。

这比“一个自定义 LSP 统一解释所有语言”更符合现有 IDE 生态。

## 5. 对当前架构的调整

### 5.1 保留的部分

- Go engine 和 `aosp-index serve --stdio`；
- `RepoPath`、Workspace View、SourceMount、PathMapper；
- metadata reader、AIDL parser/index、JNI index；
- compdb overlay 和跨平台 engine；
- JetBrains/VS Code RPC client、图形 profile UI。

### 5.2 必须新增的核心层

```text
core/buildfiles/
  Android.bp lexer/parser（受限静态 facts，不复刻完整 Soong evaluator）
  Android.mk fallback facts
  glob/source expansion

core/modules/
  ModuleCatalog
  source roots
  dependency closure
  generated/source archive ownership
  module selection

core/index/
  IndexCoordinator
  persistent snapshot
  progress/status
  invalidation/freshness

clients/jetbrains/
  module/source-root projection facade
  generated jars/srcjars projection

clients/vscode/
  Java fallback only as degraded mode
  clangd compdb integration
```

### 5.3 索引优先级

```text
Build-backed：Soong metadata + module-info + compdb + generated/source archives
       ↓
BP-backed：Android.bp 静态 module/source/dependency facts
       ↓
Source-only：文件扫描、package/import、AIDL/JNI fallback
```

有 out 时，必须优先 build-backed；没有 out 时，Android.bp 静态 facts 仍应建立 source roots
和 module catalog；只有没有 BP/metadata 时才降级到文件猜测。

## 6. 多 IDE 的正确职责

| 语言 | Android Studio/IDEA | VS Code | 共享 engine |
|---|---|---|---|
| Java/Kotlin | module/source roots、jars、srcjars 交给 IDE index | Java LS + source-root/workspace 配置 | module facts、依赖、generated mapping |
| AIDL | AIDL PSI/reference + generated peer | LSP Definition/References | AIDL parser/resolver/peers |
| C/C++ | CLion/CMake 或 clangd compdb | clangd compdb | compdb/path mapping |
| JNI | custom cross-language action/gutter | DefinitionProvider/command | JNI link index |

客户端不应该重新实现 Soong、完整 Java compiler 或 C++ semantic parser。

## 7. 新的索引状态和 UI

```text
Idle
Discovering Workspace
Reading Android.bp
Loading Soong metadata
Building Module Graph
Indexing Java/AIDL/JNI
Ready
Partial
Stale
Error
```

侧边栏必须显示：

- stage、completed/total、percent；
- modules、source roots、jars/srcjars、AIDL symbols、JNI links；
- out/metadata freshness；
- missing/ambiguous modules；
- 当前 active profile 和 selected module closure。

## 8. 结论

“打开目录后立刻所有 Java 跳转”不是单纯等待索引，而是提前生成/投影 IDE 可理解的
module model。AOSP 官方 AIDEGen、Soong compdb/CLion generator 和你旧 Gradle/CMake
脚本都证明了这一点。当前工程下一阶段应优先实现：

1. Android.bp/metadata ModuleCatalog；
2. Java source roots/classpath/generated projection；
3. IndexCoordinator progress/status；
4. JetBrains project model projection；
5. VS Code Java LS/clangd configuration。

在这些完成前，当前插件只能称为 source-only navigation prototype，不能宣称“打开完整
AOSP 后所有代码立即可跳”。
