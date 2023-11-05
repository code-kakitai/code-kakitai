package query_service

import (
	"context"
	"fmt"
	"github/code-kakitai/code-kakitai/application/product"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Fetch_Product_Query_Service(t *testing.T) {
	p := []*product.FetchProductListDto{
		{
			ID:        "01HCNYK4MQNC6G6X3F3DGXZ2J8",
			Name:      "サウナハット",
			Price:     3000,
			Stock:     20,
			OwnerID:   "01HCNYK3F7RJTWJ7GAQHPZDVE3",
			OwnerName: "古戸垣敦",
		},
	}
	tests := []struct {
		name string
		want []*product.FetchProductListDto
	}{
		{
			name: "オーナ情報を含めた商品一覧が取得できる",
			want: p,
		},
	}

	queryService := NewProductQueryService()
	resetTestData(t)
	for _, tt := range tests {

		t.Run(fmt.Sprintf(": %s", tt.name), func(t *testing.T) {
			got, _ := queryService.FetchProductList(context.Background())
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FindById() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
