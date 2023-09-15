package models

import (
	"time"
	"github.com/google/uuid"
)

type Todo struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Title string `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Done bool `gorm:"default:false"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
