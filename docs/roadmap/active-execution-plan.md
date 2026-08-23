# AOSP IDE 剩余任务分解

## 1. AOSP-IDE-RPC-01

1. 实现 LSP 兼容 `Content-Length` framing。
2. 实现 JSON-RPC request/response/error envelope。
3. 实现 initialize/shutdown lifecycle 和版本协商。
4. 暴露 profile validate、topology discover/validate。
5. 增加 daemon stdio integration test 与 Make smoke。

下一验收：`make rpc-smoke` 和 `go test -race ./protocol/...` 通过。

## 2. AOSP-IDE-JB-01

1. 固定 251 baseline 并创建 Kotlin/IntelliJ Platform 工程边界。
2. 实现 Kotlin JSON-RPC framing 与 initialize/shutdown client。
3. 加入插件 descriptor、单元测试和 `buildPlugin` 验证。
4. 后续实现 engine process、profile UI 和 Workspace Model。

下一验收：`make test-jetbrains` 通过并生成可安装插件 zip。

## 3. 开发约束

- 先改 bounded-context design，再改代码；
- 不在 Phase 0 引入 IDE API；
- 不扫描用户完整 AOSP；
- 远程能力统一使用 `remote-execution` 命名；
- 不在 dev-ide 恢复旧 Gradle/CMake 目录；
- 不以裸编译命令作为最终验证报告。
