package route

import (
	ginpkg "github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	orderApp "github/code-kakitai/code-kakitai/application/order"
	productApp "github/code-kakitai/code-kakitai/application/product"
	userApp "github/code-kakitai/code-kakitai/application/user"
	cartDomain "github/code-kakitai/code-kakitai/domain/cart"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/query_service"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/repository"
	health_handler "github/code-kakitai/code-kakitai/presentation/health_handler"
	orderPre "github/code-kakitai/code-kakitai/presentation/order"
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
	fetchQueryService := query_service.NewFetchProductQueryService()
	h := productPre.NewHandler(productApp.NewSaveProductUseCase(productRepository), productApp.NewFetchProductUseCase(fetchQueryService))
	group := r.Group("/products")
	group.GET("/", h.FetchProducts)
	group.POST("/", h.PostProducts)
}

func orderRoute(r *ginpkg.RouterGroup) {
	orderRepository := repository.NewOrderRepository()
	productRepository := repository.NewProductRepository()
	h := orderPre.NewHandler(
		orderApp.NewOrderUseCase(
			orderDomain.NewOrderDomainService(
				orderRepository,
				productRepository,
			),
			cartDomain.NewMockCartRepository(gomock.NewController(nil)), // todo impl
		),
	)
	group := r.Group("/orders")
	group.POST("/", h.OrderProducts)
}
