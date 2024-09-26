package repository

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	Id        int            `gorm:"column:id;primaryKey;autoIncrement"`
	Price     float32        `gorm:"column:price"`
	Quantity  int            `gorm:"column:quantity"`
	ProductId int            `gorm:"column:product_id"`
	UserId    int            `gorm:"column:user_id"`
	CreatedAt time.Time      `gorm:"<-:create;column:created_at;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"<-:update;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

type CartGet struct {
	Id                 int     `gorm:"column:id;primaryKey;autoIncrement"`
	Price              float32 `gorm:"column:price"`
	Quantity           int     `gorm:"column:quantity"`
	CartDate           string  `gorm:"column:created_at"`
	ProductCode        string  `gorm:"column:code"`
	ProductName        string  `gorm:"column:name"`
	ProductDescription string  `gorm:"column:description"`
	ProductRating      int     `gorm:"column:rating"`
	ProductImage       string  `gorm:"column:image"`
	UserId             int     `gorm:"column:user_id"`
}

type CartRepository interface {
	Save([]Cart) error
	GetCart(int) ([]CartGet, error)
}
