package repository

import "time"

type User struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Code      string    `gorm:"column:code"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"<-:create;column:created_at;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"<-:update;column:updated_at"`
	DeletedAt time.Time `gorm:"<-:update;column:deleted_at"`
}

type UsersRepository interface {
	CreateUser(User) error
	GetUserByName(string) ([]User, error)
}
