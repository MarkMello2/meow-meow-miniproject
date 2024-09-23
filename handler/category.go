package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type categoryHandler struct {
	cateSrv service.CatagoryService
}

func NewCategoryHandler(cateSrv service.CatagoryService) categoryHandler {
	return categoryHandler{cateSrv: cateSrv}
}

func (cate categoryHandler) GetAllCategory(c echo.Context) error {
	res, err := cate.cateSrv.GetAllCategory()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
