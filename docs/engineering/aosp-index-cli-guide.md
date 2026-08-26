# aosp-index CLI 使用指南

`aosp-index` 是共享 Go engine 的命令行入口，用于 profile/topology 检查、metadata、
AIDL/JNI、Metadata Pack 和 IDE JSON-RPC。

## 1. 构建

```bash
go build -o ./bin/aosp-index ./cmd/aosp-index
make check
```

命令不要求当前目录是 AOSP 根，路径通过参数传入。

## 2. Profile 与 Workspace

```bash
aosp-index doctor --profile /path/to/profile.json
aosp-index scan --profile /path/to/profile.json
```

`doctor` 校验 profile；`scan` 发现 Workspace View、SourceMount 和路径问题。两者不会自动
扩大扫描到完整 canonical AOSP。

## 3. Metadata 与模块图

```bash
aosp-index metadata --root /path/to/aosp/out --summary
aosp-index metadata --root /path/to/aosp/out > metadata-report.json
```

主要读取 `module-info.json`、`module_bp_java_deps.json`、`module_bp_cc_deps.json` 和
`development/ide/compdb/compile_commands.json`。缺少某类文件只产生 diagnostic。

## 4. AIDL

```bash
aosp-index aidl \
  --root /path/to/aosp/frameworks/base \
  --file android/os/IFoo.aidl \
  --type android.os.IFoo \
  --summary
```

结合 out/srcjar：

```bash
aosp-index aidl \
  --root /path/to/sources \
  --file android/os/IFoo.aidl \
  --type android.os.IFoo \
  --metadata-root /path/to/out \
  --path-map /srv/aosp/out=/Users/me/aosp/out \
  --cache-root /Users/me/.cache/aosp-workspace/srcjar
```

没有 out 时仍能 source-only 解析；不存在的 generated peer 标记为 `inferred`。

## 5. JNI

```bash
aosp-index jni \
  --java-root /path/to/frameworks/base/core/java \
  --cpp-root /path/to/frameworks/base/core/jni \
  --class android.security.Scrypt \
  --method nativeScrypt \
  --summary
```

结果包含 Java/Kotlin native declaration、静态 export、常见注册表和双向 links。

## 6. Metadata Pack 导出

默认导出 Core metadata，以及 AIDL/R/sysprop/logtags/proto 等 Java generated srcjar：

```bash
aosp-index pack export \
  --aosp-root /srv/aosp \
  --out-root /srv/aosp/out \
  --destination /srv/packs/android-16
```

额外加入 CC metadata 和 compdb：

```bash
aosp-index pack export \
  --aosp-root /srv/aosp \
  --out-root /srv/aosp/out \
  --destination /srv/packs/android-16 \
  --native
```

可重复选择目标模块及其 dependency closure：

```bash
--module framework-minus-apex --module services
```

选择模块后，AIDL 和 `--native` JNI facts 只扫描 closure 覆盖的 source roots。
模块名不存在时导出失败，不会退化为全树扫描。generated srcjar 当前仍全量保留。

首版输出：

```text
manifest.json
checksums.sha256
module-info.json
module_bp_java_deps.json
generated/srcjars/**
generated/manifest.json
aidl/index.json
jni/index.json  # 仅 --native
```

可以通过网盘、U 盘、rsync 或挂载目录复制到 Mac/Windows，不需要下载服务器。

## 7. Metadata Pack 校验

```bash
aosp-index pack validate --root /Users/me/aosp-metadata/android-16
```

校验 manifest、checksum、核心文件和 generated manifest。

## 8. Metadata Pack 查看

查看 pack ModuleCatalog：

```bash
aosp-index pack inspect --root /Users/me/aosp-metadata/android-16
aosp-index pack inspect --root /Users/me/aosp-metadata/android-16 \
  --module framework-minus-apex
```

指定 module 时返回模块 facts 和 dependency closure。

## 9. IDE 服务

```bash
aosp-index serve --stdio
```

stdin/stdout 使用 Content-Length JSON-RPC framing。stdout 只写协议，日志写 stderr。

## 10. 常见问题

- 没有 out：先用 source-only，或复制 Linux 导出的 Metadata Pack；
- Linux 绝对路径：使用 `--path-map external=local`；
- generated Java 缺失：确认 pack 中存在对应 srcjar/entry；
- compdb 路径不完整：检查 `COMPDB_PATH_UNMAPPED` 和 variant diagnostics；
- AOSP、MIUI、QCOM：使用独立 Workspace View/profile。
