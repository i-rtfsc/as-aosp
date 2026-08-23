package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/i-rtfsc/as-aosp/protocol/aosp"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

func main() {
	if err := run(os.Stdin, os.Stdout, log.New(os.Stderr, "", 0)); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

func run(reader io.Reader, writer io.Writer, logger *log.Logger) error {
	framer := jsonrpc.NewFramer(reader)
	server := &aosp.Server{}
	for {
		payload, err := framer.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		request, parseErr := jsonrpc.ParseRequest(payload)
		if parseErr != nil {
			rpcError, ok := parseErr.(*jsonrpc.Error)
			if !ok {
				return parseErr
			}
			if err := writeResponse(writer, jsonrpc.Failure(nil, rpcError)); err != nil {
				return err
			}
			return nil
		}
		response, reply := server.Handle(request)
		if reply {
			if err := writeResponse(writer, response); err != nil {
				return err
			}
		}
		if server.IsShutdown() {
			logger.Print("shutdown complete")
			return nil
		}
	}
}

func writeResponse(writer io.Writer, response jsonrpc.Response) error {
	payload, err := json.Marshal(response)
	if err != nil {
		return err
	}
	return jsonrpc.Write(writer, payload)
}
