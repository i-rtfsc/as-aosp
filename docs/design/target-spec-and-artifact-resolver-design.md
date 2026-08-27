# TargetSpec 与 ArtifactResolver 设计

> 状态：Target；定义用户目标如何解析为构建事实，不负责 IDE API。

## 1. Target kinds

| kind | 典型目标 | 主要投影 |
|---|---|---|
| `java` | Framework、Services、framework.jar | Java sources、generated、jars、srcjars |
| `android` | SystemUI.apk、Settings | Java、resources、manifest、generated、libraries |
| `native` | libui、SurfaceFlinger | source roots、headers、compdb、clangd |
| `mixed` | Framework + JNI、vendor service | Java、native、AIDL/JNI edges |

## 2. Selectors

Selectors may include:

- `modules[]`：Soong module names；
- `artifacts[]`：jar/apk/srcjar/output identity；
- `sourceRoots[]`：explicit RepoPath roots；
- `paths[]`：用户可见目录或文件；
- `includeDependencies`、`includeGeneratedSources`、`includeResources`、`includeManifest`。

Selectors are additive within one Target and do not merge unrelated profiles.

## 3. Artifact facts

An artifact fact contains:

```text
kind: jar | apk | srcjar | aar | output
name / relativePath / externalPath
sha256 / size / mtime
ownerModules[]
sourceAttachments[]
generatedRoots[]
match: exact | compatible | approximate | missing
```

APK facts include manifest/resource metadata but are not entered as Java classpath libraries.

## 4. Resolution algorithm

1. Normalize paths and target names.
2. Match exact artifact path/digest from validated Pack or local out facts.
3. Resolve owning module from module-info/Java deps/CC facts.
4. Add dependency closure only when requested.
5. Add explicit source roots and generated/archive owners.
6. Emit missing/ambiguous/approximate diagnostics.
7. Deduplicate by RepoPath and artifact identity.

No basename guessing or whole-tree expansion is allowed after an unresolved selector.

## 5. Fact source requirements

本机 out 与 Metadata Pack 都必须产出同一 ArtifactCatalog 语义。本机 out 使用
`module-info.json`、`module_bp_java_deps.json`、`module_bp_cc_deps.json` 等现有 reader
即时构建；Pack 使用导出时写入的稳定 catalog。客户端只提交 roots 与 TargetSpec，不接触
底层 Soong schema。

### Pack requirements

Pack export should include an artifact catalog:

```text
artifacts/catalog.json
artifacts/dependency-graph.json
```

Existing `modules/catalog.json`, `generated/manifest.json`, `aidl/index.json` and `jni/index.json`
remain source facts; artifact catalog adds the user-target resolution layer.

Artifact catalog consumes ModuleCatalog jars/srcjars and module-info installed outputs. Installed
`.apk` facts provide Android target anchors such as `SystemUI.apk`; they are not copied into Java
classpath or treated as jars. Clients query `aosp/artifacts/catalog` to populate graphical selectors
and may still enter an explicit custom artifact path.
