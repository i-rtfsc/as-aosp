package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type PackInfoParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
}

type PackInfoResult struct {
	SchemaVersion int                `json:"schemaVersion"`
	AOSPRoot      string             `json:"aospRoot,omitempty"`
	OutRoot       string             `json:"outRoot,omitempty"`
	Packs         []string           `json:"packs"`
	Diagnostics   []model.Diagnostic `json:"diagnostics"`
}

func (server *Server) packInfo(request jsonrpc.Request) jsonrpc.Response {
	var params PackInfoParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	result := PackInfoResult{Packs: []string{}, Diagnostics: []model.Diagnostic{}}
	root := strings.TrimSpace(params.MetadataPackRoot)
	if root == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "METADATA_PACK_ROOT_MISSING", Severity: model.SeverityError,
			Stage: "pack-rpc", Message: "metadataPackRoot is required",
		})
		return jsonrpc.Success(request.ID, result)
	}
	data, err := os.ReadFile(filepath.Join(root, "manifest.json"))
	if err != nil {
		result.Diagnostics = append(result.Diagnostics, packDiagnostic("PACK_MANIFEST_READ_FAILED", err.Error(), root))
		return jsonrpc.Success(request.ID, result)
	}
	var manifest struct {
		SchemaVersion int      `json:"schemaVersion"`
		AOSPRoot      string   `json:"aospRoot,omitempty"`
		OutRoot       string   `json:"outRoot,omitempty"`
		Packs         []string `json:"packs"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		result.Diagnostics = append(result.Diagnostics, packDiagnostic("PACK_MANIFEST_INVALID", err.Error(), root))
		return jsonrpc.Success(request.ID, result)
	}
	result.SchemaVersion = manifest.SchemaVersion
	result.AOSPRoot = manifest.AOSPRoot
	result.OutRoot = manifest.OutRoot
	result.Packs = manifest.Packs
	return jsonrpc.Success(request.ID, result)
}

func packDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "pack-rpc", Message: message, LocalPath: pathName}
}
