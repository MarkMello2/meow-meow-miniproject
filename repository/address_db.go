package repository

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type addressRepositoryDb struct {
	gorm *gorm.DB
}

func NewAddressRepositoryDb(gorm *gorm.DB) AddressRepository {
	return addressRepositoryDb{gorm: gorm}
}

func (a addressRepositoryDb) GetAddress(userId int) ([]Address, error) {
	address := []Address{}

	tx := a.gorm.Table("address").Where("user_id = ?", userId).Find(&address)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return address, nil
}

func (a addressRepositoryDb) SaveAddress(addressData Address) error {
	addressSave := Address{
		Id:        addressData.Id,
		FirstName: addressData.FirstName,
		LastName:  addressData.LastName,
		Mobile:    addressData.Mobile,
		Address:   addressData.Address,
		UserId:    addressData.UserId,
		Type:      addressData.Type,
	}

	tx := a.gorm.Table("address").Save(&addressSave)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (a addressRepositoryDb) UpdateAdress(addressData Address) error {
	tx := a.gorm.Table("address").Select("FirstName", "LastName", "Mobile", "Address", "UserId", "Type").Where("id = ?", addressData.Id).Updates(Address{
		FirstName: addressData.FirstName,
		LastName:  addressData.LastName,
		Mobile:    addressData.Mobile,
		Address:   addressData.Address,
		UserId:    addressData.UserId,
		Type:      addressData.Type,
	})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Address ID not found for update")
	}

	return nil
}

func (a addressRepositoryDb) DeleteAddress(id int) error {
	return nil
}
