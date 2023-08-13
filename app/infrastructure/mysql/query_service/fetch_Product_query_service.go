package query_service

import (
	"context"

	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
)

type fetchProductQueryService struct{}

func NewOrderRepository() product.FetchProductQueryService {
	return &fetchProductQueryService{}
}

func (q *fetchProductQueryService) Run(ctx context.Context) ([]*product.FetchProductQueryServiceDto, error) {
	query := db.GetQuery(ctx)
	productWithOwners, err := query.ProductFetchWithOwner(ctx)
	if err != nil {
		return nil, err
	}

	var productFetchServiceDtos []*product.FetchProductQueryServiceDto
	for _, productWithOwner := range productWithOwners {
		productFetchServiceDtos = append(productFetchServiceDtos, &product.FetchProductQueryServiceDto{
			ID:          productWithOwner.ID,
			Name:        productWithOwner.Name,
			Description: productWithOwner.Description,
			Price:       productWithOwner.Price,
			Stock:       int(productWithOwner.Stock),
			OwnerID:     productWithOwner.OwnerID,
			OwnerName:   productWithOwner.OwnerName.String,
		})
	}
	return productFetchServiceDtos, nil
}
