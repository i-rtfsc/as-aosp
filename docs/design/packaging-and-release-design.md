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
Go core 当前不依赖 CGO，因此同一源码可以交叉编译。插件包不同时携带六个平台二进制，
首版仍由用户配置 `aosp-index`；后续 marketplace/release 可按平台组合下载或发布。

## 2. 客户端复用

- JetBrains 使用同一 Kotlin 源码和插件 ZIP，IDE/JVM 提供平台抽象；
- VS Code 使用同一 TypeScript 源码和平台无关 VSIX；
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
