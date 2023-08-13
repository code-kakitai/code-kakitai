package route

import (
	ginpkg "github.com/gin-gonic/gin"

	productApp "github/code-kakitai/code-kakitai/application/product"
	userApp "github/code-kakitai/code-kakitai/application/user"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/repository"
	health_handler "github/code-kakitai/code-kakitai/presentation/health_handler"
	productPre "github/code-kakitai/code-kakitai/presentation/products"
	userPre "github/code-kakitai/code-kakitai/presentation/user"
)

func InitRoute(api *ginpkg.Engine) {
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)

	{
		userRoute(v1)
		productRoute(v1)
	}
}

func userRoute(r *ginpkg.RouterGroup) {
	userRepository := repository.NewUserRepository()
	h := userPre.NewHandler(
		userApp.NewFindUserUseCase(userRepository),
		userApp.NewSaveUserUseCase(userRepository),
	)
	group := r.Group("/users")
	group.GET("/:id", h.GetUserByID)
}

func productRoute(r *ginpkg.RouterGroup) {
	productRepository := repository.NewProductRepository()
	h := productPre.NewHandler(productApp.NewSaveProductUseCase(productRepository))
	group := r.Group("/products")
	group.GET("/", h.FetchProducts)
	group.POST("/", h.PostProducts)
}
