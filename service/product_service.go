package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productService struct {
	proRepo repository.ProductRepository
}

func NewProductService(proRepo repository.ProductRepository) ProductService {
	return productService{proRepo: proRepo}
}

func (p productService) GetAllProduct() ([]ProductResponse, error) {
	return nil, nil
}

func (p productService) GetProductById(productId int) (*ProductResponse, error) {
	return nil, nil
}

func (p productService) GetProductByCategoryId(categoryId int) ([]ProductResponse, error) {
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
			Image:       data.Image,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}
	return res, nil
}
