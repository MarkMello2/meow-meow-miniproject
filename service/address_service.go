package service

import (
	"meow-meow/repository"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
)

type addressService struct {
	addRepo repository.AddressRepository
}

func NewAddressService(addRepo repository.AddressRepository) AddressService {
	return addressService{addRepo: addRepo}
}

func (a addressService) GetAddressByUserId(int) ([]AddressResponse, error) {
	return nil, nil
}

func (a addressService) CreateAddress(addressRequest AddressRequest, userId int, idAddress int) error {
	if len(strings.TrimSpace(addressRequest.FirstName)) == 0 || len(strings.TrimSpace(addressRequest.LastName)) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "FirstName and LastName is require")
	}

	if utf8.RuneCountInString(addressRequest.Mobile) > 10 {
		return echo.NewHTTPError(http.StatusBadRequest, "Phone number is too long. It should not exceed 10 digits.")
	}

	if addressRequest.Type != 0 && addressRequest.Type != 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Type is require and should 0 or 1")
	}

	addressData := repository.Address{
		Id:        idAddress,
		FirstName: addressRequest.FirstName,
		LastName:  addressRequest.LastName,
		Mobile:    addressRequest.Mobile,
		Address:   addressRequest.Address,
		UserId:    userId,
		Type:      addressRequest.Type,
	}

	if idAddress != 0 {
		err := a.addRepo.UpdateAdress(addressData)
		if err != nil {
			if err.(*echo.HTTPError).Code == http.StatusNotFound {
				return err
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
		}
		return nil
	}

	err := a.addRepo.SaveAddress(addressData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return nil
}

func (a addressService) DeleteAddressById(int) error {
	return nil
}
