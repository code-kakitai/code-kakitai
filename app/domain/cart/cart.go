package cart

import (
	"time"

	"github.com/code-kakitai/go-pkg/ulid"

	errDomain "github/code-kakitai/code-kakitai/domain/error"
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
		return nil, errDomain.NewError("商品IDの値が不正です。")
	}

	if quantity < 1 {
		return nil, errDomain.NewError("購入数の値が不正です。")
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
		return nil, errDomain.NewError("ユーザーIDの値が不正です。")
	}
	return &Cart{
		userID:   userID,
		products: []cartProduct{},
	}, nil
}

func (c *Cart) UserID() string {
	return c.userID
}

func (c *Cart) Products() []cartProduct {
	return c.products
}

func (c *Cart) ProductIDs() []string {
	var productIDs []string
	for _, product := range c.products {
		productIDs = append(productIDs, product.productID)
	}
	return productIDs
}

func (c *Cart) QuantityByProductID(productID string) (int, error) {
	for _, product := range c.products {
		if product.productID == productID {
			return product.quantity, nil
		}
	}
	return 0, errDomain.NewError("カートの商品が見つかりません。")
}

func (c *Cart) AddProduct(productID string, quantity int) error {
	cp, err := newCartProduct(productID, quantity)
	if err != nil {
		return err
	}

	// 商品がすでにカートに入っている場合は更新
	for k, product := range c.products {
		if product.productID == productID {
			c.products[k] = *cp
			return nil
		}
	}

	// 商品がカートに入っていない場合は追加
	c.products = append(c.products, *cp)

	return nil
}

func (c *Cart) RemoveProduct(productID string) error {
	// 商品がカートに入っているかチェック
	var newProducts []cartProduct
	for _, product := range c.products {
		if product.productID == productID {
			continue
		}
		newProducts = append(newProducts, product)
	}

	c.products = newProducts

	return nil
}
