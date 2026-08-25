package aosp

import (
	"archive/zip"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	coremodules "github.com/i-rtfsc/as-aosp/core/modules"
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

func TestServerQueriesModuleCatalogFromPack(t *testing.T) {
	root := t.TempDir()
	pathName := filepath.Join(root, "modules", "catalog.json")
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	catalog := coremodules.Catalog{Modules: map[string]coremodules.Module{
		"framework": {Name: "framework", Dependencies: []string{"core"}},
		"core":      {Name: "core"},
	}}
	data, _ := json.Marshal(catalog)
	if err := os.WriteFile(pathName, data, 0o644); err != nil {
		t.Fatal(err)
	}
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/modules/catalog", ModulesParams{
		MetadataPackRoot: root,
	}))
	var listed ModulesCatalogResult
	decodeResult(t, response, &listed)
	if len(listed.Diagnostics) != 0 || len(listed.Modules) != 2 ||
		listed.Modules[0].Name != "core" || listed.Modules[1].Name != "framework" {
		t.Fatalf("listed=%+v", listed)
	}
	response, _ = server.Handle(makeRequest(3, "aosp/modules/closure", ModulesParams{
		MetadataPackRoot: root, Names: []string{"framework"},
	}))
	var closure ModulesClosureResult
	decodeResult(t, response, &closure)
	if len(closure.Diagnostics) != 0 || len(closure.Modules) != 2 ||
		closure.Modules[0].Name != "core" || closure.Modules[1].Name != "framework" {
		t.Fatalf("closure=%+v", closure)
	}
	response, _ = server.Handle(makeRequest(4, "aosp/modules/get", ModulesParams{
		MetadataPackRoot: root, Name: "framework",
	}))
	var got ModuleResult
	decodeResult(t, response, &got)
	if len(got.Diagnostics) != 0 || got.Module == nil || got.Module.Name != "framework" {
		t.Fatalf("got=%+v", got)
	}
}

func TestServerReportsMissingModulePack(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/modules/catalog", ModulesParams{}))
	var result ModulesCatalogResult
	decodeResult(t, response, &result)
	if len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "METADATA_PACK_ROOT_MISSING" {
		t.Fatalf("result=%+v", result)
	}
}

func TestServerBuildsIndexStatusFromMetadataPack(t *testing.T) {
	root := t.TempDir()
	pathName := filepath.Join(root, "modules", "catalog.json")
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	catalog := coremodules.Catalog{Modules: map[string]coremodules.Module{
		"framework": {
			Name: "framework", SrcJars: []metadata.PathRef{metadata.NewPathRef("out/framework.srcjar")},
		},
	}}
	data, _ := json.Marshal(catalog)
	if err := os.WriteFile(pathName, data, 0o644); err != nil {
		t.Fatal(err)
	}
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/index/start", IndexStartParams{
		Profile: model.Profile{ID: "framework", MetadataPackRoot: root},
	}))
	if response.Error != nil {
		t.Fatalf("start=%+v", response)
	}
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

func TestServerMaterializesPackSrcJars(t *testing.T) {
	pack := t.TempDir()
	archive := filepath.Join(pack, "generated", "srcjars", "framework.srcjar")
	if err := os.MkdirAll(filepath.Dir(archive), 0o755); err != nil {
		t.Fatal(err)
	}
	writeProtocolSrcJar(t, archive, "android/os/Build.java", "package android.os; class Build {}")
	manifest, _ := json.Marshal([]map[string]string{{"path": "generated/srcjars/framework.srcjar"}})
	if err := os.WriteFile(filepath.Join(pack, "generated", "manifest.json"), manifest, 0o644); err != nil {
		t.Fatal(err)
	}
	server := initializedServer(t)
	cache := t.TempDir()
	response, _ := server.Handle(makeRequest(2, "aosp/generated/materialize", GeneratedMaterializeParams{
		MetadataPackRoot: pack, CacheRoot: cache,
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

func TestServerResolvesAIDLFromPack(t *testing.T) {
	pack := t.TempDir()
	if err := os.MkdirAll(filepath.Join(pack, "aidl"), 0o755); err != nil {
		t.Fatal(err)
	}
	symbols := map[string][]aidl.Symbol{"demo.IFoo": {{CanonicalName: "demo.IFoo", Kind: aidl.SymbolInterface, Path: "demo/IFoo.aidl", Origin: "parser", Confidence: "exact"}}}
	data, _ := json.Marshal(symbols)
	if err := os.WriteFile(filepath.Join(pack, "aidl", "symbols.json"), data, 0o644); err != nil {
		t.Fatal(err)
	}
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/aidl/pack/resolve", AIDLPackResolveParams{MetadataPackRoot: pack, Name: "demo.IFoo"}))
	var result AIDLPackResolveResult
	decodeResult(t, response, &result)
	if result.State != aidl.ResolutionFound || len(result.Symbols) != 1 || len(result.Diagnostics) != 0 {
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
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{MetadataRoot: root, PathMappings: []model.PathMapping{{ExternalPrefix: "/srv/aosp", LocalPrefix: "/work/aosp"}}}))
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
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{MetadataRoot: root, PathMappings: []model.PathMapping{{ExternalPrefix: "C:\\build", LocalPrefix: "D:\\aosp"}}}))
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
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{MetadataRoot: root, PathMappings: []model.PathMapping{{ExternalPrefix: "/srv", LocalPrefix: "/short"}, {ExternalPrefix: "/srv/aosp", LocalPrefix: "/long"}}}))
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
	response, _ := server.Handle(makeRequest(2, "aosp/compdb/overlay", CompDBOverlayParams{MetadataRoot: root}))
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
