package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type bannerHandler struct {
	ban service.BannerService
}

func NewBannerHandler(ban service.BannerService) bannerHandler {
	return bannerHandler{ban: ban}
}

func (b bannerHandler) GetBannerAll(c echo.Context) error {
	res, err := b.ban.GetBannerAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
