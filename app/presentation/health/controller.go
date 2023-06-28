package presentation

import (
	"github/code-kakitai/code-kakitai/config"

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
	config.ReturnStatusOK(ctx, res)
}
