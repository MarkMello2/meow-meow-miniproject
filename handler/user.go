package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (u userHandler) UserRegister(c echo.Context) error {
	user := service.UserRequest{}
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := u.userSrv.CreateUser(user)

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, result)
}

func (u userHandler) UserLogin(c echo.Context) error {
	user := service.UserRequest{}
	err := c.Bind(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	result, err := u.userSrv.UserLogin(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
