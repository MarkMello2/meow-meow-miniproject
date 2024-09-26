package repository

import "gorm.io/gorm"

type orderRepositoryDb struct {
	gorm *gorm.DB
}

func NewOrderRepositoryDb(gorm *gorm.DB) OrderRepository {
	return orderRepositoryDb{gorm: gorm}
}

func (o orderRepositoryDb) Save(orderData []Order) error {
	err := o.gorm.Create(&orderData).Error

	return err
}

func (o orderRepositoryDb) GetOrder(int) ([]OrderGet, error) {
	return nil, nil
}
