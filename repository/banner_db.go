package repository

import "gorm.io/gorm"

type bannerRepositoryDb struct {
	gorm *gorm.DB
}

func NewBannerRepositoryDb(gorm *gorm.DB) BannerRepository {
	return bannerRepositoryDb{gorm: gorm}
}

func (b bannerRepositoryDb) GetAll() ([]Banner, error) {
	banner := []Banner{}
	tx := b.gorm.Order("id").Find(&banner)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return banner, nil
}
