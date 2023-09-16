package order

import (
	"context"
	"time"

	transactionApp "github/code-kakitai/code-kakitai/application/transaction"
	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	errDomain "github/code-kakitai/code-kakitai/domain/error"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

type SaveOrderUseCase struct {
	orderDomainService orderDomain.OrderDomainService
	cartRepo           cartDomain.CartRepository
	transactionManager transactionApp.TransactionManager
}

func NewSaveOrderUseCase(
	orderDomainService orderDomain.OrderDomainService,
	cartRepo cartDomain.CartRepository,
	transactionManager transactionApp.TransactionManager,
) *SaveOrderUseCase {
	return &SaveOrderUseCase{
		orderDomainService: orderDomainService,
		cartRepo:           cartRepo,
		transactionManager: transactionManager,
	}
}

type SaveOrderUseCaseInputDto struct {
	ProductID string
	Quantity  int
}

func (uc *SaveOrderUseCase) Run(ctx context.Context, userID string, dtos []SaveOrderUseCaseInputDto, now time.Time) (string, error) {
	// カートから商品情報を取得
	cart, err := uc.getValidCart(ctx, userID, dtos)
	if err != nil {
		return "", err
	}
	// 購入処理
	var orderID string
	if err := uc.transactionManager.RunInTransaction(ctx, func(ctx context.Context) error {
		orderID, err = uc.orderDomainService.OrderProducts(ctx, cart, now)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	// 管理者とユーザーにメール送信
	return orderID, nil
}

// カートの中身の整合性をチェック
func (uc *SaveOrderUseCase) getValidCart(ctx context.Context, userID string, dtos []SaveOrderUseCaseInputDto) (*cartDomain.Cart, error) {
	// カートから商品情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	for _, dto := range dtos {
		pq, err := cart.QuantityByProductID(dto.ProductID)
		if err != nil {
			return nil, err
		}
		// DTOで渡ってきた数量とカートの数量が一致しない場合はエラー
		if pq != dto.Quantity {
			return nil, errDomain.NewError("カートの商品数が変更されています。")
		}
	}
	return cart, nil
}
