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

func (o orderRepositoryDb) GetOrder(userId int) ([]OrderGet, error) {
	sql := "select o.id, o.price, o.quantity, o.created_at, p.code, p.name, p.description, p.rating, p.image, o.user_id from orders as o inner join products p on o.product_id = p.id where o.user_id = ? and o.deleted_at is null"
	condition := o.gorm.Raw(sql, userId)
	result := []OrderGet{}
	err := condition.Find(&result).Error

	return result, err
}
