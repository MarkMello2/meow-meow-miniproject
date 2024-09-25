package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type favoriteHandler struct {
	favSrv service.FavoriteService
}

func NewFavoriteHandler(favSrv service.FavoriteService) favoriteHandler {
	return favoriteHandler{favSrv: favSrv}
}

func (f favoriteHandler) SaveFavorite(c echo.Context) error {
	userId := c.Get("userId").(int)

	favReq := service.FavoriteRequest{}
	err := c.Bind(&favReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = f.favSrv.SaveFavorite(favReq, userId)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Create Favorite Successful")
}

func (f favoriteHandler) GetFavoriteByUserId(c echo.Context) error {
	userId := c.Get("userId").(int)

	res, err := f.favSrv.GetFavoriteByUserId(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
