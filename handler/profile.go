package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type profileHandler struct {
	proSrv service.ProfileService
}

func NewProfileHandler(proSrv service.ProfileService) profileHandler {
	return profileHandler{proSrv: proSrv}
}

func (p profileHandler) GetProfileById(c echo.Context) error {
	id := c.Get("userId").(int)

	profileRes, err := p.proSrv.GetProfileByUserId(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profileRes)
}

func (p profileHandler) CreateUserProfile(c echo.Context) error {
	id := c.Get("userId").(int)

	profileReq := service.ProfileRequest{}
	err := c.Bind(&profileReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = p.proSrv.CreateUserProfile(profileReq, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "update success")
}
