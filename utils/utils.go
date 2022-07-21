package utils

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

// ShortenString shortens the string if the string is longer than given length.
func ShortenString(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	return str[:maxLength]
}

// TryTimes will try an action up to `times` times with a delay of `delay` between each attempt
func TryTimes(handler func(attempt int) error, times int, delay time.Duration) error {
	var result error
	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for i := 0; i < times; i++ {
		if err := handler(i); err == nil {
			return nil
		} else {
			result = err
		}
		<-ticker.C
	}
	return result
}

func MapKeys(m map[string]bool) []string {
	var res []string
	for k := range m {
		res = append(res, k)
	}
	sort.Strings(res)
	return res
}

//MapToList returns a deterministic list from a map
func MapToList(m map[string]string) []string {
	result := make([]string, 0, len(m))
	for k, v := range m {
		result = append(result, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(result)
	return result
}

// ParseBoolEnvVar parses a bool env var. If nothing works, it's always false.
func ParseBoolEnvVar(name string) (value bool) {
	boolStr := os.Getenv(name)
	if len(boolStr) == 0 {
		return
	}
	value, _ = strconv.ParseBool(boolStr)
	return
}

// NormalizeJSON reorders JSON object keys in a deterministic way. The argument must either be
// raw JSON bytes or be serializable.
// See https://pkg.go.dev/encoding/json#Marshal for more details about why this orders fields correctly.
func NormalizeJSON(v interface{}) []byte {
	var b []byte
	switch val := v.(type) {
	case []byte:
		b = val
	case json.RawMessage:
		b = val
	case string:
		b = []byte(val)
	default:
		b, _ = json.Marshal(v)
	}

	var decoded interface{}
	_ = json.Unmarshal(b, &decoded)
	b, _ = json.Marshal(decoded)
	return b
}

// HashNormalizedJSON computes the hash after normalizing given input in JSON format.
func HashNormalizedJSON(v interface{}) string {
	return hex.EncodeToString(crypto.Keccak256(NormalizeJSON(v)))
}
