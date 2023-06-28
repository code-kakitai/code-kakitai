package presentation

import (
	helthPresentation "github/code-kakitai/code-kakitai/presentation/health"

	"github.com/gin-gonic/gin"
)

func InitRoute(api *gin.Engine) {
	r := api.Group("/v1")
	r.GET("/health", helthPresentation.HealthCheck)

	// ur := r.Group("/users")
	{
		// ur.GET("/", pre.GetUsers)
	}
}
