# Contributing to AOSP IDE

## Required reading

- [Documentation center](docs/README.md)
- [Commit message guidelines](docs/engineering/commit-message-guidelines.md)
- [Repository and worktree policy](docs/engineering/repository-and-worktree-policy.md)
- [Build and test guide](docs/engineering/build-and-test-guide.md)
- [Current delivery plan](docs/roadmap/delivery-plan.md)

## Development flow

1. Start from `dev-ide` or a topic branch created from it.
2. Update the bounded-context design before changing a public contract or invariant.
3. Keep each commit focused and stage exact paths only.
4. Run `make check` before submitting.
5. Report tests that were not run, especially real IDE, Windows, network mount, and remote build checks.

Commit subjects use Conventional Commits, for example:

```text
feat(topology): resolve source providers by prefix priority
```

Every commit-message line, including body and footers, must be at most
75 characters long.

Every commit requires one `Change-Id` footer. The configured commit hook
generates it automatically; contributors without that hook must install a
compatible hook or provide a valid Change-Id manually. DCO `Signed-off-by`
footers are not required. Do not add AI, tool, or generator attribution to
commit messages.
