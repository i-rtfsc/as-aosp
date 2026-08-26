# AOSP IDE 术语表

| 术语 | 定义 |
|---|---|
| RepoPath | AOSP 根目录下、使用 `/` 的规范相对路径；跨平台符号身份。 |
| Workspace View Root | 用户实际在 IDE 中打开的轻量目录。 |
| Canonical AOSP Root | 可选完整 AOSP 目录；默认只作回退和 metadata 参考。 |
| Out Root | AOSP 构建产物目录。 |
| Source Mount | `RepoPath prefix -> ViewPath/LocalPath` 的源码提供关系。 |
| Source Provider | 可为某个 RepoPath 提供内容的 mount、overlay、完整根或缺失记录。 |
| Active Provider | 冲突解析后当前生效的 Source Provider。 |
| Shadowed Provider | 存在但未生效的低优先级 Provider。 |
| ViewPath | Workspace View Root 下用户看到的相对路径。 |
| LocalPath | 当前机器可访问的操作系统路径。 |
| ExternalPath | metadata/compdb 中记录的构建机绝对路径。 |
| ProjectView | 投影给具体 IDE 的模块、源码根、生成根和依赖集合。 |
| TargetSpec | 用户定义的开发目标，可组合 artifact、Soong module、源码目录和依赖选项。 |
| Artifact | 构建产物或输出锚点，例如 jar、apk、srcjar；不等同于 Soong module。 |
| ArtifactResolver | 将 TargetSpec 的 selectors 解析为 artifact、module、source 和 generated facts 的组件。 |
| ProjectGraph | 面向 IDE 的稳定投影合同，包含 targets、modules、libraries、ownership 和 dependencies。 |
| source-only | 仅依靠可见源码与目录规则的模式。 |
| metadata-partial | 部分构建 metadata 可用的模式。 |
| build-backed | 主要查询数据来自匹配的构建产物。 |
| strict-view | 只索引 Workspace View 明确暴露内容的默认策略。 |
| Remote Execution Bridge | 可选 Linux 构建机辅助服务；目录名为 `remote-execution/`。 |
