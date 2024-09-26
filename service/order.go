package service

type OrderResponse struct {
	Id                 int     `json:"id"`
	Price              float32 `json:"price"`
	Quantity           int     `json:"quantity"`
	OrderDate          string  `json:"order_date"`
	ProductCode        string  `json:"product_code"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductRating      int     `json:"product_rating"`
	ProductImage       string  `json:"product_image"`
	UserId             int     `json:"user_id"`
}

type OrderRequest struct {
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type OrderService interface {
	SaveOrder([]OrderRequest, int) error
	GetOrderByUserId(int) ([]OrderResponse, error)
}
