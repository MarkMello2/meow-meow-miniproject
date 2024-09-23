package handler

import (
	"meow-meow/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type addressHandler struct {
	addSrv service.AddressService
}

func NewAddressHandler(addSrv service.AddressService) addressHandler {
	return addressHandler{addSrv: addSrv}
}

func (a addressHandler) CreateAddress(c echo.Context) error {
	userId := c.Get("userId").(int)

	addressReq := service.AddressRequest{}
	err := c.Bind(&addressReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = a.addSrv.CreateAddress(addressReq, userId, 0)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Create Address Successful")
}

func (a addressHandler) UpdateAddressById(c echo.Context) error {
	userId := c.Get("userId").(int)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	addressReq := service.AddressRequest{}
	err = c.Bind(&addressReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = a.addSrv.CreateAddress(addressReq, userId, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Update Address Successful")
}

func (a addressHandler) GetAddressByUserId(c echo.Context) error {
	userId := c.Get("userId").(int)

	addressRes, err := a.addSrv.GetAddressByUserId(userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, addressRes)
}
