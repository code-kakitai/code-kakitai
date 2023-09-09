package cart

import (
	"time"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type CartProduct struct {
	productID string
	quantity     int
}

var CartTimeOut = time.Minute * 30

func (cp *CartProduct) ProductID() string {
	return cp.productID
}

func (cp *CartProduct) Quantity() int {
	return cp.quantity
}

type Cart struct {
	userID   string
	products []CartProduct
}

func NewCart(userID string) (*Cart, error) {
	if !ulid.IsValid(userID) {
		return nil, errors.NewError("ユーザーIDの値が不正です。")
	}
	return &Cart{
		userID:   userID,
		products: []CartProduct{},
	}, nil
}

func (p *Cart) UserID() string {
	return p.userID
}

func (p *Cart) Products() []CartProduct {
	return p.products
}

func (p *Cart) ProductIDs() []string {
	var productIDs []string
	for _, product := range p.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

func (p *Cart) AddProduct(productID string, quantity int) error {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return errors.NewError("商品IDの値が不正です。")
	}

	// 購入数のバリデーション
	if quantity < 1 {
		return errors.NewError("購入数の値が不正です。")
	}

	// 商品がすでにカートに入っている場合は更新
	for _, product := range p.products {
		if product.productID == productID {
			product.quantity = quantity
			return nil
		}
	}

	// 商品がカートに入っていない場合は追加
	p.products = append(p.products, CartProduct{
		productID: productID,
		quantity:     quantity,
	})

	return nil
}

func (p *Cart) RemoveProduct(productID string) error {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return errors.NewError("商品IDの値が不正です。")
	}

	// 商品がカートに入っているかチェック
	var newProducts []CartProduct
	for _, product := range p.products {
		if product.productID == productID {
			continue
		}
		newProducts = append(newProducts, product)
	}

	p.products = newProducts

	return nil
}
