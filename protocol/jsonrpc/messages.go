package jsonrpc

import (
	"encoding/json"
	"fmt"
)

const Version = "2.0"

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type Response struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (err *Error) Error() string {
	return err.Message
}

func ParseRequest(payload []byte) (Request, error) {
	var request Request
	if err := json.Unmarshal(payload, &request); err != nil {
		return Request{}, &Error{Code: -32700, Message: "parse error"}
	}
	if request.JSONRPC != Version || request.Method == "" {
		return Request{}, &Error{Code: -32600, Message: "invalid request"}
	}
	return request, nil
}

func (request Request) IsNotification() bool {
	return len(request.ID) == 0
}

func Success(id json.RawMessage, result any) Response {
	payload, err := json.Marshal(result)
	if err != nil {
		return Failure(id, &Error{Code: -32603, Message: "encode result"})
	}
	return Response{JSONRPC: Version, ID: id, Result: payload}
}

func Failure(id json.RawMessage, rpcError *Error) Response {
	return Response{JSONRPC: Version, ID: id, Error: rpcError}
}

func InvalidParams(message string) *Error {
	return &Error{Code: -32602, Message: message}
}

func MethodNotFound(method string) *Error {
	return &Error{Code: -32601, Message: fmt.Sprintf("method not found: %s", method)}
}

func NotInitialized() *Error {
	return &Error{Code: -32600, Message: "server is not initialized"}
}

func AlreadyInitialized() *Error {
	return &Error{Code: -32600, Message: "server is already initialized"}
}
