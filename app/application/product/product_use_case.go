package product

import productDomain "github/code-kakitai/code-kakitai/domain/product"

type ProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewProductUseCase(
	productRepo productDomain.ProductRepository,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}
