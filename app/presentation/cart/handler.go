package cart

import (
	"github.com/gin-gonic/gin"

	cartApp "github/code-kakitai/code-kakitai/application/cart"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

type handler struct {
	cartUseCase *cartApp.CartUseCase
}

func NewHandler(cartUseCase *cartApp.CartUseCase) handler {
	return handler{
		cartUseCase: cartUseCase,
	}
}

// PostCart godoc
// @Summary カートに商品を追加する
// @Tags carts
// @Accept json
// @Produce json
// @Param request body PostCartsParams true "カートの商品"
// @Router /v1/carts/ [post]
func (h handler) PostCart(ctx *gin.Context) {
	var param PostCartsParams
	// TODO リクエストのバリデーション
	if err := ctx.ShouldBindJSON(&param); err != nil {
		settings.ReturnBadRequest(ctx, err)
	}

	// todo userIDはsession等で別途取得する
	userID := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
	dto := cartApp.CartUseCaseDto{
		ProductID: param.ProductID,
		Quantity:  param.Quantity,
		UserID:    userID,
	}
	if err := h.cartUseCase.Run(
		ctx.Request.Context(),
		dto,
	); err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
	}
	settings.ReturnStatusNoContent(ctx)
}
