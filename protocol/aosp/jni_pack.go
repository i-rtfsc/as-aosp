package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/jni"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type JNIPackPeersParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
	ClassName        string `json:"className,omitempty"`
	MethodName       string `json:"methodName,omitempty"`
	FunctionName     string `json:"functionName,omitempty"`
}

type JNIPackPeersResult struct {
	Peers       []jni.Link         `json:"peers"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func (server *Server) jniPackPeers(request jsonrpc.Request) jsonrpc.Response {
	var params JNIPackPeersParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	result := JNIPackPeersResult{Peers: []jni.Link{}, Diagnostics: []model.Diagnostic{}}
	if strings.TrimSpace(params.MetadataPackRoot) == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "JNI_PACK_ROOT_MISSING", Severity: model.SeverityError, Stage: "jni-pack-rpc", Message: "metadataPackRoot is required"})
		return jsonrpc.Success(request.ID, result)
	}
	result.Diagnostics = append(result.Diagnostics, diagnosticsToModel(server.validatePackIfDeclared(params.MetadataPackRoot))...)
	if len(result.Diagnostics) > 0 {
		return jsonrpc.Success(request.ID, result)
	}
	data, err := os.ReadFile(filepath.Join(params.MetadataPackRoot, "jni", "index.json"))
	if err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "JNI_PACK_INDEX_READ_FAILED", Severity: model.SeverityWarning, Stage: "jni-pack-rpc", Message: err.Error()})
		return jsonrpc.Success(request.ID, result)
	}
	var index jni.Index
	if err := json.Unmarshal(data, &index); err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "JNI_PACK_INDEX_INVALID", Severity: model.SeverityWarning, Stage: "jni-pack-rpc", Message: err.Error()})
		return jsonrpc.Success(request.ID, result)
	}
	if params.FunctionName != "" {
		result.Peers = index.NativePeers(params.FunctionName)
	} else {
		result.Peers = index.JavaPeers(params.ClassName, params.MethodName)
	}
	return jsonrpc.Success(request.ID, result)
}
