package models

import "time"

// AccountReceivable model
type AccountReceivable struct {
	Model
	Description    string    `json:"description"`
	ClientID       int64     `json:"-"`
	Client         *Client   `json:"client,omitempty" gorm:"foreignkey:ClientID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Amount         float64   `json:"amount"`
	CompanyID      int64     `json:"-"`
	Company        *Company  `json:"company,omitempty" gorm:"foreignkey:CompanyID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ValidationDate time.Time `json:"validationDate"`
}
