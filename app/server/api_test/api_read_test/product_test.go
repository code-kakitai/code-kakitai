//go:build integration_read

package api_read_test

import (
	"encoding/json"
	"github/code-kakitai/code-kakitai/server/api_test/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sebdah/goldie/v2"
)

func TestProduct_GetProducts(t *testing.T) {
	t.Parallel()
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
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
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

func TestProduct_GetProducts_With_Goldie(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		expectedCode int
	}{
		"正常系": {
			expectedCode: http.StatusOK,
		},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
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
			g.Assert(t, t.Name(), test_utils.FormatJSON(t, w.Body.Bytes()))
		})
	}
}
