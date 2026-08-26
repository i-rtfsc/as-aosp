# Target-driven ProjectGraph 决策

> 状态：Decision；替代“selectedModules 作为产品顶层选择器”和“单 focused module
> 承载完整 closure”的过渡模型。

## 1. 决策

用户配置的顶层对象是开发 Target，而不是单个 Soong module。Target 可以由源码目录、
Soong modules、jar、srcjar、APK、资源和 generated facts 组合而成：

```text
TargetSpec
  -> ArtifactResolver
  -> Soong module/source/generated closure
  -> ProjectGraph
  -> JetBrains modules / VS Code Java projects / clangd projects
```

Targets 是唯一用户配置；不提供 selectedModules 兼容 selector。没有 targets 时 profile
无效，客户端必须在投影前提示用户配置目标。

## 2. 为什么不是 Artifact 直接等于 Module

`framework.jar`、`services.jar` 和 `SystemUI.apk` 是用户可理解的构建目标或 artifact
anchor，不一定是一一对应的 Soong module：

- jar 需要 source roots、generated roots、dependency jars 和 source attachments；
- APK 需要 owning module、Java source、R/AIDL generated source、resources 和 manifest；
- 一个逻辑 Target 可能包含多个 Soong modules；
- 一个 Soong module 也可能被多个 Targets 共享。

APK 不作为普通 Java library；它只提供 artifact identity 和 owning/build facts anchor。

## 3. TargetSpec

```json
{
  "id": "systemui",
  "displayName": "SystemUI",
  "kind": "android",
  "selectors": {
    "modules": ["SystemUI"],
    "artifacts": [
      {"kind": "apk", "name": "SystemUI.apk"}
    ],
    "sourceRoots": ["frameworks/base/packages/SystemUI"]
  },
  "includeDependencies": true,
  "includeGeneratedSources": true,
  "includeResources": true,
  "includeManifest": true
}
```

Target resolution is deterministic and fail-closed for malformed selectors. Unknown module,
artifact or source selector produces diagnostics; it never silently expands to the full tree.

## 4. ArtifactResolver

Resolver precedence:

```text
exact artifact identity/digest
  > exact Soong output owner
  > normalized output-relative path
  > explicit user source/module selector
  > source-only fallback
```

Basename-only matching is forbidden. Resolver returns artifact kind, owner modules, output path,
source attachments, generated facts, match level and diagnostics.

## 5. ProjectGraph

```text
ProjectGraph
  targets[]
    id / kind / selectors
    resolvedModules[] / resolvedArtifacts[]
    sourceRoots[] / generatedRoots[]
    libraries[] / sourceAttachments[]
    resourceRoots[] / manifests[]
    dependencies[]
  sharedLibraries[]
  ownership[]
  duplicateSourceDecisions[]
  diagnostics[]
```

ProjectGraph is the only client projection contract. It implements the AIDEGen-equivalent
rules for duplicate source roots, framework source priority, framework srcjar separation and
module dependency order. Clients do not independently flatten ModuleCatalog.

## 6. Compatibility

- existing pack/module RPCs remain available for diagnostics and engine internals；
- current single focused module and VS Code `java.project.*` are temporary implementations of
  ProjectGraph projection, not alternate product paths；
- clients consume ProjectGraph only；
- no Gradle compatibility project is restored as product architecture.

## 7. Acceptance

- Framework target resolves `framework.jar` plus sources/generated/dependencies;
- Services target resolves `services.jar` and framework dependency;
- SystemUI APK target resolves owning module, Java/generated/resource/manifest facts without
  treating APK as a Java jar;
- duplicate physical source roots have one deterministic owner;
- JetBrains and VS Code receive equivalent ProjectGraph facts;
- missing vendor source or approximate pack match remains diagnosable and does not corrupt
  unrelated targets.
