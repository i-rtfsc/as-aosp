package aosp

import (
	"archive/zip"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

func makeRequest(id int, method string, params any) jsonrpc.Request {
	data, _ := json.Marshal(params)
	idData, _ := json.Marshal(id)
	return jsonrpc.Request{JSONRPC: jsonrpc.Version, ID: idData, Method: method, Params: data}
}

func TestServerLifecycleAndCapabilities(t *testing.T) {
	server := &Server{}
	response, reply := server.Handle(makeRequest(1, "aosp/profile/validate", map[string]any{}))
	if !reply || response.Error == nil || response.Error.Message != "server is not initialized" {
		t.Fatalf("response=%+v reply=%v", response, reply)
	}
	response, reply = server.Handle(makeRequest(2, "aosp/initialize", InitializeParams{ProtocolVersion: "1.0"}))
	if !reply || response.Error != nil {
		t.Fatalf("initialize=%+v reply=%v", response, reply)
	}
	var initialized InitializeResult
	decodeResult(t, response, &initialized)
	if !initialized.Capabilities.AIDLNavigation || !initialized.Capabilities.JNINavigation ||
		!initialized.Capabilities.ModuleCatalog {
		t.Fatalf("capabilities=%+v", initialized.Capabilities)
	}
	response, _ = server.Handle(makeRequest(3, "aosp/initialize", InitializeParams{ProtocolVersion: "1.0"}))
	if response.Error == nil {
		t.Fatalf("duplicate initialize=%+v", response)
	}
	response, reply = server.Handle(makeRequest(4, "aosp/shutdown", nil))
	if !reply || response.Error != nil || !server.shutdown {
		t.Fatalf("shutdown=%+v reply=%v", response, reply)
	}
}

func TestServerQueriesModuleCatalogFromOut(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "module-info.json", `{"framework":{"dependencies":["core"]},"core":{}}`)
	writeTestFile(t, root, "module_bp_java_deps.json", `{}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/modules/catalog", ModulesParams{
		OutRoot: root,
	}))
	var listed ModulesCatalogResult
	decodeResult(t, response, &listed)
	if len(listed.Diagnostics) != 0 || len(listed.Modules) != 2 ||
		listed.Modules[0].Name != "core" || listed.Modules[1].Name != "framework" {
		t.Fatalf("listed=%+v", listed)
	}
	response, _ = server.Handle(makeRequest(3, "aosp/modules/closure", ModulesParams{
		OutRoot: root, Names: []string{"framework"},
	}))
	var closure ModulesClosureResult
	decodeResult(t, response, &closure)
	if len(closure.Diagnostics) != 0 || len(closure.Modules) != 2 ||
		closure.Modules[0].Name != "core" || closure.Modules[1].Name != "framework" {
		t.Fatalf("closure=%+v", closure)
	}
	response, _ = server.Handle(makeRequest(4, "aosp/modules/get", ModulesParams{
		OutRoot: root, Name: "framework",
	}))
	var got ModuleResult
	decodeResult(t, response, &got)
	if len(got.Diagnostics) != 0 || got.Module == nil || got.Module.Name != "framework" {
		t.Fatalf("got=%+v", got)
	}
}

func TestServerEnhancesSourceModulesWithOutFacts(t *testing.T) {
	workspace := t.TempDir()
	out := t.TempDir()
	writeTestFile(t, workspace, "frameworks/base/Android.bp", `
java_library {
    name: "framework-minus-apex",
    srcs: ["core/java/**/*.java"],
}`)
	writeTestFile(t, out, "module-info.json", `{
  "framework-minus-apex":{"path":["frameworks/base"]}
}`)
	writeTestFile(t, out, "module_bp_java_deps.json", `{
  "framework-minus-apex":{
    "path":["frameworks/base/core/java"],
    "srcjars":["out/soong/framework.srcjar"]
  }
}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/source/modules", SourceModulesParams{
		WorkspaceViewRoot: workspace, OutRoot: out, CacheRoot: t.TempDir(),
	}))
	var result SourceModuleResult
	decodeResult(t, response, &result)
	if len(result.Modules) != 1 || result.BuildBacked != 1 ||
		result.Modules[0].Status != "build-backed" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerPersistsBuildFactCatalogInWorkspaceCache(t *testing.T) {
	out := t.TempDir()
	cache := t.TempDir()
	writeTestFile(t, out, "module-info.json", `{"framework":{}}`)
	writeTestFile(t, out, "module_bp_java_deps.json", `{"framework":{}}`)
	server := initializedServer(t)
	catalog, diagnostics := server.loadBuildFactModuleCatalogCached(out, cache)
	if len(diagnostics) != 0 || len(catalog.Modules) != 1 {
		t.Fatalf("catalog=%+v diagnostics=%+v", catalog, diagnostics)
	}
	if _, err := os.Stat(filepath.Join(cache, "build-facts", "catalog.json")); err != nil {
		t.Fatalf("build facts cache missing: %v", err)
	}
	second := initializedServer(t)
	catalog, diagnostics = second.loadBuildFactModuleCatalogCached(out, cache)
	if len(diagnostics) != 0 || len(catalog.Modules) != 1 {
		t.Fatalf("cached catalog=%+v diagnostics=%+v", catalog, diagnostics)
	}
}

