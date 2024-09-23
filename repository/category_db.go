package repository

import "gorm.io/gorm"

type categoryRepositoryDb struct {
	gorm *gorm.DB
}

func NewCategoryRepositoryDb(gorm *gorm.DB) CategoryRepository {
	return categoryRepositoryDb{gorm: gorm}
}

func (c categoryRepositoryDb) GetAll() ([]Category, error) {
	category := []Category{}
	tx := c.gorm.Order("id").Find(&category)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return category, nil
}
