# AOSP Metadata Pack 设计

> 状态：Target；导出/导入脚本先于自动下载客户端实现。

## 1. 目标

Metadata Pack 是从 Linux 完整 AOSP/out 中导出的纯数据包，供另一台只挂载部分 AOSP 源码
仓库的 macOS/Windows/Linux 使用。它不是完整 out，不包含可执行文件和编译产物，只包含
能帮助 IDE 建立项目模型和导航的事实。

```text
aosp-metadata-pack/
├── manifest.json
├── module-info.json
├── module_bp_java_deps.json
├── module_bp_cc_deps.json
├── compile_commands.json
├── aidl/
│   ├── symbols.json
│   ├── modules.json
│   └── generated-peers.json
├── generated/
│   ├── srcjars/
│   ├── java/
│   ├── aidl/
│   ├── sysprop/
│   ├── logtags/
│   ├── proto/
│   ├── headers/
│   └── manifest.json
├── modules/
│   ├── catalog.json
│   └── dependency-graph.json
└── checksums.sha256
```

压缩包可以使用 zip/tar.zst；目录格式必须保持相同，方便用户直接复制目录而不是必须
解压。插件不要求 pack 位于 AOSP 根目录。

## 2. 文件职责

### Core Java facts

- `module-info.json`：模块路径、source、generated、dependencies、classes、tags；
- `module_bp_java_deps.json`：Java/Kotlin sources、jars、srcjars、AIDL include dirs、flags；
- `modules/catalog.json`：规范化 ModuleCatalog，作为客户端首选输入；
- `modules/dependency-graph.json`：模块依赖闭包。

### Native facts

- `module_bp_cc_deps.json`：C/C++ source、include dirs、sysroot、flags；
- `compile_commands.json`：clangd/CLion 编译数据库；
- `generated/headers`：可选 generated headers manifest 或按需文件。

### AIDL/generated facts

- `aidl/symbols.json`：package、interface、parcelable、method、type ranges；
- `aidl/modules.json`：AIDL source 与 owning module；
- `aidl/generated-peers.json`：Java/C++/NDK peer 和 confidence；
- `generated/srcjars`：默认包含 Java/R/AIDL/proto/sysprop/logtags 等生成源码 archive；
  这些是 Java/AIDL 导航的基础输入，不属于可选增强；
- `generated/java`、`generated/aidl`、`generated/sysprop`、`generated/logtags`、
  `generated/proto`：默认包含从 srcjar 解出的可索引源码或 entry manifest，便于 IDE
  直接建立 generated source root；

pack 中不放：`.o`、`.so`、dex、system image、ninja 临时文件、可执行程序和与 IDE 无关的
大中间产物。

## 3. manifest

```json
{
  "schemaVersion": 1,
  "release": "android-16",
  "manifestDigest": "sha256:...",
  "buildFacts": {
    "product": "aosp_arm64",
    "variant": "userdebug",
    "architecture": "arm64"
  },
  "sourceIdentity": {
    "buildMakeRevision": "...",
    "buildSoongRevision": "...",
    "frameworksBaseRevision": "..."
  },
  "packs": ["core", "native"],
  "files": 1234,
  "bytes": 987654321
}
```

`product/variant/architecture` 用于增强和诊断，不是基础 Java 导航的启用条件。

## 4. 导出脚本

建议新增：

```bash
aosp-index pack export \
  --aosp-root /srv/aosp \
  --out-root /srv/aosp/out \
  --destination /srv/packs/android-16-aosp \
  --native
```

导出规则：

1. 读取已存在的 out metadata；
2. 复制 JSON 并归一化路径字段，保留原始 ExternalPath；
3. 收集 AIDL/JNI/generated facts；
4. 按 source/metadata/generated owner 过滤；
5. 生成 manifest 和 SHA-256；
6. 原子写入 destination；
7. 不修改 AOSP checkout 和 out。

用户手工复制方式等价于插件读取本地 pack，不要求服务端。

### 4.1 out 候选位置

导出器按候选路径查找，不要求所有文件同时存在：

```text
<out>/module-info.json
<out>/soong/module-info.json
<out>/target/product/<product>/module-info.json
<out>/module_bp_java_deps.json
<out>/soong/module_bp_java_deps.json
<out>/module_bp_cc_deps.json
<out>/soong/module_bp_cc_deps.json
<out>/development/ide/compdb/compile_commands.json
<out>/soong/development/ide/compdb/compile_commands.json
<out>/soong/.intermediates/**/{*.srcjar,*.jar,*.h,*.hpp,*.java}
```

缺失某一类只产生 diagnostic，不让整个 pack export 失败。例如只有 Java metadata 时仍可
生成 Core Pack，没有 compdb 时只缺 Native Pack。

### 4.2 导出模式

