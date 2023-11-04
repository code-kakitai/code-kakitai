package user

import (
	"context"
	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

type FindUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserUseCase(
	userRepo userDomain.UserRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

type FindUseCaseDto struct {
	ID          string
	Email       string
	PhoneNumber string
	LastName    string
	FirstName   string
	Address     string
}

func (uc *FindUserUseCase) Run(ctx context.Context, id string) (*FindUseCaseDto, error) {
	user, err := uc.userRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindUseCaseDto{
		ID:          user.ID(),
		Email:       user.Email(),
		PhoneNumber: user.PhoneNumber(),
		LastName:    user.LastName(),
		FirstName:   user.FirstName(),
		Address:     user.Pref() + user.City() + user.AddressExtra(),
	}, nil
}
