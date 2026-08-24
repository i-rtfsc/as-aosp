# AOSP IDE 文档中心

本目录是 `dev-ide` 分支的产品范围、架构、详细设计、工程规范、决策和交付计划入口。

## 1. 最短阅读路径

1. [产品概览](product/product-overview.md)和[范围与非目标](product/scope-and-non-goals.md)。
2. [系统总体架构](architecture/system-architecture.md)和
   [源码与模块边界](architecture/source-and-module-boundaries.md)。
3. [Workspace Topology 与路径映射设计](design/workspace-topology-and-path-mapping-design.md)。
4. [索引引擎设计](design/index-engine-design.md)。
5. [IDE 客户端与跨语言导航设计](design/ide-client-and-navigation-design.md)。
6. [AIDL 语言与解析器设计](design/aidl-language-and-resolver-design.md)。
7. [Generated Source Store 设计](design/generated-source-store-design.md)。
8. [JNI 跨语言索引设计](design/jni-cross-language-index-design.md)。
9. [当前状态](roadmap/current-state.md)和[当前交付计划](roadmap/delivery-plan.md)。
10. [打包与发布设计](design/packaging-and-release-design.md)。

## 2. 文档权威性

| 问题 | 权威入口 |
|---|---|
| 产品目标、范围和非目标 | `product/` |
| 稳定组件、进程和依赖边界 | `architecture/` |
| 字段、算法、状态和错误语义 | `design/` |
| 已接受且长期生效的取舍 | `decisions/` |
| 构建、测试、文件和文档规则 | `engineering/` |
| 当前代码事实与负能力 | `roadmap/current-state.md` |
| 未完成任务、顺序和验收 | `roadmap/` |
| 初始调查与完整提案背景 | 根目录两份 feasibility/design 报告 |

代码事实以当前源码和通过工程入口执行的测试为准。`Target` 设计不能写成 `Current` 能力。

## 3. 状态词

| 状态 | 含义 |
|---|---|
| `Current` | 当前工作树中存在，并有静态证据或工程测试证明。 |
| `Partial` | 已有一部分可靠实现，但关键链路或验收未闭合。 |
| `Target` | 已认可的目标设计，尚不能作为当前保证。 |
| `Decision` | 已接受并长期约束实现的取舍。 |
| `Out of scope` | 当前明确不做。 |

## 4. 目录职责

- `product/`：产品是谁、为谁解决什么问题。
- `architecture/`：稳定系统边界和依赖方向。
- `design/`：bounded context 的实现合同。
- `decisions/`：已接受取舍和被拒方案。
- `engineering/`：开发、构建、测试和文档规则。
- `roadmap/`：Current、Gap、任务顺序和验收。

## 5. 当前重要决策

- [`dev-ide` 分支与仓库根边界](decisions/dev-ide-repository-boundary.md)
- [共享 Go 索引引擎](decisions/shared-go-index-engine.md)
- [Workspace View 与完整 AOSP 根目录分离](decisions/workspace-view-separation.md)

## 6. 维护要求

- 一篇文档只承担一个 bounded context。
- Architecture 不保存周计划；Roadmap 不复制详细算法。
- Current、Target、Gap 和 Acceptance 分开写。
- 路径示例使用虚构、可替换的本地路径，不记录敏感服务器地址。
- 新代码改变 Current 时，同步更新 `roadmap/current-state.md`。
- 设计字段或算法改变时，先更新相应 Design，再修改代码。

## 7. Engineering

- [代码规范](engineering/code-guidelines.md)
- [构建与测试指南](engineering/build-and-test-guide.md)
- [Commit Message 与提交粒度规范](engineering/commit-message-guidelines.md)
- [仓库、暂存与工作树规范](engineering/repository-and-worktree-policy.md)
