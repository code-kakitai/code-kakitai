package user

import (
	"github.com/gin-gonic/gin"

	userApp "github/code-kakitai/code-kakitai/application/user"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

type handler struct {
	findUserUseCase *userApp.FindUserUseCase
	saveUserUseCase *userApp.SaveUserUseCase
}

func NewHandler(
	findUserUseCase *userApp.FindUserUseCase,
	saveUserUseCase *userApp.SaveUserUseCase,
) handler {
	return handler{
		findUserUseCase: findUserUseCase,
		saveUserUseCase: saveUserUseCase,
	}
}

// GetUserByID godoc
// @Summary ユーザーを取得する
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} getUserResponse
// @Router /v1/users/{id} [get]
func (h handler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	dto, err := h.findUserUseCase.Run(ctx, id)
	if err != nil {
		settings.ReturnError(ctx, err)
		return
	}
	res := getUserResponse{
		User: userResponseModel{
			ID:          dto.ID,
			LastName:    dto.LastName,
			FirstName:   dto.FirstName,
			Email:       dto.Email,
			PhoneNumber: dto.PhoneNumber,
			Address:     dto.Address,
		},
	}
	settings.ReturnStatusOK(ctx, res)
}
