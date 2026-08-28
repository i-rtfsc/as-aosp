# aosp-index CLI 使用指南

`aosp-index` 是共享 Go engine 的命令行入口，用于 profile/topology 检查、metadata、
AIDL/JNI、Portable Out 导出和 IDE JSON-RPC。

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

### 3.1 Android 17 CP2A 服务器构建

对于 Android 17.0.0 Release 1、Build ID 为 `CP2A.*` 的完整 AOSP checkout，推荐使用
Cuttlefish x86_64 phone 的 `eng` 配置生成开发期 facts：

```bash
cd /home/solo/code/aosp
source build/envsetup.sh
lunch aosp_cf_x86_64_phone cp2a userdebug
export SOONG_GEN_COMPDB=1
export SOONG_GEN_COMPDB_DEBUG=1
export SOONG_GEN_CMAKEFILES=1
export SOONG_GEN_CMAKEFILES_DEBUG=1
export SOONG_GEN_RUST_PROJECT=1

# 生成 Soong/Java/模块 metadata 和完整 compile_commands.json
m -j"$(nproc)" nothing

# 生成目标 Java/APK、srcjar 及其 out 中间产物
m -j"$(nproc)" framework-minus-apex services SystemUI

# 生成 APK installed/owner、module dependency 等 module-info facts
m module-info
```

`cp2a` 与当前 Android 17 release 对应；`trunk_staging` 会得到开发代号配置，不作为
该 release 的首选。若需要完整系统镜像，再执行 `m -j"$(nproc)"`，但代码导航验证通常
不需要等待完整镜像。

`module-info` 是独立构建目标。即使 Java deps、CC deps 和 compdb 已生成，也必须显式执行
`m module-info`，否则 `SystemUI.apk` 等 installed artifact 无法可靠映射到 owner module，
Pack 导出会 fail closed。

编译后检查关键文件：

```bash
find out -type f \( \
  -name module-info.json -o \
  -name module_bp_java_deps.json -o \
  -name module_bp_cc_deps.json -o \
  -name compile_commands.json \
\) -print
```

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
  --out-root /path/to/out \
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

## 6. Portable Out 导出

默认导出 Core metadata，以及 AIDL/R/sysprop/logtags/proto 等 Java generated srcjar，
并保持 OUT_DIR 相对路径：

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
  --native \
  --ide
```

服务器导出后建议直接校验并查看 artifact 统计：

```bash
aosp-index pack export \
  --aosp-root "$ANDROID_BUILD_TOP" \
  --out-root "$OUT_DIR" \
  --destination /srv/packs/android-framework \
  --module framework-minus-apex \
  --module services \
  --native

aosp-index pack validate --root /srv/packs/android-framework
aosp-index pack inspect --root /srv/packs/android-framework
du -sh /srv/packs/android-framework
```

如果还不确定服务器 out 是否生成了足够的 IDE facts，可先运行：

```bash
aosp-index pack preflight --out-root "$OUT_DIR" \
  --native \
  --ide \
  --module framework-minus-apex \
  --module services \
  --module SystemUI
```

除了确认 `module-info`、`java-deps`、`cc-deps` 和 `compdb`，输出还会给出机器可判定的
`ready`/`blockingDiagnostics`，并逐模块报告
source roots、generated、jar/srcjar 数量以及已识别的 artifact。若 compdb 显示缺失，建议在已完成
lunch 的同一 shell 中执行 `SOONG_GEN_COMPDB=1 m nothing` 后再导出；Java 导航不因
compdb 缺失而失效，但 C/C++/JNI 编译参数会降级。

preflight 默认不递归扫描整个 out；需要统计全局 jar/srcjar/APK 文件数时显式增加
`--scan-all-artifacts`。

也可以在服务器设置环境变量后使用仓库脚本：

```bash
export ANDROID_BUILD_TOP=/srv/aosp
export OUT_DIR=/srv/aosp/out
export AOSP_PACK_DESTINATION=/srv/packs/android-framework
export AOSP_PACK_MODULES=framework-minus-apex,services
export AOSP_PACK_NATIVE=1
make export-pack
```

如果完整 AOSP 已经编译过但在末尾非关键目标失败，可以导出整个 out 的所有可用事实，
而不是只选择几个模块。服务器使用独立工具目录，脚本和 engine 必须始终放在同一目录：

```text
/home/solo/code/aosp-pack-tools/
├── aosp-index
└── export-full-aosp-pack.sh
```

该目录是服务器唯一工具入口。更新 engine 或脚本时必须同时替换这两个文件；脚本通过
自身目录定位 `aosp-index`，启动时打印 SHA-256，不读取 `/tmp` 或 AOSP 根目录中的副本。

```bash
chmod +x /home/solo/code/aosp-pack-tools/export-full-aosp-pack.sh \
  /home/solo/code/aosp-pack-tools/aosp-index
