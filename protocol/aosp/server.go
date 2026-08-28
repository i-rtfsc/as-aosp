package aosp

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/jni"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

const ProtocolVersion = "1.0"

type Server struct {
	initialized           bool
	shutdown              bool
	jniSnapshot           *jniSnapshot
	aidlSnapshot          *aidlSnapshot
	moduleCatalogSnapshot *moduleCatalogSnapshot
	coordinator           *indexCoordinator
}

type jniSnapshot struct {
	root      string
	signature [32]byte
	index     jni.Index
}

type aidlSnapshot struct {
	root        string
	signature   [32]byte
	index       aidl.Index
	diagnostics []model.Diagnostic
}

func (server *Server) IsShutdown() bool {
	return server.shutdown
}

type InitializeParams struct {
	ProtocolVersion string `json:"protocolVersion"`
}

type InitializeResult struct {
	ProtocolVersion string       `json:"protocolVersion"`
	ServerInfo      ServerInfo   `json:"serverInfo"`
	Capabilities    Capabilities `json:"capabilities"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Capabilities struct {
	ProfileValidation     bool `json:"profileValidation"`
	TopologyDiscovery     bool `json:"topologyDiscovery"`
	TopologyValidation    bool `json:"topologyValidation"`
	AIDLNavigation        bool `json:"aidlNavigation"`
	JNINavigation         bool `json:"jniNavigation"`
	ModuleCatalog         bool `json:"moduleCatalog"`
	IndexCoordinator      bool `json:"indexCoordinator"`
	ProjectGraph          bool `json:"projectGraph"`
	SourceModuleDiscovery bool `json:"sourceModuleDiscovery"`
}

type ProfileResult struct {
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type TopologyDiscoverParams struct {
	WorkspaceViewRoot string `json:"workspaceViewRoot"`
	CanonicalRepoRoot string `json:"canonicalRepoRoot,omitempty"`
	MaxDepth          int    `json:"maxDepth,omitempty"`
}

type TopologyResult struct {
	Topology    topology.WorkspaceTopology `json:"topology"`
	Diagnostics []model.Diagnostic         `json:"diagnostics"`
}

type JNIPeersParams struct {
	WorkspaceViewRoot string            `json:"workspaceViewRoot,omitempty"`
	CurrentPath       string            `json:"currentPath,omitempty"`
	JavaFiles         map[string]string `json:"javaFiles,omitempty"`
	CppFiles          map[string]string `json:"cppFiles,omitempty"`
	ClassName         string            `json:"className,omitempty"`
	MethodName        string            `json:"methodName,omitempty"`
	FunctionName      string            `json:"functionName,omitempty"`
}

type JNIPeersResult struct {
	Peers       []jni.Link         `json:"peers"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func (server *Server) Handle(request jsonrpc.Request) (jsonrpc.Response, bool) {
	if request.Method == "aosp/initialize" {
		return server.initialize(request)
	}
	if request.Method == "aosp/shutdown" {
		if !server.initialized {
			return jsonrpc.Failure(request.ID, jsonrpc.NotInitialized()), !request.IsNotification()
		}
		server.shutdown = true
		return jsonrpc.Success(request.ID, nil), !request.IsNotification()
	}
	if server.shutdown {
		return jsonrpc.Failure(request.ID, &jsonrpc.Error{Code: -32600, Message: "server is shutting down"}), !request.IsNotification()
	}
	if !server.initialized {
		return jsonrpc.Failure(request.ID, jsonrpc.NotInitialized()), !request.IsNotification()
	}
	var response jsonrpc.Response
	switch request.Method {
	case "aosp/profile/validate":
		response = server.profileValidate(request)
	case "aosp/topology/discover":
		response = server.topologyDiscover(request)
	case "aosp/topology/validate":
		response = server.topologyValidate(request)
	case "aosp/jni/peers":
		response = server.jniPeers(request)
	case "aosp/aidl/resolve":
		response = server.aidlResolve(request)
	case "aosp/aidl/peers":
		response = server.aidlPeers(request)
	case "aosp/aidl/references":
		response = server.aidlReferences(request)
	case "aosp/compdb/overlay":
		response = server.compdbOverlay(request)
	case "aosp/modules/catalog":
		response = server.modulesCatalog(request)
	case "aosp/modules/get":
		response = server.moduleGet(request)
	case "aosp/modules/closure":
		response = server.modulesClosure(request)
	case "aosp/index/start":
		response = server.indexStart(request, false)
	case "aosp/index/rebuild":
		response = server.indexStart(request, true)
	case "aosp/index/status":
		response = server.indexStatus(request)
	case "aosp/index/cancel":
		response = server.indexCancel(request)
	case "aosp/index/projection":
		response = server.indexProjection(request)
	case "aosp/generated/materialize":
		response = server.generatedMaterialize(request)
	case "aosp/projects/graph":
		response = server.projectGraph(request)
	case "aosp/artifacts/catalog":
		response = server.artifactsCatalog(request)
	case "aosp/source/modules":
		response = server.sourceModules(request)
	default:
		response = jsonrpc.Failure(request.ID, jsonrpc.MethodNotFound(request.Method))
	}
	return response, !request.IsNotification()
}

func (server *Server) initialize(request jsonrpc.Request) (jsonrpc.Response, bool) {
	if server.initialized {
		return jsonrpc.Failure(request.ID, jsonrpc.AlreadyInitialized()), !request.IsNotification()
	}
	var params InitializeParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error())), !request.IsNotification()
	}
	if params.ProtocolVersion != "" && !compatibleVersion(params.ProtocolVersion, ProtocolVersion) {
		return jsonrpc.Failure(request.ID, &jsonrpc.Error{Code: -32600, Message: "incompatible protocol version"}), !request.IsNotification()
	}
	server.initialized = true
	return jsonrpc.Success(request.ID, InitializeResult{
		ProtocolVersion: ProtocolVersion,
		ServerInfo:      ServerInfo{Name: "aosp-index", Version: "dev"},
		Capabilities: Capabilities{
			ProfileValidation:     true,
			TopologyDiscovery:     true,
			TopologyValidation:    true,
			AIDLNavigation:        true,
			JNINavigation:         true,
			ModuleCatalog:         true,
			IndexCoordinator:      true,
			ProjectGraph:          true,
			SourceModuleDiscovery: true,
		},
	}), !request.IsNotification()
}