func TestServerPagesSourceModuleResults(t *testing.T) {
	workspace := t.TempDir()
	writeTestFile(t, workspace, "Android.bp", `
java_library { name: "one" }
java_library { name: "two" }
`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/source/modules", SourceModulesParams{
		WorkspaceViewRoot: workspace, Limit: 1,
	}))
	var result SourceModuleResult
	decodeResult(t, response, &result)
	if result.Total != 2 || len(result.Modules) != 1 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerMaterializesLocalRustProject(t *testing.T) {
	workspace := t.TempDir()
	out := t.TempDir()
	cache := t.TempDir()
	writeTestFile(t, out, "soong/rust-project.json", `{
  "sysroot":"prebuilts/rust",
  "crates":[{"root_module":"system/lib.rs","env":{"OUT_DIR":"out/soong/gen"}}]
}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/rust/project", RustProjectParams{
		WorkspaceViewRoot: workspace, OutRoot: out, CacheRoot: cache,
	}))
	var result RustProjectResult
	decodeResult(t, response, &result)
	if result.Crates != 1 || result.ProjectPath == "" || len(result.Diagnostics) != 0 {
		t.Fatalf("result=%+v", result)
	}
	if _, err := os.Stat(result.ProjectPath); err != nil {
		t.Fatal(err)
	}
}

func TestServerBuildsProjectGraphFromOut(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "module-info.json", `{"framework":{"path":["frameworks/base/core/java"]}}`)
	writeTestFile(t, root, "module_bp_java_deps.json", `{"framework":{"jars":["out/soong/framework.jar"]}}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/projects/graph", ProjectGraphParams{
		OutRoot: root,
		Targets: []model.TargetSpec{{
			ID: "framework", Kind: model.TargetJava,
			Selectors: model.TargetSelectors{Artifacts: []model.ArtifactSelector{{Kind: model.ArtifactJar, Name: "framework.jar"}}},
		}},
	}))
	var result ProjectGraphResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Graph.Targets) != 1 ||
		len(result.Graph.Targets[0].ResolvedArtifacts) != 1 ||
		result.Graph.Targets[0].ResolvedArtifacts[0].Name != "framework.jar" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerBuildsProjectGraphFromLocalOut(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "module-info.json", `{
		"SystemUI":{"path":["frameworks/base/packages/SystemUI"],"installed":["out/target/product/demo/system_ext/priv-app/SystemUI/SystemUI.apk"]}
	}`)
	writeTestFile(t, root, "module_bp_java_deps.json", `{
		"SystemUI":{"path":["frameworks/base/packages/SystemUI"]}
	}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/projects/graph", ProjectGraphParams{
		OutRoot: root,
		Targets: []model.TargetSpec{{
			ID: "systemui", Kind: model.TargetAndroid,
			Selectors: model.TargetSelectors{Artifacts: []model.ArtifactSelector{{Kind: model.ArtifactAPK, Name: "SystemUI.apk"}}},
		}},
	}))
	var result ProjectGraphResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Graph.Targets) != 1 {
		t.Fatalf("result=%+v", result)
	}
	target := result.Graph.Targets[0]
	if len(target.ResolvedArtifacts) != 1 || len(target.ResolvedModules) != 1 ||
		len(target.SourceRoots) != 1 || len(target.Libraries) != 0 {
		t.Fatalf("target=%+v", target)
	}
}

func TestServerListsArtifactsByKindAndQuery(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "module-info.json", `{"SystemUI":{"installed":["out/target/product/demo/SystemUI.apk"]}}`)
	writeTestFile(t, root, "module_bp_java_deps.json", `{}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/artifacts/catalog", ArtifactsCatalogParams{
		OutRoot: root, Kind: "apk", Query: "system",
	}))
	var result ArtifactsCatalogResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Artifacts) != 1 || result.Artifacts[0].Name != "SystemUI.apk" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerReadsCompdbFromOut(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "compile_commands.json", `[{"directory":"/srv/aosp","arguments":["clang","-I","/srv/aosp/include"],"file":"/srv/aosp/demo.cpp"}]`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{
		OutRoot:           root,
		WorkspaceViewRoot: t.TempDir(),
		PathMappings:      []model.PathMapping{{ExternalPrefix: "/srv/aosp", LocalPrefix: "/workspace/aosp"}},
	}))
	var result CompDBOverlayResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Entries) != 1 || result.Entries[0].File != filepath.FromSlash("/workspace/aosp/demo.cpp") {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerReportsMissingModulePack(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/modules/catalog", ModulesParams{}))
	var result ModulesCatalogResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "BUILD_FACTS_ROOT_MISSING" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerBuildsIndexStatusFromOut(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "module-info.json", `{}`)
	writeTestFile(t, root, "module_bp_java_deps.json", `{"framework":{"srcjars":["out/soong/framework.srcjar"]}}`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/index/start", IndexStartParams{
		Profile: model.Profile{ID: "framework", OutRoot: root},
	}))
	if response.Error != nil {
		t.Fatalf("start=%+v", response)
	}
	result := waitForIndexState(t, server, "ready")
	if result.ModuleCount != 1 || result.GeneratedCount != 1 || result.Progress != 100 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerBuildsIndexStatusFromLocalOut(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "module-info.json", `{"framework":{"path":["frameworks/base"]}}`)
	writeTestFile(t, root, "module_bp_java_deps.json", `{"framework":{"srcjars":["out/soong/framework.srcjar"]}}`)
	server := initializedServer(t)
	server.Handle(makeRequest(2, "aosp/index/start", IndexStartParams{
		Profile: model.Profile{ID: "framework", OutRoot: root},
	}))
	result := waitForIndexState(t, server, "ready")
	if result.ModuleCount != 1 || result.GeneratedCount != 1 || result.Progress != 100 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerKeepsMissingPackIndexPartial(t *testing.T) {
	server := initializedServer(t)
	server.Handle(makeRequest(2, "aosp/index/start", IndexStartParams{
		Profile: model.Profile{ID: "source-only"},
	}))
	result := waitForIndexState(t, server, "partial")
	if result.Progress != 100 || len(result.Diagnostics) == 0 || result.ModuleCount != 0 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerMaterializesOutSrcJars(t *testing.T) {
	pack := t.TempDir()
	archive := filepath.Join(pack, "soong", "framework.srcjar")
	if err := os.MkdirAll(filepath.Dir(archive), 0o755); err != nil {
		t.Fatal(err)
	}
	writeProtocolSrcJar(t, archive, "android/os/Build.java", "package android.os; class Build {}")
	server := initializedServer(t)
	cache := t.TempDir()
	response, _ := server.Handle(makeRequest(2, "aosp/generated/materialize", GeneratedMaterializeParams{
		OutRoot: pack, Archives: []string{"soong/framework.srcjar"}, CacheRoot: cache,
	}))
	var result GeneratedMaterializeResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.SourceRoots) != 1 || result.SourceRoots[0].Entries != 1 {
		t.Fatalf("result=%+v", result)
	}
	if _, err := os.Stat(filepath.Join(result.SourceRoots[0].Root, "entries", "android", "os", "Build.java")); err != nil {
		t.Fatalf("materialized source missing: %v", err)
	}
}

