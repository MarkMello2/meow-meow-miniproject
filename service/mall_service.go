package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type mallService struct {
	mallRepo repository.MallRepository
	minioSrv MinioService
}

func NewMallService(mallRepo repository.MallRepository, minioSrv MinioService) MallService {
	return mallService{mallRepo: mallRepo, minioSrv: minioSrv}
}

func (m mallService) GetAllShoppingMall() ([]MallResponse, error) {
	mallData, err := m.mallRepo.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []MallResponse{}

	for _, data := range mallData {
		newUrlImg, err := m.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, MallResponse{
			Id:          data.Id,
			Name:        data.Name,
			Description: data.Description,
			Image:       newUrlImg,
		})
	}

	return res, nil
}
