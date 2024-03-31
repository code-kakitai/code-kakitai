package test_utils

import (
	"bytes"
	"encoding/json"
	"testing"
)

// Jsonのフォーマットを整える
func FormatJSON(t *testing.T, b []byte) []byte {
	t.Helper()

	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return out.Bytes()
}