func TestServerMaterializesSelectedLocalOutSrcJar(t *testing.T) {
	out := t.TempDir()
	archive := filepath.Join(out, "soong", "framework.srcjar")
	if err := os.MkdirAll(filepath.Dir(archive), 0o755); err != nil {
		t.Fatal(err)
	}
	writeProtocolSrcJar(t, archive, "android/os/Build.java", "package android.os; class Build {}")
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/generated/materialize", GeneratedMaterializeParams{
		OutRoot: out, Archives: []string{"soong/framework.srcjar"}, CacheRoot: t.TempDir(),
	}))
	var result GeneratedMaterializeResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.SourceRoots) != 1 || result.SourceRoots[0].Entries != 1 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerRejectsLocalSrcJarOutsideOut(t *testing.T) {
	out := t.TempDir()
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/generated/materialize", GeneratedMaterializeParams{
		OutRoot: out, Archives: []string{"../outside.srcjar"}, CacheRoot: t.TempDir(),
	}))
	var result GeneratedMaterializeResult
	decodeResult(t, response, &result)
	if len(result.SourceRoots) != 0 || len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "GENERATED_ARCHIVE_PATH_INVALID" {
		t.Fatalf("result=%+v", result)
	}
}

func writeProtocolSrcJar(t *testing.T, pathName, entry, source string) {
	t.Helper()
	file, err := os.Create(pathName)
	if err != nil {
		t.Fatal(err)
	}
	archive := zip.NewWriter(file)
	writer, err := archive.Create(entry)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := writer.Write([]byte(source)); err != nil {
		t.Fatal(err)
	}
	if err := archive.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}

