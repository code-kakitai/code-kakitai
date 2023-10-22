package query_service

import (
	"context"

	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
)

type productQueryService struct{}

func NewProductQueryService() product.ProductQueryService {
	return &productQueryService{}
}

func (q *productQueryService) FetchProductList(ctx context.Context) ([]*product.FetchProductListDto, error) {
	query := db.GetReadQuery()
	productWithOwners, err := query.ProductFetchWithOwner(ctx)
	if err != nil {
		return nil, err
	}

	var productFetchServiceDtos []*product.FetchProductListDto
	for _, productWithOwner := range productWithOwners {
		productFetchServiceDtos = append(productFetchServiceDtos, &product.FetchProductListDto{
			ID:        productWithOwner.ID,
			Name:      productWithOwner.Name,
			Price:     productWithOwner.Price,
			Stock:     int(productWithOwner.Stock),
			OwnerID:   productWithOwner.OwnerID,
			OwnerName: productWithOwner.OwnerName.String,
		})
	}
	return productFetchServiceDtos, nil
}
