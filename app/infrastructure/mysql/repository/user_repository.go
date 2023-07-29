package repository

import (
	"context"
	"github/code-kakitai/code-kakitai/domain/user"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
)

type userRepository struct {
	query *dbgen.Queries
}

func NewUserRepository(query *dbgen.Queries) user.UserRepository {
	return &userRepository{query: query}
}

func (r *userRepository) FindById(ctx context.Context, id string) (*user.User, error) {
	u, err := r.query.UserFindById(ctx, id)
	if err != nil {
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
	if err := r.query.UpsertUser(ctx, dbgen.UpsertUserParams{
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
