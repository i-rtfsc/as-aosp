# AOSP IDE 代码规范

## 1. 通用规则

- 一个 package/file 只承担一个可定位的职责簇；
- facade 只组合，不重新拥有子模块算法；
- domain 使用稳定值类型，不传递 IDE、HTTP、shell 或 OS 对象；
- 错误必须带可编程 code、来源和上下文，不能只返回 free-text；
- 路径、profile、metadata 和 RPC 字段必须有边界校验；
- 大仓扫描必须有 scope、depth、budget、cancel 和 backpressure；
- 日志不输出用户真实路径、凭据或完整编译命令中的敏感参数。

## 2. Go

- package 按 bounded context 命名，不设 `utils`/`common` 垃圾包；
- `cmd/` 只做参数解析、IO 和 composition；
- `_test.go` 与被测 package 同目录；跨 package fixture 进入 `testdata/`；
- 文件系统访问通过小接口注入，测试不得依赖真实 AOSP；
- 所有 goroutine 有 owner、cancel 和退出路径；
- daemon 请求有最大消息、超时和并发边界。

## 3. Kotlin/TypeScript

- 客户端只做 IDE API 适配；
- engine RPC DTO 不与 IDE PSI/VS Code URI 混用；
- UI 状态来自 engine snapshot/diagnostics，不自行推断 Current；
- 进程重启后客户端必须重新 initialize，不复用旧 session 状态。

## 4. 文件和生成物

- 生成代码、srcjar 解压物、compdb overlay 进入缓存或 `out/`，不进入源码目录；
- 任何跨语言 schema 先在 `protocol/` 定义，再生成 DTO；
- 不手工维护大量按 AOSP 版本复制的路径清单。
