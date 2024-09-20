package repository

import "time"

type Profile struct {
	Id        int        `gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string     `gorm:"column:first_name"`
	LastName  string     `gorm:"column:last_name"`
	Mobile    string     `gorm:"column:mobile"`
	Sex       string     `gorm:"column:sex"`
	Status    string     `gorm:"column:status"`
	Image     string     `gorm:"column:image"`
	UserId    int        `gorm:"<-:create;column:user_id"`
	CreatedAt time.Time  `gorm:"<-:create;column:created_at;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"<-:update;column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

type ProfileRepository interface {
	GetProfileById(int) (*Profile, error)
	CreateProfile(Profile) error
}
