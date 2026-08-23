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

- [Feasibility report](docs/aosp-ide-plugin-feasibility-report.md)
- [Detailed design](docs/aosp-ide-plugin-design.md)
- [Contributing guide](CONTRIBUTING.md)

Implementation will follow the directory boundaries defined in the detailed design. Remote build support, if added later, will live under remote-execution/; the project does not use that ambiguous term as a component or directory name.
