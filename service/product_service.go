package service

import (
	"errors"
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type productService struct {
	proRepo  repository.ProductRepository
	minioSrv MinioService
}

func NewProductService(proRepo repository.ProductRepository, minioSrv MinioService) ProductService {
	return productService{proRepo: proRepo, minioSrv: minioSrv}
}

func (p productService) GetAllProduct() ([]ProductResponse, error) {

	productDataDb, err := p.proRepo.GetAll()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		newUrlImg, err := p.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       newUrlImg,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}

	return res, nil
}

func (p productService) GetProductById(productId int) ([]ProductResponse, error) {
	productDataDb, err := p.proRepo.GetById(productId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "Product not found")
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		newUrlImg, err := p.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       newUrlImg,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}

	return res, nil
}

func (p productService) GetProductByCategoryId(categoryId int) ([]ProductResponse, error) {
	productDataDb, err := p.proRepo.GetByCategoryId(categoryId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		newUrlImg, err := p.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       newUrlImg,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}
	return res, nil
}

func (p productService) GetProductByMallId(mallId int) ([]ProductResponse, error) {
	productDataDb, err := p.proRepo.GetByMallId(mallId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponse{}

	for _, data := range productDataDb {
		newUrlImg, err := p.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, ProductResponse{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       newUrlImg,
			CategoryId:  data.CategoryId,
			MallId:      data.MallId,
		})
	}
	return res, nil
}

func (p productService) GetProductRecommended() ([]ProductResponseRec, error) {
	productDataDb, err := p.proRepo.GetRecommended()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []ProductResponseRec{}

	for _, data := range productDataDb {
		newUrlImg, err := p.minioSrv.getUrlImagePath(data.Image)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, ProductResponseRec{
			Id:          data.Id,
			Code:        data.Code,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Rating:      data.Rating,
			Image:       newUrlImg,
		})
	}
	return res, nil
}
