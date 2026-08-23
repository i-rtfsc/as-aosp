package jsonrpc

import (
	"encoding/json"
	"testing"
)

func TestParseRequestPreservesIDAndNotifications(t *testing.T) {
	request, err := ParseRequest([]byte(`{"jsonrpc":"2.0","id":7,"method":"test"}`))
	if err != nil || string(request.ID) != "7" || request.IsNotification() {
		t.Fatalf("request=%+v err=%v", request, err)
	}
	notification, err := ParseRequest([]byte(`{"jsonrpc":"2.0","method":"test"}`))
	if err != nil || !notification.IsNotification() {
		t.Fatalf("notification=%+v err=%v", notification, err)
	}
	encoded, err := json.Marshal(Success(request.ID, map[string]bool{"ok": true}))
	if err != nil || string(encoded) == "" {
		t.Fatalf("encoded=%s err=%v", encoded, err)
	}
	nullResult, err := json.Marshal(Success(request.ID, nil))
	if err != nil || string(nullResult) != `{"jsonrpc":"2.0","id":7,"result":null}` {
		t.Fatalf("null result=%s err=%v", nullResult, err)
	}
}
