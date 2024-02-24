package route

import (
	ginpkg "github.com/gin-gonic/gin"

	cartApp "github.com/yumekumo/sauna-shop/application/cart"
	orderApp "github.com/yumekumo/sauna-shop/application/order"
	productApp "github.com/yumekumo/sauna-shop/application/product"
	userApp "github.com/yumekumo/sauna-shop/application/user"
	orderDomain "github.com/yumekumo/sauna-shop/domain/order"
	"github.com/yumekumo/sauna-shop/infrastructure/mysql/query_service"
	"github.com/yumekumo/sauna-shop/infrastructure/mysql/repository"
	redisRepo "github.com/yumekumo/sauna-shop/infrastructure/redis/repository"
	cartPre "github.com/yumekumo/sauna-shop/presentation/cart"
	health_handler "github.com/yumekumo/sauna-shop/presentation/health_handler"
	orderPre "github.com/yumekumo/sauna-shop/presentation/order"
	productPre "github.com/yumekumo/sauna-shop/presentation/products"
	"github.com/yumekumo/sauna-shop/presentation/settings"
	userPre "github.com/yumekumo/sauna-shop/presentation/user"
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
