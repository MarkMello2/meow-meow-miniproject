package service

import (
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type favoriteService struct {
	favRepo  repository.FavoriteRepository
	minioSrv MinioService
}

func NewFavoriteService(favRepo repository.FavoriteRepository, minioSrv MinioService) FavoriteService {
	return favoriteService{favRepo: favRepo, minioSrv: minioSrv}
}

func (f favoriteService) GetFavoriteByUserId(userId int) ([]FavoriteResponse, error) {
	favData, err := f.favRepo.GetById(userId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	res := []FavoriteResponse{}

	for _, v := range favData {
		newUrlImg, err := f.minioSrv.getUrlImagePath(v.ProductImage)

		if err != nil {
			return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}

		res = append(res, FavoriteResponse{
			Id:                 v.Id,
			Price:              v.Price,
			Quantity:           v.Quantity,
			FavoriteDate:       v.FavoriteDate,
			ProductCode:        v.ProductCode,
			ProductName:        v.ProductName,
			ProductDescription: v.ProductDescription,
			ProductRating:      v.ProductRating,
			ProductImage:       newUrlImg,
			UserId:             v.UserId,
		})
	}
	return res, nil
}

func (f favoriteService) SaveFavorite(favReq FavoriteRequest, userId int) error {

	if favReq.ProductId == 0 || favReq.Price == 0 || favReq.Quantity == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "ProductId, Price, and Quantity is require")
	}

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

func (f favoriteService) DeleteFavoriteById(id int) error {
	err := f.favRepo.DeleteById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}
