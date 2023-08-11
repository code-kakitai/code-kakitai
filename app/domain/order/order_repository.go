//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package order

import (
	"context"
)

type OrderHistoryRepository interface {
	Save(ctx context.Context, history *Order) error
}
