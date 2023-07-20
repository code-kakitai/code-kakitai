//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package owner

import "context"

type OwnerRepository interface {
	Save(ctx context.Context) error
	FindById(ctx context.Context, id string) (*Owner, error)
}
