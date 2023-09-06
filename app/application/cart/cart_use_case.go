package cart

import (
	"context"

	"github.com/code-kakitai/go-pkg/errors"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

type CartUseCase struct {
	cartRepo    cartDomain.CartRepository
	productRepo productDomain.ProductRepository
}

func NewCartUseCase(
	cartRepo cartDomain.CartRepository,
	productRepo productDomain.ProductRepository,
) *CartUseCase {
	return &CartUseCase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

type CartUseCaseDto struct {
	UserID    string
	ProductID string
	Quantity  int
}

func (uc *CartUseCase) Run(ctx context.Context, dto CartUseCaseDto) error {
	// 現在のカート情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, dto.UserID)
	if err != nil {
		return err
	}

	// 在庫情報を取得し、在庫を確認
	product, err := uc.productRepo.FindByID(ctx, dto.ProductID)
	if err != nil {
		return err
	}
	// todo 「NoRowsエラーの場合は商品が見つからないので、エラーを返す」という形に変えること
	if product == nil {
		return errors.NewError("商品が見つかりません。")
	}
	if product.Consume(dto.Quantity); err != nil {
		return err
	}

	// カートの更新
	if err := uc.updateCart(cart, dto); err != nil {
		return err
	}

	// カートの永続化
	if err := uc.cartRepo.Save(ctx, cart); err != nil {
		return err
	}
	return nil
}

func (uc *CartUseCase) updateCart(cart *cartDomain.Cart, dto CartUseCaseDto) error {
	// 商品数が0の時はカートから商品を削除、それ以外は追加・更新
	if dto.Quantity == 0 {
		if err := cart.RemoveProduct(dto.ProductID); err != nil {
			return err
		}
		return nil
	}
	// カートに商品を追加
	if err := cart.AddProduct(dto.ProductID, dto.Quantity); err != nil {
		return err
	}
	return nil
}
