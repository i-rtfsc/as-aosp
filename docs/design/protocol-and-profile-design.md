# RPC、Profile 与 Diagnostics 设计

> 状态：Partial；profile JSON 和 CLI 已实现，首版 stdio daemon 进入交付。

## 1. Profile

Profile 至少包含：

- schemaVersion、id；
- workspaceViewRoot、canonicalRepoRoot、outRoot；
- buildTarget product/release/variant；
- scope include/exclude；
- targets：用户定义的 TargetSpec 集合，可按 artifact、module、source root 组合；
- sourceMounts；
- pathMappings；
- visibilityPolicy；
- buildIntegration mode。

本机绝对路径只进入用户本地配置，不进入共享 fixture。

`targets` 必须至少包含一个 TargetSpec。客户端只请求 ProjectGraph；ModuleCatalog/closure
仅供 engine 内部和诊断命令使用，不作为客户端投影 API。缺少 targets 的 profile 无效。

## 2. 进程与目录边界

共享协议分三层：

```text
protocol/jsonrpc
  JSON-RPC 2.0 envelope、Content-Length framing、错误和 server loop

protocol/aosp
  AOSP 方法名、参数、结果、版本和 capability

cmd/aosp-index
  CLI commands and `serve --stdio` lifecycle wiring
```

`core` 不依赖 `protocol`，JetBrains/VS Code 客户端不解析 CLI 文本。

## 3. Transport

首版使用 stdin/stdout JSON-RPC 2.0，并采用 LSP 兼容 framing：

```text
Content-Length: <UTF-8 byte length>\r\n
\r\n
<JSON payload>
```

约束：

- header 名大小写不敏感；
- 必须且只能有一个有效 `Content-Length`；
- payload 上限首版为 16 MiB；
- stdout 只写协议帧，日志写 stderr；
- request id 保留 JSON string/number/null 原始表示；
- notification 没有 response；
- malformed frame 终止当前进程，合法 JSON-RPC 错误返回 error response；
- 首版串行处理 request，后续并发必须保持同一 id 的单响应不变量。

## 4. Lifecycle

客户端必须首先调用 `aosp/initialize`。服务端返回协议版本和 capability。

```text
created -> initialized -> shutdown-requested -> exited
```

- initialize 前的业务方法返回 `SERVER_NOT_INITIALIZED`；
- 重复 initialize 返回 `INVALID_REQUEST`；
- `aosp/shutdown` 返回 `null`，之后拒绝新业务请求；
- stdin EOF 或 shutdown 后 EOF 退出；
- 首版不支持跨进程 session 恢复。

## 5. 版本与 capability

首版协议版本为 `1.0`。initialize 结果至少包含：

```json
{
  "protocolVersion": "1.0",
  "serverInfo": {
    "name": "aosp-index",
    "version": "dev"
  },
  "capabilities": {
    "profileValidation": true,
    "topologyDiscovery": true,
    "topologyValidation": true,
    "moduleCatalog": true,
    "aidlNavigation": true,
    "jniNavigation": true
  }
}
```

客户端传入不兼容 major version 时 initialize fail closed。minor version 可以向后兼容。

## 6. RPC 方法

首个可独立验收的 daemon slice：

```text
aosp/initialize
aosp/shutdown
aosp/profile/validate
aosp/topology/discover
aosp/topology/validate
```

Build-Facts ModuleCatalog 查询 slice：

```text
aosp/modules/catalog
aosp/modules/get
aosp/modules/closure
aosp/artifacts/catalog
aosp/projects/graph
```

`aosp/projects/graph` 接收非空 TargetSpec 集合，返回统一 ProjectGraph。客户端不得
独立 flatten ModuleCatalog。

`aosp/artifacts/catalog` 返回按 kind/name/path 稳定排序的 artifact facts，并支持可选
kind/query 过滤，用于两个客户端的图形化 Target selector。

三个方法都显式携带 `outRoot`，不依赖 daemon 当前工作目录或
进程级全局 profile。`catalog` 返回按模块名排序的摘要，供 UI 选择模块和
显示规模；`get` 返回一个模块的完整 facts；`closure` 返回选中模块及其传递
依赖的完整模块列表，顺序按模块名稳定排序。

