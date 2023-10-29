package order

import (
	"time"

	validator "github.com/code-kakitai/go-pkg/validator"
	"github.com/gin-gonic/gin"

	orderApp "github/code-kakitai/code-kakitai/application/order"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

type handler struct {
	saveOrderUseCase *orderApp.SaveOrderUseCase
}

func NewHandler(saveOrderUseCase *orderApp.SaveOrderUseCase) handler {
	return handler{
		saveOrderUseCase: saveOrderUseCase,
	}
}

// PostOrders godoc
// @Summary 注文をする
// @Tags orders
// @Accept json
// @Produce json
// @Param request body []PostOrdersParams true "注文商品"
// @Success 200 {int} id
// @Router /v1/orders [post]
func (h handler) PostOrders(ctx *gin.Context) {
	var params []*PostOrdersParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		settings.ReturnBadRequest(ctx, err)
	}
	validate := validator.GetValidator()
	if err := validate.Struct(&params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
	}
	// todo userIDはsession等で別途取得する
	userID := "test_user_id"
	dtos := make([]orderApp.SaveOrderUseCaseInputDto, len(params))
	for _, param := range params {
		dtos = append(dtos, orderApp.SaveOrderUseCaseInputDto{
			ProductID: param.ProductID,
			Quantity:  param.Quantity,
		})
	}
	id, err := h.saveOrderUseCase.Run(
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
