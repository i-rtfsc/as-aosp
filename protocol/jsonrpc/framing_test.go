package jsonrpc

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestFramerReadsConsecutiveMessages(t *testing.T) {
	input := "Content-Length: 7\r\nX-Test: ignored\r\n\r\n{\"a\":1}" +
		"Content-Length: 7\n\n{\"b\":2}"
	framer := NewFramer(&chunkReader{data: []byte(input), size: 3})
	first, err := framer.Read()
	if err != nil || string(first) != `{"a":1}` {
		t.Fatalf("first=%q err=%v", first, err)
	}
	second, err := framer.Read()
	if err != nil || string(second) != `{"b":2}` {
		t.Fatalf("second=%q err=%v", second, err)
	}
	if _, err := framer.Read(); err != io.EOF {
		t.Fatalf("expected EOF, got %v", err)
	}
}

type chunkReader struct {
	data []byte
	size int
}

func (reader *chunkReader) Read(target []byte) (int, error) {
	if len(reader.data) == 0 {
		return 0, io.EOF
	}
	count := reader.size
	if count > len(target) {
		count = len(target)
	}
	if count > len(reader.data) {
		count = len(reader.data)
	}
	copy(target, reader.data[:count])
	reader.data = reader.data[count:]
	return count, nil
}

func TestFramerRejectsInvalidHeaders(t *testing.T) {
	for _, input := range []string{
		"Content-Length: nope\r\n\r\n",
		"Content-Length: 0\r\nContent-Length: 0\r\n\r\n",
		"X-Test: ignored\r\n\r\n",
	} {
		if _, err := NewFramer(strings.NewReader(input)).Read(); err == nil {
			t.Fatalf("accepted invalid input: %q", input)
		}
	}
}

func TestWriteUsesUTF8ByteLength(t *testing.T) {
	var output bytes.Buffer
	if err := Write(&output, []byte(`{"message":"中文"}`)); err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(output.String(), "Content-Length: 20\r\n\r\n") {
		t.Fatalf("output=%q", output.String())
	}
}
