//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package purchase

import (
	"context"
)

type PurchaseHistoryRepository interface {
	Save(ctx context.Context, history *PurchaseHistory) error
}
