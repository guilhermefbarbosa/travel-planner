package structs

import "github.com/google/uuid"

type Travel struct {
	Metadata
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid REFERENCES users(id)"`
	StartsAt Date      `json:"starts_at" gorm:"type:date"`
	EndsAt   Date      `json:"ends_at" gorm:"type:date"`
	Budget   Money     `json:"budget" gorm:"type:text"`
}

func (t Travel) TableName() string {
	return "travels"
}
