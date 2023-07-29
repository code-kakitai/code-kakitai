package presentation

import (
	"github/code-kakitai/code-kakitai/infrastructure/mysql/db/dbgen"
	health_handler "github/code-kakitai/code-kakitai/presentation/health_handler"
	user "github/code-kakitai/code-kakitai/presentation/user"

	ginpkg "github.com/gin-gonic/gin"
)

func InitRoute(api *ginpkg.Engine, query *dbgen.Queries) {
	v1 := api.Group("/v1")
	v1.GET("/health", health_handler.HealthCheck)

	{
		user.Route(v1, query)
	}
}
