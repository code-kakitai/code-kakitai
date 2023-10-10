package health_handler

import (
	"github/code-kakitai/code-kakitai/presentation/presenter"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary ヘルスチェック
// @Description ヘルスチェック
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /v1/health [get]
func HealthCheck(ctx *gin.Context) {
	res := HealthResponse{
		Status: "ok",
	}
	presenter.ReturnStatusOK(ctx, res)
}
