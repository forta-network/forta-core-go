package utils

import (
	"bytes"
	"compress/gzip"
	"io"
)

// GzipEncode compresses bytes.
func GzipEncode(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	_, _ = w.Write(b)
	w.Close()
	return buf.Bytes(), nil
}

// GzipDecode decompresses bytes.
func GzipDecode(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}
