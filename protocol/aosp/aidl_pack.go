package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type AIDLPackResolveParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
	Name             string `json:"name"`
}

type AIDLPackResolveResult struct {
	Name        string               `json:"name"`
	State       aidl.ResolutionState `json:"state"`
	Symbols     []aidl.Symbol        `json:"symbols,omitempty"`
	Diagnostics []model.Diagnostic   `json:"diagnostics"`
}

func (server *Server) aidlPackResolve(request jsonrpc.Request) jsonrpc.Response {
	var params AIDLPackResolveParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	result := AIDLPackResolveResult{Name: strings.TrimSpace(params.Name), State: aidl.ResolutionNotFound, Symbols: []aidl.Symbol{}, Diagnostics: []model.Diagnostic{}}
	if strings.TrimSpace(params.MetadataPackRoot) == "" || result.Name == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "AIDL_PACK_PARAMS_MISSING", Severity: model.SeverityError, Stage: "aidl-pack-rpc", Message: "metadataPackRoot and name are required"})
		return jsonrpc.Success(request.ID, result)
	}
	data, err := os.ReadFile(filepath.Join(params.MetadataPackRoot, "aidl", "symbols.json"))
	if err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "AIDL_PACK_SYMBOLS_READ_FAILED", Severity: model.SeverityWarning, Stage: "aidl-pack-rpc", Message: err.Error()})
		return jsonrpc.Success(request.ID, result)
	}
	var symbols map[string][]aidl.Symbol
	if err := json.Unmarshal(data, &symbols); err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "AIDL_PACK_SYMBOLS_INVALID", Severity: model.SeverityWarning, Stage: "aidl-pack-rpc", Message: err.Error()})
		return jsonrpc.Success(request.ID, result)
	}
	result.Symbols = symbols[result.Name]
	switch len(result.Symbols) {
	case 1:
		result.State = aidl.ResolutionFound
	case 0:
		result.State = aidl.ResolutionNotFound
	default:
		result.State = aidl.ResolutionAmbiguous
	}
	return jsonrpc.Success(request.ID, result)
}
