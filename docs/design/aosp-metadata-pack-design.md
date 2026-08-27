# Portable Out 设计

> 状态：Decision；替代旧 Metadata Pack 私有目录格式。

## 1. 定义

Portable Out 是 AOSP `OUT_DIR` 的稀疏只读镜像。它只挑选 IDE 导航需要的 AOSP 构建
产物，但必须保持每个文件相对于 `OUT_DIR` 的原始路径。

```text
服务器 OUT_DIR                          本机 Portable Out
out/soong/module_bp_java_deps.json   -> out/soong/module_bp_java_deps.json
out/target/product/x/module-info.json -> out/target/product/x/module-info.json
out/soong/.intermediates/a.srcjar    -> out/soong/.intermediates/a.srcjar
```

完整 out 和 Portable Out 对 engine 来说都是 `Out Root`，不存在不同 reader、不同 RPC
参数或不同客户端配置。

## 2. 内容

Core Java 导航至少包含：

```text
soong/module_bp_java_deps.json
target/product/<product>/module-info.json
soong/.intermediates/**/*.srcjar
```

Native 导航增加：

```text
soong/module_bp_cc_deps.json
soong/development/ide/compdb/compile_commands.json
```

可以包含 Java classpath jar：

```text
soong/.intermediates/**/*.jar
```

`.srcjar` 是 generated Java source，默认复制；`.jar` 是 classpath/ABI 编译产物，可以按
导出策略选择，但不能改写路径。

Portable Out 不包含插件私有文件：

```text
manifest.json
checksums.sha256
modules/catalog.json
artifacts/catalog.json
aidl/index.json
jni/index.json
generated/manifest.json
```

ModuleCatalog、ArtifactCatalog、AIDL/JNI index 和 srcjar 解压目录由本机 engine 按需生成，
只写 IDE cache。

## 3. 导出

```bash
aosp-index pack export \
  --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" \
  --destination /srv/aosp-packs/android-17-out \
  --native
```

导出器只读加载原始 out metadata，计算 selected module closure，并按原始 out-relative
路径复制 metadata、srcjar 和所需 jar。它原子替换 destination，不修改 AOSP checkout/out，
也不在 destination 内生成插件私有索引。

ZIP 和完整性校验属于传输层，放在 out 目录外：

```text
android-17-out.zip
android-17-out.zip.sha256
android-17-out.json
```

## 4. 导入

用户可以把 Portable Out 解压到 `<AOSP_ROOT>/out`，也可以放在任意目录并在 UI 中选择为
`Out Root`。软链接、NFS、SMB、只读挂载和部分仓库 Workspace 都使用同一目录合同。

## 5. 厂商 out

MIUI、OPPO、QCOM 等厂商 out 仍是 Out Root。reader 按候选位置加载标准 AOSP facts，
厂商额外文件不会被当成 Pack 标识，也不会和插件私有文件冲突。

## 6. 验收

- 完整 out 与其 Portable Out 产生相同 ModuleCatalog/ArtifactCatalog；
- 解压到 `<AOSP_ROOT>/out` 后无需额外格式配置；
- 不出现 `out/out/soong`；
- 不出现 `generated/srcjars` 重排；
- 插件卸载后 out 中不残留 AOSP Workspace 文件；
- Windows/macOS/Linux 路径均通过统一 PathRef 解析。
