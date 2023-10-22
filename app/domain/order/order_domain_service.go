package order

import (
	"context"
	"time"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	errDomain "github/code-kakitai/code-kakitai/domain/error"
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

func (ds *orderDomainService) OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error) {
	// 購入対象の商品を取得
	ps, err := ds.productRepo.FindByIDs(ctx, cart.ProductIDs())
	if err != nil {
		return "", err
	}
	productMap := make(map[string]*productDomain.Product)
	for _, p := range ps {
		productMap[p.ID()] = p
	}

	// 購入処理
	ops := make([]OrderProduct, 0, len(cart.ProductIDs()))
	for _, cp := range cart.Products() {
		p, ok := productMap[cp.ProductID()]
		op, err := NewOrderProduct(cp.ProductID(), p.Price(), cp.Quantity())
		if err != nil {
			return "", err
		}
		ops = append(ops, *op)
		if !ok {
			// 購入した商品の商品詳細が見つからない場合はエラー
			// 商品を購入すると同時に、商品が削除された場合等に発生
			return "", errDomain.NewError("商品が見つかりません。")
		}
		if err := p.Consume(cp.Quantity()); err != nil {
			return "", err
		}
		if err := ds.productRepo.Save(ctx, p); err != nil {
			return "", err
		}
	}

	// 注文履歴保存
	o, err := NewOrder(cart.UserID(), OrderProducts(ops).TotalAmount(), ops, now)
	if err != nil {
		return "", err
	}
	if err := ds.orderRepo.Save(ctx, o); err != nil {
		return "", err
	}
	return o.ID(), nil
}
