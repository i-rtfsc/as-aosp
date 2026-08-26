package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/metadata"
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

type AIDLPackReferencesParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
	Name             string `json:"name"`
}

type AIDLPackReferencesResult struct {
	References  []aidl.Reference   `json:"references"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type AIDLPackPeersParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
	CacheRoot        string `json:"cacheRoot,omitempty"`
	Name             string `json:"name"`
}

type AIDLPackPeersResult struct {
	Peers       []aidl.GeneratedPeer `json:"peers"`
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
	if len(result.Symbols) == 0 && !strings.Contains(result.Name, ".") {
		for canonical, values := range symbols {
			if strings.HasSuffix(canonical, "."+result.Name) {
				result.Symbols = append(result.Symbols, values...)
			}
		}
	}
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

func loadAIDLPackIndex(root string) (aidl.Index, []model.Diagnostic) {
	data, err := os.ReadFile(filepath.Join(root, "aidl", "index.json"))
	if err != nil {
		return aidl.NewIndex(), []model.Diagnostic{{Code: "AIDL_PACK_INDEX_READ_FAILED", Severity: model.SeverityWarning, Stage: "aidl-pack-rpc", Message: err.Error()}}
	}
	var index aidl.Index
	if err := json.Unmarshal(data, &index); err != nil || index.Documents == nil || index.Symbols == nil {
		return aidl.NewIndex(), []model.Diagnostic{{Code: "AIDL_PACK_INDEX_INVALID", Severity: model.SeverityWarning, Stage: "aidl-pack-rpc", Message: "aidl/index.json is invalid"}}
	}
	return index, nil
}

func (server *Server) aidlPackReferences(request jsonrpc.Request) jsonrpc.Response {
	var params AIDLPackReferencesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	index, diagnostics := loadAIDLPackIndex(params.MetadataPackRoot)
	return jsonrpc.Success(request.ID, AIDLPackReferencesResult{References: index.References(strings.TrimSpace(params.Name)), Diagnostics: diagnostics})
}

func (server *Server) aidlPackPeers(request jsonrpc.Request) jsonrpc.Response {
	var params AIDLPackPeersParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	index, diagnostics := loadAIDLPackIndex(params.MetadataPackRoot)
	peers := make([]aidl.GeneratedPeer, 0)
	archiveEntries := []aidl.ArchiveEntry{}
	sourceArchives := []metadata.PathRef{}
	if strings.TrimSpace(params.CacheRoot) != "" {
		_, entries, generatedDiagnostics := materializeGeneratedSources(params.MetadataPackRoot, params.CacheRoot)
		archiveEntries = entries
		diagnostics = append(diagnostics, generatedDiagnostics...)
		seen := map[string]bool{}
		for _, entry := range entries {
			if !seen[entry.Archive.Raw] {
				seen[entry.Archive.Raw] = true
				sourceArchives = append(sourceArchives, entry.Archive)
			}
		}
	}
	for _, document := range index.Documents {
		resolution := index.Resolve(document, params.Name)
		if resolution.State != aidl.ResolutionFound || len(resolution.Symbols) != 1 {
			continue
		}
		report := index.GeneratedPeers(document, resolution.Name, aidl.PeerOptions{ArchiveEntries: archiveEntries, SourceArchives: sourceArchives})
		peers = append(peers, report.Peers...)
		diagnostics = append(diagnostics, report.Diagnostics...)
	}
	return jsonrpc.Success(request.ID, AIDLPackPeersResult{Peers: peers, Diagnostics: diagnostics})
}
