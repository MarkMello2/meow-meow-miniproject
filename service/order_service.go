package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return orderService{orderRepo: orderRepo}
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
	return nil, nil
}
