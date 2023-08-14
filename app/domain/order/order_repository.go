//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package order

import (
	"context"
)

type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
}
