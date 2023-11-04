package cart

import (
	"context"
	"log"
	"testing"

	"github.com/code-kakitai/go-pkg/ulid"
	"go.uber.org/mock/gomock"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

func TestAddCartUseCase_Run(t *testing.T) {
	// usecase準備
	ctrl := gomock.NewController(t)
	mockCartRepo := cartDomain.NewMockCartRepository(ctrl)
	mockProductRepo := productDomain.NewMockProductRepository(ctrl)
	uc := NewCartUseCase(mockCartRepo, mockProductRepo)

	// テストに必要な変数準備
	userID := ulid.NewULID()
	ownerID := ulid.NewULID()
	cart, err := cartDomain.NewCart(userID)
	if err != nil {
		log.Println(err)
	}
	product, err := productDomain.NewProduct(ownerID, "test", "description", 100, 10)
	if err != nil {
		log.Println(err)
	}

	tests := []struct {
		name     string
		dto      AddCartUseCaseInputDto
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "work",
			dto: AddCartUseCaseInputDto{
				UserID:    userID,
				ProductID: product.ID(),
				Quantity:  1,
			},
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), gomock.Any()).Return(cart, nil),
					mockProductRepo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(product, nil),
					mockCartRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, cart *cartDomain.Cart) error {
						cart.AddProduct(product.ID(), 1)
						return nil
					}).Return(nil),
				)
			},
		},
		{
			name: "すでにcartに商品が入っており、dtoの商品数の値が0の時はカートから商品を削除",
			dto: AddCartUseCaseInputDto{
				UserID:    userID,
				ProductID: product.ID(),
				Quantity:  0,
			},
			mockFunc: func() {
				gomock.InOrder(
					mockCartRepo.EXPECT().FindByUserID(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, userID string) (*cartDomain.Cart, error) {
						cart.AddProduct(product.ID(), 1)
						return cart, nil
					}),
					mockProductRepo.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(product, nil),
					mockCartRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Do(func(ctx context.Context, cart *cartDomain.Cart) error {
						cart.RemoveProduct(product.ID())
						return nil
					}).Return(nil),
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			if err := uc.Run(context.Background(), tt.dto); (err != nil) != tt.wantErr {
				t.Errorf("AddCartUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
