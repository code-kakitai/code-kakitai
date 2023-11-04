package product

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/mock/gomock"

	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

func TestSaveProductUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockProductRepo := productDomain.NewMockProductRepository(ctrl)
	uc := NewSaveProductUseCase(mockProductRepo)

	tests := []struct {
		name     string
		input    SaveProductUseCaseInputDto
		mockFunc func()
		want     *SaveProductUseCaseOutputDto
		wantErr  bool
	}{
		{
			name: "商品を保存し、DTOを返却すること",
			input: SaveProductUseCaseInputDto{
				OwnerID:     "01HCNYK0PKYZWB0ZT1KR0EPWGQ",
				Name:        "商品1",
				Description: "商品1の説明",
				Price:       1000,
				Stock:       10,
			},
			mockFunc: func() {
				mockProductRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			want: &SaveProductUseCaseOutputDto{
				ID:          "01HCNYK0PKYZWB0ZT1KR0EPWGP",
				OwnerID:     "01HCNYK0PKYZWB0ZT1KR0EPWGQ",
				Name:        "商品1",
				Description: "商品1の説明",
				Price:       1000,
				Stock:       10,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()

			got, err := uc.Run(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got, cmpopts.IgnoreFields(*got, "ID")); diff != "" {
				t.Errorf("Run() diff = %v", diff)
			}
		})
	}
}
