package repository

import "gorm.io/gorm"

type mallRepositoryDb struct {
	gorm *gorm.DB
}

func NewMallRepositoryDb(gorm *gorm.DB) MallRepository {
	return mallRepositoryDb{gorm: gorm}
}

func (m mallRepositoryDb) GetAll() ([]Mall, error) {
	mall := []Mall{}
	tx := m.gorm.Order("id").Find(&mall)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return mall, nil
}
