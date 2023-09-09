package cart

type PostCartsParams struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
