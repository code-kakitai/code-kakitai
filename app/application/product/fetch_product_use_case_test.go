package product

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"
)

func TestFetchProductUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductQueryService := NewMockProductQueryService(ctrl)
	uc := NewFetchProductUseCase(mockProductQueryService)

	tests := []struct {
		name     string
		mockFunc func()
		want     []*FetchProductUseCaseDto
		wantErr  bool
	}{
		{
			name: "商品一覧を取得し、DTOを返却すること",
			mockFunc: func() {
				mockProductQueryService.EXPECT().FetchProductList(gomock.Any()).Return([]*FetchProductListDto{
					{
						ID:        "01HCNYK0PKYZWB0ZT1KR0EPWGP",
						Name:      "商品1",
						Price:     1000,
						Stock:     10,
						OwnerID:   "01HCNYK0PKYZWB0ZT1KR0EPWGQ",
						OwnerName: "オーナー1",
					},
				}, nil)
			},
			want: []*FetchProductUseCaseDto{
				{
					ID:        "01HCNYK0PKYZWB0ZT1KR0EPWGP",
					Name:      "商品1",
					Price:     1000,
					Stock:     10,
					OwnerID:   "01HCNYK0PKYZWB0ZT1KR0EPWGQ",
					OwnerName: "オーナー1",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()

			got, err := uc.Run(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}
