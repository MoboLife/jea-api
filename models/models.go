package models

import (
	"time"
)

// Model basic model
type Model struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
