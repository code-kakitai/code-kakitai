package test_utils

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sebdah/goldie/v2"
)

func WithIgnoreMapKeys(t *testing.T, ignoreKeys ...string) goldie.Option {
	t.Helper()

	equalFn := func(actual, expected []byte) bool {
		var actualInterface, expectedInterface interface{}

		if err := json.Unmarshal(actual, &actualInterface); err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(expected, &expectedInterface); err != nil {
			t.Fatal(err)
		}

		// 対象のフィールドを削除
		removeMapKeys(actualInterface, ignoreKeys)
		removeMapKeys(expectedInterface, ignoreKeys)
		return cmp.Equal(actualInterface, expectedInterface)
	}
	return goldie.WithEqualFn(equalFn)
}

func removeMapKeys(data interface{}, keys []string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for _, field := range keys {
			delete(v, field)
		}
		for _, _v := range v {
			switch __v := _v.(type) {
			case map[string]any, []any:
				removeMapKeys(__v, keys)
			}
		}
	case []interface{}:
		for i := range v {
			removeMapKeys(v[i], keys)
		}
	}
}
