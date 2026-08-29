# Built-in AOSP Profile Catalog and Storage Design

> Status: Decision; built-in profiles provide source-first navigation knowledge.
> Out metadata is optional runtime enrichment, not a plugin installation
> prerequisite.

## 1. Why This Design

The previous Gradle/CMake scripts worked because they converted AOSP build
facts into an IDE project model before the IDE started indexing. They did not
make the IDE discover AOSP semantics at every definition request.

The plugin keeps that advantage without requiring users to run scripts:

```text
development-time AOSP metadata
  -> curated built-in profiles
  -> plugin Project Model Compiler
  -> native Java/C++/Rust IDE projects
```

The plugin must not require every imported project to contain a complete Out.
Built-in profiles cover common AOSP and vendor-derived layouts. A local Out,
when present, adds exact jars, srcjars, generated files, compile commands,
CMake projects, and product facts.

## 2. Profile Count and UI Groups

The first built-in catalog contains 28 user-facing development profiles. The
UI shows these profiles, while the engine expands their module anchors,
families, source roots, and dependencies.

### 2.1 Java and system services

```text
Java Foundation
System Server
ActivityManager
WindowManager
System Server Input
SystemUI
SystemUI Shell
SystemUI Keyguard
SystemUI Compose
Common System Apps
```

Core anchors include:

```text
core-oj, core-libart, core-all, ext, framework,
framework-minus-apex, framework-all, framework-annotations,
android_stubs_current, android_system_stubs_current,
android_system_server_stubs_current, android_module_lib_stubs_current,
services, services.impl, services.core, services.core.priorityboosted,
services.core.unboosted, services.accessibility, services.appwidget,
services.devicepolicy, services.net, services.permission, services.power,
services.storage, services.usb, services.wifi, service-art,
service-blobstore, service-jobscheduler, service-permission
```

The `services.*` and `service-*` families are preferred over a fixed list so
new service modules in AOSP or vendor trees are included automatically.

### 2.2 Graphics, input, and native foundation

```text
Native Foundation
Graphics
SurfaceFlinger
RenderEngine
Input Stack
Display HAL
Power HAL
Sensors
USB
```

Core native anchors include:

```text
libbase, libutils, liblog, libcutils, libbinder, libbinder_ndk,
libnativehelper, libprocessgroup, libprocessgroup_setup, libprocinfo,
libvintf, libhidlbase, libhwbinder, libhardware, libhardware_legacy,
libnativebridge, libnativeloader, libunwindstack, libmemunreachable,
libziparchive, libz, libsqlite,
libui, libgui, libnativewindow, libhwui, libandroidfw,
surfaceflinger, librenderengine, libskia_renderengine,
libinput, libinputflinger, libinputdispatcher, libinputreader,
libinputreporter, libservices, libservices.core,
```

Typical source roots include:

```text
system/libbase
system/core/libutils
system/logging/liblog
system/core/libcutils
frameworks/native/libs/binder
frameworks/native/libs/ui
frameworks/native/libs/gui
frameworks/native/libs/nativewindow
frameworks/native/services/surfaceflinger
frameworks/native/services/inputflinger
frameworks/base/services/core/jni
frameworks/base/libs/services
frameworks/base/libs/hwui
```

### 2.3 Media, audio, and camera

```text
Media Framework
Stagefright
MediaPlayerService
MediaCodec
MediaProvider
Audio Framework
AudioFlinger
AudioPolicy
Audio HAL
Camera Framework
CameraService
Camera Provider
Camera HAL
```

Core anchors include:

```text
MediaProvider, MediaProviderLegacy,
libmedia, libmediandk, libmediautils, libmediametrics,
libstagefright, libstagefright_foundation, libstagefright_codec2,
libcodec2, libmediaextractor, libmediaplayerservice,
libaudioclient, libaudioflinger, libaudiopolicy, libaudioutils,
libaudiohal, libaudioprocessing, libcamera_client,
libcameraservice, libcamera_metadata, libcamera2ndk,
```

Families for generated interfaces:

```text
android.hardware.audio.*
android.hardware.audio.common.*
android.hardware.audio.core.*
android.hardware.audio.effect.*
android.hardware.audio.sounddose.*
android.hardware.camera.common.*
android.hardware.camera.device.*
android.hardware.camera.metadata.*
android.hardware.camera.provider.*
```

### 2.4 Connectivity and device modules

```text
Connectivity
Tethering
NetworkStack
DNS Resolver
Netd
Wi-Fi Framework
Wi-Fi Service
Wi-Fi HAL
Bluetooth Framework
Bluetooth Service
Bluetooth HAL
Bluetooth Rust
Telephony Framework
Radio HAL
IMS
```

