# Commit Message 与提交粒度规范

> 文档职责：定义 commit message、提交粒度、footer 和历史修正规则。暂存、共享工作树和分支操作见
> [仓库与工作树规范](repository-and-worktree-policy.md)。

## 1. 基本格式

本仓库采用 Conventional Commits 风格：

```text
<type>(<scope>): <subject>

<body>

<footer>
```

示例：

```text
feat(topology): resolve source providers by prefix priority

Workspace views can expose the same AOSP RepoPath through symlinks,
overlays, and the canonical source tree. Selecting by discovery order
made the active source depend on filesystem traversal order.

Resolve providers by longest RepoPath prefix and explicit priority.
Return a conflict when equally ranked providers expose different
physical content, and retain lower-ranked candidates as shadowed.

Validate with make test-core and make check.
```

## 2. 必填语义

除纯拼写修复等极小提交外，正文必须说明：

1. 为什么需要变更；
2. 改了什么模块和关键不变量；
3. 对协议、兼容性、路径、缓存、运行或后续工作的影响；
4. 使用哪些工程入口验证，结果如何；
5. 哪些风险或外部测试尚未覆盖。

正文描述设计和行为，不逐行复述 diff。

## 3. Type

| Type | 用途 |
|---|---|
| `feat` | 新增用户或开发者可见能力 |
| `fix` | 修复错误行为或不变量 |
| `docs` | 只修改文档 |
| `refactor` | 不改变对外行为的结构调整 |
| `test` | 测试、fixture 和测试工具 |
| `build` | Make、Go module、Gradle、打包和 toolchain |
| `ci` | GitHub Actions 和自动化 gate |
| `perf` | 可复验的性能改进 |
| `chore` | 不属于以上类型的仓库维护 |

不要使用 `update`、`changes`、`misc`、`fix stuff` 等无法表达结果的 type 或 subject。

## 4. Scope

Scope 使用真实 bounded context 或交付面：

| Scope | 适用范围 |
|---|---|
| `model` | RepoPath、Profile、Diagnostic 等基础值类型 |
| `topology` | Workspace View、SourceMount、provider 和 cycle/conflict |
| `pathmap` | View/Local/External 与 RepoPath 转换 |
| `metadata` | Soong、module-info 和 compdb reader |
| `generated` | srcjar 和 generated source cache |
| `aidl` | AIDL parser、resolver 和 generated peers |
| `jni` | JNI descriptor、registration 和 cross-language links |
| `graph` | ModuleGraph、dependency closure 和 ProjectView |
| `index` | snapshot、incremental update 和 symbol query |
| `rpc` | Profile/RPC schema 和 daemon protocol |
| `cli` | `aosp-index` 和命令行组合 |
| `jetbrains` | Android Studio/IntelliJ IDEA 客户端 |
| `vscode` | VS Code 客户端 |
| `remote-execution` | 可选 Linux 构建桥接 |
| `packaging` | 多平台二进制和插件打包 |
| `build` | 仓库构建入口 |
| `ci` | CI gate |
| `docs` | 文档体系 |
| `repo` | 分支、目录和仓库维护 |

一次变更跨越多个 scope 时，优先拆分提交。确实不能拆分时使用最主要的 owner，不使用 `all`、`misc`
或多 scope 拼接。

## 5. Subject

- 使用英文，便于 GitHub、release notes 和跨语言协作统一消费；
- 使用小写开头的祈使/结果动词，例如 `add`、`reject`、`normalize`、`resolve`、`split`；
- 描述提交完成后的结果，不描述“正在做什么”；
- 不写句号；
- 不超过 72 个字符；
- 不加入 issue 编号、作者名或工具名。

推荐：

```text
fix(pathmap): preserve Windows drive paths during repo mapping
```

禁止：

```text
Update files
fix stuff
feat: changes for path
```

## 6. Body

- 正文可以使用英文或中文，但一个 commit 内保持同一种语言；
- 公共产品和跨团队提交优先使用英文；
- 每一行都不得超过 75 个字符，正文尽量在 72 字符前自然换行；
- 同一逻辑段连续换行，段落之间保留一个空行；
- 行为、协议、数据、构建、兼容性或安全边界变化必须有正文；
- 验证必须写工程入口，例如 `make check`，不能只写 `tests passed`；
- 未运行真实 IDE、Windows 或远程挂载测试时必须明确说明。

