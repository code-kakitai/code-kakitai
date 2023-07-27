package owner

import ownerDomain "github/code-kakitai/code-kakitai/domain/owner"

type OwnerUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewOwnerUseCase(
	ownerRepo ownerDomain.OwnerRepository,
) *OwnerUseCase {
	return &OwnerUseCase{
		ownerRepo: ownerRepo,
	}
}