Core anchors include:

```text
framework-connectivity, framework-connectivity-t, Tethering,
TetheringNext, NetworkStack, CaptivePortalLogin,
wifi-service-pre-jarjar, WifiTrackerLib, wifi-modules-utils,
Bluetooth, BluetoothServices, BluetoothLib, BluetoothMidiService,
telephony-common, TelephonyProvider, TelephonyStatsLib, ims-common,
libnetd_client, libnetdutils, libnetwork, libnetworkstats,
libdnsresolver, netd, libwifi-hal, libwifi-system, libwifi,
libril, libreference-ril, libtelephony, libims,
```

### 2.5 Package, security, storage, and runtime

```text
Package Manager
Permission Controller
Settings
Device Policy
Keystore
SELinux
Storage Framework
Vold
File System Manager
Incremental FS
ART Runtime
Dex2Oat
libcore
OpenJDK
Native Bridge
```

Core anchors include:

```text
Settings, SettingsLib, SettingsProvider, PackageInstaller,
PermissionController, SafetyCenter, RoleController, DocumentsUI,
vold, libvold, libstorage, libfs_mgr, libincremental, libdm,
libdmabufheap, libsnapshot, libsparse, liblp, libext4_utils,
libf2fs, libkeystore2, libkeystore2_crypto, libkeymaster_messages,
libkeymaster3device, libselinux, libsepol, libavb,
libart, libartbase, libartpalette, libdexfile,
libdexfile_support, libnativeloader, libnativebridge,
libopenjdk, libjavacore, libandroidio, dex2oat, dexdump,
dexlist, profman, oatdump, odrefresh, artd,
```

### 2.6 Automotive and common applications

```text
Automotive Framework
CarService
CarSystemUI
Vehicle HAL
Car Settings
```

Core anchors include:

```text
android.car, android.car.builtin, android.car-module,
CarService, CarServiceUpdatable, CarSystemUI, CarSettings,
CarFrameworkPackageStubs, car-frameworks-service,
car-frameworks-service-module, car-admin-ui-lib,
libvehiclehal, libvhalclient, libcarwatchdog,
```

Common application anchors include:

```text
SystemUI, Settings, SettingsLib, SettingsProvider,
PackageInstaller, DocumentsUI, Launcher3, Launcher3QuickStep,
PermissionController, MediaProvider, CaptivePortalLogin,
NetworkStack, Bluetooth, TelephonyProvider, ContactsProvider,
CalendarProvider, ManagedProvisioning, PrintSpooler,
DownloadProvider, ExternalStorageProvider,
```

### 2.7 AIDL and HIDL families

Versioned interface modules are represented by families rather than a fixed
version list:

```text
android.hardware.*-V*-java
android.hardware.*-V*-cpp
android.hardware.*-V*-ndk
android.frameworks.*-V*-java
android.frameworks.*-V*-cpp
android.frameworks.*-V*-ndk
android.system.*-V*-java
android.system.*-V*-cpp
android.system.*-V*-ndk
*-aidl-java
*-aidl-cpp
*-aidl-ndk
*-aidl-rust
```

The active source tree and optional Out determine which versions exist.

## 3. Module Configuration Shape

An individual built-in entry is declarative and compact:

```json
{
  "id": "system-server-wms",
  "kind": "java",
  "moduleNames": ["services", "services.impl", "services.core"],
  "modulePatterns": ["services.*", "service-*"],
  "sourceRoots": ["frameworks/base/services"],
  "sourcePatterns": ["frameworks/base/services/**/java"],
  "dependencies": ["java-foundation"],
  "generatedCapabilities": ["aidl", "sysprop", "aconfig"],
  "recommended": true
}
```

Rules can be overridden per workspace. User overrides add or replace module
names, roots, dependencies, CMake paths, include paths, and generated roots.
They never modify the built-in catalog.

## 4. Source-First and Out Enhancement

```text
Android.bp / Android.mk
  -> discover the actual modules visible in this checkout
  -> validate built-in module names and source roots

Built-in Profile
  -> provide common module relationships and fallback paths

Out metadata, if present
  -> add exact jars, srcjars, generated files, artifacts,
     compile commands, CMake projects, and product facts
```

A source module remains visible if Out is missing or incompatible. An Out-only
entry is marked `build-only` and does not become a default code module. `ETC`,
`filegroup`, `genrule`, `se_build_files`, and `TESTS` remain hidden from normal
views; tests require the explicit Advanced `Include test modules` switch.

## 5. Storage and Plugin Size

The plugin must not embed raw build facts. Current full AOSP examples are large:

