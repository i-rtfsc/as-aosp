# AOSP IDE 源码与模块边界

> 文档职责：固定仓库目录、依赖方向和模块 owner。物理目录即当前 `dev-ide` 仓库根。

## 1. 顶层目录

| 目录 | owner | 不承载 |
|---|---|---|
| `core/` | Go 索引领域和应用 | IDE API、UI、平台特例 |
| `cmd/` | CLI/daemon composition | 解析算法和领域状态 |
| `clients/jetbrains/` | Kotlin JetBrains adapter | Soong JSON reader |
| `clients/vscode/` | TypeScript VS Code adapter | Go domain duplicate |
| `protocol/` | RPC/profile schema | 业务实现 |
| `remote-execution/` | 可选 Linux 构建桥接 | 核心索引 owner |
| `packaging/` | 多平台二进制和插件打包 | runtime 逻辑 |
| `scripts/` | 仓库质量 gate 和薄维护入口 | 领域解析和产品逻辑 |
| `testdata/` | fixture、metadata 和目录拓扑样本 | 生产源码 |
| `docs/` | 文档权威入口 | 代码事实快照 |

## 2. Core bounded contexts

| Package | 责任 |
|---|---|
| `core/model` | RepoPath、profile、symbol、diagnostic 基础值 |
| `core/topology` | Workspace View、SourceMount、provider priority、cycle/conflict |
| `core/pathmap` | View/Local/External 与 RepoPath 转换 |
| `core/metadata` | Soong JSON、module-info、compdb 和 ide_query reader |
| `core/modules` | ModuleCatalog、source roots、generated/classpath facts 和 dependency closure |
| `core/buildfiles` | Android.bp/Android.mk fallback scanner |
| `core/generated` | srcjar、generated file cache 和 depfile |
| `core/aidl` | AIDL parser、resolver 和 generated mapping |
| `core/jni` | Java native、C++ registration 和 cross-language links |
| `core/graph` | ModuleGraph、dependency closure 和 ProjectView |
| `core/index` | snapshot、incremental update、symbol query |
| `core/rpc` | versioned JSON-RPC server/DTO |
| `core/app` | scan/query orchestration |

依赖方向：

~~~text
model <- topology/pathmap/metadata/buildfiles/generated/aidl/jni
      <- graph/index
      <- rpc/app
      <- cmd
clients -> rpc only
~~

## 3. 禁止的反向依赖

- `core` 不依赖 `clients`；
- `clients` 不直接读取 Soong JSON；
- `cmd` 不复制 core 解析逻辑；
- `remote-execution` 不拥有 ModuleGraph authority；
- `testdata` 不从用户真实 AOSP 读取文件；
- 生成缓存不写入 AOSP checkout。

## 4. 产物边界

- `aosp-index`：诊断、单次查询和 `serve --stdio` JSON-RPC daemon；
- JetBrains plugin：平台插件包；
- VS Code extension：VSIX；
- engine binaries：按 OS/arch 携带，不进入 AOSP out。
