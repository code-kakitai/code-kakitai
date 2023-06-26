package presentation

import (
	pre "github/code-kakitai/code-kakitai/presentation/health"

	"github.com/gin-gonic/gin"
)

func InitRoute(api *gin.Engine) {
	r := api.Group("/v1")
	r.GET("/health", pre.HealthCheck)
	
	// ur := r.Group("/users")
	{
		// ur.GET("/users", pre.GetUsers)
	}
}
