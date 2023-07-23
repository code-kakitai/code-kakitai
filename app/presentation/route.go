package presentation

import (
	health_handler "github/code-kakitai/code-kakitai/presentation/health_handler"
	user "github/code-kakitai/code-kakitai/presentation/user"

	ginpkg "github.com/gin-gonic/gin"
)

func InitRoute(api *ginpkg.Engine) {
	r := api.Group("/v1")
	r.GET("/health", health_handler.HealthCheck)

	{
		user.Route(r)
	}
}
