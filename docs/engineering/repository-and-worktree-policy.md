# 仓库、暂存与工作树规范

> 文档职责：定义分支、精确暂存、共享工作树、敏感文件和危险 Git 操作边界。

## 1. 分支边界

- `aosp` 保留历史 Gradle/CMake 方案；
- `dev-ide` 只承载新插件方案；
- 不把旧目录复制回 `dev-ide` 作为兼容实现；
- 不建立新的 Git 仓库承载新产品；
- 大功能使用从 `dev-ide` 创建的短生命周期 topic branch，完成后通过 PR 合并。

## 2. 精确暂存

暂存时必须显式列出当前 commit 的文件：

```bash
git add docs/engineering/commit-message-guidelines.md \
  docs/engineering/repository-and-worktree-policy.md \
  docs/README.md CONTRIBUTING.md
```

禁止：

```bash
git add .
git add -A
git commit -a
```

原因：当前分支重建包含大量删除和多个 bounded context，宽泛暂存极易把不相关实现、用户文件或下一提交
内容混入。

## 3. 多人或自动化共享工作树

- 开始前记录分支、HEAD、`git status --short`、修改路径和验证入口；
- 只修改获得任务范围的路径；
- 发现他人修改时不回滚、不覆盖，先判断是否影响当前任务；
- 共享文件（go.mod、Makefile、schema、README、CI）采用短时串行修改；
- 协作者分别拥有独立提交，不由一个人代替其他人提交；
- 需要并行开发时优先使用独立 `git worktree`。

## 4. 禁止的工作树操作

未获用户或维护者明确授权时，禁止：

- `git reset --hard`；
- `git clean -fdx`；
- `git checkout -- <path>` 或 `git restore <path>` 回滚他人修改；
- `git stash` 隐藏共享工作树状态；
- `git pull --rebase`、交互式 rebase 或 force push；
- 切换到其他分支导致未提交工作丢失；
- 删除用户本地 AOSP、out、挂载点或软链接。

## 5. 本地和敏感文件

不得提交：

- `local.properties`；
- `.idea/`、`.vscode/` 和本地 workspace cache；
- `.aosp-ide/*.local.json`；
- 用户真实 AOSP/out 绝对路径；
- token、SSH key、cookie、内部地址和私有 metadata；
- srcjar 解压缓存、compdb overlay 和 engine runtime cache。

可以提交的 testdata 必须是最小合成 fixture，不复制真实厂商源码或构建产物。

## 6. 修改前检查

```bash
git branch --show-current
git rev-parse --short HEAD
git status --short
git diff --stat
```

如工作树已 dirty，先确认哪些修改属于当前任务；不得通过清理命令获得“干净状态”。

## 7. 提交前检查

```bash
git status --short
git diff --cached --name-status
git diff --cached --check
git diff --cached
make check
```

提交完成后检查：

```bash
git show --stat --oneline HEAD
git status --short
```

剩余未提交文件应属于明确的后续 commit，不应是意外遗漏。

## 8. PR 与合并

- PR 标题遵循 commit subject 风格；
- PR 正文说明问题、设计、不变量、验证和未覆盖环境；
- 一个 PR 可以包含多个按主题拆分的 commit，但不能混入无关清理；
- 合并策略由维护者选择 merge/rebase/squash；开发者不自行重写已评审历史；
- breaking profile/RPC 变更必须在 PR 中给出迁移说明。
