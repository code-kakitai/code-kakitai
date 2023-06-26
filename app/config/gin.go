package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type"}

	router.Use(cors.New(config))
	return router
}
