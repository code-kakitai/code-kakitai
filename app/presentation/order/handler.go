package order

import (
	"time"

	"github.com/gin-gonic/gin"

	orderApp "github/code-kakitai/code-kakitai/application/order"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

type handler struct {
	orderUseCase *orderApp.OrderUseCase
}

func NewHandler(orderUseCase *orderApp.OrderUseCase) handler {
	return handler{
		orderUseCase: orderUseCase,
	}
}

// OrderProducts godoc
// @Summary 注文をする
// @Tags orders
// @Accept json
// @Produce json
// @Param request body []OrderParams true "注文商品"
// @Success 200 {int} id
// @Router /v1/orders [post]
func (h handler) OrderProducts(ctx *gin.Context) {
	var params []*OrderParams
	// TODO リクエストのバリデーション
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		settings.ReturnBadRequest(ctx, err)
	}
	// todo userIDはsession等で別途取得する
	userID := "test_user_id"
	dtos := make([]orderApp.OrderUseCaseDto, len(params))
	for _, param := range params {
		dtos = append(dtos, orderApp.OrderUseCaseDto{
			ProductID: param.ProductID,
			Count:     param.Count,
		})
	}
	id, err := h.orderUseCase.Run(
		ctx.Request.Context(),
		userID,
		dtos,
		time.Now(),
	)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
	}

	settings.ReturnStatusCreated(ctx, id)
}
