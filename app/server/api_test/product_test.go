package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProduct_GetProducts(t *testing.T) {
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

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			resetTestData(t)
			req := httptest.NewRequest(http.MethodGet, "/v1/products", nil)

			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			f := func(w *httptest.ResponseRecorder) bool {
				if w.Code != tt.expectedCode {
					t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
					// ステータスが異なる場合は以降の比較を行わない
					return false
				}

				var actualBody []map[string]interface{}
				if err := json.Unmarshal(w.Body.Bytes(), &actualBody); err != nil {
					t.Errorf("failed to unmarshal response body: %v", err)
					return false
				}
				if diff := cmp.Diff(tt.expectedBody, actualBody); diff != "" {
					t.Errorf("response body mismatch (-want +got):\n%s", diff)
				}

				return true
			}

			if !f(w) {
				t.Fail()
			}
		})
	}
}
