# AOSP IDE 剩余任务分解

## 1. AOSP-IDE-CORE-01

1. 加入 Go module 和 Makefile。
2. 实现 profile schema 和严格字段校验。
3. 实现 RepoPath/ViewPath/LocalPath/ExternalPath。
4. 实现 WorkspaceTopology 与 SourceMount priority。
5. 实现 cycle、duplicate、shadowed 和 conflict diagnostics。
6. 实现跨平台 PathMapper。
7. 加入 topology fixture 和测试。

下一验收：`make test-core` 通过，所有结果以 RepoPath 为身份。

## 2. AOSP-IDE-CLI-01

1. 实现 `aosp-index doctor`。
2. 实现 `aosp-index scan`。
3. 输出 profile、topology、scope 和 diagnostics JSON。
4. 增加 smoke fixture。

下一验收：没有真实 AOSP 时也能完成完整 fixture smoke。

## 3. 开发约束

- 先改 bounded-context design，再改代码；
- 不在 Phase 0 引入 IDE API；
- 不扫描用户完整 AOSP；
- 远程能力统一使用 `remote-execution` 命名；
- 不在 dev-ide 恢复旧 Gradle/CMake 目录；
- 不以裸编译命令作为最终验证报告。
