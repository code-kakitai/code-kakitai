package health_handler

import (
	"github/code-kakitai/code-kakitai/presentation/settings"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary ヘルスチェック
// @Description　ヘルスチェック
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /v1/health [get]
func HealthCheck(ctx *gin.Context) {
	res := HealthResponse{
		Status: "ok",
	}
	settings.ReturnStatusOK(ctx, res)
}
