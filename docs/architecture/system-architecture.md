# AOSP IDE 系统总体架构

> 文档职责：说明逻辑层、进程边界、部署形态和数据流。字段和算法见 `../design/`。

## 1. 架构结论

1. Shared Index Engine 是 AOSP 事实和索引查询的唯一核心 owner。
2. JetBrains 和 VS Code 只是客户端，不直接解析 Soong JSON，不拥有另一套符号索引。
3. Workspace Topology 负责“用户打开的目录”和“真实 AOSP/构建目录”的映射。
4. IDE ProjectView 是索引图的投影，不是 AOSP 构建图的替代品。
5. C/C++ 类型语义由 clangd/IDE 负责，JNI 只作为共享跨语言边。
6. 远程执行服务是可选的 Linux 构建辅助，不是核心索引必需进程。

## 2. 逻辑架构

~~~text
L1 Client Adapters
  JetBrains Plugin / VS Code Extension
  UI, PSI/LSP, IDE project model, clangd integration

L2 Index Application
  profile, scan, query, cache, diagnostics, RPC

L3 Domain Index
  WorkspaceTopology, RepoPath, SourceProvider
  Soong metadata, generated files, AIDL symbols, JNI links

L4 External Sources
  AOSP source, out, srcjar, module metadata, compdb, optional remote execution
~~~

依赖方向只能从 L1 到 L2/L3，再由 L3 读取 L4。L4 不反向依赖 IDE 客户端。

## 3. 进程视图

### Current / Phase 0

单次 CLI 直接调用 Go package，便于验证路径和 topology。不存在 IDE plugin 或远程服务 Current 能力。

### Target

~~~text
Android Studio / IDEA ─┐
                       ├─ stdio JSON-RPC ─> aosp-index serve --stdio
VS Code Extension ─────┘                         │
                                                ├─ source view
                                                ├─ out metadata
                                                └─ optional remote-execution bridge
~~~

engine 崩溃不应破坏 IDE；客户端负责重启、重新初始化和状态投影。

## 4. 关键数据流

1. 客户端提交 Workspace View Root 和 profile。
2. engine 发现 SourceMount，构建 active provider 表。
3. engine 读取 metadata；不可用时执行 scope 限制的源码扫描。
4. engine 生成 ModuleGraph、SymbolIndex、AidlIndex 和 JniIndex。
5. 客户端请求 Definition/References/Symbols。
6. engine 返回 RepoPath 和来源/置信度；客户端映射回 IDE URI。
7. generated source 写入缓存，只有 ProjectView 需要的文件进入 IDE。

## 5. 部署形态

| 形态 | 状态 | 说明 |
|---|---|---|
| 本机 CLI + 源码 view | Current target | Phase 0 验证形态 |
| JetBrains + 本机 engine | Target | 首个用户可见客户端 |
| VS Code + 本机 engine | Target | 第二客户端 |
| 客户端 + Linux remote-execution | Target | 解决 Linux-only 查询和远程构建 |

## 6. 负能力边界

- engine 不拥有 AOSP build graph 的全部求值能力；
- 客户端不保证完整 Android resource/Kotlin 项目体验；
- source-only 不证明当前 product/variant；
- 一条 compdb 记录不证明全部 variant；
- remote-execution 不属于第一阶段完成条件。
