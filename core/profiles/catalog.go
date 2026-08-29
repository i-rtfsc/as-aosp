package profiles

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

type Baseline struct {
	AndroidRelease     string `json:"androidRelease"`
	AndroidTag         string `json:"androidTag"`
	GeneratedFrom      string `json:"generatedFrom"`
	ProfileDataVersion string `json:"profileDataVersion"`
}

type Profile struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name"`
	Kind                  string   `json:"kind"`
	ModuleNames           []string `json:"moduleNames,omitempty"`
	ModulePatterns        []string `json:"modulePatterns,omitempty"`
	SourceRoots           []string `json:"sourceRoots,omitempty"`
	SourcePatterns        []string `json:"sourcePatterns,omitempty"`
	NativeRoots           []string `json:"nativeRoots,omitempty"`
	IncludeRoots          []string `json:"includeRoots,omitempty"`
	Dependencies          []string `json:"dependencies,omitempty"`
	GeneratedCapabilities []string `json:"generatedCapabilities,omitempty"`
	Recommended           bool     `json:"recommended"`
	Default               bool     `json:"default"`
	Priority              int      `json:"priority"`
}

type Catalog struct {
	SchemaVersion int       `json:"schemaVersion"`
	Baseline      Baseline  `json:"baseline"`
	Profiles      []Profile `json:"profiles"`
}

type ListDelta struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

type ProfileDelta struct {
	ModuleNames    ListDelta `json:"moduleNames,omitempty"`
	ModulePatterns ListDelta `json:"modulePatterns,omitempty"`
	SourceRoots    ListDelta `json:"sourceRoots,omitempty"`
	SourcePatterns ListDelta `json:"sourcePatterns,omitempty"`
	NativeRoots    ListDelta `json:"nativeRoots,omitempty"`
	IncludeRoots   ListDelta `json:"includeRoots,omitempty"`
	Dependencies   ListDelta `json:"dependencies,omitempty"`
}

type ReleaseDelta struct {
	SchemaVersion  int                     `json:"schemaVersion"`
	Base           string                  `json:"base"`
	AndroidRelease string                  `json:"androidRelease"`
	Profiles       map[string]ProfileDelta `json:"profiles,omitempty"`
}

type CompileReport struct {
	Catalog         Catalog            `json:"catalog"`
	Diagnostics     []model.Diagnostic `json:"diagnostics"`
	ProfilesChecked int                `json:"profilesChecked"`
	AnchorsChecked  int                `json:"anchorsChecked"`
	AnchorsFound    int                `json:"anchorsFound"`
}

