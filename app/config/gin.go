package config

import (
	"net/http"

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

func Return200[T any](ctx *gin.Context, body T) {
	ctx.JSON(200, &body)
}

func Return201[T any](ctx *gin.Context, body T) {
	ctx.JSON(201, &body)
}

func Return204(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func Return400(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 400, err)
}

func ReturnBadRequest(ctx *gin.Context, err error) {
	Return400(ctx, err)
}

func Return401(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 401, err)
}

func ReturnUnauthorized(ctx *gin.Context, err error) {
	Return401(ctx, err)
}

func Return403(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 403, err)
}

func ReturnForbidden(ctx *gin.Context, err error) {
	Return403(ctx, err)
}

func Return404(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 404, err)
}

func ReturnNotFound(ctx *gin.Context, err error) {
	Return404(ctx, err)
}

func Return500(ctx *gin.Context, err error) {
	returnAbortWith(ctx, 500, err)
}

func returnAbortWith(ctx *gin.Context, code int, err error) {
	var msg string
	if err != nil {
		msg = err.Error()
	}

	ctx.AbortWithStatusJSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
}
