# 打包与发布设计

> 状态：Partial；首个本地跨平台产物切片进入实现，签名与商店发布仍为 Target。

## 1. 产物边界

```text
dist/
  engine/<os>-<arch>/aosp-index[.exe]
  jetbrains/aosp-ide-jetbrains-<version>.zip
  vscode/aosp-ide-vscode-<version>.vsix
```

engine 构建矩阵为 `darwin/linux/windows × amd64/arm64`。Windows 文件名带 `.exe`；
Go core 当前不依赖 CGO，因此同一源码可以交叉编译。JetBrains ZIP 携带六个平台二进制；
VS Code 按平台发布六个 target-specific VSIX，每个 VSIX 只携带一个对应平台二进制。

客户端 engine 解析优先级：用户显式 executable、用户公共 engine store 的
`engines/aosp-index`、扩展目录内当前 `os-arch` bundled engine、PATH 中的
`aosp-index`。JetBrains ZIP 与 VSIX 均可携带六平台 engine；客户端只执行当前
平台目录中的一个文件。在
manifest 签名和下载信任模型完成前，客户端不自动联网下载二进制。

JetBrains 与 VS Code 使用相同的 bundled layout：
`engine/<os>-<arch>/aosp-index[.exe]`。JetBrains 从已安装插件 descriptor 获取 plugin root，
不依赖当前工作目录；OS/arch 归一化集中在 engine resolver，Windows 公共入口文件名为
`~/engines/aosp-index.exe`。

公共 store 的稳定入口必须保持为 `engines/aosp-index`，不能把版本目录作为唯一入口。
客户端可在该文件旁读取版本/哈希元数据；更新时先写临时文件并原子替换，不能覆盖
正在运行的 executable。默认位置为 macOS `~/Library/Application Support/AOSP IDE`、
Linux `~/.local/share/aosp-ide`、Windows `%LOCALAPPDATA%/AOSP IDE`。

本机开发安装使用 `make install-vscode`：脚本自动识别当前 OS/架构、构建对应 engine、
生成只包含当前平台 engine 的 VSIX，并调用 `code --install-extension` 安装。用户安装
后只需打开 Workspace View，不需要手动复制 `engines/aosp-index` 或配置 executable；公共 engine 入口仍
保留给跨 IDE 复用和开发者自定义场景。

## 2. 客户端复用

- JetBrains 使用同一 Kotlin 源码和一个跨平台插件 ZIP，IDE/JVM 提供平台抽象；
- VS Code 使用同一 TypeScript 源码和六个 target-specific VSIX，由 Marketplace 根据
  当前平台选择包；
- 两个客户端共享 `com.rtfsc.aosp` 产品命名和 JSON-RPC；
- 插件安装包不同，因为 JetBrains 与 VS Code 扩展格式不兼容。

## 3. 安全与可复验性

- 所有产物只写 `dist/`；
- 不写 AOSP checkout、out 或用户全局目录；
- 打包前执行对应 test/build gate；
- 首版不声称代码签名、公证、JetBrains Marketplace 或 VS Code Marketplace 已完成；
- VSIX 使用 `.vscodeignore` 排除 TypeScript 源码、测试、node_modules 和 source map；
- 每次本地打包生成 `dist/manifest.sha256` 和 `dist/manifest.json`。JSON manifest 包含
version、client/engine kind、相对路径和 SHA-256；engine 条目还包含显式 os、arch 和
executable 文件名。两个 manifest 都不包含自身，
  签名设计另行评审。

## 4. Acceptance

- 六个平台 engine 文件名与 executable 格式正确；
- Windows 只生成 `.exe`；
- JetBrains ZIP 与 VSIX 能由仓库入口生成；
- 打包脚本不依赖用户真实 AOSP；
- `dist/` 不进入 Git。
- manifest 能覆盖 engine、JetBrains ZIP 和 VSIX，且不包含临时文件。
- JSON manifest 的 version 与 `AOSP_IDE_VERSION` 一致。
