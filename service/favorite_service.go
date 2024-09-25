package service

import (
	"meow-meow/repository"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type favoriteService struct {
	favRepo repository.FavoriteRepository
}

func NewFavoriteService(favRepo repository.FavoriteRepository) FavoriteService {
	return favoriteService{favRepo: favRepo}
}

func (f favoriteService) GetFavoriteByUserId(userId int) ([]FavoriteResponse, error) {
	pathImg := os.Getenv("IMG_PATH_LOCAL")

	favData, err := f.favRepo.GetById(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []FavoriteResponse{}

	for _, v := range favData {
		res = append(res, FavoriteResponse{
			Id:                 v.Id,
			Price:              v.Price,
			Quantity:           v.Quantity,
			FavoriteDate:       v.FavoriteDate,
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

func (f favoriteService) SaveFavorite(favReq FavoriteRequest, userId int) error {
	favData := repository.Favorite{
		ProductId: favReq.ProductId,
		Price:     favReq.Price,
		Quantity:  favReq.Quantity,
		UserId:    userId,
	}

	err := f.favRepo.Save(favData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}
