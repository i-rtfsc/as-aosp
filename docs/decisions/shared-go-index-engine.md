# 共享 Go 索引引擎决策

> 状态：Decision；接受日期：2026-08-23。

## 1. Decision

共享索引核心使用 Go，实现为可执行文件和可复用 package。JetBrains Kotlin 客户端与 VS Code
TypeScript 客户端通过版本化 JSON-RPC 调用同一引擎。

## 2. Why

- 单文件跨平台二进制便于随插件分发；
- 标准库覆盖 JSON、zip/srcjar、路径、进程和缓存；
- AOSP `ide_query` 也是 Go，数据模型容易对照；
- 避免 JVM 和 Node 各自复制 Soong、AIDL、JNI 与路径逻辑；
- 用户不需要预装 Python 或额外 runtime。

## 3. Alternatives considered

### Kotlin-only

JetBrains 集成直接，但 VS Code 需要重复核心或启动 JVM。Rejected。

### TypeScript-only

VS Code 集成直接，但 JetBrains 和大规模文件索引复用较弱。Rejected。

### Python

开发快，但运行时分发、环境差异和当前历史失败已经证明不适合作为硬依赖。Rejected。

### Rust

性能和类型安全良好，但当前团队已有 AOSP Go 数据源，首版开发和跨平台打包成本更高。Deferred。

## 4. Consequences

- Go package 不依赖 IDE API；
- 客户端不得直接解析 Soong metadata；
- RPC schema 必须版本化；
- C++ 完整语义继续委托 clangd，不在 Go 中重写编译器。
