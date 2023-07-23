package user

import (
	"github/code-kakitai/code-kakitai/presentation/settings"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(ctx *gin.Context) {
	res := GetUsersResponse{
		Status: "ok",
	}
	settings.ReturnStatusOK(ctx, res)
}
