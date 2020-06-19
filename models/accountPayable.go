package models

import "time"

// AccountPayable model
type AccountPayable struct {
	Model
	Description    string    `json:"description"`
	ClientID       int64     `json:"-" `
	Client         *Client   `json:"client,omitempty" gorm:"foreignkey:ClientId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Amount         float64   `json:"amount"`
	CompanyID      int64     `json:"-"`
	Company        *Company  `json:"company,omitempty" gorm:"foreignkey:CompanyId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ValidationDate time.Time `json:"validationDate"`
}