func waitForIndexState(t *testing.T, server *Server, expected string) IndexSnapshot {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		response, _ := server.Handle(makeRequest(9, "aosp/index/status", nil))
		var result IndexSnapshot
		decodeResult(t, response, &result)
		if result.State == expected {
			return result
		}
		time.Sleep(time.Millisecond)
	}
	t.Fatalf("index did not reach state %q", expected)
	return IndexSnapshot{}
}

func TestServerRejectsIncompatibleVersion(t *testing.T) {
	server := &Server{}
	response, _ := server.Handle(makeRequest(1, "aosp/initialize", InitializeParams{ProtocolVersion: "2.0"}))
	if response.Error == nil || server.initialized {
		t.Fatalf("response=%+v initialized=%v", response, server.initialized)
	}
}

func TestServerValidatesProfilesAndTopologies(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/profile/validate", map[string]any{
		"profile": model.Profile{},
	}))
	var profile ProfileResult
	decodeResult(t, response, &profile)
	if len(profile.Diagnostics) == 0 {
		t.Fatalf("profile=%+v", profile)
	}
	response, _ = server.Handle(makeRequest(3, "aosp/topology/validate", map[string]any{
		"topology": topology.WorkspaceTopology{},
	}))
	var view TopologyResult
	decodeResult(t, response, &view)
	if len(view.Diagnostics) != 0 {
		t.Fatalf("topology=%+v", view)
	}
}

func TestServerDiscoversAnEmptyWorkspace(t *testing.T) {
	root := t.TempDir()
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/topology/discover", TopologyDiscoverParams{
		WorkspaceViewRoot: root,
		MaxDepth:          2,
	}))
	var result TopologyResult
	decodeResult(t, response, &result)
	expected, err := os.Stat(root)
	if err != nil || !expected.IsDir() || string(result.Topology.WorkspaceViewRoot) == "" {
		t.Fatalf("result=%+v err=%v", result, err)
	}
}

func TestServerReturnsJNIPeersFromVisibleSources(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/jni/peers", JNIPeersParams{
		JavaFiles: map[string]string{
			"demo/Bridge.java": "package demo; class Bridge { native void close(); }",
		},
		CppFiles: map[string]string{
			"demo/bridge.cpp": `
static const JNINativeMethod methods[] = {
  {"close", "()V", (void*)closeNative},
};
int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", methods, 1);
}`,
		},
		ClassName:  "demo.Bridge",
		MethodName: "close",
	}))
	var result JNIPeersResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Peers) != 1 {
		t.Fatalf("result=%+v", result)
	}
	if result.Peers[0].Native.FunctionName != "closeNative" || result.Peers[0].Confidence != "exact" {
		t.Fatalf("peers=%+v", result.Peers)
	}
}

