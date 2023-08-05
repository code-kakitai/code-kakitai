package purchase

import (
	"context"
	"time"

	purchaseDomain "github/code-kakitai/code-kakitai/domain/purchase"
)

type PurchaseUseCase struct {
	purchaseDomainService purchaseDomain.PurchaseDomainService
}

func NewPurchaseUseCase(
	purchaseDomainService purchaseDomain.PurchaseDomainService,
) *PurchaseUseCase {
	return &PurchaseUseCase{
		purchaseDomainService: purchaseDomainService,
	}
}

type PurchaseUseCaseDto struct {
	ProductID string
	Count     int
}

func (uc *PurchaseUseCase) Run(ctx context.Context, userID string, dtos []PurchaseUseCaseDto, now time.Time) error {
	// dtoからPurchaseProductへ変換
	var pps []purchaseDomain.PurchaseProduct
	for _, dto := range dtos {
		p, err := purchaseDomain.NewPurchaseProduct(dto.ProductID, dto.Count)
		if err != nil {
			return err
		}
		pps = append(pps, *p)
	}
	// 購入処理
	if err := uc.purchaseDomainService.PurchaseProducts(ctx, userID, pps, now); err != nil {
		return nil
	}
	// 管理者とユーザーにメール送信
	return nil
}
