package repository

import "gorm.io/gorm"

type productRepositoryDb struct {
	gorm *gorm.DB
}

func NewProductRepositoryDb(gorm *gorm.DB) ProductRepository {
	return productRepositoryDb{gorm: gorm}
}

func (p productRepositoryDb) GetAll() ([]Product, error) {
	return nil, nil
}

func (p productRepositoryDb) GetById(productId int) (*Product, error) {
	return nil, nil
}

func (p productRepositoryDb) GetByCategoryId(categoryId int) ([]Product, error) {
	product := []Product{}

	tx := p.gorm.Where("category_id = ?", categoryId).Order("id").Find(&product)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return product, nil
}
