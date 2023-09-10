package user

import (
	"context"
	userDomain "github/code-kakitai/code-kakitai/domain/user"
)

type SaveUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewSaveUserUseCase(
	userRepo userDomain.UserRepository,
) *SaveUserUseCase {
	return &SaveUserUseCase{
		userRepo: userRepo,
	}
}

type SaveUseCaseDto struct {
	LastName     string
	FirstName    string
	Email        string
	PhoneNumber  string
	Prefecture   string
	City         string
	AddressExtra string
}

func (uc *SaveUserUseCase) Run(ctx context.Context, dto SaveUseCaseDto) error {
	// dtoからuserへ変換
	user, err := userDomain.NewUser(dto.Email, dto.PhoneNumber, dto.LastName, dto.FirstName, dto.Prefecture, dto.City, dto.AddressExtra)
	if err != nil {
		return err
	}
	return uc.userRepo.Save(ctx, user)
}