## 7. Footer

允许：

```text
BREAKING CHANGE: profile schema v2 requires sourceMounts.repoPrefix
Refs: #123
Closes: #123
```

当前开发环境配置了 Conventional Commit 校验和 `Change-Id` 生成 hook，因此：

- 每个 commit 必须包含唯一的 `Change-Id: I<40 个小写十六进制字符>`；
- hook 自动生成时必须保留最终 message 中唯一的 Change-Id；
- 没有 hook 的贡献者需要在提交前安装兼容 hook，或手工提供合法 Change-Id；
- Change-Id 必须是最后一个 footer；
- **不强制 `Signed-off-by`**，除非仓库后续明确启用 DCO；
- 不添加 `Co-authored-by`、AI 工具、生成器或模型署名；
- 不在 footer 中包含 token、内部服务器地址、用户目录或私有数据。

如果未来移除 Change-Id 流程或引入 DCO，必须先修改本文和自动化 gate，再改变 footer 要求。

## 8. 提交粒度

一个 commit 只包含一个可独立解释、验证和回滚的主题。

应该拆分：

- 文档决策与实现可以分提交；
- 纯文件移动与行为修改分提交；
- formatter/机械重写与功能改动分提交；
- 构建工具升级与业务实现分提交；
- metadata reader、AIDL、JNI 和 IDE 客户端按 owner 分提交；
- 大量历史文件删除与新功能实现分提交。

可以同提交：

- 实现与直接证明该实现的不变量单元测试；
- schema 变更与其生成代码/兼容 fixture；
- 小范围代码变更与必须同步的 Current/Design 文档。

## 9. 当前 dev-ide 重建建议

当前工作树同时包含旧方案删除、文档体系和 Phase 0 代码。提交时建议拆成以下主题，而不是形成一个超大
commit：

1. `chore(repo): replace the legacy tree on dev-ide`
   - 删除旧 Gradle/CMake/脚本和旧资源；
   - 新增根 README 与通用 `.gitignore`；
   - 不包含 Go engine 实现。
2. `docs(architecture): define workspace product and system boundaries`
   - 产品、架构、决策、工程规范和 Roadmap；
   - 不包含生产代码。
3. `feat(topology): add workspace topology and path mapping core`
   - Go module、model/topology/pathmap 和单元测试。
4. `feat(cli): add profile doctor and topology scan commands`
   - CLI、Makefile smoke 和 testdata fixture。

若实际文件依赖使拆分后的某个中间 commit 无法通过自身验收，应调整顺序或把不可分割部分合并，但仍要
保持每个提交可构建、可测试和可解释。

## 10. 提交前检查

```bash
git status --short
git diff --check
git diff --cached --check
git diff --cached --stat
git diff --cached
make check
```

检查待推送提交的 message 行宽：

```bash
for commit in $(git rev-list --reverse origin/dev-ide..HEAD); do
  git show -s --format=%B "$commit" | awk -v commit="$commit" '
    length($0) > 75 {
      printf "%s:%d exceeds 75 characters: %s\n", commit, NR, $0
      failed = 1
    }
    END { exit failed }
  '
done
```

没有 upstream 时，将 `origin/dev-ide..HEAD` 替换为明确的本地范围，
例如 `HEAD~6..HEAD`。

仓库入口：

```bash
make commit-message-check
COMMIT_RANGE=origin/dev-ide..HEAD make commit-message-check
```

还必须确认：

- 暂存区只包含当前主题；
- 没有 `local.properties`、真实 AOSP 路径、凭据或生成缓存；
- 删除旧方案只发生在 `dev-ide`；
- 文档链接和 Current 状态与代码一致；
- 未运行项已在正文说明。

## 11. 历史修正

- 已共享分支禁止擅自 rebase、squash、amend 或 force push；
- 只修改尚未推送且由自己独占的本地提交时，可以使用 `commit --amend` 或交互式 rebase；
- 已推送到共享远端的历史需要维护者明确批准后才能重写；
- 修复旧提交 message 不能顺带改代码树；
- 重写后必须比较 tree、重新运行受影响 gate，并明确通知协作者；
- 不采用宽泛 `filter-branch` 作为日常 message 修复手段。

本仓库是 GitHub 协作仓库，优先通过新 commit 修复公开历史中的内容问题，而不是重写已有用户基线。
