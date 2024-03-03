package api_test

import (
	"bytes"
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

func TestProduct_PostProducts(t *testing.T) {
	tests := map[string]struct {
		requestBody  map[string]any
		expectedCode int
		expectedBody []map[string]any
	}{
		"正常系": {
			requestBody: map[string]any{
				"owner_id":    "01HCNYK3F7RJTWJ7GAQHPZDVE3",
				"name":        "サウナハット2",
				"description": "今治タオルを素材としたこだわりのサウナハット",
				"price":       3000,
				"stock":       2,
			},
			expectedCode: http.StatusOK,
			expectedBody: []map[string]any{
				{
					"description": "",
					"id":          "01HCNYK4MQNC6G6X3F3DGXZ2J8",
					"name":        "サウナハット",
					"price":       float64(3000),
					"stock":       float64(20),
					"owner_id":    "01HCNYK3F7RJTWJ7GAQHPZDVE3",
				},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			resetTestData(t)
			b, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Errorf("failed to marshal err: %v", err)
				t.Fail()
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewBuffer(b))

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
