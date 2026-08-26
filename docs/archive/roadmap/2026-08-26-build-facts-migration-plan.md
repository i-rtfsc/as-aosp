# Build Facts 迁移实施计划

> 状态：Target；本计划用于替换 prototype 主路径，不要求一次性推翻已有客户端。

## Phase 1：ModuleCatalog

- 复用现有 metadata readers；首个 `core/modules` ModuleCatalog slice 已实现；
- 定义 `ModuleCatalog`、SourceRoot、GeneratedSource、Classpath、NativeCompileUnit；
- module-info/Java deps/CC deps 合并为稳定模块 facts；
- 增加 module ownership、source root 和 dependency closure 查询；
- 通过 `aosp/modules/catalog|get|closure` 向两个 IDE 客户端提供统一查询；
- 验收：framework fixture 能回答 `android.os.*` 所属 source root/module。

## Phase 2：Metadata Pack export/import

- `aosp-index pack export` 和 `pack validate` 首个 Core/native slice 已实现，并输出规范化 ModuleCatalog/dependency graph；
- 默认收集 AIDL/R/sysprop/logtags/proto 等 Java generated srcjar 和 entry manifest；
- 实现 manifest/checksum/目录校验；
- 支持从 Linux out 导出并在另一个 Workspace View 导入；
- 验收：复制 pack 到 fixture 外部目录后仍能解析 module/source/generated facts。

## Phase 3：IndexCoordinator

- profile 保存后启动后台 coordinator；
- VS Code workspace 激活时对已有 profile 自动启动，配置变化 debounce 后重建；
- 阶段：discover、metadata、module graph、Java/AIDL/JNI、generated、compdb；
- 增加 `aosp/index/status`、`aosp/index/rebuild`、`aosp/index/cancel`；
- 首版先完成可查询的阶段状态和 Metadata Pack catalog/generate 计数，再接 IDE 原生索引；
- snapshot 按 profile/topology/metadata/build facts fingerprint 失效；
- 验收：侧边栏显示 progress、ready/partial/stale/error。

## Phase 4：JetBrains project projection

- re-audit 后增加 `ProjectGraph` 作为 ModuleCatalog 到 IDE 的稳定投影层；
- 实现 duplicate source ownership、framework srcjars 和 dependency module 规则；

- 使用 Workspace Model/ModuleRootModel facade 投影 source roots、jars、srcjars、generated；
- 当前 JetBrains slice 先复用 ModuleRootModel 投影 source/generated roots；libraries 和
  Workspace Model closure 后续接入；
- 保留用户现有模块，不把 canonical root 注册为 content root；
- 对 selected module closure 创建 focused modules；
- 验收：无需旧 Gradle 脚本，framework Java source root 和 generated peer 可由 IDE 原生索引。
- 当前切片已完成 VS Code `java.project.sourcePaths` 与 generated srcjar cache 投影；
  Java LS 完整 classpath 作为下一验收。

## Phase 5：VS Code language integration

- 当前 `java.project.sourcePaths` 保留为 fallback；目标改为 JDT LS 多 project/importer；

- Java：生成 workspace-scoped source paths/referenced libraries，优先交给 Java LS；
- AIDL：继续使用 shared engine provider；
- C/C++：使用 compdb overlay + clangd；
- JNI：shared DefinitionProvider/command；
- 验收：打开 AOSP Workspace 后 Java、AIDL、C/C++、JNI 分别走正确 language authority。

## Phase 6：增量与厂商树

- Android.bp/Android.mk 指纹失效；
- 本地修改模块只重建受影响 closure；
- MIUI/OPPO/QCOM 允许 baseline pack + local facts 合并；
- 冲突显示 diagnostics，不静默覆盖；
- 验收：同机多个 profile 独立、挂载缺失可诊断、90% 公共 framework 导航保持可用。

## Prototype 迁移规则

- 当前 source-only AIDL/JNI/Java fallback 保留为 degraded mode；
- Java/AIDL generated srcjar 属于 Core Pack 默认内容，不需要额外 `--generated` 才导出；
- 不再把 fallback 结果称为 build-backed；
- 先接 ModuleCatalog，再替换 Java 路径猜测；
- 现有 RPC/VS Code/JetBrains 生命周期可复用；
- 每个 phase 先更新 bounded-context design，再实现和提交。
