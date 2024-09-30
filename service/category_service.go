package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type categoryService struct {
	cateRepo repository.CategoryRepository
	minioSrv MinioService
}

func NewCategoryService(cateRepo repository.CategoryRepository, minioSrv MinioService) CatagoryService {
	return categoryService{cateRepo: cateRepo, minioSrv: minioSrv}
}

func (c categoryService) GetAllCategory() ([]CategoryResponse, error) {
	cateData, err := c.cateRepo.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []CategoryResponse{}

	for _, data := range cateData {
		newUrlImg, err := c.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, CategoryResponse{
			Id:          data.Id,
			Name:        data.Name,
			Description: data.Description,
			Image:       newUrlImg,
		})
	}

	return res, nil
}
