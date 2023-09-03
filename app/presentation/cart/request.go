package cart

type PostAddCartParams struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
