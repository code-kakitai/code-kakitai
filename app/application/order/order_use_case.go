package order

import (
	"context"
	"time"

	orderDomain "github/code-kakitai/code-kakitai/domain/order"
)

type OrderUseCase struct {
	orderDomainService orderDomain.OrderDomainService
}

func NewOrderUseCase(
	orderDomainService orderDomain.OrderDomainService,
) *OrderUseCase {
	return &OrderUseCase{
		orderDomainService: orderDomainService,
	}
}

type OrderUseCaseDto struct {
	ProductID string
	Count     int
}

func (uc *OrderUseCase) Run(ctx context.Context, userID string, dtos []OrderUseCaseDto, now time.Time) error {
	// dtoからOrderProductへ変換
	var pps []orderDomain.OrderProduct
	for _, dto := range dtos {
		p, err := orderDomain.NewOrderProduct(dto.ProductID, dto.Count)
		if err != nil {
			return err
		}
		pps = append(pps, *p)
	}
	// 購入処理
	if err := uc.orderDomainService.OrderProducts(ctx, userID, pps, now); err != nil {
		return nil
	}
	// 管理者とユーザーにメール送信
	return nil
}
