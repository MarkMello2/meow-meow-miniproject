package service

type CartResponse struct {
	Id                 int     `json:"id"`
	Price              float32 `json:"price"`
	Quantity           int     `json:"quantity"`
	CartDate           string  `json:"cart_date"`
	ProductCode        string  `json:"product_code"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductRating      int     `json:"product_rating"`
	ProductImage       string  `json:"product_image"`
	UserId             int     `json:"user_id"`
}

type CartRequest struct {
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type CartService interface {
	SaveCart([]CartRequest, int) error
	GetCartByUserId(int) ([]CartResponse, error)
}