```text
module_bp_java_deps.json   approximately 57 MB
module_bp_cc_deps.json    approximately 154 MB
module-info.json          approximately 108 MB
rust-project.json         approximately 5 MB
clion CMake tree          approximately 1.7 GB
```

Embedding these files would make the plugin platform packages very large,
release-specific, and impossible to maintain for vendor trees.

The plugin embeds only:

```text
common profile rules
release delta rules
module aliases and family patterns
stable source roots
capability and role metadata
```

Expected size for the initial catalog:

```text
150-300 module anchors
50-100 family/path rules
28 profile descriptors
uncompressed: a few hundred KB to about 1 MB
gzip-compressed: typically tens to a few hundred KB
```

## 5.1 Profile File Layout

There is one common base and small release deltas:

```text
profiles/
  common.json
  android-17.json
  android-18.json
```

`common.json` is the base catalog, not a second copy of every Android release.
It contains the current baseline release metadata so the origin of the rules
is explicit:

```json
{
  "schemaVersion": 1,
  "baseline": {
    "androidRelease": "android-17",
    "androidTag": "<validated Android 17 tag>",
    "generatedFrom": "AOSP IDE metadata pack",
    "profileDataVersion": "2026-08-29"
  },
  "profiles": {}
}
```

`android-17.json` and future release files contain only differences from the
base. They never repeat the full common catalog:

```json
{
  "schemaVersion": 1,
  "base": "common",
  "androidRelease": "android-17",
  "profiles": {
    "system-server": {
      "moduleNames": {"add": ["services.new"]},
      "sourceRoots": {"add": ["frameworks/base/services/new/java"]}
    }
  }
}
```

Delta arrays use explicit `add`/`remove` operations instead of generic JSON
array replacement. This prevents a small release change from accidentally
discarding common modules, roots, or dependencies. The engine applies:

```text
common.json
  -> selected release delta
  -> vendor/workspace overrides
  -> source validation and optional Out enrichment
```

If a release has no differences, its delta file is omitted and `common.json`
is used directly. The baseline release in `common.json` is always reported in
the UI and engine diagnostics.

## 5.2 Profile Generation and Maintenance

The development-time profile compiler reads representative Android 17 metadata
and emits the compact common catalog plus only detected release differences.
It never copies raw `module_bp_*`, `module-info`, compdb, or CMake data into the
runtime package. A new Android release requires a new delta review, not a new
full catalog.

The shared Go engine owns this data using one of these equivalent packaging
forms:

1. `go:embed` of minified JSON or gzip-compressed JSON. This guarantees that
   every engine platform package contains the same catalog and avoids runtime
   downloads.
2. A versioned `profiles/` asset beside the engine binary. This allows profile
   data updates without rebuilding the engine, while the package still works
   offline.

The preferred release form is compressed JSON embedded in the Go engine plus a
human-readable source copy in the repository. VS Code and JetBrains do not
duplicate the catalog; they ask the engine for the merged profile catalog
through RPC.

The repository source files are copied into `core/profiles/data` by
`go generate ./core/profiles` and embedded with `go:embed`. Release builds
therefore work without the checkout, while profile maintenance remains a
reviewable JSON change under `profiles/`. `make profile-compile` regenerates
the baseline and embedded copies; `make profile-validate` checks the baseline
schema and profile identifiers.

User overrides and derived caches are stored separately:

```text
built-in profiles       engine package, read-only
workspace overrides     IDE workspace configuration/state
source catalog cache    IDE extension/system cache
Out facts               user-provided Out Root, read-only
generated material      IDE cache or controlled .aosp-ide projection
```

There is no network service requirement and no plugin download button for
profiles or Out. A future release can ship a small profile update without
shipping any source code or build output.

## 6. Project Generation

The Go Project Model Compiler consumes the compact profile, source catalog, and
optional Out facts. It emits a neutral model that clients convert to native IDE
projects:

```text
ProjectModel
  Java targets       -> IntelliJ modules or Eclipse/JDT projects
  Native targets     -> CMake projects and clangd compilation databases
  Rust targets       -> rust-project.json
  Generated targets  -> selected srcjar/AIDL/sysprop source roots
```

The compiler does not implement ordinary Java or C++ definition semantics.
IDE language engines remain responsible for indexing and navigation.

## 7. Acceptance

- A clean source checkout with no Out can select and generate the built-in
  profiles.
- A matching Out enhances the same profile without changing source identity.
- A partial or incompatible Out only lowers capability for affected modules.
- The plugin package does not contain raw `module_bp_*`, full `module-info`,
  full `compile_commands`, or the complete CMake tree.
- VS Code and JetBrains use the same built-in profile catalog through the Go
  engine.
- Vendor trees can add custom modules and roots without modifying the shipped
  catalog.
