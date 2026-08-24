package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"testing"

	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

func TestServeJSONRPCSmoke(t *testing.T) {
	var input bytes.Buffer
	writeServeRequest(t, &input, map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "aosp/initialize",
		"params":  map[string]string{"protocolVersion": "1.0"},
	})
	writeServeRequest(t, &input, map[string]any{
		"jsonrpc": "2.0",
		"id":      2,
		"method":  "aosp/shutdown",
		"params":  nil,
	})
	var output bytes.Buffer
	if err := runJSONRPC(&input, &output, log.New(io.Discard, "", 0)); err != nil {
		t.Fatal(err)
	}
	framer := jsonrpc.NewFramer(&output)
	initialize, err := framer.Read()
	if err != nil || !bytes.Contains(initialize, []byte(`"protocolVersion":"1.0"`)) {
		t.Fatalf("initialize=%s err=%v", initialize, err)
	}
	shutdown, err := framer.Read()
	if err != nil || !bytes.Contains(shutdown, []byte(`"result":null`)) {
		t.Fatalf("shutdown=%s err=%v", shutdown, err)
	}
}

func writeServeRequest(t *testing.T, writer io.Writer, value any) {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := jsonrpc.Write(writer, payload); err != nil {
		t.Fatal(err)
	}
}
