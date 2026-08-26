# Target-driven ProjectGraph 实施计划

> 状态：Target；替代 archived Build Facts migration plan 的 selected-module 顶层模型。

## Phase 1：Target 与 Artifact domain

- 在 `core/model` 定义 TargetSpec、TargetKind、ArtifactSelector；
- 在 `core/artifacts` 建 artifact catalog 和 resolver；
- 支持 module、jar、apk、srcjar、source root selectors；
- targets 是唯一用户输入，不提供 selectedModules 兼容转换；
- 验收：framework.jar、services.jar、SystemUI.apk fixture 可解析 owner/modules/sources。

## Phase 2：ProjectGraph

- 新增 `core/projects`；
- 构建 target/module/artifact/source/generated/library/dependency graph；
- 实现 duplicate source ownership 和 framework priority；
- framework srcjars 与 shared dependencies 分层；
- 验收：Framework + Services + SystemUI 同时解析且无重复 source root。

## Phase 3：Pack 与 RPC

- pack 增加 `artifacts/catalog.json` 和 target resolution facts；
- 新增 target resolve/project graph RPC；
- 当前首个 RPC 名称为 `aosp/projects/graph`；
- manifest 记录 target closure 与 artifact coverage；
- 验收：Linux 导出后 Mac 可离线解析相同 ProjectGraph。

## Phase 4：JetBrains

- ProjectGraph targets 投影为多个稳定 modules；
- dependencies/shared jars/framework srcjars 使用 ModuleOrderEntry/library；
- Android target 投影 manifest/resource roots，APK 只作为 artifact anchor；
- 保留 current focused module fallback；
- 验收：Framework、Services、SystemUI 可独立导航且依赖顺序正确。

## Phase 5：VS Code

- current `java.project.*` 继续作为 fallback；
- 实现 JDT LS 多 project/importer 或 workspace-scoped project descriptions；
- target-scoped source/generated/classpath；
- target-scoped clangd compdb overlay；
- 验收：三个 targets 不被扁平为一个 unmanaged Java project。

## Phase 6：真实验证与增量

- 服务器真实 AOSP pack；
- macOS/Windows/Linux 与 symlink/mount/partial source；
- target/artifact/Android.bp/out 指纹失效；
- affected ProjectGraph 局部重建；
- Android Studio、IDEA、VS Code、clangd/CLion 验证矩阵。
