package aosp

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

func makeRequest(id int, method string, params any) jsonrpc.Request {
	data, _ := json.Marshal(params)
	idData, _ := json.Marshal(id)
	return jsonrpc.Request{JSONRPC: jsonrpc.Version, ID: idData, Method: method, Params: data}
}

func TestServerLifecycleAndCapabilities(t *testing.T) {
	server := &Server{}
	response, reply := server.Handle(makeRequest(1, "aosp/profile/validate", map[string]any{}))
	if !reply || response.Error == nil || response.Error.Message != "server is not initialized" {
		t.Fatalf("response=%+v reply=%v", response, reply)
	}
	response, reply = server.Handle(makeRequest(2, "aosp/initialize", InitializeParams{ProtocolVersion: "1.0"}))
	if !reply || response.Error != nil {
		t.Fatalf("initialize=%+v reply=%v", response, reply)
	}
	response, _ = server.Handle(makeRequest(3, "aosp/initialize", InitializeParams{ProtocolVersion: "1.0"}))
	if response.Error == nil {
		t.Fatalf("duplicate initialize=%+v", response)
	}
	response, reply = server.Handle(makeRequest(4, "aosp/shutdown", nil))
	if !reply || response.Error != nil || !server.shutdown {
		t.Fatalf("shutdown=%+v reply=%v", response, reply)
	}
}

func TestServerRejectsIncompatibleVersion(t *testing.T) {
	server := &Server{}
	response, _ := server.Handle(makeRequest(1, "aosp/initialize", InitializeParams{ProtocolVersion: "2.0"}))
	if response.Error == nil || server.initialized {
		t.Fatalf("response=%+v initialized=%v", response, server.initialized)
	}
}

func TestServerValidatesProfilesAndTopologies(t *testing.T) {
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/profile/validate", map[string]any{
		"profile": model.Profile{},
	}))
	var profile ProfileResult
	decodeResult(t, response, &profile)
	if len(profile.Diagnostics) == 0 {
		t.Fatalf("profile=%+v", profile)
	}
	response, _ = server.Handle(makeRequest(3, "aosp/topology/validate", map[string]any{
		"topology": topology.WorkspaceTopology{},
	}))
	var view TopologyResult
	decodeResult(t, response, &view)
	if len(view.Diagnostics) != 0 {
		t.Fatalf("topology=%+v", view)
	}
}

func TestServerDiscoversAnEmptyWorkspace(t *testing.T) {
	root := t.TempDir()
	server := initializedServer(t)
	response, _ := server.Handle(makeRequest(2, "aosp/topology/discover", TopologyDiscoverParams{
		WorkspaceViewRoot: root,
		MaxDepth:          2,
	}))
	var result TopologyResult
	decodeResult(t, response, &result)
	expected, err := os.Stat(root)
	if err != nil || !expected.IsDir() || string(result.Topology.WorkspaceViewRoot) == "" {
		t.Fatalf("result=%+v err=%v", result, err)
	}
}

func initializedServer(t *testing.T) *Server {
	t.Helper()
	server := &Server{}
	response, _ := server.Handle(makeRequest(1, "aosp/initialize", InitializeParams{
		ProtocolVersion: ProtocolVersion,
	}))
	if response.Error != nil {
		t.Fatalf("initialize=%+v", response)
	}
	return server
}

func decodeResult(t *testing.T, response jsonrpc.Response, target any) {
	t.Helper()
	if response.Error != nil {
		t.Fatalf("response error=%+v", response.Error)
	}
	if err := json.Unmarshal(response.Result, target); err != nil {
		t.Fatalf("decode result: %v", err)
	}
}
