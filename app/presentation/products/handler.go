package products

import (
	"github/code-kakitai/code-kakitai/application/product"
	"github/code-kakitai/code-kakitai/presentation/presenter"

	validator "github.com/code-kakitai/go-pkg/validator"
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
		saveProductUseCase:  saveProductUseCase,
		fetchProductUseCase: fetchProductUseCase,
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
		presenter.ReturnBadRequest(ctx, err)
	}
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		presenter.ReturnStatusBadRequest(ctx, err)
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
		presenter.ReturnError(ctx, err)
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
	presenter.ReturnStatusCreated(ctx, response)
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
		presenter.ReturnError(ctx, err)
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

	presenter.ReturnStatusCreated(ctx, products)
}
