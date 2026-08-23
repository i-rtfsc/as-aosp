# Workspace View 与完整 AOSP 根目录分离决策

> 状态：Decision；接受日期：2026-08-23。

## 1. Decision

Workspace View Root、Canonical AOSP Root、Out Root 和 External Build Root 是四个独立概念。
Canonical AOSP Root 的存在不得自动扩大 IDE content root 或索引范围。

默认可见性策略为 `strict-view`：只索引 Workspace View 中 active 且被 scope 选择的 Source Mount。

## 2. Why

用户可能拥有完整 AOSP，但只希望打开由软链接、junction、mount 或 overlay 拼装的目录 A。将
realpath 后的完整根目录自动加入 IDE 会破坏轻量工作区目标，并造成重复身份和巨大索引开销。

## 3. Consequences

- RepoPath 是唯一跨平台文件身份；
- ViewPath、LocalPath、Physical Target 和 ExternalPath 分别保存；
- 同一 RepoPath 只能有一个 active provider；
- hidden dependency 优先使用 jar、srcjar、generated header 或 compdb；
- 打开未暴露源码时必须遵守可见性策略，不递归扩大 source root；
- 目录循环、断链、overlay 冲突和重复物理文件必须可诊断。

## 4. Alternatives considered

### 直接打开完整 AOSP

简单，但违背性能和部分仓库需求。Rejected。

### 所有路径先 realpath

易去除表面别名，但会丢失 Workspace View 语义并暴露完整根。Rejected。

### 复制所需源码到新目录

避免链接问题，但同步、修改归属和磁盘成本高。Supported as a Source Mount origin，不作为唯一方案。
