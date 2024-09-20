package service

import (
	"errors"
	"meow-meow/repository"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type profileService struct {
	proRepo repository.ProfileRepository
}

func NewProfielService(proRepo repository.ProfileRepository) ProfileService {
	return profileService{proRepo: proRepo}
}

func (p profileService) GetProfileByUserId(userId int) (*ProfileResponse, error) {

	if userId == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "User id is require")
	}

	proRawData, err := p.proRepo.GetProfileById(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &ProfileResponse{}, nil
		}
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	resProfile := ProfileResponse{
		Id:        proRawData.Id,
		FirstName: proRawData.FirstName,
		LastName:  proRawData.LastName,
		Mobile:    proRawData.Mobile,
		Sex:       proRawData.Sex,
		Status:    proRawData.Status,
		Image:     proRawData.Image,
		UserId:    proRawData.UserId,
	}

	return &resProfile, nil
}

func (p profileService) SaveProfileByUserId(profileReq ProfileRequest) (string, error) {
	return "", nil
}
