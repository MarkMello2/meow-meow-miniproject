package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bannerService struct {
	ban      repository.BannerRepository
	minioSrv MinioService
}

func NewBannerService(ban repository.BannerRepository, minioSrv MinioService) BannerService {
	return bannerService{ban: ban, minioSrv: minioSrv}
}

func (b bannerService) GetBannerAll() ([]BannerResponse, error) {
	bannerData, err := b.ban.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []BannerResponse{}

	for _, data := range bannerData {
		newUrlImg, err := b.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, BannerResponse{
			Id:          data.Id,
			Name:        data.Name,
			Description: data.Description,
			Image:       newUrlImg,
		})
	}

	return res, nil
}
