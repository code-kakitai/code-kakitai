package purchase

import (
	"context"

	"github.com/code-kakitai/go-pkg/errors"

	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

type purchaseDomainService struct {
	purchaseHistoryRepo PurchaseHistoryRepository
	productRepo         productDomain.ProductRepository
}

func NewPurchaseDomainService(
	purchaseHistoryRepo PurchaseHistoryRepository,
	productRepo productDomain.ProductRepository,
) PurchaseDomainService {
	return &purchaseDomainService{
		purchaseHistoryRepo: purchaseHistoryRepo,
		productRepo:         productRepo,
	}
}

func (ds *purchaseDomainService) PurchaseProducts(ctx context.Context, purchaseProducts []PurchaseProduct) error {
	totalAmount := int64(0)
	for _, purchaseProduct := range purchaseProducts {
		p, err := ds.productRepo.FindByID(ctx, purchaseProduct.ProductID())
		if err != nil {
			return err
		}
		totalAmount += p.Price() * int64(purchaseProduct.Count())
	}
	// 購入可能かチェック
	for _, purchaseProduct := range purchaseProducts {
		if err := ds.canPurchase(ctx, purchaseProduct); err != nil {
			return err
		}
	}

	// 購入処理
	for _, purchaseProduct := range purchaseProducts {
		p, err := ds.productRepo.FindByID(ctx, purchaseProduct.ProductID())
		if err != nil {
			return err
		}
		if err := p.UpdateStock(purchaseProduct.Count()); err != nil {
			return err
		}
		if err := ds.productRepo.Store(ctx, p); err != nil {
			return err
		}
	}
	// 購入履歴保存
	ph, err := NewPurchaseHistory(totalAmount, purchaseProducts)
	if err != nil {
		return err
	}
	if err := ds.purchaseHistoryRepo.Save(ctx, ph); err != nil {
		return err
	}
	return nil
}

func (ds *purchaseDomainService) canPurchase(ctx context.Context, purchaseProduct PurchaseProduct) error {
	// 購入可能かチェック
	p, err := ds.productRepo.FindByID(ctx, purchaseProduct.ProductID())
	if err != nil {
		return err
	}
	if p.Stock() < purchaseProduct.Count() {
		return errors.NewError("在庫が足りません。")
	}

	return nil
}
