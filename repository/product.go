package repository

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          int            `gorm:"column:id;primaryKey;autoIncrement"`
	Code        string         `gorm:"column:code"`
	Name        string         `gorm:"column:name"`
	Description string         `gorm:"column:description"`
	Price       float32        `gorm:"column:price"`
	Rating      int            `gorm:"column:rating"`
	Image       string         `gorm:"column:image"`
	CategoryId  int            `gorm:"column:category_id"`
	MallId      int            `gorm:"column:mall_id"`
	CreatedAt   time.Time      `gorm:"<-:create;column:created_at;default:current_timestamp"`
	UpdatedAt   time.Time      `gorm:"<-:update;column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(int) ([]Product, error)
	GetByCategoryId(int) ([]Product, error)
	GetByMallId(int) ([]Product, error)
}
