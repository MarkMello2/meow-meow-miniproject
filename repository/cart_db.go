package repository

import "gorm.io/gorm"

type cartRepositoryDb struct {
	gorm *gorm.DB
}

func NewCartRepositoryDb(gorm *gorm.DB) CartRepository {
	return cartRepositoryDb{gorm: gorm}
}

func (c cartRepositoryDb) Save(cartData []Cart) error {
	err := c.gorm.Create(&cartData).Error

	return err
}

func (c cartRepositoryDb) GetCart(userId int) ([]CartGet, error) {
	sql := "select s.id, s.price, s.quantity, s.created_at, p.code, p.name, p.description, p.rating, p.image, s.user_id from carts as s inner join products p on s.product_id = p.id where s.user_id = ? and s.deleted_at is null"
	condition := c.gorm.Raw(sql, userId)
	result := []CartGet{}
	err := condition.Find(&result).Error

	return result, err
}
