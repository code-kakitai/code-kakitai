package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github/code-kakitai/code-kakitai/domain/error"
	"github/code-kakitai/code-kakitai/domain/user"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type userRepository struct{}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	query := db.GetQuery(ctx)
	u, err := query.UserFindById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NotFoundErr
		}
		return nil, err
	}
	ud, err := user.Reconstruct(
		u.ID,
		u.Email,
		u.PhoneNumber,
		u.LastName,
		u.FirstName,
		u.Prefecture,
		u.City,
		u.AddressExtra,
	)
	if err != nil {
		return nil, err
	}
	return ud, nil
}
func (r *userRepository) Save(ctx context.Context, u *user.User) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:           u.ID(),
		Email:        u.Email(),
		PhoneNumber:  u.PhoneNumber(),
		LastName:     u.LastName(),
		FirstName:    u.FirstName(),
		City:         u.City(),
		Prefecture:   u.Pref(),
		AddressExtra: u.AddressExtra(),
	}); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	query := db.GetQuery(ctx)
	us, err := query.UserFindAll(ctx)
	if err != nil {
		return nil, err
	}
	var users []*user.User
	for _, u := range us {
		user, err := user.Reconstruct(
			u.ID,
			u.Email,
			u.PhoneNumber,
			u.LastName,
			u.FirstName,
			u.Prefecture,
			u.City,
			u.AddressExtra,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