func (server *Server) jniPeers(request jsonrpc.Request) jsonrpc.Response {
	var params JNIPeersParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	javaFiles, diagnostics := normalizeJNIInputs(params.JavaFiles, "java")
	cppFiles, cppDiagnostics := normalizeJNIInputs(params.CppFiles, "cpp")
	diagnostics = append(diagnostics, cppDiagnostics...)
	if strings.TrimSpace(params.WorkspaceViewRoot) != "" {
		workspaceJava, workspaceCpp, workspaceDiagnostics := scanJNIWorkspace(params.WorkspaceViewRoot)
		diagnostics = append(diagnostics, workspaceDiagnostics...)
		for path, source := range workspaceJava {
			javaFiles[path] = source
		}
		for path, source := range workspaceCpp {
			cppFiles[path] = source
		}
	}
	if len(diagnostics) > 0 {
		return jsonrpc.Success(request.ID, JNIPeersResult{Peers: []jni.Link{}, Diagnostics: diagnostics})
	}
	index := server.buildJNIIndex(params.WorkspaceViewRoot, javaFiles, cppFiles)
	peers := make([]jni.Link, 0)
	if params.ClassName != "" || params.MethodName != "" {
		peers = append(peers, index.JavaPeers(params.ClassName, params.MethodName)...)
	}
	if params.FunctionName != "" {
		peers = append(peers, index.NativePeers(params.FunctionName)...)
	}
	if params.ClassName == "" && params.MethodName == "" && params.FunctionName == "" {
		peers = append(peers, index.Links...)
	}
	if params.CurrentPath != "" {
		current, err := model.NormalizeRepoPath(params.CurrentPath)
		if err != nil {
			return jsonrpc.Success(request.ID, JNIPeersResult{
				Peers: []jni.Link{},
				Diagnostics: []model.Diagnostic{{
					Code: "JNI_INVALID_REPO_PATH", Severity: model.SeverityError,
					Stage: "jni-rpc", Message: "current path is invalid: " + params.CurrentPath,
				}},
			})
		}
		peers = filterJNIPeersByPath(peers, current)
	}
	return jsonrpc.Success(request.ID, JNIPeersResult{Peers: peers, Diagnostics: index.Diagnostics})
}

func (server *Server) buildJNIIndex(root string, javaFiles, cppFiles map[model.RepoPath]string) jni.Index {
	signature := hashJNISources(javaFiles, cppFiles)
	if server.jniSnapshot != nil && server.jniSnapshot.root == root && server.jniSnapshot.signature == signature {
		return server.jniSnapshot.index
	}
	index := jni.BuildIndex(javaFiles, cppFiles)
	if root != "" {
		server.jniSnapshot = &jniSnapshot{root: root, signature: signature, index: index}
	}
	return index
}

