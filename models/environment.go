package models

// Environment model
type Environment struct {
	Model
	EID           string  `json:"eid" gorm:"column:eid"`
	Client        *Client `json:"client,omitempty" gorm:"foreignkey:PurchaserID;association_foreignkey:ID"`
	ClientID      int64   `json:"-"`
	StructureType string  `json:"structure_type"`
	Created       bool    `json:"created"`
}