```bash
# 最小 framework/Java 导航包
aosp-index pack export --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" --destination /tmp/aosp-pack --core

# 加入 clangd/CLion 数据
aosp-index pack export --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" --destination /tmp/aosp-pack \
  --core --native

# 只保留 selected module closure，Java generated/srcjar 仍默认包含
aosp-index pack export --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" --destination /tmp/aosp-pack \
  --native \
  --module framework-minus-apex --module services
```

`--core` 是默认且推荐模式，并默认包含 Java/AIDL 相关 generated srcjar、entry manifest
和必要的解出源码。Native headers、超大非 Java generated 文件和未被 selected module
引用的 archive 才是可选项；不能把 `--generated` 作为 Java/AIDL 基础能力的开关。

建议命令：

```bash
# 默认：Core + Java/AIDL/generated source archives
aosp-index pack export --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" --destination /tmp/aosp-pack

# 额外加入 C/C++ compdb、headers 和 native facts
aosp-index pack export --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" --destination /tmp/aosp-pack --native
```

### 4.3 复制到另一台机器

```bash
# 普通目录复制
rsync -a --delete /tmp/aosp-pack/ \
  user@mac:/Users/user/aosp-metadata/android-16/

# 压缩后通过网盘/U盘
tar -C /tmp -caf aosp-pack-android16.tar.zst aosp-pack

# Mac 解压到任意只读/可写目录
mkdir -p /Users/user/aosp-metadata/android-16
tar -C /Users/user/aosp-metadata/android-16 \
  --strip-components=1 -xaf aosp-pack-android16.tar.zst
```

插件 profile 只保存本机 pack path：

```json
{
  "workspaceViewRoot": "/Users/user/code/aosp-view",
  "metadataPackRoot": "/Users/user/aosp-metadata/android-16"
}
```

不要求 pack 位于 Workspace View 或 canonical AOSP root 内。

### 4.4 部分源码仓库场景

Linux 构建机可以拥有完整 AOSP/out，而 Mac 只包含：

```text
frameworks/base
system/core
packages/modules/Connectivity
```

导入 pack 后：

- ModuleCatalog 可以保留完整 module/dependency facts；
- 只把本机 Workspace 可见 source root 投影给 IDE；
- 不可见 module 标记 missing provider；
- jars/srcjars/generated 默认作为 library/source attachment；
- 不把 Linux 绝对路径直接写进 Mac IDE，必须经 RepoPath/PathMapping；
- Java/framework baseline 优先启用，不因 C++ product/arch 不匹配而整体禁用。

## 5. 导入与匹配

匹配等级：

```text
exact       manifest digest / 核心 revision 完全匹配
compatible  release 相同，核心 revision 或 product 有差异
approximate release 接近或只有部分 identity
incompatible schema 不兼容或明显不同 release
```

使用优先级：

```text
本地 out exact facts
> 本地 Metadata Pack exact
> 本地 out/pack compatible
> Android.bp 静态 facts
> source-only fallback
```

即使 `compatible/approximate`，framework Java source roots 仍可以启用；generated/compdb
结果必须显示等级并允许关闭。

### 5.1 Framework-first 宽松策略

为了实现用户希望的“90% 公共代码先可跳”，导入时分域判断，不使用一个总开关：

| 域 | compatible/approximate pack 默认行为 |
|---|---|
| Framework Java source roots | 启用，Workspace 本地源码优先 |
| Java jars/srcjars | 默认启用并显示 pack match |
| AIDL source symbols | 启用；本地 AIDL 优先，generated 标记 match |
| JNI source links | 启用 source-only/baseline link |
| C/C++ compdb | compatible 可启用，approximate 默认提示确认 |
| product generated/R/sysprop | 仅 exact 默认启用，否则提示或降级 |
| system image/runtime artifacts | 不进入 pack |

这保证 product/variant/arch 是“精度与增强维度”，而不是阻止 Java/framework 导航的门槛。

## 6. 路径与挂载

pack 的 identity 使用 RepoPath，不使用构建机绝对路径。ExternalPath 通过用户 profile
`pathMappings` 转换。用户可以只挂载：

```text
frameworks/base
system/core
vendor/qcom
```

pack 仍提供完整 module facts；不可见 source provider 标记 missing，不自动扫描 canonical
root。generated source 只有在本机可见或 pack 自带时标记 exact，否则保留 inferred。

## 7. 安全

- pack 只包含数据；
- checksum 校验；
- 解压拒绝绝对路径、`..`、symlink 和超大 entry；
- 写入 `.aosp-ide/cache`，不写源码；
- 导入失败不替换旧 pack；
- 未签名 pack 显示 `unverified`，但允许用户显式启用；
- 公共公益分发可用网盘/GitHub Release，动态下载服务不是前置条件。
