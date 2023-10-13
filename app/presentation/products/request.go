package products

type PostProductsParams struct {
	OwnerID     string `json:"owner_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int64  `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required"`
}
