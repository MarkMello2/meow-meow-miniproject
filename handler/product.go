package handler

import (
	"meow-meow/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	proSrv service.ProductService
}

func NewProductHandler(proSrv service.ProductService) productHandler {
	return productHandler{proSrv: proSrv}
}

func (p productHandler) GetProductByCategoryId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := p.proSrv.GetProductByCategoryId(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
