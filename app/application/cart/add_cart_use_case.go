package cart

import (
	"context"

	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

type AddCartUseCase struct {
	cartRepo    cartDomain.CartRepository
	productRepo productDomain.ProductRepository
}

func NewCartUseCase(
	cartRepo cartDomain.CartRepository,
	productRepo productDomain.ProductRepository,
) *AddCartUseCase {
	return &AddCartUseCase{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

type AddCartUseCaseInputDto struct {
	UserID    string
	ProductID string
	Quantity  int
}

func (uc *AddCartUseCase) Run(ctx context.Context, dto AddCartUseCaseInputDto) error {
	// 現在のカート情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, dto.UserID)
	if err != nil {
		return err
	}

	// 在庫情報を取得
	product, err := uc.productRepo.FindByID(ctx, dto.ProductID)
	if err != nil {
		return err
	}
	if err := product.Consume(dto.Quantity); err != nil {
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

func (uc *AddCartUseCase) updateCart(cart *cartDomain.Cart, dto AddCartUseCaseInputDto) error {
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
