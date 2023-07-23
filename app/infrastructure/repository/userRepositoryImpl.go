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

func (r *userRepositoryImpl) FindById(ctx context.Context, id string) (*user.User, error) {
	u, err := r.query.UserFindById(ctx, id)
	if err != nil {
		return nil, err
	}
	ud, err := user.BuilderUser(
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
func (r *userRepositoryImpl) Save(ctx context.Context, u *user.User) error {
	return nil
}