cd /home/solo/code/aosp-pack-tools
./export-full-aosp-pack.sh
```

默认输出：

```text
/home/solo/code/aosp-packs/android-17-cp2a-aosp_cf_x86_64_phone-full
/home/solo/code/aosp-packs/android-17-cp2a-aosp_cf_x86_64_phone-full.zip
/home/solo/code/aosp-packs/android-17-cp2a-aosp_cf_x86_64_phone-full.zip.sha256
```

脚本不传 `--module`，因此会导出所有已生成的 Java/CC metadata、jar、srcjar 和完整
compdb。它会先执行 `m module-info`，再运行 preflight；任何 blocking diagnostic 都会阻止生成
不完整压缩包。可通过 `AOSP_ROOT`、`OUT_ROOT`、`AOSP_PACK_NAME`、`PACK_ROOT` 和
`ZIP_ROOT` 覆盖默认路径。engine 不允许通过环境变量覆盖，始终使用脚本同目录的
`aosp-index`，脚本启动时会打印 engine SHA-256。默认名称由 Android 版本、release、product
和 full pack 组成，不使用 `xx` 之类的占位名称。
脚本检测到 Soong/Ninja/Siso 仍在构建时会拒绝导出，避免复制正在变化的 out；通常不应使用
`ALLOW_ACTIVE_BUILD=1` 绕过该保护。

Portable Out 会把 ModuleCatalog 明确引用且实际存在的 Java jars 复制到
`soong/.intermediates/**/*.jar`。复制到 Mac 后，直接选择 Portable Out 目录作为 Out Root；
它与本机完整 AOSP out 使用完全相同的相对路径和 reader。

```bash
# Portable Out 可以放在任意目录，不需要放进 AOSP Workspace View
aosp-index pack validate --root /Users/me/aosp-packs/android-framework --native
```

可重复选择目标模块及其 dependency closure：

```bash
--module framework-minus-apex --module services
```

选择模块后，AIDL 和 `--native` JNI facts 只扫描 closure 覆盖的 source roots，
generated srcjar 只保留 closure 明确引用的 archive，并直接按引用路径复制，不遍历整个
`out/soong/.intermediates`。模块名不存在时导出失败，
不会退化为全树扫描或按 basename 猜 owner。

Portable Out 输出：

```text
target/product/<product>/module-info.json
soong/module_bp_java_deps.json
soong/module_bp_cc_deps.json
soong/development/ide/compdb/compile_commands.json
soong/.intermediates/**/*.srcjar
soong/.intermediates/**/*.jar
soong/rust-project.json
development/ide/clion/**/CMakeLists.txt
```

可以通过网盘、U 盘、rsync 或挂载目录复制到 Mac/Windows，不需要下载服务器。

## 7. Portable Out 校验

```bash
aosp-index pack validate --root /Users/me/aosp-metadata/android-16 --native
```

校验原始 module-info、Java deps，以及 Native 模式下的 CC deps/compdb；同时解析
ModuleCatalog 和 artifact facts。

## 8. Portable Out 查看

查看 pack ModuleCatalog：

```bash
aosp-index pack inspect --root /Users/me/aosp-metadata/android-16
aosp-index pack inspect --root /Users/me/aosp-metadata/android-16 \
  --module framework-minus-apex
```

输出包含 module、jar、srcjar 数量和字节；指定 module 时额外返回模块 facts 和
dependency closure。服务器复制前建议依次执行：

```bash
aosp-index pack validate --root /srv/packs/framework
aosp-index pack inspect --root /srv/packs/framework \
  --module framework-minus-apex
du -sh /srv/packs/framework
```

## 9. IDE 服务

```bash
aosp-index serve --stdio
```

stdin/stdout 使用 Content-Length JSON-RPC framing。stdout 只写协议，日志写 stderr。

## 10. 常见问题

- 没有 out：先用 source-only；也可以复制 Linux 导出的 Portable Out；
- Linux 绝对路径：使用 `--path-map external=local`；
- generated Java 缺失：确认 pack 中存在对应 srcjar/entry；
- compdb 路径不完整：检查 `COMPDB_PATH_UNMAPPED` 和 variant diagnostics；
- AOSP、MIUI、QCOM：使用独立 Workspace View/profile。
