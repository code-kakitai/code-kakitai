package product

import (
	"context"
	productDomain "github/code-kakitai/code-kakitai/domain/product"
)

type SaveProductUseCase struct {
	productRepo productDomain.ProductRepository
}

func NewSaveProductUseCase(
	productRepo productDomain.ProductRepository,
) *SaveProductUseCase {
	return &SaveProductUseCase{
		productRepo: productRepo,
	}
}

type SaveProductUseCaseDto struct {
	ID          string
	OwnerID     string
	Name        string
	Description string
	Price       int64
	Stock       int
}

func (uc *SaveProductUseCase) Run(
	ctx context.Context,
	OwnerID string,
	Name string,
	Description string,
	Price int64,
	Stock int,
) (*SaveProductUseCaseDto, error) {
	p, err := productDomain.NewProduct(OwnerID, Name, Description, Price, Stock)
	if err != nil {
		return nil, err
	}
	err = uc.productRepo.Save(ctx, p)
	if err != nil {
		return nil, err
	}
	return &SaveProductUseCaseDto{
		ID:          p.ID(),
		OwnerID:     p.OwnerID(),
		Name:        p.Name(),
		Description: p.Description(),
		Price:       p.Price(),
		Stock:       p.Stock(),
	}, nil
}
