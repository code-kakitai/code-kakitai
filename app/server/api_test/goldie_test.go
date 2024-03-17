package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sebdah/goldie/v2"
)

func TestProduct_GetProducts_With_Goldie(t *testing.T) {
	// GET処理なので、冒頭でのみテストデータを初期化する
	// 書き込み処理の場合は、テストケースごとに初期化する
	resetTestData(t)

	tests := map[string]struct {
		expectedCode int
		expectedBody []map[string]any
	}{
		"正常系": {
			expectedCode: http.StatusOK,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)
			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			// ステータスコードの期待値と比較
			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			// レスポンスボディの期待値と比較
			// レスポンスボディが変わった時は、-updateフラグをつけてテストを実行する
			g := goldie.New(t,
				goldie.WithNameSuffix(".golden.json"),
				goldie.WithFixtureDir("testdata/product_test"),
			)
			g.Assert(t, t.Name(), formatJSON(t, w.Body.Bytes()))
		})
	}
}

// Jsonのフォーマットを整える
func formatJSON(t *testing.T, b []byte) []byte {
	t.Helper()

	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	return out.Bytes()
}
