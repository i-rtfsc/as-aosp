package aosp

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	corepack "github.com/i-rtfsc/as-aosp/core/pack"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type packValidationStamp struct {
	rootTime, manifestTime, checksumsTime int64
	manifestSize, checksumsSize           int64
}

type packValidationSnapshot struct {
	stamp  packValidationStamp
	result corepack.Result
}

type packValidationCache struct {
	mu      sync.Mutex
	entries map[string]packValidationSnapshot
}

func (server *Server) validatePackIfDeclared(root string) []corepack.Diagnostic {
	stamp, declared := currentPackStamp(root)
	if !declared {
		return nil
	}
	server.packValidation.mu.Lock()
	defer server.packValidation.mu.Unlock()
	if snapshot, ok := server.packValidation.entries[root]; ok && snapshot.stamp == stamp {
		return snapshot.result.Diagnostics
	}
	result := corepack.Validate(root)
	if server.packValidation.entries == nil {
		server.packValidation.entries = make(map[string]packValidationSnapshot)
	}
	server.packValidation.entries[root] = packValidationSnapshot{stamp: stamp, result: result}
	return result.Diagnostics
}

func currentPackStamp(root string) (packValidationStamp, bool) {
	rootInfo, rootErr := os.Stat(root)
	manifestInfo, manifestErr := os.Stat(filepath.Join(root, "manifest.json"))
	checksumsInfo, checksumsErr := os.Stat(filepath.Join(root, "checksums.sha256"))
	if os.IsNotExist(manifestErr) {
		return packValidationStamp{}, false
	}
	if rootErr != nil || manifestErr != nil || checksumsErr != nil {
		return packValidationStamp{}, true
	}
	return packValidationStamp{
		rootTime: rootInfo.ModTime().UnixNano(), manifestTime: manifestInfo.ModTime().UnixNano(),
		checksumsTime: checksumsInfo.ModTime().UnixNano(), manifestSize: manifestInfo.Size(),
		checksumsSize: checksumsInfo.Size(),
	}, true
}

type PackValidateParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
}

type PackValidateResult struct {
	Valid       bool                  `json:"valid"`
	Diagnostics []corepack.Diagnostic `json:"diagnostics"`
}

func (server *Server) packValidate(request jsonrpc.Request) jsonrpc.Response {
	var params PackValidateParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root := strings.TrimSpace(params.MetadataPackRoot)
	if root == "" {
		return jsonrpc.Success(request.ID, PackValidateResult{Diagnostics: []corepack.Diagnostic{{Code: "PACK_ROOT_MISSING", Message: "metadataPackRoot is required"}}})
	}
	result := corepack.Validate(root)
	if stamp, declared := currentPackStamp(root); declared {
		server.packValidation.mu.Lock()
		if server.packValidation.entries == nil {
			server.packValidation.entries = make(map[string]packValidationSnapshot)
		}
		server.packValidation.entries[root] = packValidationSnapshot{stamp: stamp, result: result}
		server.packValidation.mu.Unlock()
	}
	return jsonrpc.Success(request.ID, PackValidateResult{Valid: result.Valid, Diagnostics: result.Diagnostics})
}
