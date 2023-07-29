package purchase

import (
	"time"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type PurchaseHistory struct {
	id          string
	totalAmount int64
	products    []PurchaseProduct
	purchasedAt time.Time
}

func NewPurchaseHistory(totalAmount int64, products []PurchaseProduct, now time.Time) (*PurchaseHistory, error) {
	return newPurchaseHistory(
		"",
		totalAmount,
		products,
		now,
	)
}

func Reconstruct(id string, totalAmount int64, products []PurchaseProduct, purchasedAt time.Time) (*PurchaseHistory, error) {
	return newPurchaseHistory(
		id,
		totalAmount,
		products,
		purchasedAt,
	)
}

func newPurchaseHistory(
	id string,
	totalAmount int64,
	products []PurchaseProduct,
	purchasedAt time.Time,
) (*PurchaseHistory, error) {
	// idが空文字の時は新規作成
	if id == "" {
		id = ulid.NewULID()
	}

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
		purchasedAt: purchasedAt,
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
