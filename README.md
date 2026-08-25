# AOSP IDE

This branch contains the new AOSP IDE integration.

The project will provide:

- a shared cross-platform indexing engine;
- an Android Studio and IntelliJ IDEA plugin;
- a VS Code extension;
- AIDL and JNI cross-language navigation;
- source-only and build-backed indexing;
- lightweight workspace views assembled from directories, symlinks, junctions, mounts, and overlays.

The current repository root is the new project root. The historical Gradle/CMake implementation remains on the aosp branch and is not part of dev-ide.

## Design

The active design is split by responsibility. Start with the
[documentation center](docs/README.md), then read the current architecture,
decisions, design contracts, and roadmap entries.

The original feasibility report and monolithic design are retained in
[`docs/archive/`](docs/archive/README.md) for historical context only.
- [Contributing guide](CONTRIBUTING.md)

Implementation will follow the directory boundaries defined in the detailed design. Remote build support, if added later, will live under remote-execution/; the project does not use that ambiguous term as a component or directory name.
