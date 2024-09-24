package service

type MallResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type MallService interface {
	GetAllShoppingMall() ([]MallResponse, error)
}
