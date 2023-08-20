package order

import (
	"context"
	"time"

	"github.com/code-kakitai/go-pkg/errors"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

type OrderUseCase struct {
	orderDomainService orderDomain.OrderDomainService
	cartRepo           cartDomain.CartRepository
}

func NewOrderUseCase(
	orderDomainService orderDomain.OrderDomainService,
	cartRepo cartDomain.CartRepository,
) *OrderUseCase {
	return &OrderUseCase{
		orderDomainService: orderDomainService,
		cartRepo:           cartRepo,
	}
}

type OrderUseCaseDto struct {
	ProductID string
	Count     int
}

func (uc *OrderUseCase) Run(ctx context.Context, userID string, dtos []OrderUseCaseDto, now time.Time) error {
	// カートから商品情報を取得
	cart, err := uc.getValidCart(ctx, userID, dtos)
	if err != nil {
		return err
	}
	// 購入処理
	if err := uc.orderDomainService.OrderProducts(ctx, cart, now); err != nil {
		return nil
	}
	// 管理者とユーザーにメール送信
	return nil
}

// カートの中身の整合性をチェック
func (uc *OrderUseCase) getValidCart(ctx context.Context, userID string, dtos []OrderUseCaseDto) (*cartDomain.Cart, error) {
	// カートから商品情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	cartProductMap := make(map[string]cartDomain.CartProduct)
	for _, cp := range cart.Products() {
		cartProductMap[cp.ProductID()] = cp
	}
	for _, dto := range dtos {
		cp, ok := cartProductMap[dto.ProductID]
		if !ok {
			return nil, errors.NewError("カートの商品が見つかりません。")
		}
		if cp.Count() != dto.Count {
			return nil, errors.NewError("カートの商品数が変更されています。")
		}
	}
	return cart, nil
}
