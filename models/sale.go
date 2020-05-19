package models

import "github.com/jinzhu/gorm"

type Sale struct {
	Model
	Purchaser	*Client		`json:"purchaser,omitempty" gorm:"foreignkey:PurchaserId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	PurchaserId	int32		`json:"-"`
	Products	[]*Product	`json:"products,omitempty" gorm:"many2many:sale_products;association_autoupdate:false;association_autocreate:false"`
	Total		float32		`json:"total"`
	Discount	float32		`json:"discount"`
	Increase	float32		`json:"increase"`
}

func (s *Sale) Setup(db *gorm.DB) {
}
