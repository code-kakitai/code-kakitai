package purchase

import (
	"context"
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	purchaseDomain "github/code-kakitai/code-kakitai/domain/purchase"
)

func TestPurchaseUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPurchaseDomainService := purchaseDomain.NewMockPurchaseDomainService(ctrl)
	uc := NewPurchaseUseCase(mockPurchaseDomainService)
	now := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	userID := ulid.NewULID()
	type args struct {
		dtos []PurchaseUseCaseDto
		now  time.Time
	}
	tests := []struct {
		name    string
		dtos    []PurchaseUseCaseDto
		wantErr bool
	}{
		{
			name: "work",
			dtos: []PurchaseUseCaseDto{
				{
					ProductID: ulid.NewULID(),
					Count:     1,
				},
				{
					ProductID: ulid.NewULID(),
					Count:     3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gomock.InOrder(
				mockPurchaseDomainService.EXPECT().PurchaseProducts(
					context.Background(),
					userID,
					gomock.Any(),
					now,
				).Do(
					func(ctx context.Context, userID string, pps []purchaseDomain.PurchaseProduct, now time.Time) {
						var ps []purchaseDomain.PurchaseProduct
						for _, dto := range tt.dtos {
							p, _ := purchaseDomain.NewPurchaseProduct(dto.ProductID, dto.Count)
							ps = append(ps, *p)
						}
						diff := cmp.Diff(ps, pps,
							cmp.AllowUnexported(purchaseDomain.PurchaseProduct{}))
						if diff != "" {
							t.Errorf("PurchaseUseCase.Run() got diff: %s", diff)
						}
					},
				).Return(nil),
			)

			if err := uc.Run(context.Background(), userID, tt.dtos, now); (err != nil) != tt.wantErr {
				t.Errorf("PurchaseUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
