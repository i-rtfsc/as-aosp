# AOSP IDE 文档归档

本目录保存已经被当前 Build-Facts 设计、架构拆分和路线图取代的历史
提案。归档文档用于追溯早期判断、研究证据和迁移背景，不是 Current、
Target 或实施状态的权威来源。

当前事实以 [`../README.md`](../README.md)、`architecture/`、`design/`、
`decisions/`、`engineering/` 和 `roadmap/` 下的活动文档为准。

## 2026-08-25

| 文档 | 归档原因 | 当前替代入口 |
|---|---|---|
| [`design/2026-08-25-aosp-ide-plugin-design.md`](design/2026-08-25-aosp-ide-plugin-design.md) | 原始单体设计已拆分为架构、索引、拓扑、AIDL、JNI、UI 和打包合同；Build-Facts 驱动导航已成为主路径。 | [`../architecture/system-architecture.md`](../architecture/system-architecture.md)、[`../design/aosp-metadata-pack-design.md`](../design/aosp-metadata-pack-design.md)、[`../decisions/target-driven-project-graph.md`](../decisions/target-driven-project-graph.md) |
| [`research/2026-08-25-aosp-ide-plugin-feasibility-report.md`](research/2026-08-25-aosp-ide-plugin-feasibility-report.md) | 初始可行性报告中的 Gradle/脚本背景和 source-only 过渡结论已被后续研究与 Build-Facts 决策更新。 | [`../research/aosp-ide-navigation-research.md`](../research/aosp-ide-navigation-research.md)、[`../decisions/build-facts-driven-navigation.md`](../decisions/build-facts-driven-navigation.md) |

归档不表示历史内容无效；它表示阅读者必须结合当前替代文档解释其中的
方案、状态和验收结论。

## 2026-08-26

| 文档 | 归档原因 | 当前替代入口 |
|---|---|---|
| [`research/2026-08-26-aosp-ide-navigation-reaudit.md`](research/2026-08-26-aosp-ide-navigation-reaudit.md) | 重新审计正确识别 ProjectGraph 缺口，但仍把 selected Soong module 作为顶层用户模型；旧分支复核证明用户需要 artifact/module/source 组合 Target。 | [`../decisions/target-driven-project-graph.md`](../decisions/target-driven-project-graph.md) |
| [`roadmap/2026-08-26-build-facts-migration-plan.md`](roadmap/2026-08-26-build-facts-migration-plan.md) | 旧迁移计划以 ModuleCatalog closure 和单 focused module 为主，已被 TargetSpec、ArtifactResolver、ProjectGraph 顺序替代。 | [`../design/target-spec-and-artifact-resolver-design.md`](../design/target-spec-and-artifact-resolver-design.md)、[`../design/project-graph-and-ide-projection-design.md`](../design/project-graph-and-ide-projection-design.md) |
