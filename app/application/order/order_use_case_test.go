package order

import (
	"context"
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/mock/gomock"

	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

func TestOrderUseCase_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockOrderDomainService := orderDomain.NewMockOrderDomainService(ctrl)
	uc := NewOrderUseCase(mockOrderDomainService)
	now := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	userID := ulid.NewULID()
	type args struct {
		dtos []OrderUseCaseDto
		now  time.Time
	}
	tests := []struct {
		name    string
		dtos    []OrderUseCaseDto
		wantErr bool
	}{
		{
			name: "work",
			dtos: []OrderUseCaseDto{
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
				mockOrderDomainService.EXPECT().OrderProducts(
					context.Background(),
					userID,
					gomock.Any(),
					now,
				).Do(
					func(ctx context.Context, userID string, pps []orderDomain.OrderProduct, now time.Time) {
						var ps []orderDomain.OrderProduct
						for _, dto := range tt.dtos {
							p, _ := orderDomain.NewOrderProduct(dto.ProductID, dto.Count)
							ps = append(ps, *p)
						}
						diff := cmp.Diff(ps, pps,
							cmp.AllowUnexported(orderDomain.OrderProduct{}))
						if diff != "" {
							t.Errorf("OrderUseCase.Run() got diff: %s", diff)
						}
					},
				).Return(nil),
			)

			if err := uc.Run(context.Background(), userID, tt.dtos, now); (err != nil) != tt.wantErr {
				t.Errorf("OrderUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
