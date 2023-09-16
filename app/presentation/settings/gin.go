package settings

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	router.Use(cors.New(config))
	return router
}

func ReturnStatusOK[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusOK, &body)
}

func ReturnStatusCreated[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusCreated, &body)
}

func ReturnStatusNoContent(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusNoContent)
}

func ReturnStatusBadRequest(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusBadRequest, err)
}

func ReturnBadRequest(ctx *gin.Context, err error) {
	ReturnStatusBadRequest(ctx, err)
}

func ReturnStatusUnauthorized(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusUnauthorized, err)
}

func ReturnUnauthorized(ctx *gin.Context, err error) {
	ReturnStatusUnauthorized(ctx, err)
}

func ReturnStatusForbidden(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusForbidden, err)
}

func ReturnForbidden(ctx *gin.Context, err error) {
	ReturnStatusForbidden(ctx, err)
}

func ReturnStatusNotFound(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusNotFound, err)
}

func ReturnNotFound(ctx *gin.Context, err error) {
	ReturnStatusNotFound(ctx, err)
}

func ReturnStatusInternalServerError(ctx *gin.Context, err error) {
	returnAbortWith(ctx, http.StatusInternalServerError, err)
}

func ReturnError(ctx *gin.Context, err error) {
	ctx.Error(err)
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
