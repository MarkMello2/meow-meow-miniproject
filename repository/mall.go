package repository

import (
	"time"

	"gorm.io/gorm"
)

type Mall struct {
	Id          int            `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	Image       string         `gorm:"column:image"`
	CreatedAt   time.Time      `gorm:"<-:create;column:created_at;default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"<-:update;column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

type MallRepository interface {
	GetAll() ([]Mall, error)
}
