package aosp

import (
	"strings"
	"sync"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type IndexStartParams struct {
	Profile model.Profile `json:"profile"`
}

type IndexProjectionParams struct {
	ProfileID      string             `json:"profileId"`
	SourceRoots    int                `json:"sourceRoots"`
	GeneratedRoots int                `json:"generatedRoots"`
	Libraries      int                `json:"libraries"`
	Skipped        int                `json:"skipped"`
	Diagnostics    []model.Diagnostic `json:"diagnostics"`
}

type IndexSnapshot struct {
	ProfileID          string             `json:"profileId"`
	State              string             `json:"state"`
	Stage              string             `json:"stage"`
	CompletedStages    []string           `json:"completedStages"`
	TotalStages        int                `json:"totalStages"`
	Progress           int                `json:"progress"`
	ModuleCount        int                `json:"moduleCount"`
	GeneratedCount     int                `json:"generatedCount"`
	SourceRootCount    int                `json:"sourceRootCount"`
	GeneratedRootCount int                `json:"generatedRootCount"`
	LibraryCount       int                `json:"libraryCount"`
	SkippedCount       int                `json:"skippedCount"`
	Diagnostics        []model.Diagnostic `json:"diagnostics"`
}

type indexCoordinator struct {
	mu       sync.RWMutex
	snapshot IndexSnapshot
	cancel   chan struct{}
}

func (server *Server) indexStart(request jsonrpc.Request, rebuild bool) jsonrpc.Response {
	var params IndexStartParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	profileID := strings.TrimSpace(params.Profile.ID)
	if profileID == "" {
		profileID = "default"
	}
	if server.coordinator == nil {
		server.coordinator = &indexCoordinator{}
	}
	server.coordinator.mu.Lock()
	if server.coordinator.cancel != nil {
		close(server.coordinator.cancel)
	}
	cancel := make(chan struct{})
	server.coordinator.cancel = cancel
	server.coordinator.snapshot = IndexSnapshot{
		ProfileID: profileID, State: "running", Stage: "profile",
		CompletedStages: []string{}, TotalStages: 4, Progress: 0,
		Diagnostics: []model.Diagnostic{},
	}
	server.coordinator.mu.Unlock()
	go server.runIndex(profileID, params.Profile, cancel, rebuild)
	return jsonrpc.Success(request.ID, server.indexStatusResult())
}

func (server *Server) indexStatus(request jsonrpc.Request) jsonrpc.Response {
	if server.coordinator == nil {
		return jsonrpc.Success(request.ID, IndexSnapshot{
			State: "idle", Stage: "", TotalStages: 4,
			CompletedStages: []string{}, Diagnostics: []model.Diagnostic{},
		})
	}
	return jsonrpc.Success(request.ID, server.indexStatusResult())
}

func (server *Server) indexProjection(request jsonrpc.Request) jsonrpc.Response {
	var params IndexProjectionParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	if server.coordinator == nil {
		return jsonrpc.Success(request.ID, IndexSnapshot{State: "idle", TotalStages: 4, CompletedStages: []string{}, Diagnostics: []model.Diagnostic{}})
	}
	server.coordinator.mu.Lock()
	defer server.coordinator.mu.Unlock()
	if params.ProfileID != "" && server.coordinator.snapshot.ProfileID != params.ProfileID {
		return jsonrpc.Success(request.ID, server.coordinator.snapshot)
	}
	server.coordinator.snapshot.Stage = "projection"
	server.coordinator.snapshot.SourceRootCount = params.SourceRoots
	server.coordinator.snapshot.GeneratedRootCount = params.GeneratedRoots
	server.coordinator.snapshot.LibraryCount = params.Libraries
	server.coordinator.snapshot.SkippedCount = params.Skipped
	server.coordinator.snapshot.Diagnostics = append(server.coordinator.snapshot.Diagnostics, params.Diagnostics...)
	server.coordinator.snapshot.CompletedStages = appendUniqueStage(server.coordinator.snapshot.CompletedStages, "projection")
	server.coordinator.snapshot.Progress = 100
	if len(params.Diagnostics) > 0 {
		server.coordinator.snapshot.State = "partial"
	} else if server.coordinator.snapshot.State == "running" || server.coordinator.snapshot.State == "ready" {
		server.coordinator.snapshot.State = "ready"
	}
	return jsonrpc.Success(request.ID, server.coordinator.snapshot)
}

func appendUniqueStage(values []string, stage string) []string {
	for _, value := range values {
		if value == stage {
			return values
		}
	}
	return append(values, stage)
}

func (server *Server) indexCancel(request jsonrpc.Request) jsonrpc.Response {
	if server.coordinator == nil {
		return jsonrpc.Success(request.ID, IndexSnapshot{
			State: "idle", TotalStages: 4, CompletedStages: []string{},
			Diagnostics: []model.Diagnostic{},
		})
	}
	server.coordinator.mu.Lock()
	if server.coordinator.cancel != nil {
		close(server.coordinator.cancel)
		server.coordinator.cancel = nil
	}
	server.coordinator.snapshot.State = "cancelled"
	server.coordinator.snapshot.Stage = "cancelled"
	server.coordinator.mu.Unlock()
	return jsonrpc.Success(request.ID, server.indexStatusResult())
}

func (server *Server) indexStatusResult() IndexSnapshot {
	server.coordinator.mu.RLock()
	defer server.coordinator.mu.RUnlock()
	snapshot := server.coordinator.snapshot
	snapshot.CompletedStages = append([]string(nil), snapshot.CompletedStages...)
	snapshot.Diagnostics = append([]model.Diagnostic(nil), snapshot.Diagnostics...)
	return snapshot
}

func (server *Server) runIndex(profileID string, profile model.Profile, cancel chan struct{}, rebuild bool) {
	_ = rebuild
	server.updateIndex(profileID, "profile", 1, nil, cancel)
	if !server.indexStillRunning(cancel) {
		return
	}
	catalog, diagnostics := server.loadBuildFactModuleCatalog(profile.OutRoot)
	if !server.updateIndex(profileID, "catalog", 2, diagnostics, cancel) {
		return
	}
	if !server.indexStillRunning(cancel) {
		return
	}
	generated := 0
	for _, module := range catalog.Modules {
		generated += len(module.Generated) + len(module.SrcJars)
	}
	server.updateIndexCounts(profileID, len(catalog.Modules), generated, cancel)
	if !server.updateIndex(profileID, "generated", 3, nil, cancel) {
		return
	}
	server.updateIndex(profileID, "ready", 4, nil, cancel)
}

func (server *Server) indexStillRunning(cancel chan struct{}) bool {
	select {
	case <-cancel:
		return false
	default:
		return true
	}
}

func (server *Server) updateIndex(profileID, stage string, progress int, diagnostics []model.Diagnostic, cancel chan struct{}) bool {
	if !server.indexStillRunning(cancel) {
		return false
	}
	server.coordinator.mu.Lock()
	defer server.coordinator.mu.Unlock()
	if server.coordinator.snapshot.ProfileID != profileID {
		return false
	}
	server.coordinator.snapshot.Stage = stage
	server.coordinator.snapshot.Progress = progress * 100 / server.coordinator.snapshot.TotalStages
	for _, completed := range server.coordinator.snapshot.CompletedStages {
		if completed == stage {
			return true
		}
	}
	if stage != "ready" {
		server.coordinator.snapshot.CompletedStages = append(server.coordinator.snapshot.CompletedStages, stage)
	}
	server.coordinator.snapshot.Diagnostics = append(server.coordinator.snapshot.Diagnostics, diagnostics...)
	if len(diagnostics) > 0 {
		server.coordinator.snapshot.State = "partial"
	} else if stage == "ready" && server.coordinator.snapshot.State != "partial" {
		server.coordinator.snapshot.State = "ready"
	}
	return true
}

func (server *Server) updateIndexCounts(profileID string, modules, generated int, cancel chan struct{}) {
	if !server.indexStillRunning(cancel) {
		return
	}
	server.coordinator.mu.Lock()
	defer server.coordinator.mu.Unlock()
	if server.coordinator.snapshot.ProfileID == profileID {
		server.coordinator.snapshot.ModuleCount = modules
		server.coordinator.snapshot.GeneratedCount = generated
	}
}
