package repository

import "gorm.io/gorm"

type profileRepositoryDb struct {
	gorm *gorm.DB
}

func NewProfileRepositoryDb(gorm *gorm.DB) ProfileRepository {
	return profileRepositoryDb{gorm: gorm}
}

func (p profileRepositoryDb) GetProfileById(userId int) (*Profile, error) {
	profile := Profile{}
	tx := p.gorm.Where("user_id = ?", userId).First(&profile)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &profile, nil
}

func (p profileRepositoryDb) SaveProfile(profile Profile) (string, error) {
	return "", nil
}
