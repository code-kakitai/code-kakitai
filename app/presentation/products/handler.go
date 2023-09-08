package products

import (
	"github.com/go-playground/validator/v10"
	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/presentation/settings"

	"github.com/gin-gonic/gin"
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
		saveProductUseCase: saveProductUseCase,
	}
}

type PostProductsParams struct {
	OwnerID     string `json:"owner_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
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
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		settings.ReturnBadRequest(ctx, err)
	}
	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
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
		settings.ReturnStatusInternalServerError(ctx, err)
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

// FetchProduct godoc
// @Summary 商品一覧を取得する
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} fetchProductResponse
// @Router /v1/products [get]
func (h handler) FetchProducts(ctx *gin.Context) {
	dtos, err := h.fetchProductUseCase.Run(ctx)
	if err != nil {
		settings.ReturnStatusInternalServerError(ctx, err)
	}

	var products []productsWithOwnerModel
	for _, dto := range dtos {
		products = append(products, productsWithOwnerModel{
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

	settings.ReturnStatusCreated(ctx, products)
}
