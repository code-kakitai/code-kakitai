//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE
package user

import "context"

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindById(ctx context.Context, id string) (*User, error)
	FindAll(ctx context.Context) ([]*User, error)
}
