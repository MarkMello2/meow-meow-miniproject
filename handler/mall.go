package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type mallHandler struct {
	mallSrv service.MallService
}

func NewMallHandler(mallSrv service.MallService) mallHandler {
	return mallHandler{mallSrv: mallSrv}
}

func (m mallHandler) GetAllShoppingMall(c echo.Context) error {
	res, err := m.mallSrv.GetAllShoppingMall()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
