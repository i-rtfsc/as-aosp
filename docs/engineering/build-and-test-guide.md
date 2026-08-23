# AOSP IDE 构建与测试指南

## 1. 当前工程入口

Phase 0 采用仓库根 Makefile 作为统一入口。不得把裸 `go build`、`go test` 作为最终报告入口。

| 目标 | 入口 |
|---|---|
| 查看帮助 | `make help` |
| 格式检查 | `make fmt-check` |
| Go 静态检查 | `make vet` |
| Commit message 检查 | `make commit-message-check` |
| Core 测试 | `make test-core` |
| Protocol 测试 | `make test-protocol` |
| 全部本地检查 | `make check` |
| CLI smoke | `make smoke` |
| JetBrains 验证 | `make test-jetbrains` |
| VS Code 验证 | `make test-vscode`（后续添加） |

## 2. Phase 0 最低验收

- `make fmt-check` 通过；
- `make vet` 通过；
- `make test-core` 通过；
- `make smoke` 能解析 fixture profile 并输出 topology；
- 真实 AOSP 不是测试输入，除非命令显式指定路径；
- 测试输出不暴露用户绝对路径或敏感 metadata。

JetBrains 首次构建默认下载锁定的 IntelliJ Platform SDK。开发者也可以复用本机已安装 IDE：

```text
JAVA_HOME_21=<jdk-21> \
INTELLIJ_PLATFORM_PATH=<idea-or-android-studio-contents> \
make test-jetbrains
```

`INTELLIJ_PLATFORM_PATH` 只属于本机环境，不能写入版本控制。Wrapper 锁定 Gradle 9.1.0，
Linux、macOS 和 Windows 使用相同 Gradle/Kotlin 源码；具体 IDE build 兼容仍由 verifier 证明。

## 3. 测试边界

| 测试 | 可以证明 | 不能证明 |
|---|---|---|
| topology fixture | mount、overlay、cycle、conflict 算法 | 真实网络 FS 性能 |
| metadata fixture | JSON schema 和容错 | 当前 AOSP 全仓兼容 |
| source-only fixture | fallback parser | Soong variant 正确性 |
| CLI smoke | 组合和序列化 | IDE Project Model |
| JetBrains test | PSI/Workspace adapter | VS Code 行为 |
| VS Code test | LSP/client adapter | Android Studio 行为 |

## 4. 失败报告模板

```text
入口：make test-core
结果：FAIL
失败 fixture：...
首个根因：...
复现性：stable/intermittent
与本次变更关系：introduced/pre-existing/unknown
未运行项及原因：...
```
