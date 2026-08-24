package aosp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/jni"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

const ProtocolVersion = "1.0"

type Server struct {
	initialized bool
	shutdown    bool
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
	ProfileValidation  bool `json:"profileValidation"`
	TopologyDiscovery  bool `json:"topologyDiscovery"`
	TopologyValidation bool `json:"topologyValidation"`
	AIDLNavigation     bool `json:"aidlNavigation"`
	JNINavigation      bool `json:"jniNavigation"`
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
			ProfileValidation:  true,
			TopologyDiscovery:  true,
			TopologyValidation: true,
			JNINavigation:      true,
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
	index := jni.BuildIndex(javaFiles, cppFiles)
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
	diagnostics := make([]model.Diagnostic, 0)
	absolute, err := filepath.Abs(root)
	if err != nil {
		return javaFiles, cppFiles, []model.Diagnostic{{
			Code: "JNI_WORKSPACE_INVALID", Severity: model.SeverityError,
			Stage: "jni-workspace", Message: err.Error(),
		}}
	}
	info, err := os.Stat(absolute)
	if err != nil || !info.IsDir() {
		message := "workspace view root is not a directory"
		if err != nil {
			message = err.Error()
		}
		return javaFiles, cppFiles, []model.Diagnostic{{
			Code: "JNI_WORKSPACE_INVALID", Severity: model.SeverityError,
			Stage: "jni-workspace", Message: message,
		}}
	}
	visited := make(map[string]bool)
	var walk func(string, string) error
	walk = func(current, relative string) error {
		real, evalErr := filepath.EvalSymlinks(current)
		if evalErr == nil {
			real = filepath.Clean(real)
			if visited[real] {
				return nil
			}
			visited[real] = true
		}
		entries, readErr := os.ReadDir(current)
		if readErr != nil {
			return readErr
		}
		for _, entry := range entries {
			childRelative := filepath.Join(relative, entry.Name())
			if isJNIExcluded(childRelative) {
				continue
			}
			child := filepath.Join(current, entry.Name())
			info, statErr := os.Stat(child)
			if statErr != nil {
				diagnostics = append(diagnostics, jniWorkspaceDiagnostic("JNI_WORKSPACE_READ", statErr.Error()))
				continue
			}
			if info.IsDir() {
				if err := walk(child, childRelative); err != nil {
					diagnostics = append(diagnostics, jniWorkspaceDiagnostic("JNI_WORKSPACE_READ", err.Error()))
				}
				continue
			}
			extension := strings.ToLower(filepath.Ext(entry.Name()))
			if !isJNISourceExtension(extension) {
				continue
			}
			source, readErr := os.ReadFile(child)
			if readErr != nil {
				diagnostics = append(diagnostics, jniWorkspaceDiagnostic("JNI_WORKSPACE_READ", readErr.Error()))
				continue
			}
			repoPath, pathErr := model.NormalizeRepoPath(filepath.ToSlash(childRelative))
			if pathErr != nil {
				diagnostics = append(diagnostics, jniWorkspaceDiagnostic("JNI_INVALID_REPO_PATH", pathErr.Error()))
				continue
			}
			if extension == ".java" || extension == ".kt" {
				javaFiles[repoPath] = string(source)
			} else {
				cppFiles[repoPath] = string(source)
			}
		}
		return nil
	}
	err = walk(absolute, "")
	if err != nil {
		diagnostics = append(diagnostics, jniWorkspaceDiagnostic("JNI_WORKSPACE_READ", err.Error()))
	}
	return javaFiles, cppFiles, diagnostics
}

func isJNIExcluded(relative string) bool {
	for _, part := range strings.Split(filepath.ToSlash(relative), "/") {
		if part == ".git" || part == ".repo" || part == "out" || part == ".aosp-ide" {
			return true
		}
	}
	return false
}

func isJNISourceExtension(extension string) bool {
	switch extension {
	case ".java", ".kt", ".c", ".cc", ".cpp", ".cxx", ".h", ".hh", ".hpp":
		return true
	default:
		return false
	}
}

func jniWorkspaceDiagnostic(code, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "jni-workspace", Message: message}
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