func TestServerReturnsAllJNIPeersWithoutSelector(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/jni/peers", JNIPeersParams{
		JavaFiles: map[string]string{
			"demo/Bridge.java": "package demo; class Bridge { native void close(); }",
		},
		CppFiles: map[string]string{
			"demo/bridge.cpp": `
static const JNINativeMethod methods[] = {
  {"close", "()V", (void*)closeNative},
};
int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", methods, 1);
}`,
		},
	}))
	var result JNIPeersResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Peers) != 1 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerScansWorkspaceViewAndFiltersCurrentPath(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "frameworks/base/demo/Bridge.java",
		"package demo; class Bridge { native void close(); }")
	writeTestFile(t, root, "frameworks/base/demo/bridge.cpp", `
static const JNINativeMethod methods[] = {
  {"close", "()V", (void*)closeNative},
};
int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", methods, 1);
}`)
	writeTestFile(t, root, "out/generated/ignored.cpp",
		"void Java_demo_Bridge_close(JNIEnv*, jobject) {}")

	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/jni/peers", JNIPeersParams{
		WorkspaceViewRoot: root,
		CurrentPath:       "frameworks/base/demo/Bridge.java",
	}))
	var result JNIPeersResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Peers) != 1 {
		t.Fatalf("result=%+v", result)
	}
	if result.Peers[0].Native.Path != "frameworks/base/demo/bridge.cpp" {
		t.Fatalf("peers=%+v", result.Peers)
	}
}

func TestServerReusesUnchangedJNIWorkspaceSnapshot(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "demo/Bridge.java",
		"package demo; class Bridge { native void close(); }")
	writeTestFile(t, root, "demo/bridge.cpp", `
void Java_demo_Bridge_close(JNIEnv*, jobject) {}
`)
	server := initializedServer(t)
	params := JNIPeersParams{WorkspaceViewRoot: root}
	first, _ := server.Handle(makeRequest(2, "aosp/jni/peers", params))
	second, _ := server.Handle(makeRequest(3, "aosp/jni/peers", params))
	var firstResult, secondResult JNIPeersResult
	decodeResult(t, first, &firstResult)
	decodeResult(t, second, &secondResult)
	if server.jniSnapshot == nil || len(firstResult.Peers) != len(secondResult.Peers) {
		t.Fatalf("snapshot=%+v first=%+v second=%+v", server.jniSnapshot, firstResult, secondResult)
	}
	signature := server.jniSnapshot.signature
	second, _ = server.Handle(makeRequest(4, "aosp/jni/peers", params))
	if server.jniSnapshot == nil || server.jniSnapshot.signature != signature {
		t.Fatalf("snapshot changed for unchanged workspace: %+v", server.jniSnapshot)
	}
}

func TestServerScansSymlinkedWorkspaceDirectory(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("symlink fixture requires platform permission on Windows")
	}
	root := t.TempDir()
	target := t.TempDir()
	writeTestFile(t, target, "demo/Bridge.java",
		"package demo; class Bridge { native void close(); }")
	writeTestFile(t, target, "demo/bridge.cpp", `
void Java_demo_Bridge_close(JNIEnv*, jobject) {}
`)
	if err := os.Symlink(filepath.Join(target, "demo"), filepath.Join(root, "demo")); err != nil {
		t.Fatal(err)
	}
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/jni/peers", JNIPeersParams{
		WorkspaceViewRoot: root,
		CurrentPath:       "demo/Bridge.java",
	}))
	var result JNIPeersResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Peers) != 1 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerResolvesAIDLSourceRangeFromWorkspace(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "demo/IFoo.aidl", "package demo; interface IFoo {}")
	writeTestFile(t, root, "demo/Client.aidl", "package demo; import demo.IFoo; interface Client { IFoo value(); }")
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/aidl/resolve", AIDLResolveParams{
		WorkspaceViewRoot: root,
		CurrentPath:       "demo/Client.aidl",
		Name:              "IFoo",
	}))
	var result AIDLResolveResult
	decodeResult(t, response, &result)
	if result.Resolution.State != aidl.ResolutionFound || len(result.Resolution.Symbols) != 1 {
		t.Fatalf("result=%+v", result)
	}
	if result.Resolution.Symbols[0].Path != "demo/IFoo.aidl" {
		t.Fatalf("symbols=%+v", result.Resolution.Symbols)
	}
}

