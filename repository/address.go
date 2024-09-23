package repository

import "time"

type Address struct {
	Id        int        `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	Mobile    string     `gorm:"column:mobile"`
	Address   string     `gorm:"column:address"`
	Type      int        `gorm:"column:type"`
	UserId    int        `gorm:"<-:create;column:user_id"`
	CreatedAt time.Time  `gorm:"<-:create;column:created_at;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"<-:update;column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

type AddressRepository interface {
	GetAddress(int) ([]Address, error)
	SaveAddress(Address) error
	UpdateAdress(Address) error
	DeleteAddress(int) error
}