func DefaultCatalog(release, tag string) Catalog {
	if strings.TrimSpace(release) == "" {
		release = "android-17"
	}
	if strings.TrimSpace(tag) == "" {
		tag = release
	}
	return Catalog{
		SchemaVersion: 1,
		Baseline: Baseline{
			AndroidRelease: release, AndroidTag: tag,
			GeneratedFrom:      "AOSP IDE metadata pack",
			ProfileDataVersion: "2026-08-29",
		},
		Profiles: []Profile{
			{ID: "java-foundation", Name: "Java Foundation", Kind: "java", Default: true, Recommended: true, Priority: 700,
				ModuleNames:    []string{"core-oj", "core-libart", "ext", "framework", "framework-minus-apex"},
				SourceRoots:    []string{"libcore", "frameworks/base/core/java"},
				SourcePatterns: []string{"frameworks/base/**/java", "frameworks/base/**/src"}},
			{ID: "system-server-wms", Name: "System Server & WMS", Kind: "java", Default: true, Recommended: true, Priority: 650,
				ModuleNames: []string{"services", "services.impl", "services.core"}, ModulePatterns: []string{"services.*", "service-*"},
				SourceRoots: []string{"frameworks/base/services"}, SourcePatterns: []string{"frameworks/base/services/**/java"}, Dependencies: []string{"java-foundation"},
				GeneratedCapabilities: []string{"aidl", "sysprop", "aconfig"}},
			{ID: "activity-manager", Name: "ActivityManager", Kind: "java", Recommended: true, Priority: 620,
				ModuleNames: []string{"services.core"}, SourceRoots: []string{"frameworks/base/services/core/java/com/android/server/am"}, Dependencies: []string{"system-server-wms"}},
			{ID: "window-manager", Name: "WindowManager", Kind: "mixed", Recommended: true, Priority: 610,
				ModuleNames: []string{"services.core", "WindowManager-Shell"}, SourceRoots: []string{"frameworks/base/services/core/java/com/android/server/wm"}, NativeRoots: []string{"frameworks/native/libs/ui", "frameworks/native/libs/gui"}, Dependencies: []string{"system-server-wms", "graphics"}},
			{ID: "system-server-input", Name: "System Server Input", Kind: "mixed", Recommended: true, Priority: 605,
				ModuleNames: []string{"services.core"}, SourceRoots: []string{"frameworks/base/services/core/java/com/android/server/input"}, NativeRoots: []string{"frameworks/base/services/core/jni", "frameworks/native/services/inputflinger"}, Dependencies: []string{"system-server-wms", "input-stack"}},
			{ID: "system-ui", Name: "SystemUI", Kind: "android", Default: true, Recommended: true, Priority: 500,
				ModuleNames: []string{"SystemUI", "SystemUI-application", "SystemUI-core"}, ModulePatterns: []string{"SystemUI*"}, SourceRoots: []string{"frameworks/base/packages/SystemUI"}, SourcePatterns: []string{"frameworks/base/packages/SystemUI/**/src"}, Dependencies: []string{"java-foundation", "window-manager"}},
			{ID: "system-ui-shell", Name: "SystemUI Shell", Kind: "mixed", Recommended: true, Priority: 490,
				ModuleNames: []string{"WindowManager-Shell", "SystemUISharedLib"}, NativeRoots: []string{"frameworks/native/libs/gui", "frameworks/native/services/surfaceflinger"}, Dependencies: []string{"system-ui", "graphics"}},
			{ID: "system-ui-keyguard", Name: "SystemUI Keyguard", Kind: "java", Recommended: true, Priority: 480,
				ModuleNames: []string{"SystemUI-core"}, SourceRoots: []string{"frameworks/base/packages/SystemUI/src/com/android/systemui/keyguard"}, Dependencies: []string{"system-ui"}},
			{ID: "system-ui-compose", Name: "SystemUI Compose", Kind: "android", Recommended: true, Priority: 470,
				ModuleNames: []string{"SystemUI-core"}, SourceRoots: []string{"frameworks/base/packages/SystemUI/compose"}, Dependencies: []string{"system-ui"}},
			{ID: "native-foundation", Name: "Native Foundation", Kind: "native", Default: true, Recommended: true, Priority: 400,
				ModuleNames: []string{"libbase", "libutils", "liblog", "libcutils", "libbinder", "libbinder_ndk", "libnativehelper"}, NativeRoots: []string{"system/libbase", "system/core/libutils", "system/logging/liblog", "system/core/libcutils", "frameworks/native/libs/binder"}},
			{ID: "graphics", Name: "Graphics", Kind: "native", Recommended: true, Priority: 350,
				ModuleNames: []string{"libui", "libgui", "libnativewindow", "libhwui", "libandroidfw"}, NativeRoots: []string{"frameworks/native/libs/ui", "frameworks/native/libs/gui", "frameworks/native/libs/nativewindow", "frameworks/base/libs/hwui", "frameworks/base/libs/androidfw"}, Dependencies: []string{"native-foundation"}},
			{ID: "graphics-surfaceflinger", Name: "Graphics & SurfaceFlinger", Kind: "native", Default: true, Recommended: true, Priority: 340,
				ModuleNames: []string{"libui", "libgui", "libnativewindow", "surfaceflinger", "librenderengine", "libskia_renderengine"}, NativeRoots: []string{"frameworks/native/libs/ui", "frameworks/native/libs/gui", "frameworks/native/libs/nativewindow", "frameworks/native/services/surfaceflinger"}, Dependencies: []string{"native-foundation"}},
			{ID: "render-engine", Name: "RenderEngine", Kind: "native", Recommended: true, Priority: 330, ModuleNames: []string{"librenderengine", "libskia_renderengine"}, NativeRoots: []string{"frameworks/native/services/surfaceflinger/RenderEngine"}, Dependencies: []string{"graphics"}},
			{ID: "input-stack", Name: "Input Stack", Kind: "native", Default: true, Recommended: true, Priority: 300,
				ModuleNames: []string{"libinput", "libinputflinger", "libinputdispatcher", "libinputreader", "libinputreporter"}, NativeRoots: []string{"frameworks/native/services/inputflinger"}, Dependencies: []string{"native-foundation"}},
			{ID: "media-framework", Name: "Media Framework", Kind: "mixed", Recommended: true, Priority: 280, ModuleNames: []string{"libmedia", "libmediandk", "MediaProvider"}, NativeRoots: []string{"frameworks/av/media/libmedia", "frameworks/av/media/libstagefright"}, Dependencies: []string{"native-foundation"}},
			{ID: "audio", Name: "Audio", Kind: "mixed", Recommended: true, Priority: 270, ModuleNames: []string{"libaudioclient", "libaudioflinger", "libaudiopolicy", "libaudiohal"}, ModulePatterns: []string{"android.hardware.audio.*"}, NativeRoots: []string{"frameworks/av/services/audioflinger", "frameworks/av/services/audiopolicy", "frameworks/av/media/libaudioclient"}, Dependencies: []string{"media-framework"}},
			{ID: "camera", Name: "Camera", Kind: "mixed", Recommended: true, Priority: 260, ModuleNames: []string{"libcamera_client", "libcameraservice", "libcamera_metadata"}, ModulePatterns: []string{"android.hardware.camera.*"}, NativeRoots: []string{"frameworks/av/camera", "frameworks/av/services/camera"}, Dependencies: []string{"graphics", "native-foundation"}},
			{ID: "connectivity", Name: "Connectivity", Kind: "mixed", Recommended: true, Priority: 240, ModuleNames: []string{"framework-connectivity", "Tethering", "NetworkStack", "CaptivePortalLogin"}, ModulePatterns: []string{"connectivity*", "framework-connectivity*"}, SourceRoots: []string{"packages/modules/Connectivity", "packages/modules/NetworkStack"}, NativeRoots: []string{"system/netd"}},
			{ID: "wifi", Name: "Wi-Fi", Kind: "mixed", Recommended: true, Priority: 230, ModuleNames: []string{"wifi-service-pre-jarjar", "WifiTrackerLib", "wifi-modules-utils"}, ModulePatterns: []string{"wifi-*", "Wifi*"}, SourceRoots: []string{"packages/modules/Wifi", "frameworks/opt/net/wifi"}, NativeRoots: []string{"hardware/interfaces/wifi"}},
			{ID: "bluetooth", Name: "Bluetooth", Kind: "mixed", Recommended: true, Priority: 220, ModuleNames: []string{"Bluetooth", "BluetoothServices", "BluetoothLib"}, ModulePatterns: []string{"Bluetooth*", "bluetooth-*"}, SourceRoots: []string{"packages/modules/Bluetooth"}, NativeRoots: []string{"packages/modules/Bluetooth/system"}},
			{ID: "telephony", Name: "Telephony", Kind: "mixed", Recommended: true, Priority: 210, ModuleNames: []string{"telephony-common", "TelephonyProvider", "ims-common"}, ModulePatterns: []string{"android.hardware.radio.*"}, SourceRoots: []string{"packages/modules/Telephony", "frameworks/opt/telephony", "frameworks/opt/net/ims"}, NativeRoots: []string{"hardware/ril"}},
			{ID: "package-manager", Name: "Package Manager", Kind: "mixed", Recommended: true, Priority: 190, ModuleNames: []string{"PackageInstaller", "services.core"}, SourceRoots: []string{"frameworks/base/packages/PackageInstaller", "frameworks/base/services/core/java/com/android/server/pm"}, Dependencies: []string{"system-server-wms"}},
			{ID: "permission-security", Name: "Permission and Security", Kind: "mixed", Recommended: true, Priority: 180, ModuleNames: []string{"PermissionController", "Settings", "SettingsProvider"}, SourceRoots: []string{"packages/modules/Permission", "packages/apps/Settings", "packages/providers/SettingsProvider"}, NativeRoots: []string{"system/security", "external/selinux"}},
			{ID: "storage-vold", Name: "Storage and Vold", Kind: "mixed", Recommended: true, Priority: 170, ModuleNames: []string{"MediaProvider", "DocumentsUI", "vold"}, SourceRoots: []string{"packages/providers/MediaProvider", "packages/apps/DocumentsUI"}, NativeRoots: []string{"system/vold", "system/core/fs_mgr"}},
			{ID: "art-libcore", Name: "ART and libcore", Kind: "mixed", Recommended: true, Priority: 160, ModuleNames: []string{"core-oj", "core-libart", "libart", "libartbase", "libdexfile", "libnativebridge"}, SourceRoots: []string{"art", "libcore", "dalvik"}, NativeRoots: []string{"art/runtime", "art/libartbase", "art/libdexfile", "art/libnativebridge"}},
			{ID: "aidl-hidl", Name: "AIDL and HIDL", Kind: "mixed", Recommended: true, Priority: 150, ModulePatterns: []string{"android.hardware.*-V*-java", "android.hardware.*-V*-cpp", "android.hardware.*-V*-ndk", "android.frameworks.*-V*", "*-aidl-*"}, SourceRoots: []string{"hardware/interfaces", "frameworks/hardware/interfaces"}, Dependencies: []string{"java-foundation", "native-foundation"}, GeneratedCapabilities: []string{"aidl"}},
			{ID: "automotive", Name: "Automotive", Kind: "mixed", Recommended: true, Priority: 100, ModuleNames: []string{"android.car", "CarService", "CarServiceUpdatable", "CarSystemUI", "CarSettings"}, ModulePatterns: []string{"Car*", "car-*"}, SourceRoots: []string{"packages/services/Car", "packages/apps/Car"}, NativeRoots: []string{"hardware/interfaces/automotive"}},
			{ID: "common-system-apps", Name: "Common System Apps", Kind: "android", ModuleNames: []string{"Settings", "SettingsProvider", "PackageInstaller", "DocumentsUI", "Launcher3", "PermissionController", "MediaProvider", "NetworkStack", "Bluetooth", "TelephonyProvider"}, SourceRoots: []string{"packages/apps", "packages/providers", "packages/modules"}},
		},
	}
}

