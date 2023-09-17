//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package transaction

import (
	"context"
)

type TransactionManager interface {
	RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
