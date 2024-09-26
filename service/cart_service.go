package service

import (
	"meow-meow/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return cartService{cartRepo: cartRepo}
}

func (c cartService) SaveCart(cartReq []CartRequest, userId int) error {

	cartData := []repository.Cart{}

	for _, v := range cartReq {

		if v.ProductId == 0 || v.Price == 0 || v.Quantity == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "ProductId, Price, and Quantity is require")
		}

		cartData = append(cartData, repository.Cart{
			Price:     v.Price,
			Quantity:  v.Quantity,
			ProductId: v.ProductId,
			UserId:    userId,
		})
	}

	err := c.cartRepo.Save(cartData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}

func (c cartService) GetCartByUserId(userId int) ([]CartResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	cartData, err := c.cartRepo.GetCart(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []CartResponse{}

	for _, v := range cartData {
		res = append(res, CartResponse{
			Id:                 v.Id,
			Price:              v.Price,
			Quantity:           v.Quantity,
			CartDate:           v.CartDate,
			ProductCode:        v.ProductCode,
			ProductName:        v.ProductName,
			ProductDescription: v.ProductDescription,
			ProductRating:      v.ProductRating,
			ProductImage:       pathImg + v.ProductImage,
			UserId:             v.UserId,
		})
	}

	return res, nil
}
