package service

import (
	"errors"
	"meow-meow/repository"
	"net/http"
	"strings"
	"unicode/utf8"

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

func (p profileService) CreateUserProfile(profileReq ProfileRequest, userId int, isInitPro bool) error {
	if (len(strings.TrimSpace(profileReq.FirstName)) == 0 || len(strings.TrimSpace(profileReq.LastName)) == 0) && !isInitPro {
		return echo.NewHTTPError(http.StatusBadRequest, "FirstName and LastName is require")
	}

	if utf8.RuneCountInString(profileReq.Mobile) > 10 && !isInitPro {
		return echo.NewHTTPError(http.StatusBadRequest, "Phone number is too long. It should not exceed 10 digits.")
	}

	if profileReq.Sex != "M" && profileReq.Sex != "F" && !isInitPro {
		return echo.NewHTTPError(http.StatusBadRequest, "Sex is require and should M or F")
	}

	proFileDataId, err := p.GetProfileByUserId(userId)
	if err != nil {
		return err
	}

	profileData := repository.Profile{
		Id:        proFileDataId.Id,
		FirstName: profileReq.FirstName,
		LastName:  profileReq.LastName,
		Mobile:    profileReq.Mobile,
		Sex:       profileReq.Sex,
		Status:    "A",
		UserId:    userId,
	}

	err = p.proRepo.CreateProfile(profileData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}
