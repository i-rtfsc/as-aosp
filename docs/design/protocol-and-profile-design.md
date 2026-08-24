# RPC、Profile 与 Diagnostics 设计

> 状态：Partial；profile JSON 和 CLI 已实现，首版 stdio daemon 进入交付。

## 1. Profile

Profile 至少包含：

- schemaVersion、id；
- workspaceViewRoot、canonicalRepoRoot、outRoot；
- buildTarget product/release/variant；
- scope include/exclude；
- sourceMounts；
- pathMappings；
- visibilityPolicy；
- buildIntegration mode。

本机绝对路径只进入用户本地配置，不进入共享 fixture。

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
    "aidlNavigation": false,
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

请求携带当前 source 的 `RepoPath`、`language` 和 `content`，以及同一 focused
view 中可见的候选 source documents。engine 使用 `core/jni` 建立临时 source-only
index，返回 `Link[]`；不读取客户端不可见的 canonical root，也不要求 `out`。
JetBrains 和 VS Code 只消费同一个 peer schema，不各自复制 descriptor、静态导出
或动态注册解析。首版是按需查询，长期 snapshot 和增量失效属于后续 index slice。

后续索引 slice：

```text
aosp/index/status
aosp/index/rebuild
aosp/definition
aosp/references
aosp/aidl/resolve
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
