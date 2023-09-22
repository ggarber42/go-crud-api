package models

import (
	"time"
)

type Todo struct {
	ID  uint     `json:"id" gorm:"primaryKey"`
	Title string `gorm:"uniqueIndex;not null" json:"title,omitempty"`
	Done bool `gorm:"default:false"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}
