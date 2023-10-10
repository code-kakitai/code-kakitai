package cart

import (
	"github.com/gin-gonic/gin"

	cartApp "github/code-kakitai/code-kakitai/application/cart"
	"github/code-kakitai/code-kakitai/presentation/presenter"
)

type handler struct {
	addCartUseCase *cartApp.AddCartUseCase
}

func NewHandler(addCartUseCase *cartApp.AddCartUseCase) handler {
	return handler{
		addCartUseCase: addCartUseCase,
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
		presenter.ReturnBadRequest(ctx, err)
	}

	// todo userIDはsession等で別途取得する
	userID := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
	dto := cartApp.AddCartUseCaseInputDto{
		ProductID: param.ProductID,
		Quantity:  param.Quantity,
		UserID:    userID,
	}
	if err := h.addCartUseCase.Run(
		ctx.Request.Context(),
		dto,
	); err != nil {
		presenter.ReturnError(ctx, err)
	}
	presenter.ReturnStatusNoContent(ctx)
}
