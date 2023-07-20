package purchase

import (
	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type PurchaseHistory struct {
	id          string
	totalAmount int64
	products    []PurchaseProduct
}

func NewPurchaseHistory(totalAmount int64, products []PurchaseProduct) (*PurchaseHistory, error) {
	// 購入金額のバリデーション
	// 割引等で合計金額が0円になることはあるため、0円以上を許容とする
	if totalAmount < 0 {
		return nil, errors.NewError("購入金額の値が不正です。")
	}

	// 購入商品のバリデーション
	if len(products) < 1 {
		return nil, errors.NewError("購入商品がありません。")
	}
	return &PurchaseHistory{
		id:          ulid.NewULID(),
		totalAmount: totalAmount,
		products:    products,
	}, nil
}

func (p *PurchaseHistory) TotalAmount() int64 {
	return p.totalAmount
}

func (p *PurchaseHistory) Products() []PurchaseProduct {
	return p.products
}

func (p *PurchaseHistory) ProductIDs() []string {
	var productIDs []string
	for _, product := range p.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}
