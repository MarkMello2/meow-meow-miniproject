package service

type CategoryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type CatagoryService interface {
	GetAllCategory() ([]CategoryResponse, error)
}
