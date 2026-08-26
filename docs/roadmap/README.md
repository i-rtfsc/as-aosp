# AOSP IDE Roadmap

本目录只记录 Current、Target、Gap、任务顺序和下一验收。实现证据在代码和测试入口，历史完成流水不写入 active roadmap。

## 阅读顺序

1. [当前状态](current-state.md)
2. [当前交付计划](delivery-plan.md)
3. [剩余任务分解](active-execution-plan.md)
4. Target-driven ProjectGraph 迁移以 `docs/decisions/target-driven-project-graph.md`
   和 `docs/design/project-graph-and-ide-projection-design.md` 为准。
5. [Target-driven ProjectGraph 实施计划](target-driven-project-graph-plan.md)

## 状态词

| 状态 | 含义 |
|---|---|
| `Current` | 当前工作树有代码和可复验证据 |
| `Partial` | 原语存在但产品链或验收未闭合 |
| `Ready` | 技术前置已满足，可以执行 |
| `Queued` | 等待另一个本地前置 |
| `External` | 需要真实 IDE、网络、远程主机或平台环境 |

## 规则

- 不写完成百分比、提交账本或临时 owner；
- Target 不得写成 Current；
- 任务必须有下一验收和负能力边界；
- 设计变更同步更新 Current/Gap 和交付计划。
