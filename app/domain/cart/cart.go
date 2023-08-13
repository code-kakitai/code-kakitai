package cart

import (
	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type CartProduct struct {
	productID string
	count     int
}

func (cp *CartProduct) ProductID() string {
	return cp.productID
}

func (cp *CartProduct) Count() int {
	return cp.count
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
		userID: userID,
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

func (p *Cart) AddProduct(productID string, count int) error {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return errors.NewError("商品IDの値が不正です。")
	}

	// 購入数のバリデーション
	if count < 1 {
		return errors.NewError("購入数の値が不正です。")
	}

	// 商品がすでにカートに入っているかチェック
	for _, product := range p.products {
		if product.productID == productID {
			return errors.NewError("商品がすでにカートに入っています。")
		}
	}

	p.products = append(p.products, CartProduct{
		productID: productID,
		count:     count,
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
