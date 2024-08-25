package product

import (
	"unicode/utf8"

	"github.com/code-kakitai/go-pkg/ulid"

	errDomain "github/code-kakitai/code-kakitai/domain/error"
)

type Product struct {
	id          string
	ownerID     string
	name        string
	description string
	price       int64
	stock       int
}

func Reconstruct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		id,
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		ulid.NewULID(),
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

func newProduct(
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	// ownerIDのバリデーション
	if !ulid.IsValid(ownerID) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("商品名の値が不正です。")
	}

	// 商品説明のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("商品説明の値が不正です。")
	}

	// 価格のバリデーション
	if price < 1 {
		return nil, errDomain.NewError("価格の値が不正です。")
	}

	// 在庫数のバリデーション
	// 在庫はないけど、商品は登録したい等あるため、0は許容する
	if stock < 0 {
		return nil, errDomain.NewError("在庫数の値が不正です。")
	}
	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) OwnerID() string {
	return p.ownerID
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Description() string {
	return p.description
}

func (p *Product) Price() int64 {
	return p.price
}

func (p *Product) Stock() int {
	return p.stock
}

func (p *Product) Consume(quantity int) error {
	if quantity < 0 {
		return errDomain.NewError("在庫数の値が不正です。")
	}

	if p.stock-quantity < 0 {
		return errDomain.NewError("在庫数が不足しています。")
	}
	p.stock -= quantity
	return nil
}

const (
	// 名前の最大値/最小値
	nameLengthMin = 1
	nameLengthMax = 100

	// 説明の最大値/最小値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)
