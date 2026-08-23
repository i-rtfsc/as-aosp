# Generated Source Store 设计

> 状态：Target；首个 consumer 是 AIDL srcjar peer，后续可复用于 proto、R、sysprop 和 logtags。

## 1. 责任

Generated Source Store 负责安全读取 srcjar、按内容 hash 建缓存、原子发布解压目录和查询 entry。
它不负责决定 AIDL owning module，也不修改 AOSP checkout 或 out。

## 2. 缓存身份

```text
<cache-root>/srcjar/<sha256>/
  entries/...
  manifest.json
```

SHA-256 基于 archive 内容，不使用绝对路径或 mtime。相同内容在不同构建目录只保存一份。

## 3. 安全限制

默认限制：

- 最大 10,000 entries；
- 单 entry 最大 16 MiB；
- 解压总量最大 256 MiB；
- 拒绝绝对 entry；
- 拒绝 `..` 目录逃逸；
- 拒绝 symlink 和非普通文件；
- 只在 cache root 下创建临时目录；
- 任一 entry 失败则删除临时目录，不发布部分结果。

限制必须可在测试中缩小，但生产默认不能由 archive 内容修改。

## 4. 原子发布

1. 计算 archive hash；
2. 已有完成 manifest 时复用；
3. 在 cache root 创建临时目录；
4. 校验并写入全部 entry；
5. 写 manifest；
6. rename 到 hash 目录；
7. 并发进程已发布同一 hash 时复用现有结果。

## 5. AIDL peer 接线

- Java expected entry 为 `package/path/Outer.java`；
- nested declaration 仍匹配外层 Java entry；
- srcjar 包含 entry 时 peer origin 为 `build-srcjar`；
- confidence 为 `exact`，path 指向 cache 中的 LocalPath；
- archive 不存在、entry 缺失或解压失败时保留 naming-rule peer；
- 解压错误通过 stable diagnostic 返回，不让 AIDL source symbol 消失。

## 6. Diagnostics

| Code | 含义 |
|---|---|
| `GENERATED_ARCHIVE_OPEN_FAILED` | archive 不可读 |
| `GENERATED_ARCHIVE_LIMIT` | entry 数或字节限制超出 |
| `GENERATED_ENTRY_UNSAFE` | entry 为绝对路径或逃逸路径 |
| `GENERATED_ENTRY_WRITE_FAILED` | entry 写入失败 |
| `GENERATED_ENTRY_NOT_FOUND` | 目标 Java entry 不在 srcjar |

## 7. Acceptance

- 相同 archive 内容复用同一 hash 目录；
- Java entry 可升级 AIDL peer 为 `build-srcjar`/`exact`；
- zip-slip fixture 被拒绝且 cache root 外无文件；
- entry/size 限制可测试；
- 中断或失败不留下可被误认为完成的缓存目录。
