package cart

import (
	validator "github.com/code-kakitai/go-pkg/validator"
	"github.com/gin-gonic/gin"

	cartApp "github/code-kakitai/code-kakitai/application/cart"
	"github/code-kakitai/code-kakitai/presentation/settings"
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
// @Router /v1/carts [post]
func (h handler) PostCart(ctx *gin.Context) {
	var param PostCartsParams
	if err := ctx.ShouldBindJSON(&param); err != nil {
		settings.ReturnBadRequest(ctx, err)
	}
	validate := validator.GetValidator()
	if err := validate.Struct(&param); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
	}

	// todo userIDはsession等で別途取得する
	userID := "01HCNYK0PKYZWB0ZT1KR0EPWGP"
	dto := cartApp.AddCartUseCaseInputDto{
		ProductID: param.ProductID,
		Quantity:  param.Quantity,
		UserID:    userID,
	}
	if err := h.addCartUseCase.Run(
		ctx.Request.Context(),
		dto,
	); err != nil {
		settings.ReturnError(ctx, err)
	}
	settings.ReturnStatusNoContent(ctx)
}
