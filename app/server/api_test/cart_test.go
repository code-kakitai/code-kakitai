package api_test

import (
	"bytes"
	"encoding/json"
	"github/code-kakitai/code-kakitai/presentation/cart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrder_PostCart(t *testing.T) {
	tests := map[string]struct {
		requestBody  cart.PostCartsParams
		expectedCode int
		expectedBody map[string]any
	}{
		"正常系": {
			requestBody: cart.PostCartsParams{
				ProductID: "01HCNYK4MQNC6G6X3F3DGXZ2J8",
				Quantity:  1,
			},
			expectedCode: http.StatusOK,
			expectedBody: map[string]any{},
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

			var actualBody map[string]interface{}
			if err := json.Unmarshal(w.Body.Bytes(), &actualBody); err != nil {
				t.Errorf("failed to unmarshal response body: %v", err)
			}
			if diff := cmp.Diff(tt.expectedBody, actualBody); diff != "" {
				t.Errorf("response body mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
