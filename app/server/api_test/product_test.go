package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProduct_GetProducts(t *testing.T) {
	// GET処理なので、冒頭でのみテストデータを初期化する
	// 書き込み処理の場合は、テストケースごとに初期化する
	resetTestData(t)

	tests := map[string]struct {
		expectedCode int
		expectedBody []map[string]any
	}{
		"正常系": {
			expectedCode: http.StatusOK,
			expectedBody: []map[string]any{
				{
					"description": "",
					"id":          "01HCNYK4MQNC6G6X3F3DGXZ2J8",
					"name":        "サウナハット",
					"price":       float64(3000),
					"stock":       float64(20),
					"owner_id":    "01HCNYK3F7RJTWJ7GAQHPZDVE3",
					"owner_name":  "古戸垣敦",
				},
			},
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

			// レスポンスボディのパース
			var responseBody []map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
				t.Fatalf("failed to unmarshal response body: %v", err)
			}

			// レスポンスボディの期待値と比較
			if diff := cmp.Diff(tt.expectedBody, responseBody); diff != "" {
				t.Errorf("response body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