func TestServerReusesUnchangedAIDLWorkspaceSnapshot(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "demo/IFoo.aidl", "package demo; interface IFoo {}")
	writeTestFile(t, root, "demo/Client.aidl", "package demo; import demo.IFoo; interface Client { IFoo value(); }")
	server := initializedServer(t)
	params := AIDLResolveParams{WorkspaceViewRoot: root, CurrentPath: "demo/Client.aidl", Name: "IFoo"}
	first, _ := server.Handle(makeRequest(2, "aosp/aidl/resolve", params))
	var firstResult AIDLResolveResult
	decodeResult(t, first, &firstResult)
	if server.aidlSnapshot == nil || firstResult.Resolution.State != aidl.ResolutionFound {
		t.Fatalf("snapshot=%+v result=%+v", server.aidlSnapshot, firstResult)
	}
	signature := server.aidlSnapshot.signature
	second, _ := server.Handle(makeRequest(3, "aosp/aidl/resolve", params))
	var secondResult AIDLResolveResult
	decodeResult(t, second, &secondResult)
	if server.aidlSnapshot == nil || server.aidlSnapshot.signature != signature ||
		secondResult.Resolution.State != aidl.ResolutionFound {
		t.Fatalf("snapshot=%+v result=%+v", server.aidlSnapshot, secondResult)
	}
}

func TestServerReturnsAIDLGeneratedPeersWithInferredFallback(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "demo/IFoo.aidl", "package demo; interface IFoo {}")
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/aidl/peers", AIDLPeersParams{
		WorkspaceViewRoot: root,
		CurrentPath:       "demo/IFoo.aidl",
		Name:              "demo.IFoo",
	}))
	var result AIDLPeersResult
	decodeResult(t, response, &result)
	if len(result.Peers) != 6 {
		t.Fatalf("result=%+v", result)
	}
	for _, peer := range result.Peers {
		if peer.Confidence != "inferred" || peer.Exists {
			t.Fatalf("peer=%+v", peer)
		}
	}
}

func TestServerReturnsAIDLReferences(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "demo/IFoo.aidl", "package demo; interface IFoo {}")
	writeTestFile(t, root, "demo/Client.aidl", "package demo; import demo.IFoo; interface Client { IFoo value(); }")
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/aidl/references", AIDLReferencesParams{WorkspaceViewRoot: root, CurrentPath: "demo/IFoo.aidl", Name: "demo.IFoo"}))
	var result AIDLReferencesResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.References) != 2 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerMapsCompDBOverlayPaths(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "compile_commands.json", `[{"directory":"/srv/aosp/out","arguments":["clang++","-I/srv/aosp/include"],"file":"/srv/aosp/frameworks/base/foo.cpp","output":"/srv/aosp/out/foo.o"}]`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{OutRoot: root, PathMappings: []model.PathMapping{{ExternalPrefix: "/srv/aosp", LocalPrefix: "/work/aosp"}}}))
	var result CompDBOverlayResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Entries) != 1 {
		t.Fatalf("result=%+v", result)
	}
	if result.Entries[0].File != "/work/aosp/frameworks/base/foo.cpp" || result.Entries[0].Arguments[1] != "-I/work/aosp/include" {
		t.Fatalf("entry=%+v", result.Entries[0])
	}
}

func TestServerMapsWindowsCompDBPathsOnAnyHost(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "compile_commands.json", `[{"directory":"C:\\build\\out","arguments":["clang++","-IC:\\build\\include"],"file":"C:\\build\\src\\foo.cpp"}]`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{OutRoot: root, PathMappings: []model.PathMapping{{ExternalPrefix: "C:\\build", LocalPrefix: "D:\\aosp"}}}))
	var result CompDBOverlayResult
	decodeResult(t, response, &result)
	if len(result.Entries) != 1 || filepath.ToSlash(result.Entries[0].File) != "D:/aosp/src/foo.cpp" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerMapsSeparatedAndLongestCompDBArguments(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "compile_commands.json", `[{"directory":"/srv/aosp/out","arguments":["clang++","-I","/srv/aosp/include","--sysroot=/srv/aosp/sysroot","@/srv/aosp/args.rsp"],"file":"/srv/aosp/src/foo.cpp"}]`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{OutRoot: root, PathMappings: []model.PathMapping{{ExternalPrefix: "/srv", LocalPrefix: "/short"}, {ExternalPrefix: "/srv/aosp", LocalPrefix: "/long"}}}))
	var result CompDBOverlayResult
	decodeResult(t, response, &result)
	if result.Entries[0].Arguments[2] != "/long/include" || result.Entries[0].Arguments[3] != "--sysroot=/long/sysroot" || result.Entries[0].Arguments[4] != "@/long/args.rsp" {
		t.Fatalf("result=%+v", result)
	}
}

