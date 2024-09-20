package handler

import (
	"meow-meow/service"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type profileHandler struct {
	proSrv service.ProfileService
}

func NewProfileHandler(proSrv service.ProfileService) profileHandler {
	return profileHandler{proSrv: proSrv}
}

func (p profileHandler) GetProfileById(c echo.Context) error {
	id, err := getUserIdFromToken(c)
	if err != nil {
		return err
	}

	profileRes, err := p.proSrv.GetProfileByUserId(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, profileRes)
}

func getUserIdFromToken(c echo.Context) (int, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"].(float64)

	if !ok {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	id := int(userId)
	return id, nil
}
