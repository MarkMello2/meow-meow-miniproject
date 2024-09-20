package repository

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type usersRepositoryDb struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) UsersRepository {
	return usersRepositoryDb{gorm: gorm}
}

func (u usersRepositoryDb) CreateUser(userData User) (string, error) {

	code := uuid.New().String()

	users := User{
		Code:      code,
		Email:     userData.Email,
		Password:  userData.Password,
		CreatedAt: time.Now(),
	}

	tx := u.gorm.Create(&users)

	if tx.Error != nil {
		return "", tx.Error
	}

	return "create user successfully", nil
}

func (u usersRepositoryDb) GetUserByName(email string) ([]User, error) {

	users := []User{}
	tx := u.gorm.Where("email = ?", email).First(&users)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}
