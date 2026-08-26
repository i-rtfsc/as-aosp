package aosp

import (
	"strings"

	corepack "github.com/i-rtfsc/as-aosp/core/pack"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

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
	return jsonrpc.Success(request.ID, PackValidateResult{Valid: result.Valid, Diagnostics: result.Diagnostics})
}
