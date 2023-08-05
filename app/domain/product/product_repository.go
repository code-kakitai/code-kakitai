//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package product

import (
	"context"
)

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	FindByID(ctx context.Context, id string) (*Product, error)
	FindByIDs(ctx context.Context, ids []string) ([]*Product, error)
}
