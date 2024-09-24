package service

import (
	"errors"
	"meow-meow/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type productService struct {
	proRepo repository.ProductRepository
}

func NewProductService(proRepo repository.ProductRepository) ProductService {
	return productService{proRepo: proRepo}
}

func (p productService) GetAllProduct() ([]ProductResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	productDataDb, err := p.proRepo.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       pathImg + data.Image,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}
	return res, nil
}

func (p productService) GetProductById(productId int) ([]ProductResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	productDataDb, err := p.proRepo.GetById(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Product not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       pathImg + data.Image,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}

	return res, nil
}

func (p productService) GetProductByCategoryId(categoryId int) ([]ProductResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	productDataDb, err := p.proRepo.GetByCategoryId(categoryId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       pathImg + data.Image,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}
	return res, nil
}

func (p productService) GetProductByMallId(mallId int) ([]ProductResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	productDataDb, err := p.proRepo.GetByMallId(mallId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       pathImg + data.Image,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}
	return res, nil
}