func TestResponseFileTokenizerPreservesQuotedPaths(t *testing.T) {
	parts := tokenizeResponseFile(`-I "/build/include with spaces" --sysroot='/build/sys root' escaped\\ path`)
	if len(parts) != 5 || parts[1] != "/build/include with spaces" || parts[2] != "--sysroot=/build/sys root" || parts[3] != "escaped\\" || parts[4] != "path" {
		t.Fatalf("parts=%q", parts)
	}
}

func TestServerReportsDuplicateCompDBVariants(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "compile_commands.json", `[{"directory":".","arguments":["clang++","-DHOST"],"file":"foo.cpp"},{"directory":".","arguments":["clang++","-DDEVICE"],"file":"foo.cpp"}]`)
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{OutRoot: root}))
	var result CompDBOverlayResult
	decodeResult(t, response, &result)
	if len(result.Entries) != 2 || len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "COMPDB_VARIANT_DUPLICATE" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerUpgradesVisibleAIDLGeneratedJavaPeer(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "demo/IFoo.aidl", "package demo; interface IFoo {}")
	writeTestFile(t, root, "demo/IFoo.java", "package demo; final class IFoo {}")
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/aidl/peers", AIDLPeersParams{
		WorkspaceViewRoot: root,
		CurrentPath:       "demo/IFoo.aidl",
		Name:              "demo.IFoo",
	}))
	var result AIDLPeersResult
	decodeResult(t, response, &result)
	for _, peer := range result.Peers {
		if peer.Language == aidl.PeerJava {
			if peer.Confidence != "exact" || !peer.Exists {
				t.Fatalf("peer=%+v", peer)
			}
			return
		}
	}
	t.Fatal("Java peer not found")
}

func TestServerKeepsDistinctSymlinkAliases(t *testing.T) {
	if os.PathSeparator == '\\' {
		t.Skip("symlink fixture requires platform permission on Windows")
	}
	root := t.TempDir()
	target := t.TempDir()
	writeTestFile(t, target, "demo/Bridge.java",
		"package demo; class Bridge { native void close(); }")
	writeTestFile(t, target, "demo/bridge.cpp", `
void Java_demo_Bridge_close(JNIEnv*, jobject) {}
`)
	if err := os.Symlink(filepath.Join(target, "demo"), filepath.Join(root, "one")); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(target, "demo"), filepath.Join(root, "two")); err != nil {
		t.Fatal(err)
	}
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/jni/peers", JNIPeersParams{
		WorkspaceViewRoot: root,
	}))
	var result JNIPeersResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 0 || len(result.Peers) != 4 {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerRejectsAbsoluteJNIPaths(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/jni/peers", JNIPeersParams{
		JavaFiles: map[string]string{"/tmp/Bridge.java": "class Bridge {}"},
	}))
	var result JNIPeersResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "JNI_INVALID_REPO_PATH" {
		t.Fatalf("result=%+v", result)
	}
}

func initializedServer(t *testing.T) *Server {
	t.Helper()
	server := &Server{}
	response, _ := server.Handle(makeRequest(1, "aosp/initialize", InitializeParams{
		ProtocolVersion: ProtocolVersion,
	}))
	if response.Error != nil {
		t.Fatalf("initialize=%+v", response)
	}
	return server
}

func decodeResult(t *testing.T, response jsonrpc.Response, target any) {
	t.Helper()
	if response.Error != nil {
		t.Fatalf("response error=%+v", response.Error)
	}
	if err := json.Unmarshal(response.Result, target); err != nil {
		t.Fatalf("decode result: %v", err)
	}
}

func writeTestFile(t *testing.T, root, relative, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(relative))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
