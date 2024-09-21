package main

import (
	"log"
	"meow-meow/handler"
	"meow-meow/repository"
	"meow-meow/service"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Email    string
	Password string
}

func main() {
	e := echo.New()

	db := initDatabase()

	userRepositoryDb := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepositoryDb)
	userHandler := handler.NewUserHandler(userService)

	profileRepositoryDb := repository.NewProfileRepositoryDb(db)
	profileService := service.NewProfielService(profileRepositoryDb)
	profileHandler := handler.NewProfileHandler(profileService)

	e.POST("/user/register", userHandler.UserRegister)
	e.POST("/user/login", userHandler.UserLogin)

	r := e.Group("/")

	r.Use(echojwt.JWT([]byte("meow-meow")))
	r.Use(userIDMiddleware)

	r.GET("profile", profileHandler.GetProfileById)
	r.PATCH("profile", profileHandler.CreateUserProfile)

	err := e.Start(":8080")

	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initDatabase() *gorm.DB {

	dsn := "host=localhost user=user password=password dbname=shopping_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial)

	if err != nil {
		panic(err)
	}

	return db
}

func userIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
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
