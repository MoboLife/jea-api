package models

import "github.com/jinzhu/gorm"

type Sale struct {
	Model
	Purchaser		*Client			`json:"purchaser,omitempty" gorm:"foreignkey:PurchaserId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	PurchaserId		int32			`json:"-"`
	Status			int				`json:"status"`
	Products		[]*SaleProduct	`json:"products,omitempty" gorm:"foreignkey:SaleId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	CompanyId		int64			`json:"-"`
	Company			*Company		`json:"company,omitempty" gorm:"foreignkey:CompanyId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	SellerId		int64			`json:"-"`
	Seller			*Employer		`json:"seller,omitempty" gorm:"foreignkey:SellerId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Total			float32			`json:"total"`
	Discount		float32			`json:"discount"`
	Increase		float32			`json:"increase"`
	PaymentType		int				`json:"paymentType"`
}

type SaleProduct struct {
	Model
	SaleId		int64		`json:"-"`
	Sale		*Sale		`json:"sale,omitempty" gorm:"foreignkey:SaleId;association_foreignkey:ID;association_autoupdate:false;assoaciation_autocreate:false"`
	ProductId	int64		`json:"-"`
	Product		*Product	`json:"product,omitempty" gorm:"foreignkey:ProductId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Quantity	int			`json:"quantity"`
	Increase	float64		`json:"increase"`
	Discount	float64		`json:"discount"`
}

func (s *Sale) Setup(db *gorm.DB) {
}
