package user

import (
	"github/code-kakitai/code-kakitai/presentation/settings"

	"github.com/gin-gonic/gin"
)

type handler struct {
}

func newHandler() handler {
	return handler{} // TODO: 依存関係を追加する
}

func Route(r *gin.RouterGroup) {
	h := newHandler()

	group := r.Group("/user")
	group.GET("/", h.GetUsers)
}

func (h handler) GetUsers(ctx *gin.Context) {
	res := GetUsersResponse{
		Status: "ok",
	}

	settings.ReturnStatusOK(ctx, res)
}
