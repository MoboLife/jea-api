package models

import (
	"time"
)

type Model struct {
	Id			int64		`json:"id" gorm:"primary_key"`
	CreatedAt	time.Time	`json:"createdAt"`
	UpdatedAt	*time.Time	`json:"updatedAt,omitempty"`
}
