package route

import (
	ginpkg "github.com/gin-gonic/gin"

	cartApp "github/code-kakitai/code-kakitai/application/cart"
	orderApp "github/code-kakitai/code-kakitai/application/order"
	productApp "github/code-kakitai/code-kakitai/application/product"
	userApp "github/code-kakitai/code-kakitai/application/user"
	orderDomain "github/code-kakitai/code-kakitai/domain/order"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/query_service"
	"github/code-kakitai/code-kakitai/infrastructure/mysql/repository"
	redisRepo "github/code-kakitai/code-kakitai/infrastructure/redis/repository"
	cartPre "github/code-kakitai/code-kakitai/presentation/cart"
	health_handler "github/code-kakitai/code-kakitai/presentation/health_handler"
	orderPre "github/code-kakitai/code-kakitai/presentation/order"
	productPre "github/code-kakitai/code-kakitai/presentation/products"
	"github/code-kakitai/code-kakitai/presentation/settings"
	userPre "github/code-kakitai/code-kakitai/presentation/user"
)

func InitRoute(api *ginpkg.Engine) {
	api.Use(settings.ErrorHandler())
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)

	{
		userRoute(v1)
		productRoute(v1)
		cartRoute(v1)
		orderRoute(v1)
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
	fetchQueryService := query_service.NewProductQueryService()
	h := productPre.NewHandler(
		productApp.NewSaveProductUseCase(productRepository),
		productApp.NewFetchProductUseCase(fetchQueryService),
	)
	group := r.Group("/products")
	group.GET("", h.GetProducts)
	group.POST("", h.PostProducts)
}

func orderRoute(r *ginpkg.RouterGroup) {
	orderRepository := repository.NewOrderRepository()
	productRepository := repository.NewProductRepository()
	transactionManager := repository.NewTransactionManager()
	h := orderPre.NewHandler(
		orderApp.NewSaveOrderUseCase(
			orderDomain.NewOrderDomainService(
				orderRepository,
				productRepository,
			),
			redisRepo.NewCartRepository(),
			transactionManager,
		),
	)
	group := r.Group("/orders")
	group.POST("", h.PostOrders)
}

func cartRoute(r *ginpkg.RouterGroup) {
	cartRepository := redisRepo.NewCartRepository()
	productRepository := repository.NewProductRepository()
	h := cartPre.NewHandler(
		cartApp.NewCartUseCase(
			cartRepository,
			productRepository,
		),
	)
	group := r.Group("/carts")
	group.POST("", h.PostCart)
}
