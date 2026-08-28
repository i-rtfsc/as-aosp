# AOSP Workspace 当前状态

> 复核日期：2026-08-28。

## 已完成

- `dev-ide` 使用单一 `Out Root` 模型。完整 AOSP `out` 与稀疏 Portable Out
  对 engine 完全等价，Portable Out 保持原始 `OUT_DIR` 相对路径。
- Go engine 已实现 profile、拓扑、路径映射、Build Facts、ModuleCatalog、
  ArtifactCatalog、ProjectGraph、AIDL 和 JNI source index。
- `aosp-index serve --stdio` 已提供初始化、索引状态、模块/制品/项目图、
  AIDL、JNI、compdb 和 generated srcjar materialize RPC。
- Portable Out 导出、校验、inspect、preflight 及服务器脚本已实现。导出结果
  不包含 `manifest.json`、`modules/`、`aidl/`、`jni/`、`generated/` 等插件私有目录。
- `.srcjar` 按真实 `soong/.intermediates/...` 路径复制；Java classpath jar
  也按 Build Facts 引用路径复制。派生索引和解压目录只写本机 IDE cache。
- VS Code 与 JetBrains 均有图形化 profile/target UI、Out Root 配置、索引状态、
  artifact 选择、target 预览和 AIDL/JNI 导航入口。
- 源码模块发现和工作区隔离的最终设计已确定：Android.bp 负责模块发现与无 out
  fallback，Out Root 负责构建事实增强；模块勾选后由 engine 生成内部 TargetSpec。
- `core/sourcecatalog` 已提供 Android.bp/Android.mk 发现、推荐标记、fingerprint cache
  和 `aosp/source/modules` RPC；VS Code 侧边栏与 JetBrains Tool Window 已接入模块列表。
- Source Catalog 已与 Out ModuleCatalog 合并，UI 可显示 `build-backed` 和 Java/native/
  generated/AIDL/artifact 能力；Build Facts 派生 catalog 写 workspace/project IDE cache。
- `aosp/rust/project` 已能把 Soong Rust project 归一化到本机 cache；VS Code 将其加入
  workspace-scoped rust-analyzer linked projects，缺少 Rust metadata 时不影响其他语言。
- Workbench/Manage Modules、双语 UI、presets 和 Selection Planner 去重合同已确定；
  相同 source root、srcjar 和 library 必须只投影一次，Select All 不创建海量 IDE modules。
- VS Code 已实现紧凑 Workbench、Manage Modules、默认 Platform presets、中英文资源和
  grouped Selection Plan；旧 workspace 的一模块一 Target 配置会在首次启动时自动迁移。
- engine 支持 macOS、Linux、Windows 的 bundled/shared/PATH 解析；JetBrains
  插件和 VS Code VSIX 有跨平台 engine 打包入口。
- JetBrains 打包已区分 thin 开发包和 full 发布包；full ZIP 内置六平台 engine，
  thin ZIP 不允许作为用户发布包，VS Code 继续使用六个 target-specific VSIX。
- 真实本地 Portable Out 已完成目录迁移并通过 `pack validate --native`，
  ModuleCatalog 约 126k 模块、ArtifactCatalog 约 44k 制品。

## 当前限制

- `aosp-index pack` 命令名称暂时保留，但其输入输出语义已经是 Portable Out；
  不存在旧版私有 Pack reader。
- JetBrains Settings 仍保留高级 Target 编辑器；日常入口已切换为 Tool Window 的模块发现
  列表。Android.bp parser 当前是有限属性解析器，复杂 Soong 表达式仍返回部分结果或诊断。
- JetBrains 尚未切换到 Workbench/Manage Modules 和 ResourceBundle；Selection Planner
  尚未实现 pinned targets、Select All 影响预估和 product-aware Automotive 默认策略。
- 真实 IDE host 的完整 Java Language Server/Workspace Model 接线仍在完善，
  目前以 focused source roots、generated roots 和 libraries 投影为主。
- JNI 采用静态源码和常见注册表分析；复杂宏、运行时 class 拼接和完整 C++
  语义会返回 inferred/unresolved。
- AOSP 构建机绝对路径不会自动写入本机配置；跨机器场景需要显式 path mapping。
- 尚未提供远程执行服务、插件市场发布签名和公证流程。

## 下一步

1. 实现 Selection Planner、preset grouped targets、custom buckets 和 pinned targets。
2. 实现紧凑 Workbench、Manage Modules 和 VS Code/JetBrains 中英文资源。
3. 完成真实宿主中的 ProjectGraph、Java classpath 和 IDE index 状态验证。
4. 增强 Android.bp defaults、namespace、glob 和 vendor 属性解析。
5. 补充 framework、services、SystemUI、AIDL、JNI 和 Rust 导航验收夹具。

## 验证命令

```bash
go test ./core/... ./protocol/... ./cmd/aosp-index
cd clients/vscode && npm run check
make test-jetbrains
bash -n scripts/export-aosp-pack.sh scripts/export-full-aosp-pack.sh
```
