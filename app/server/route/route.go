package route

import (
	ginpkg "github.com/gin-gonic/gin"

	health_handler "github/code-kakitai/code-kakitai/presentation/health_handler"
)

func InitRoute(api *ginpkg.Engine) {
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)

	{
		userRoute(v1)
	}
}
