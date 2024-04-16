//go:build integration_read

package api_read_test

import (
	"encoding/json"
	"fmt"
	"github/code-kakitai/code-kakitai/server/api_test/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sebdah/goldie/v2"
)

func TestUser_GetUserByID(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		id           string
		expectedCode int
		expectedBody map[string]any
	}{
		"正常系": {
			id:           "01HCNYK0PKYZWB0ZT1KR0EPWGP",
			expectedCode: http.StatusOK,
			expectedBody: map[string]any{
				"users": map[string]any{
					"id":           "01HCNYK0PKYZWB0ZT1KR0EPWGP",
					"email":        "example@test.com",
					"phone_number": "08011112222",
					"last_name":    "山田",
					"first_name":   "太郎",
					"address":      "東京都渋谷区広尾2-2",
				},
			},
		},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/users/%s", tt.id), nil)
			w := httptest.NewRecorder()
			api.ServeHTTP(w, req)

			// ステータスコードの期待値と比較
			if w.Code != tt.expectedCode {
				t.Errorf("expected status code %d, got %d", tt.expectedCode, w.Code)
			}

			// レスポンスボディのパース
			var responseBody map[string]interface{}
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

func TestUser_GetUserByID_With_Goldie(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		id           string
		expectedCode int
		expectedBody map[string]any
	}{
		"正常系": {
			id:           "01HCNYK0PKYZWB0ZT1KR0EPWGP",
			expectedCode: http.StatusOK,
		},
	}

	for testName, tt := range tests {
		tt := tt
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/users/%s", tt.id), nil)
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
				goldie.WithFixtureDir("testdata/user_test"),
			)
			g.Assert(t, t.Name(), test_utils.FormatJSON(t, w.Body.Bytes()))
		})
	}
}
