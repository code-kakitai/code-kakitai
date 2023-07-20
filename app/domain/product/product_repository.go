//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package product

import (
	"context"
)

type ProductRepository interface {
	Store(ctx context.Context, product *Product) error
	FindByOwnerID(ctx context.Context, ownerID string) ([]*Product, error)
	FindByID(ctx context.Context, id string) (*Product, error)
}
