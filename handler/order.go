package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderSrv service.OrderService
}

func NewOrderHandler(orderSrv service.OrderService) orderHandler {
	return orderHandler{orderSrv: orderSrv}
}

func (o orderHandler) SaveOrder(c echo.Context) error {
	userId := c.Get("userId").(int)

	orderReq := []service.OrderRequest{}
	err := c.Bind(&orderReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = o.orderSrv.SaveOrder(orderReq, userId)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Product Order Successful")
}
