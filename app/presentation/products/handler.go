package products

import (
	"github.com/gin-gonic/gin"
	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

type handler struct {
	saveProductUseCase *product.SaveProductUseCase
}

func NewHandler(saveProductUseCase *product.SaveProductUseCase) handler {
	return handler{
		saveProductUseCase: saveProductUseCase,
	}
}

type PostProductsParams struct {
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}

// PostProducts godoc
// @Summary 商品を保存する
// @Tags products
// @Accept json
// @Produce json
// @Param request body PostProductsParams true "登録商品"
// @Success 201 {object} postProductResponse
// @Router /v1/products [post]
func (h handler) PostProducts(ctx *gin.Context) {
	var params PostProductsParams
	// TODO リクエストのバリデーション
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		settings.ReturnBadRequest(ctx, err)
	}
	p, err := h.saveProductUseCase.Run(
		ctx,
		params.OwnerID,
		params.Name,
		params.Description,
		params.Price,
		params.Stock)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
	}
	response := postProductResponse{
		productResponseModel{
			Id:          p.ID(),
			OwnerID:     p.OwnerID(),
			Name:        p.Name(),
			Description: p.Description(),
			Price:       p.Price(),
			Stock:       p.Stock(),
		},
	}
	settings.ReturnStatusCreated(ctx, response)
}
