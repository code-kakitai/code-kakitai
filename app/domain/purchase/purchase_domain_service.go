package purchase

import (
	"context"
	"time"

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

func (ds *purchaseDomainService) PurchaseProducts(ctx context.Context, userID string, purchaseProducts []PurchaseProduct, now time.Time) error {
	// 購入商品のIDを取得
	productIDs := make([]string, 0, len(purchaseProducts))
	for _, purchaseProduct := range purchaseProducts {
		productIDs = append(productIDs, purchaseProduct.ProductID())
	}

	// todo ここからトランザクション & 行ロック
	// 購入対象の商品を取得
	ps, err := ds.productRepo.FindByIDs(ctx, productIDs)
	if err != nil {
		return err
	}
	productMap := make(map[string]*productDomain.Product)
	for _, p := range ps {
		productMap[p.ID()] = p
	}

	// 購入処理
	var totalAmount int64
	for _, pp := range purchaseProducts {
		p, ok := productMap[pp.ProductID()]
		if !ok {
			// 購入した商品の商品詳細が見つからない場合はエラー（商品を購入すると同時に、商品が削除された場合等に発生）
			return errors.NewError("商品が見つかりません。")
		}
		// 購入金額計算
		totalAmount += p.Price() * int64(pp.Count())
		if err := p.Consume(pp.Count()); err != nil {
			return err
		}
		if err := ds.productRepo.Save(ctx, p); err != nil {
			return err
		}
	}

	// 購入履歴保存
	ph, err := NewPurchaseHistory(userID, totalAmount, purchaseProducts, now)
	if err != nil {
		return err
	}
	if err := ds.purchaseHistoryRepo.Save(ctx, ph); err != nil {
		return err
	}
	return nil
}
