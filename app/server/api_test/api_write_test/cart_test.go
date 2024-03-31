//go:build integration_write

package api_write_test

import (
	"bytes"
	"encoding/json"
	"github/code-kakitai/code-kakitai/presentation/cart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrder_PostCart(t *testing.T) {
	tests := map[string]struct {
		requestBody  cart.PostCartsParams
		expectedCode int
	}{
		"正常系": {
			requestBody: cart.PostCartsParams{
				ProductID: "01HCNYK4MQNC6G6X3F3DGXZ2J8",
				Quantity:  1,
			},
			expectedCode: http.StatusNoContent,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			resetTestData(t)

			b, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Errorf("failed to marshal err: %v", err)
				t.Fail()
			}
			req := httptest.NewRequest(http.MethodPost, "/v1/carts", bytes.NewBuffer(b))
			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}
		})
	}
}
