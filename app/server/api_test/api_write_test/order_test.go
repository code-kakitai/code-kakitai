//go:build integration_write

package api_write_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestOrder_PostOrders(t *testing.T) {
	tests := map[string]struct {
		postCartsParams map[string]any
		requestBody     []map[string]any
		expectedCode    int
		expectedBody    map[string]any
	}{
		"正常系": {
			postCartsParams: map[string]any{
				"product_id": "01HCNYK4MQNC6G6X3F3DGXZ2J8",
				"quantity":   1,
			},
			requestBody: []map[string]any{
				{
					"product_id": "01HCNYK4MQNC6G6X3F3DGXZ2J8",
					"quantity":   1,
				},
			},
			expectedCode: http.StatusCreated,
			expectedBody: map[string]any{
				"id": "01HSQWHYEVKRH80JJPBRQ5DCGW",
			},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			resetTestData(t)

			// TODO: Redisの初期データを追加できたら、この処理は不要
			// Redisに値を入れるために、カートに商品を追加
			b, err := json.Marshal(tt.postCartsParams)
			if err != nil {
				t.Fatalf("failed to marshal err: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/carts", bytes.NewBuffer(b))
			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			// 注文を作成
			b, err = json.Marshal(tt.requestBody)
			if err != nil {
				t.Fatalf("failed to marshal err: %v", err)
			}
			req = httptest.NewRequest(http.MethodPost, "/v1/orders", bytes.NewBuffer(b))
			w = httptest.NewRecorder()
			api.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			var actualBody map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &actualBody); err != nil {
				t.Fatalf("failed to unmarshal response body: %v", err)
			}
			// ULIDはランダムな文字列なため、形式のみチェック
			if _, err := ulid.ParseStrict(actualBody["id"].(string)); err != nil {
				t.Errorf("id is not a valid ULID: %v", actualBody["id"])
			}
		})
	}
}
