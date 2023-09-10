package cart

import (
	"time"

	"github.com/code-kakitai/go-pkg/errors"
	"github.com/code-kakitai/go-pkg/ulid"
)

type cartProduct struct {
	productID string
	quantity  int
}

func (cp *cartProduct) ProductID() string {
	return cp.productID
}

func (cp *cartProduct) Quantity() int {
	return cp.quantity
}

func newCartProduct(productID string, quantity int) (*cartProduct, error) {
	if !ulid.IsValid(productID) {
		return nil, errors.NewError("商品IDの値が不正です。")
	}

	if quantity < 1 {
		return nil, errors.NewError("購入数の値が不正です。")
	}

	return &cartProduct{
		productID: productID,
		quantity:  quantity,
	}, nil
}

var CartTimeOut = time.Minute * 30

type Cart struct {
	userID   string
	products []cartProduct
}

func NewCart(userID string) (*Cart, error) {
	if !ulid.IsValid(userID) {
		return nil, errors.NewError("ユーザーIDの値が不正です。")
	}
	return &Cart{
		userID:   userID,
		products: []cartProduct{},
	}, nil
}

func (p *Cart) UserID() string {
	return p.userID
}

func (p *Cart) Products() []cartProduct {
	return p.products
}

func (p *Cart) ProductIDs() []string {
	var productIDs []string
	for _, product := range p.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

func (p *Cart) HasProduct(productID string) bool {
	for _, product := range p.products {
		if product.productID == productID {
			return true
		}
	}
	return false
}

func (p *Cart) QuantityByProductID(productID string) (int, error) {
	for _, product := range p.products {
		if product.productID == productID {
			return product.quantity, nil
		}
	}
	return 0, errors.NewError("カートの商品が見つかりません。")
}

func (p *Cart) AddProduct(productID string, quantity int) error {
	cp, err := newCartProduct(productID, quantity)
	if err != nil {
		return err
	}

	// 商品がすでにカートに入っている場合は更新
	for k, product := range p.products {
		if product.productID == productID {
			p.products[k] = *cp
			return nil
		}
	}

	// 商品がカートに入っていない場合は追加
	p.products = append(p.products, *cp)

	return nil
}

func (p *Cart) RemoveProduct(productID string) error {
	// 商品IDのバリデーション
	if !ulid.IsValid(productID) {
		return errors.NewError("商品IDの値が不正です。")
	}

	// 商品がカートに入っているかチェック
	var newProducts []cartProduct
	for _, product := range p.products {
		if product.productID == productID {
			continue
		}
		newProducts = append(newProducts, product)
	}

	p.products = newProducts

	return nil
}
