package cart

type PostCartParams struct {
	UserID       string         `json:"user_id"`
	CartProducts []*CartProduct `json:"cart_products"`
}

type CartProduct struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
