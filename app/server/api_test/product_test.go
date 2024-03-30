package api_test

import (
	"bytes"
	"encoding/json"
	"github/code-kakitai/code-kakitai/server/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/oklog/ulid/v2"
	"github.com/sebdah/goldie/v2"
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
					// IDはランダムな文字列なため、Diffでは比較しない
					// "id":          "01HCNYK3F7RJTWJ7GAQHPZDVE3",
					"name":     "サウナハット青",
					"price":    float64(3000),
					"stock":    float64(2),
					"owner_id": "01HCNYK3F7RJTWJ7GAQHPZDVE3",
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
			// ULIDはランダムな文字列なため、形式のみチェック
			if _, err := ulid.ParseStrict(product["id"].(string)); err != nil {
				t.Errorf("id is not a valid ULID: %v", product["id"])
			}
			// IDはランダムな文字列なため、Diffでは比較しない
			delete(actualBody["product"].(map[string]any), "id")
			if diff := cmp.Diff(tt.expectedBody, actualBody); diff != "" {
				t.Errorf("response body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestProduct_GetProducts_With_Goldie(t *testing.T) {
	// GET処理なので、冒頭でのみテストデータを初期化する
	// 書き込み処理の場合は、テストケースごとに初期化する
	resetTestData(t)

	tests := map[string]struct {
		expectedCode int
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
				t.Fatalf("expected status code %d, got %d", tt.expectedCode, w.Code)
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
				utils.WithIgnoreMapKeys(t, "id"),
			)
			g.Assert(t, t.Name(), formatJSON(t, w.Body.Bytes()))
		})
	}
}
