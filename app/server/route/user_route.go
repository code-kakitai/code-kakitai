package route

import (
	"github.com/gin-gonic/gin"

	userApp "github/code-kakitai/code-kakitai/application/user"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/repository"
	userPre "github/code-kakitai/code-kakitai/presentation/user"
)

func userRoute(r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepository),
		userApp.NewSaveUserUseCase(userRepository),
	)
	group := r.Group("/users")
	group.GET("/:id", h.GetUserByID)
}