func hashJNISources(javaFiles, cppFiles map[model.RepoPath]string) [32]byte {
	hash := sha256.New()
	keys := make([]string, 0, len(javaFiles)+len(cppFiles))
	for path := range javaFiles {
		keys = append(keys, "java:"+string(path))
	}
	for path := range cppFiles {
		keys = append(keys, "cpp:"+string(path))
	}
	sort.Strings(keys)
	for _, key := range keys {
		_, _ = hash.Write([]byte(key))
		if strings.HasPrefix(key, "java:") {
			_, _ = hash.Write([]byte(javaFiles[model.RepoPath(strings.TrimPrefix(key, "java:"))]))
		} else {
			_, _ = hash.Write([]byte(cppFiles[model.RepoPath(strings.TrimPrefix(key, "cpp:"))]))
		}
	}
	return sha256.Sum256(hash.Sum(nil))
}

func filterJNIPeersByPath(peers []jni.Link, current model.RepoPath) []jni.Link {
	result := make([]jni.Link, 0)
	for _, peer := range peers {
		if peer.Java.Path == current || peer.Native.Path == current {
			result = append(result, peer)
		}
	}
	return result
}

func scanJNIWorkspace(root string) (map[model.RepoPath]string, map[model.RepoPath]string, []model.Diagnostic) {
	javaFiles := make(map[model.RepoPath]string)
	cppFiles := make(map[model.RepoPath]string)
	scan := scanWorkspaceSources(
		root,
		"jni-workspace",
		"JNI_WORKSPACE_INVALID",
		"JNI_WORKSPACE_READ",
		isJNISourceExtension,
	)
	for repoPath, source := range scan.Files {
		extension := strings.ToLower(filepath.Ext(string(repoPath)))
		if extension == ".java" || extension == ".kt" {
			javaFiles[repoPath] = source
		} else {
			cppFiles[repoPath] = source
		}
	}
	return javaFiles, cppFiles, scan.Diagnostics
}

func isJNISourceExtension(extension string) bool {
	switch extension {
	case ".java", ".kt", ".c", ".cc", ".cpp", ".cxx", ".h", ".hh", ".hpp":
		return true
	default:
		return false
	}
}

func normalizeJNIInputs(inputs map[string]string, language string) (map[model.RepoPath]string, []model.Diagnostic) {
	result := make(map[model.RepoPath]string, len(inputs))
	diagnostics := make([]model.Diagnostic, 0)
	for rawPath, source := range inputs {
		path, err := model.NormalizeRepoPath(rawPath)
		if err != nil {
			diagnostics = append(diagnostics, model.Diagnostic{
				Code:     "JNI_INVALID_REPO_PATH",
				Severity: model.SeverityError,
				Stage:    "jni-rpc",
				Message:  language + " file path is invalid: " + rawPath,
			})
			continue
		}
		result[path] = source
	}
	return result, diagnostics
}

func (server *Server) profileValidate(request jsonrpc.Request) jsonrpc.Response {
	var params struct {
		Profile model.Profile `json:"profile"`
	}
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	return jsonrpc.Success(request.ID, ProfileResult{Diagnostics: params.Profile.Validate()})
}

func (server *Server) topologyDiscover(request jsonrpc.Request) jsonrpc.Response {
	var params TopologyDiscoverParams
	if err := decodeParams(request, &params); err != nil || strings.TrimSpace(params.WorkspaceViewRoot) == "" {
		if err == nil {
			err = fmt.Errorf("workspaceViewRoot is required")
		}
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	view, diagnostics, err := topology.Discover(topology.DiscoverOptions{
		WorkspaceViewRoot: params.WorkspaceViewRoot,
		CanonicalRepoRoot: params.CanonicalRepoRoot,
		MaxDepth:          params.MaxDepth,
	})
	if err != nil {
		return jsonrpc.Failure(request.ID, &jsonrpc.Error{Code: -32603, Message: err.Error()})
	}
	return jsonrpc.Success(request.ID, TopologyResult{Topology: view, Diagnostics: diagnostics})
}

func (server *Server) topologyValidate(request jsonrpc.Request) jsonrpc.Response {
	var params struct {
		Topology topology.WorkspaceTopology `json:"topology"`
	}
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	return jsonrpc.Success(request.ID, TopologyResult{Topology: params.Topology, Diagnostics: params.Topology.Validate()})
}

func decodeParams(request jsonrpc.Request, target any) error {
	if len(request.Params) == 0 {
		return fmt.Errorf("params are required")
	}
	if err := json.Unmarshal(request.Params, target); err != nil {
		return err
	}
	return nil
}

func compatibleVersion(client, server string) bool {
	clientMajor := strings.SplitN(client, ".", 2)[0]
	serverMajor := strings.SplitN(server, ".", 2)[0]
	_, clientErr := strconv.Atoi(clientMajor)
	_, serverErr := strconv.Atoi(serverMajor)
	return clientErr == nil && serverErr == nil && clientMajor == serverMajor
}
