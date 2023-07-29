package user

import (
	userApp "github/code-kakitai/code-kakitai/application/user"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/repository"
	"github/code-kakitai/code-kakitai/presentation/settings"

	"github.com/gin-gonic/gin"
)

type handler struct {
	findUserUseCase *userApp.FindUserUseCase
	saveUserUseCase *userApp.SaveUserUseCase
}

func newHandler(
	findUserUseCase *userApp.FindUserUseCase,
	saceUserUseCase *userApp.SaveUserUseCase,
) handler {
	return handler{
		findUserUseCase: findUserUseCase,
		saveUserUseCase: saceUserUseCase,
	}
}

func Route(r *gin.RouterGroup, query *dbgen.Queries) {
	userRepository := repository.NewUserRepository(query)
	h := newHandler(
		userApp.NewFindUserUseCase(userRepository),
		userApp.NewSaveUserUseCase(userRepository),
	)
	group := r.Group("/users")
	group.GET("/:id", h.GetUserByID)
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
		settings.ReturnNotFound(ctx, err)
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
