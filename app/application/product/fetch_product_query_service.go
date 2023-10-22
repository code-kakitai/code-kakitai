//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package product

import "context"

type FetchProductListDto struct {
	ID        string
	Name      string
	Price     int64
	Stock     int
	OwnerID   string
	OwnerName string
}

type ProductQueryService interface {
	FetchProductList(ctx context.Context) ([]*FetchProductListDto, error)
}
