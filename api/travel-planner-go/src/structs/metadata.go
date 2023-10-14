package structs

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Metadata struct {
	ID        uuid.UUID      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
