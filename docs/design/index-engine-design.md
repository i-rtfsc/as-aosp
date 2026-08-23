# AOSP Index Engine 设计

> 状态：Target；实现顺序从 topology/pathmap 开始，暂不接 IDE。

## 1. 责任

engine 负责 profile、topology、metadata、fallback scan、generated cache、module graph、AIDL/JNI index 和 RPC 查询。它不负责 IDE UI、完整 C++ 类型系统或 AOSP build graph 求值。

## 2. Index modes

| Mode | 输入 | 结果边界 |
|---|---|---|
| `source-only` | active source mounts | 源码级/启发式 |
| `metadata-partial` | 部分 Soong metadata | 部分精确、部分降级 |
| `build-backed` | 匹配 metadata、generated、compdb | 构建数据来源，仍可能有 variant 去重 |
| `stale` | 过期 fingerprint | 可浏览，提示刷新 |

## 3. Pipeline

```text
Profile
  -> topology discover/validate
  -> metadata readers
  -> generated source store
  -> module graph
  -> language indexes
  -> snapshot
  -> query/diagnostics
```

每个阶段独立返回 diagnostics。Java metadata 失败不能阻止 topology 和 AIDL source-only；compdb 失败不能阻止 Java 查询。

## 4. Metadata readers

读取顺序：

1. `module_bp_java_deps.json`；
2. `module_bp_cc_deps.json`；
3. 产品输出或 Soong 的 `module-info`；
4. `development/ide/compdb/compile_commands.json`；
5. metadata 引用的 srcjar/generated file。

reader 必须忽略未知字段，保留 source path、external path 和 fingerprint。不要把绝对路径直接当 symbol identity。

## 5. Generated Source Store

- srcjar 解压到 IDE cache；
- 使用输入 hash 去重；
- 临时目录写入后原子 rename；
- 记录 owner module、RepoPath、entry、mtime、size、hash；
- 不写 AOSP checkout；
- 缺失 generated file 只产生 diagnostics。

## 6. Snapshot 与失效

snapshot key 至少包含 profile id、topology fingerprint、metadata fingerprint、build target 和 engine schema。

失效条件：profile/path mapping/topology、Android.bp/Android.mk、metadata、srcjar、compdb、AIDL/JNI source 或 engine schema 改变。

## 7. Resource limits

- 默认 scope 必须非空或由用户显式选择 full visible tree；
- topology discovery 有深度和文件数预算；
- srcjar 解压有 entry 数、总字节数和单文件大小限制；
- RPC 有消息大小、并发请求和 cancellation；
- index 扫描使用有界 worker pool。

## 8. Acceptance

- 使用 fixture 完成 profile -> topology -> snapshot；
- unknown metadata fields 不破坏读取；
- source-only 不需要 out；
- generated 缓存不污染源码；
- 相同逻辑 RepoPath 在不同 OS 得到相同 symbol key。
