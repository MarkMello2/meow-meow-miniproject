package main

import (
	"log"
	"meow-meow/handler"
	"meow-meow/repository"
	"meow-meow/service"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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
	initEnv()

	e := echo.New()

	db := initDatabase()

	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	userRepositoryDb := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepositoryDb)
	userHandler := handler.NewUserHandler(userService)

	profileRepositoryDb := repository.NewProfileRepositoryDb(db)
	profileService := service.NewProfielService(profileRepositoryDb)
	profileHandler := handler.NewProfileHandler(profileService)

	addressRepositoryDb := repository.NewAddressRepositoryDb(db)
	addressService := service.NewAddressService(addressRepositoryDb)
	addressHandler := handler.NewAddressHandler(addressService)

	categoryRepositoryDb := repository.NewCategoryRepositoryDb(db)
	categoryService := service.NewCategoryService(categoryRepositoryDb)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepositoryDb := repository.NewProductRepositoryDb(db)
	productService := service.NewProductService(productRepositoryDb)
	productHandler := handler.NewProductHandler(productService)

	mallRepositoryDb := repository.NewMallRepositoryDb(db)
	mallService := service.NewMallService(mallRepositoryDb)
	mallHandler := handler.NewMallHandler(mallService)

	bannerRepositoryDb := repository.NewBannerRepositoryDb(db)
	bannerService := service.NewBannerService(bannerRepositoryDb)
	bannerHandler := handler.NewBannerHandler(bannerService)

	e.Static("/static", "assets")

	e.POST("/user/register", userHandler.UserRegister)
	e.POST("/user/login", userHandler.UserLogin)

	e.GET("/product", productHandler.GetAllProduct)
	e.GET("/product/:id", productHandler.GetProductById)
	e.GET("/product/category/:id", productHandler.GetProductByCategoryId)
	e.GET("/product/category", categoryHandler.GetAllCategory)

	e.GET("/product/shopping-mall", mallHandler.GetAllShoppingMall)
	e.GET("/product/shopping-mall/:id", productHandler.GetProductByMallId)

	e.GET("/product/banner", bannerHandler.GetBannerAll)

	e.GET("/product/recommended", productHandler.GetProductRecommended)

	r := e.Group("/")

	r.Use(echojwt.JWT([]byte(jwtSecret)))
	r.Use(userIDMiddleware)

	r.GET("profile", profileHandler.GetProfileById)
	r.PATCH("profile", profileHandler.CreateUserProfile)

	r.GET("address", addressHandler.GetAddressByUserId)
	r.POST("address", addressHandler.CreateAddress)
	r.PATCH("address/:id", addressHandler.UpdateAddressById)
	r.DELETE("address/:id", addressHandler.DeleteAddressById)

	err := e.Start(":8080")

	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initDatabase() *gorm.DB {

	dsn := "host=localhost user=user password=password dbname=shopping_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"
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

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
