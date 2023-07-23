package user

import (
	"github.com/gin-gonic/gin"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func Route(r *gin.RouterGroup) {
	h := newHandler()

	group := r.Group("/user")
	group.GET("/", h.GetUsers)
}
