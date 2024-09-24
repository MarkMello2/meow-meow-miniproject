package service

import (
	"meow-meow/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type bannerService struct {
	ban repository.BannerRepository
}

func NewBannerService(ban repository.BannerRepository) BannerService {
	return bannerService{ban: ban}
}

func (b bannerService) GetBannerAll() ([]BannerResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	bannerData, err := b.ban.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []BannerResponse{}

	for _, data := range bannerData {
		res = append(res, BannerResponse{
			Id:          data.Id,
			Name:        data.Name,
			Description: data.Description,
			Image:       pathImg + data.Image,
		})
	}

	return res, nil
}
