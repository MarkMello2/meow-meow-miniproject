package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartSrv service.CartService
}

func NewCartHandler(cartSrv service.CartService) cartHandler {
	return cartHandler{cartSrv: cartSrv}
}

func (ca cartHandler) SaveCart(c echo.Context) error {
	userId := c.Get("userId").(int)

	cartReq := []service.CartRequest{}
	err := c.Bind(&cartReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = ca.cartSrv.SaveCart(cartReq, userId)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Add Product Cart Successful")
}

func (ca cartHandler) GetCartByUserId(c echo.Context) error {
	userId := c.Get("userId").(int)

	res, err := ca.cartSrv.GetCartByUserId(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
