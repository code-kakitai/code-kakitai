package repository

import (
	"context"
	"github/code-kakitai/code-kakitai/domain/user"
	"github/code-kakitai/code-kakitai/infrastructure/db/dbgen"
)

type userRepositoryImpl struct {
	query dbgen.Queries
}

func NewPlayerRepositoryImpl(query dbgen.Queries) user.UserRepository {
	return &userRepositoryImpl{query: query}
}

func (r *userRepositoryImpl) FindById(id string) (*user.User, error) {
	ctx := context.Background()
	u, err := r.query.UserFindById(ctx, id)
	if err != nil {
		return nil, err
	}

	address, err := user.NewAddress(
		u.Prefecture.String,
		u.City.String,
		u.AddressExtra.String,
	)
	if err != nil {
		return nil, err
	}

	return user.Reconstruct(
		u.ID,
		u.Email,
		u.PhoneNumber.String,
		u.Name.String,
		u.Name.String,
		address,
	), nil
}
func (r *userRepositoryImpl) Save(u *user.User) error {
	return nil
}
