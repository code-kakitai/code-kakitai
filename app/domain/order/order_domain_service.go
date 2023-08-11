package order

import (
	"context"
	"time"

	"github.com/code-kakitai/go-pkg/errors"

	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

type orderDomainService struct {
	orderRepo   OrderRepository
	productRepo productDomain.ProductRepository
}

func NewOrderDomainService(
	orderRepo OrderRepository,
	productRepo productDomain.ProductRepository,
) OrderDomainService {
	return &orderDomainService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (ds *orderDomainService) OrderProducts(ctx context.Context, userID string, OrderProducts []OrderProduct, now time.Time) error {
	// 購入商品のIDを取得
	productIDs := make([]string, 0, len(OrderProducts))
	for _, OrderProduct := range OrderProducts {
		productIDs = append(productIDs, OrderProduct.ProductID())
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
	for _, pp := range OrderProducts {
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
	ph, err := NewOrder(userID, totalAmount, OrderProducts, now)
	if err != nil {
		return err
	}
	if err := ds.orderRepo.Save(ctx, ph); err != nil {
		return err
	}
	return nil
}
