package products

import (
	validator "github.com/code-kakitai/go-pkg/validator"
	"github.com/gin-gonic/gin"

	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/presentation/settings"
)

type handler struct {
	saveProductUseCase  *product.SaveProductUseCase
	fetchProductUseCase *product.FetchProductUseCase
}

func NewHandler(
	saveProductUseCase *product.SaveProductUseCase,
	fetchProductUseCase *product.FetchProductUseCase,
) handler {
	return handler{
		saveProductUseCase:  saveProductUseCase,
		fetchProductUseCase: fetchProductUseCase,
	}
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
	if err := ctx.ShouldBindJSON(&params); err != nil {
		settings.ReturnBadRequest(ctx, err)
		return
	}
	validate := validator.GetValidator()
	if err := validate.Struct(params); err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
		return
	}

	input := product.SaveProductUseCaseInputDto{
		OwnerID:     params.OwnerID,
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
		Stock:       params.Stock,
	}

	dto, err := h.saveProductUseCase.Run(ctx, input)
	if err != nil {
		settings.ReturnError(ctx, err)
		return
	}
	response := postProductResponse{
		productResponseModel{
			Id:          dto.ID,
			OwnerID:     dto.OwnerID,
			Name:        dto.Name,
			Description: dto.Description,
			Price:       dto.Price,
			Stock:       dto.Stock,
		},
	}
	settings.ReturnStatusCreated(ctx, response)
}

// GetProducts godoc
// @Summary 商品一覧を取得する
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} getProductsResponse
// @Router /v1/products [get]
func (h handler) GetProducts(ctx *gin.Context) {
	dtos, err := h.fetchProductUseCase.Run(ctx)
	if err != nil {
		settings.ReturnError(ctx, err)
	}

	var products []getProductsResponse
	for _, dto := range dtos {
		products = append(products, getProductsResponse{
			productResponseModel: &productResponseModel{
				Id:      dto.ID,
				OwnerID: dto.OwnerID,
				Name:    dto.Name,
				Price:   dto.Price,
				Stock:   dto.Stock,
			},
			OwnerName: dto.OwnerName,
		})
	}

	settings.ReturnStatusOK(ctx, products)
}
