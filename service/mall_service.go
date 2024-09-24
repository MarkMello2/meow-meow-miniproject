package service

import (
	"meow-meow/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type mallService struct {
	mallRepo repository.MallRepository
}

func NewMallService(mallRepo repository.MallRepository) MallService {
	return mallService{mallRepo: mallRepo}
}

func (m mallService) GetAllShoppingMall() ([]MallResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	mallData, err := m.mallRepo.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []MallResponse{}

	for _, data := range mallData {
		res = append(res, MallResponse{
			Id:          data.Id,
			Name:        data.Name,
			Description: data.Description,
			Image:       pathImg + data.Image,
		})
	}

	return res, nil
}
