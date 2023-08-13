package query_service

import (
	"context"
	"fmt"

	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
)

type fetchProductQueryService struct{}

func NewOrderRepository() product.FetchProductQueryService {
	return &fetchProductQueryService{}
}

func (q *fetchProductQueryService) Run(ctx context.Context) (*product.FetchProductQueryServiceDto, error) {
	query := db.GetQuery(ctx)
	fmt.Printf("query: %v\n", query)

	return &product.FetchProductQueryServiceDto{}, nil
}