Out Root 缺失、Build Facts JSON 损坏或模块不存在属于业务诊断，结果仍是
成功的 JSON-RPC response。服务端不得回退扫描整个 Workspace View 来伪造
Build-Facts 结果。首版每次请求按 Out Root 下 Build Facts 文件的路径、大小和修改时间
判断是否复用 daemon 内存快照；文件被原子替换后必须重新加载。

IndexCoordinator 状态方法：

```text
aosp/index/start
aosp/index/status
aosp/index/rebuild
aosp/index/cancel
aosp/index/projection
```

`start` 和 `rebuild` 接收完整 `profile`，并返回 coordinator snapshot。
状态至少包含 `state`（`idle`、`running`、`ready`、`partial`、`stale`、
`error`、`cancelled`）、当前 `stage`、`completedStages`、`totalStages`、
`progress`、`moduleCount`、`generatedCount` 和 diagnostics。首版 coordinator
在 daemon 进程内运行并以有界阶段顺序执行；即使某个 metadata reader 缺失，
也必须保留已完成阶段并进入 `partial`，不能伪装成 `ready`。`cancel` 只影响
当前 profile 的 coordinator，不删除旧 snapshot。

`projection` 由 IDE 客户端在本机 source/generated/library projection 完成后调用，
携带 `profileId`、各类 root/library counts 和 diagnostics。daemon 将 projection
diagnostics 合并到 snapshot；任一 projection diagnostic 会把状态保持为 `partial`。

Generated source materialization：

```text
aosp/generated/materialize
```

请求携带 `outRoot` 和本机 `cacheRoot`。engine 只读取 Out Root 中由 ProjectGraph
明确引用的 `.srcjar`，通过 `core/generated` 按 archive hash 安全解压，并返回已发布的
source roots 与 diagnostics。结果不写入 AOSP 源码或 out；重复请求复用相同 hash 目录。

Rust project materialization：

```text
aosp/rust/project
```

请求携带 Workspace View Root、Out Root、cacheRoot 和 pathMappings。engine 把 Soong
`rust-project.json` 中的源码、out、sysroot 和 proc macro 路径归一化到本机，并只在
IDE cache 中发布可供 rust-analyzer 使用的 project JSON。

进程入口统一为一个跨平台二进制：

```text
aosp-index doctor|scan|metadata|aidl|jni
aosp-index serve --stdio
```

插件只配置 `aosp-index` executable，固定追加 `serve --stdio` 参数；不再分发独立
独立 daemon 文件。

首个 JNI navigation slice：

```text
aosp/jni/peers
```

请求优先携带 `workspaceViewRoot + currentPath`。engine 只在这个用户明确配置的
Workspace View 内扫描 Java/Kotlin/C/C++，跳过 `.git`、`.repo`、`out` 和
`.aosp-ide`，使用 `core/jni` 建立 source-only index 并返回 `Link[]`。为支持
受控客户端或远程桥接，协议仍接受显式 `javaFiles/cppFiles`，但客户端不应把整棵
focused view 序列化进普通导航请求。不读取客户端不可见的 canonical root，也不要求
`out`。JetBrains 和 VS Code 只消费同一个 peer schema，不各自复制 JNI 解析。
daemon 会话内按规范化 Workspace View Root 保存 JNI snapshot。每次请求遍历可见
候选文件并按 view-relative path 与内容生成指纹；指纹未变时复用已经构建的
`core/jni.Index`，变化时整体重建。snapshot 不写磁盘，也不跨进程恢复。目录遍历只
阻止当前祖先链上的 symlink cycle；指向同一物理目录的两个合法 view alias 保留为
不同 RepoPath。内容 hash、单文件增量更新和持久 cache 属于后续 index slice。

首个 AIDL resolve slice：

```text
aosp/aidl/resolve
```

