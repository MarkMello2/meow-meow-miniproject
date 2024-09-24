package service

type ProductResponse struct {
	Id          int     `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Rating      int     `json:"rating"`
	Image       string  `json:"image"`
	CategoryId  int     `json:"category_id"`
	MallId      int     `json:"mall_id"`
}

type ProductService interface {
	GetAllProduct() ([]ProductResponse, error)
	GetProductById(int) ([]ProductResponse, error)
	GetProductByCategoryId(int) ([]ProductResponse, error)
}
