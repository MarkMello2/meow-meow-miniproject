package service

type BannerResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type BannerService interface {
	GetBannerAll() ([]BannerResponse, error)
}
