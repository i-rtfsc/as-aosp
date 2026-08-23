package jsonrpc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const MaxPayloadSize = 16 << 20

type Framer struct {
	reader *bufio.Reader
}

func NewFramer(reader io.Reader) *Framer {
	return &Framer{reader: bufio.NewReader(reader)}
}

func (f *Framer) Read() ([]byte, error) {
	contentLength := -1
	readHeader := false
	for {
		line, err := f.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF && line == "" && !readHeader {
				return nil, io.EOF
			}
			return nil, fmt.Errorf("read JSON-RPC header: %w", err)
		}
		readHeader = true
		line = strings.TrimSuffix(strings.TrimSuffix(line, "\n"), "\r")
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("malformed JSON-RPC header %q", line)
		}
		name := strings.ToLower(strings.TrimSpace(parts[0]))
		value := strings.TrimSpace(parts[1])
		if name != "content-length" {
			continue
		}
		if contentLength >= 0 {
			return nil, fmt.Errorf("duplicate Content-Length header")
		}
		length, parseErr := strconv.Atoi(value)
		if parseErr != nil || length < 0 {
			return nil, fmt.Errorf("invalid Content-Length %q", value)
		}
		if length > MaxPayloadSize {
			return nil, fmt.Errorf("JSON-RPC payload exceeds %d bytes", MaxPayloadSize)
		}
		contentLength = length
	}
	if contentLength < 0 {
		return nil, fmt.Errorf("missing Content-Length header")
	}
	payload := make([]byte, contentLength)
	if _, err := io.ReadFull(f.reader, payload); err != nil {
		return nil, fmt.Errorf("read JSON-RPC payload: %w", err)
	}
	return payload, nil
}

func Write(writer io.Writer, payload []byte) error {
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(payload))
	if _, err := io.WriteString(writer, header); err != nil {
		return err
	}
	if _, err := io.Copy(writer, bytes.NewReader(payload)); err != nil {
		return err
	}
	return nil
}
