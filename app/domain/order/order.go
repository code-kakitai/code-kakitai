//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package order

import (
	"context"
	"time"

	"github.com/code-kakitai/go-pkg/ulid"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	errDomain "github/code-kakitai/code-kakitai/domain/error"
)

type Order struct {
	id          string
	userID      string
	totalAmount int64
	products    OrderProducts
	orderedAt   time.Time
}

func NewOrder(userID string, totalAmount int64, products []OrderProduct, now time.Time) (*Order, error) {
	return newOrder(
		ulid.NewULID(),
		userID,
		totalAmount,
		products,
		now,
	)
}

func Reconstruct(id string, userID string, totalAmount int64, products []OrderProduct, OrderedAt time.Time) (*Order, error) {
	return newOrder(
		id,
		userID,
		totalAmount,
		products,
		OrderedAt,
	)
}

func newOrder(
	id string,
	userID string,
	totalAmount int64,
	products []OrderProduct,
	orderedAt time.Time,
) (*Order, error) {
	// userIDのバリデーション
	if !ulid.IsValid(userID) {
		return nil, errDomain.NewError("ユーザーIDの値が不正です。")
	}

	// 購入金額のバリデーション
	// 割引等で合計金額が0円になることはあるため、0円以上を許容とする
	if totalAmount < 0 {
		return nil, errDomain.NewError("購入金額の値が不正です。")
	}

	// 購入商品のバリデーション
	if len(products) < 1 {
		return nil, errDomain.NewError("購入商品がありません。")
	}
	return &Order{
		id:          id,
		userID:      userID,
		totalAmount: totalAmount,
		products:    products,
		orderedAt:   orderedAt,
	}, nil
}

func (p *Order) ID() string {
	return p.id
}

func (p *Order) UserID() string {
	return p.userID
}

func (p *Order) TotalAmount() int64 {
	return p.totalAmount
}

func (p *Order) Products() []OrderProduct {
	return p.products
}

func (p *Order) OrderedAt() time.Time {
	return p.orderedAt
}

func (p *Order) ProductIDs() []string {
	var productIDs []string
	for _, product := range p.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

type OrderProducts []OrderProduct

func (p OrderProducts) ProductIDs() []string {
	var productIDs []string
	for _, product := range p {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

// 合計金額計算
func (p OrderProducts) TotalAmount() int64 {
	var totalAmount int64
	for _, product := range p {
		totalAmount += product.price * int64(product.quantity)
	}
	return totalAmount
}

type OrderProduct struct {
	productID string
	price     int64
	quantity  int
}

func NewOrderProduct(productID string, price int64, quantity int) (*OrderProduct, error) {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return nil, errDomain.NewError("商品IDの値が不正です。")
	}

	// 購入数のバリデーション
	if quantity < 1 {
		return nil, errDomain.NewError("購入数の値が不正です。")
	}

	return &OrderProduct{
		productID: productID,
		price:     price,
		quantity:  quantity,
	}, nil
}

func (p *OrderProduct) ProductID() string {
	return p.productID
}

func (p *OrderProduct) Quantity() int {
	return p.quantity
}

func (p *OrderProduct) Price() int64 {
	return p.price
}

type OrderDomainService interface {
	OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error)
}