func Compile(outRoot, release, tag string) (CompileReport, error) {
	catalog := DefaultCatalog(release, tag)
	report := CompileReport{Catalog: catalog, Diagnostics: []model.Diagnostic{}}
	if strings.TrimSpace(outRoot) == "" {
		return report, nil
	}
	metadataReport := metadata.LoadModuleDirectory(outRoot)
	buildCatalog := modules.Build(metadataReport.Snapshot)
	present := make(map[string]bool, len(buildCatalog.Modules))
	for name := range buildCatalog.Modules {
		present[name] = true
	}
	for index := range report.Catalog.Profiles {
		profile := &report.Catalog.Profiles[index]
		report.ProfilesChecked++
		for _, name := range profile.ModuleNames {
			report.AnchorsChecked++
			if present[name] {
				report.AnchorsFound++
				continue
			}
			report.Diagnostics = append(report.Diagnostics, model.Diagnostic{
				Code: "PROFILE_MODULE_MISSING", Severity: model.SeverityWarning,
				Stage: "profile-compile", Message: fmt.Sprintf("profile %s anchor is not in Out: %s", profile.ID, name),
			})
		}
	}
	return report, nil
}

func WriteCatalog(path string, catalog Catalog) error {
	data, err := json.MarshalIndent(catalog, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	temporary := path + ".tmp"
	if err := os.WriteFile(temporary, append(data, '\n'), 0o644); err != nil {
		return err
	}
	return os.Rename(temporary, path)
}

func LoadCatalog(path string) (Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Catalog{}, err
	}
	var catalog Catalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return Catalog{}, err
	}
	return catalog, nil
}

func ValidateCatalog(catalog Catalog) []model.Diagnostic {
	diagnostics := make([]model.Diagnostic, 0)
	seen := make(map[string]bool, len(catalog.Profiles))
	for _, profile := range catalog.Profiles {
		if strings.TrimSpace(profile.ID) == "" {
			diagnostics = append(diagnostics, model.Diagnostic{Code: "PROFILE_ID_EMPTY", Severity: model.SeverityError, Stage: "profile-validate", Message: "profile id is required"})
		}
		if seen[profile.ID] {
			diagnostics = append(diagnostics, model.Diagnostic{Code: "PROFILE_ID_DUPLICATE", Severity: model.SeverityError, Stage: "profile-validate", Message: "duplicate profile: " + profile.ID})
		}
		seen[profile.ID] = true
	}
	sort.Slice(diagnostics, func(i, j int) bool { return diagnostics[i].Message < diagnostics[j].Message })
	return diagnostics
}
