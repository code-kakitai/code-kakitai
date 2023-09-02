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
// @Param request body PostCartParams true "カートの商品"
// @Router /v1/carts/ [post]
func (h handler) PostCart(ctx *gin.Context) {
	var params PostCartParams
	// TODO リクエストのバリデーション
	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnBadRequest(ctx, err)
	}

	// todo sessionに保存しているuserIDとcartのuserIDが一致するか確認する
	dto := make([]cartApp.CartUseCaseDto, 0, len(params.CartProducts))
	for _, param := range params.CartProducts {
		dto = append(dto, cartApp.CartUseCaseDto{
			ProductID: param.ProductID,
			Count:     param.Quantity,
		})
	}
	// todo userIDはsession等で別途取得する
	userID := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
	if _, err := h.cartUseCase.Run(
		ctx.Request.Context(),
		userID,
		dto,
	); err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
	}
	settings.ReturnStatusCreated(ctx, "")
}