请求携带 `workspaceViewRoot`、当前 AIDL `currentPath` 和光标处的 `name`。engine
只扫描 Workspace View 内可见的 `.aidl`，使用 `core/aidl` 的 package/import/type
resolver 返回 `Resolution` 与精确 source range；固定排除目录和 symlink cycle 规则
与 JNI workspace 扫描一致。首版不返回 generated peers，也不启用 references。

首个 AIDL generated peer slice：

```text
aosp/aidl/peers
```

请求同样携带 `workspaceViewRoot`、`currentPath` 和 `name`。engine 先解析当前
source declaration，再调用 `core/aidl.GeneratedPeers` 生成 Java/C++/NDK 命名候选。
候选路径相对 Workspace View 且实际存在时升级为 `exact`；只有命名规则时保持
`inferred`，不声明 generated 文件一定存在。srcjar/out metadata 接入后沿用同一
`GeneratedPeer` schema，不改变客户端路径和 confidence 语义。

首个 AIDL references slice：

```text
aosp/aidl/references
```

请求携带 `workspaceViewRoot`、当前 `currentPath` 和 `name`。engine 先解析目标
类型，再扫描可见 AIDL 的 import、返回类型、参数类型和字段类型。parser 保存
`typeRange` 时返回具体类型 token，`confidence` 为 `exact`；无法恢复 token span 时
才降级到所属声明 range 和 `inferred`，不改变 references DTO。

首个 compdb overlay slice：

```text
aosp/compdb/overlay
```

请求携带 `outRoot`、`workspaceViewRoot` 和显式
`pathMappings`。engine 读取
compilation database，映射 `directory`、`file`、`output` 以及 argv 中的构建机绝对
路径；`@response-file` 会被读取、转换并作为 overlay-side response file 返回。原始
out/compdb 不被修改。不能证明映射的参数保留原值并标记 `COMPDB_PATH_UNMAPPED`，
同一 `file` 的多条 variant 全部保留并标记 `COMPDB_VARIANT_DUPLICATE`，
response file 无法读取时标记 `COMPDB_RESPONSE_READ_FAILED`，不会猜测本机路径。
Portable Out 中的构建机绝对路径仍必须通过显式 mapping。

后续索引 slice：

```text
aosp/index/status
aosp/index/rebuild
aosp/definition
aosp/references
aosp/path/explain
aosp/diagnostics/list
```

### 6.1 Profile validate

请求直接携带 `model.Profile`，结果返回稳定 diagnostics。该方法不读取源码树。

### 6.2 Topology discover

参数包含 workspace view root、可选 canonical root 和 max depth。该方法只读本机可见目录，
不自动扫描 canonical AOSP 全树，也不写 profile。

### 6.3 Topology validate

请求直接携带 `WorkspaceTopology`，用于客户端保存配置前检查 mount conflict 和缺失字段。

## 7. JSON-RPC errors

标准错误：

| code | 含义 |
|---:|---|
| `-32700` | JSON parse error |
| `-32600` | invalid request 或 lifecycle 不合法 |
| `-32601` | method not found |
| `-32602` | invalid params |
| `-32603` | internal error |

业务诊断不放入 JSON-RPC error。能完成方法调用时，profile/topology 问题作为 result 中的
`diagnostics` 返回。

## 8. Diagnostics

每条诊断包含：

```text
code
severity
stage
repoPath/viewPath/localPath/externalPath, when known
origin
confidence
message
suggestedAction
```

诊断 code 必须稳定，message 可以改进。客户端不应按 message 文本分支。

## 9. Compatibility

- unknown JSON fields 忽略；
- 缺少新字段使用明确默认值；
- schemaVersion 不兼容时 fail closed；
- engine 返回 protocolVersion；
- 客户端 initialize 时协商能力；
- RepoPath 永远比 absolute path 更优先。

## 10. Acceptance

- framing fixture 能处理分片读取、连续消息和 UTF-8 byte length；
- malformed header、超限 payload 和非法 JSON 有稳定错误；
- initialize 前业务请求被拒绝；
- initialize、profile validate、topology discover/validate 和 shutdown 可通过 stdio smoke；
- `go test -race ./protocol/...` 不报告 data race；
- 不声称索引、AIDL 或 JNI RPC capability 已启用。
