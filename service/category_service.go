package service

import (
	"meow-meow/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type categoryService struct {
	cateRepo repository.CategoryRepository
}

func NewCategoryService(cateRepo repository.CategoryRepository) CatagoryService {
	return categoryService{cateRepo: cateRepo}
}

func (c categoryService) GetAllCategory() ([]CategoryResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	cateData, err := c.cateRepo.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []CategoryResponse{}

	for _, data := range cateData {
		res = append(res, CategoryResponse{
			Id:          data.Id,
			Name:        data.Name,
			Description: data.Description,
			Image:       pathImg + data.Image,
		})
	}

	return res, nil
}
