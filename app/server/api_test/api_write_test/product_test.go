//go:build integration_write

package api_write_test

import (
	"bytes"
	"encoding/json"
	"github/code-kakitai/code-kakitai/server/api_test/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/oklog/ulid/v2"
	"github.com/sebdah/goldie/v2"
)

func TestProduct_PostProducts(t *testing.T) {
	tests := map[string]struct {
		requestBody  map[string]any
		expectedCode int
		expectedBody map[string]any
	}{
		"正常系": {
			requestBody: map[string]any{
				"owner_id":    "01HCNYK3F7RJTWJ7GAQHPZDVE3",
				"name":        "サウナハット青",
				"description": "今治タオルを素材としたこだわりのサウナハット",
				"price":       3000,
				"stock":       2,
			},
			expectedCode: http.StatusCreated,
			expectedBody: map[string]any{
				"product": map[string]any{
					"description": "今治タオルを素材としたこだわりのサウナハット",
					"id":          "01HCNYK3F7RJTWJ7GAQHPZDVE3",
					"name":        "サウナハット青",
					"price":       float64(3000),
					"stock":       float64(2),
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
				t.Fatalf("failed to marshal err: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewBuffer(b))

			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			var actualBody map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &actualBody); err != nil {
				t.Fatalf("failed to unmarshal response body: %v", err)
			}

			product, ok := actualBody["product"].(map[string]any)
			if !ok {
				t.Fatalf("failed to parse response body: %v", actualBody)
			}
			// ULIDの形式をチェック
			if _, err := ulid.ParseStrict(product["id"].(string)); err != nil {
				t.Errorf("id is not a valid ULID: %v", product["id"])
			}

			opts := []cmp.Option{
				cmpopts.IgnoreMapEntries(func(key string, value any) bool {
					return key == "id"
				}),
			}
			if diff := cmp.Diff(tt.expectedBody, actualBody, opts...); diff != "" {
				t.Errorf("response body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProduct_PostProducts_With_Goldie(t *testing.T) {
	tests := map[string]struct {
		requestBody  map[string]any
		expectedCode int
		expectedBody map[string]any
	}{
		"正常系": {
			requestBody: map[string]any{
				"owner_id":    "01HCNYK3F7RJTWJ7GAQHPZDVE3",
				"name":        "サウナハット青",
				"description": "今治タオルを素材としたこだわりのサウナハット",
				"price":       3000,
				"stock":       2,
			},
			expectedCode: http.StatusCreated,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			resetTestData(t)
			b, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Fatalf("failed to marshal err: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/products", bytes.NewBuffer(b))

			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			var actualBody map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &actualBody); err != nil {
				t.Fatalf("failed to unmarshal response body: %v", err)
			}
			// レスポンスボディの期待値と比較
			// レスポンスボディが変わった時は、-updateフラグをつけてテストを実行する
			g := goldie.New(t,
				goldie.WithNameSuffix(".golden.json"),
				goldie.WithFixtureDir("testdata/product_test"),
				test_utils.WithIgnoreMapKeys(t, "id"),
			)
			g.Assert(t, t.Name(), test_utils.FormatJSON(t, w.Body.Bytes()))
		})
	}
}
