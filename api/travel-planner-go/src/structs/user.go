package structs

import (
	"time"
)

type User struct {
	Metadata
	Email     string     `gorm:"unique; not null"`
	PwdHash   string     `gorm:"not null"`
	LastLogin *time.Time `gorm:"default:null"`
	IsActive  bool       `gorm:"default:true"`
}

func (u User) TableName() string {
	return "users"
}
