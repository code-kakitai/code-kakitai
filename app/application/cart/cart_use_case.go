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
	ProductID string
	Quantity  int
}

func (uc *CartUseCase) Run(ctx context.Context, userID string, dto CartUseCaseDto) (string, error) {
	// カートから商品情報を取得
	cart, err := uc.cartRepo.FindByUserID(ctx, userID)
	if err != nil {
		return "", err
	}

	// 在庫情報を取得
	product, err := uc.productRepo.FindByID(ctx, dto.ProductID)
	if err != nil {
		return "", err
	}
	if product == nil {
		return "", errors.NewError("商品が見つかりません。")
	}

	// 商品数が0の時はカートから商品を削除、それ以外は追加・更新
	if dto.Quantity == 0 {
		if err := cart.RemoveProduct(dto.ProductID); err != nil {
			return "", err
		}
	} else {
		if product.Consume(dto.Quantity); err != nil {
			return "", err
		}
		// カートに商品を追加
		if err := cart.AddProduct(dto.ProductID, dto.Quantity); err != nil {
			return "", err
		}
	}

	// カートの永続化
	if err := uc.cartRepo.Save(ctx, cart); err != nil {
		return "", err
	}
	return "", nil
}
