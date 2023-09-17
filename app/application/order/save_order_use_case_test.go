package order

import (
	"context"
	"testing"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"
	"go.uber.org/mock/gomock"

	transactionApp "github/code-kakitai/code-kakitai/application/transaction"
	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

func TestSaveOrderUseCase_Run(t *testing.T) {
	// usecase準備
	ctrl := gomock.NewController(t)
	mockOrderDomainService := orderDomain.NewMockOrderDomainService(ctrl)
	mockCartRepo := cartDomain.NewMockCartRepository(ctrl)
	mockTransactionManager := transactionApp.NewMockTransactionManager(ctrl)
	uc := NewSaveOrderUseCase(mockOrderDomainService, mockCartRepo, mockTransactionManager)

	// 各種テストデータ準備
	now := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local)
	userID := ulid.NewULID()
	dtos := []SaveOrderUseCaseInputDto{
		{
			ProductID: ulid.NewULID(),
			Quantity:  1,
		},
		{
			ProductID: ulid.NewULID(),
			Quantity:  3,
		},
	}
	cart, _ := cartDomain.NewCart(userID)
	for _, dto := range dtos {
		cart.AddProduct(dto.ProductID, dto.Quantity)
	}

	tests := []struct {
		name     string
		dtos     []SaveOrderUseCaseInputDto
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "work",
			dtos: dtos,
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), userID).Return(cart, nil),

					mockTransactionManager.EXPECT().RunInTransaction(gomock.Any(), gomock.Any()).
						Do(func(ctx context.Context, fn func(ctx context.Context) error) {
							if err := fn(ctx); err != nil {
								t.Errorf("Error executing fn: %v", err)
							}
						}).Return(nil),
					mockOrderDomainService.EXPECT().OrderProducts(gomock.Any(), cart, now).Return("", nil),
				)
			},
			wantErr: false,
		},
		{
			name: "cartの中身とdtosの中身が一致しない",
			dtos: []SaveOrderUseCaseInputDto{
				{
					ProductID: ulid.NewULID(),
					Quantity:  1,
				},
			},
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), userID).Return(cart, nil),
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			tt.mockFunc()
			_, err := uc.Run(context.Background(), userID, tt.dtos, now)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveOrderUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
