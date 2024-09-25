package service

type FavoriteResponse struct {
	Id                 int     `json:"id"`
	Price              float32 `json:"price"`
	Quantity           int     `json:"quantity"`
	FavoriteDate       string  `json:"favorite_date"`
	ProductCode        string  `json:"product_code"`
	ProductName        string  `json:"product_name"`
	ProductDescription string  `json:"product_description"`
	ProductRating      int     `json:"product_rating"`
	ProductImage       string  `json:"product_image"`
	UserId             int     `json:"user_id"`
}

type FavoriteRequest struct {
	ProductId int     `json:"product_id"`
	Price     float32 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type FavoriteService interface {
	GetFavoriteByUserId(int) ([]FavoriteResponse, error)
	SaveFavorite(FavoriteRequest, int) error
	DeleteFavoriteById(int) error
}
