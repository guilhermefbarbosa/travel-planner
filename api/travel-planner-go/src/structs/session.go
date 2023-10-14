package structs

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	Metadata
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid REFERENCES users(id)"`
	StartsAt time.Time `json:"starts_at" gorm:"type:timestamp"`
	EndsAt   time.Time `json:"ends_at" gorm:"type:timestamp"`
}
