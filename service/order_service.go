package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderService struct {
	orderRepo repository.OrderRepository
	minioSrv  MinioService
}

func NewOrderService(orderRepo repository.OrderRepository, minioSrv MinioService) OrderService {
	return orderService{orderRepo: orderRepo, minioSrv: minioSrv}
}

func (o orderService) SaveOrder(orderReq []OrderRequest, userId int) error {
	cartData := []repository.Order{}

	for _, v := range orderReq {

		if v.ProductId == 0 || v.Price == 0 || v.Quantity == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "ProductId, Price, and Quantity is require")
		}

		cartData = append(cartData, repository.Order{
			Price:     v.Price,
			Quantity:  v.Quantity,
			ProductId: v.ProductId,
			UserId:    userId,
		})
	}

	err := o.orderRepo.Save(cartData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}

func (o orderService) GetOrderByUserId(userId int) ([]OrderResponse, error) {
	orderData, err := o.orderRepo.GetOrder(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []OrderResponse{}

	for _, v := range orderData {
		newUrlImg, err := o.minioSrv.getUrlImagePath(v.ProductImage)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, OrderResponse{
			Id:                 v.Id,
			Price:              v.Price,
			Quantity:           v.Quantity,
			OrderDate:          v.OrderDate,
			ProductCode:        v.ProductCode,
			ProductName:        v.ProductName,
			ProductDescription: v.ProductDescription,
			ProductRating:      v.ProductRating,
			ProductImage:       newUrlImg,
			UserId:             v.UserId,
		})
	}

	return res, nil
}
