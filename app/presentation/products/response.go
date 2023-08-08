package products

type postProductResponse struct {
	Product productResponseModel `json:"product"`
}
type productResponseModel struct {
	Id          string `json:"id"`
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}
