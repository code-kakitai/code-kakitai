package product

import (
	"context"
)

type FetchProductUseCase struct {
	productQueryService ProductQueryService
}

func NewFetchProductUseCase(
	productQueryService ProductQueryService,
) *FetchProductUseCase {
	return &FetchProductUseCase{
		productQueryService: productQueryService,
	}
}

type FetchProductUseCaseDto struct {
	ID        string
	Name      string
	Price     int64
	Stock     int
	OwnerID   string
	OwnerName string
}

func (uc *FetchProductUseCase) Run(ctx context.Context) ([]*FetchProductUseCaseDto, error) {
	qsDtos, err := uc.productQueryService.FetchProductList(ctx)
	if err != nil {
		return nil, err
	}
	var ucDtos []*FetchProductUseCaseDto

	for _, qsDto := range qsDtos {
		ucDtos = append(ucDtos, &FetchProductUseCaseDto{
			ID:        qsDto.ID,
			Name:      qsDto.Name,
			Price:     qsDto.Price,
			Stock:     qsDto.Stock,
			OwnerID:   qsDto.OwnerID,
			OwnerName: qsDto.OwnerName,
		})
	}
	return ucDtos, err
}
