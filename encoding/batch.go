package encoding

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func gzipBytes(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write(b)
	if err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func gunzipBytes(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}

func EncodeGzippedProto(msg proto.Message) (string, error) {
	b, err := proto.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal msg: %v", err)
	}
	zipped, err := gzipBytes(b)
	if err != nil {
		return "", fmt.Errorf("failed to gzip msg: %v", err)
	}
	return base64.StdEncoding.EncodeToString(zipped), nil
}

func DecodeGzippedProto(encoded string, target proto.Message) error {
	zipped, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return fmt.Errorf("failed to base64decode: %v", err)
	}
	b, err := gunzipBytes(zipped)
	if err != nil {
		return err
	}
	return proto.Unmarshal(b, target)
}
