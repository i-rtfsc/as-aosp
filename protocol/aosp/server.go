package aosp

import (
	"encoding/json"
	"fmt"
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
	JavaFiles    map[string]string `json:"javaFiles,omitempty"`
	CppFiles     map[string]string `json:"cppFiles,omitempty"`
	ClassName    string            `json:"className,omitempty"`
	MethodName   string            `json:"methodName,omitempty"`
	FunctionName string            `json:"functionName,omitempty"`
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
	return jsonrpc.Success(request.ID, JNIPeersResult{Peers: peers, Diagnostics: index.Diagnostics})
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
