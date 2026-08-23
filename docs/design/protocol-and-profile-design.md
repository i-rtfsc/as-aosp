# RPC、Profile 与 Diagnostics 设计

> 状态：Target；Phase 0 先实现 profile JSON 和 CLI 输出，daemon RPC 后续接入。

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

## 2. RPC 方法

```text
aosp/initialize
aosp/profile/validate
aosp/topology/discover
aosp/topology/validate
aosp/index/status
aosp/index/rebuild
aosp/definition
aosp/references
aosp/aidl/resolve
aosp/jni/peers
aosp/path/explain
aosp/diagnostics/list
```

## 3. Diagnostics

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

## 4. Compatibility

- unknown JSON fields 忽略；
- 缺少新字段使用明确默认值；
- schemaVersion 不兼容时 fail closed；
- engine 返回 protocolVersion；
- 客户端 initialize 时协商能力；
- RepoPath 永远比 absolute path 更优先。
