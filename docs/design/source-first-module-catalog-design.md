# Source-First Module Catalog Design

> Status: Decision; `Android.bp` and `Android.mk` are the authoritative module
> inventory. Out metadata is an optional enhancement and must never replace the
> source inventory.

## 1. Non-Negotiable Rule

When a workspace is opened, module discovery starts from the visible
`Android.bp` and `Android.mk` files. This remains true when:

- no `out` directory exists;
- only part of AOSP is mounted or exposed through symlinks;
- the Out directory belongs to another machine, product, or variant;
- the source tree is a vendor tree with incomplete or incompatible build facts.

Out facts may enrich an existing source module with build-backed status,
artifacts, generated sources, AIDL metadata, JNI facts, jars, and compiler
commands. Out must not create the primary module inventory or hide a module
that was found in source configuration.

The precedence is therefore:

```text
Android.bp / Android.mk
  -> source module inventory
  -> source type, source paths, declared dependencies
Out facts, when available
  -> build status, artifacts, generated material, language facts
Unified catalog
  -> recommendations, selection, ProjectGraph
```

## 2. Source Module Parsing

The parser reads module name, module type, build file path, namespace/package,
source roots, source files, declared dependencies, generated inputs, AIDL
properties, native include hints, Android resources, and manifests. It does
not execute Soong and does not require Linux.

Unknown properties are ignored with a diagnostic. An invalid file does not
discard valid modules from other files. Source discovery obeys the active
Workspace View and SourceMount boundaries and never recursively scans Out as
source configuration.

## 3. Module Roles

The catalog assigns independent build and navigation dimensions:

```text
Build role: product | library | generator | build-support | test
Capability: java | android | native | rust | aidl | jni | generated
Evidence: source-only | build-backed | build-only
```

Typical source/build classes are mapped as follows:

| Source or module-info class | Typical role | Default module list |
|---|---|---|
| `JAVA_LIBRARIES` | product or library | yes when Java-capable |
| `APPS` | product | yes when Android/Java-capable |
| `SHARED_LIBRARIES` | library | yes when native/JNI-capable |
| `STATIC_LIBRARIES` | library | yes when native-capable |
| `EXECUTABLES` | product or library | yes when source navigation is available |
| `ETC` | build-support or data | no |
| `TESTS` | test | no by default |
| `filegroup`, `genrule`, `se_build_files` | generator/build-support | no |

The class names above are evidence for classification, not the source of
truth. Without Out, the parser uses the `Android.bp`/`Android.mk` module type
and capabilities. With Out, `module-info.json` can confirm classes and
installed outputs.

## 4. Buildable Targets vs Navigation Targets

Any valid Soong module may be accepted by `m`, including a sepolicy or
configuration module. That does not make it a useful default navigation
target. For example:

```text
m framework-minus-apex
m services
m SystemUI
```

are common code-oriented targets, while:

```text
m 202604.board.ignore.map
```

may be valid but is a build-support operation for sepolicy input data. The
second category remains searchable through an advanced build-target view, but
must not pollute the default code module list.

## 5. Default UI Contract

The default module view contains product modules and modules with a meaningful
navigation capability. It excludes `TESTS` and build-support modules by
default. The UI provides an explicit advanced view for all source-discovered
build modules, so source-first discovery is never lost.

Recommended views:

```text
Recommended
Java
Android
Native
Rust
Generated / AIDL
All code modules
Advanced build modules
```

Each result shows the source module name, module type, build role, capability,
source/build evidence, output kind when known, and the Android.bp/Android.mk
path. Product artifacts such as `framework.jar`, `services.jar`,
`SystemUI.apk`, shared libraries, and executables are shown as output facts,
not as replacement module identities.

`TESTS` modules are retained in the catalog and searchable from Advanced build
modules, but are not in presets or default results.

## 6. Merge and Precedence

Merge identity uses module name plus namespace/build-file path and proven source
paths. A source module always survives an Out merge. Out-only entries are
marked `build-only` and are never silently treated as source modules.

Lists are unioned and deduplicated. Out values fill missing source facts but do
not overwrite a valid source declaration merely because the Out value is
present. Conflicting providers produce diagnostics.

## 7. No-Out Behavior

Without Out, the catalog still supports:

- source module browsing and filtering;
- preset selection based on source module names and types;
- source-only Java/native/Rust capability classification;
- source AIDL and JNI fallback navigation.

Selecting a module without Out creates a source-only TargetSpec. Build-backed
status, jars, srcjars, compdb, and generated artifact attachments are simply
absent or marked unavailable.

## 8. Acceptance Criteria

- A workspace with no Out still lists modules from `Android.bp`/`Android.mk`.
- Adding or removing Out changes enrichment, not the source inventory.
- `framework-minus-apex`, `services`, and `SystemUI` are first-class default
  code targets when present in source configuration.
- `202604.board.ignore.map` and other `ETC`/`se_build_files` entries are not
  default modules but remain available in Advanced build modules.
- `TESTS` modules are hidden from defaults and presets.
- The same source module has one identity whether or not matching Out facts
  are available.
