# Workspace Topology 与路径映射设计

> 状态：Target；本设计冻结 Phase 0 的数据模型和算法边界。

## 1. 问题与非目标

用户可能在完整 AOSP 之外打开目录 A。目录 A 可以由普通目录、软链接、junction、mount、复制目录和 overlay 组成。设计必须保留用户视图，又能解析真实 AOSP RepoPath。

本设计不负责 IDE Project Model；它只负责发现、解析和选择 Source Provider。

## 2. 核心类型

```text
RepoPath      AOSP 根目录相对路径，使用 /
ViewPath      Workspace View Root 下的相对路径
LocalPath     当前机器路径
ExternalPath  metadata/compdb 中的构建机路径
```

```go
type SourceMount struct {
    Name       string
    RepoPrefix RepoPath
    ViewPath   ViewPath
    LocalPath  LocalPath
    Origin     MountOrigin
    Priority   int
    ReadOnly   bool
}

type WorkspaceTopology struct {
    WorkspaceViewRoot LocalPath
    CanonicalRepoRoot *LocalPath
    OutRoot           *LocalPath
    SourceMounts      []SourceMount
    PathMappings      []PathMapping
}
```

## 3. Provider 选择

1. 选择最长匹配 `RepoPrefix`；
2. 同前缀按 `Priority` 降序；
3. 同优先级不同内容返回 `conflict`；
4. 未生效候选保留为 `shadowed`；
5. Canonical Root 只能作为 fallback；
6. `strict-view` 下，未被 active mount 暴露的 RepoPath 不进入索引。

## 4. 自动发现

1. 使用 `lstat` 读取 Workspace View Root；
2. 只在受控深度内识别目录、软链接、junction 和 mount；
3. view 子路径像 AOSP RepoPath 且 target 位于 Canonical Root 时生成候选；
4. 候选一致自动接受，不一致要求用户确认；
5. 每个 mount 记录 traversal budget、visited identity 和诊断。

## 5. 异常与恢复

| 异常 | 行为 |
|---|---|
| dangling link | 保留 mount，状态 `missing`，不崩溃 |
| cycle | 停止该分支，状态 `cycle` |
| duplicate physical target | RepoPath 只建一份 symbol identity，记录 aliases |
| provider conflict | 不猜测，状态 `conflict`，要求 priority 或显式选择 |
| case-only collision | Windows 状态 `conflict` |
| mixed release | topology 可用，build-backed 降级为 mixed/source-only |

## 6. Acceptance

- fixture 能表达完整 AOSP + 轻量 view；
- 任意别名需要显式映射；
- 真实 root 不自动成为 content root；
- 循环和断链不会无限递归；
- 同一个 RepoPath 不生成重复 symbol owner；
- Linux、macOS、Windows 风格路径归一化结果一致。
