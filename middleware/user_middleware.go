package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func UserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userId, ok := claims["user_id"].(float64)

		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		c.Set("userId", int(userId))

		return next(c)
	}
}
