package models

// Sale model
type Sale struct {
	Model
	Purchaser   *Client        `json:"purchaser,omitempty" gorm:"foreignkey:PurchaserId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	PurchaserID int32          `json:"-"`
	Status      int            `json:"status"`
	Products    []*SaleProduct `json:"products,omitempty" gorm:"foreignkey:SaleId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	CompanyID   int64          `json:"-"`
	Company     *Company       `json:"company,omitempty" gorm:"foreignkey:CompanyId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	SellerID    int64          `json:"-"`
	Seller      *Employer      `json:"seller,omitempty" gorm:"foreignkey:SellerId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Total       float32        `json:"total"`
	Discount    float32        `json:"discount"`
	Increase    float32        `json:"increase"`
	PaymentType int            `json:"paymentType"`
}

// SaleProduct model
type SaleProduct struct {
	Model
	SaleID    int64    `json:"-"`
	Sale      *Sale    `json:"sale,omitempty" gorm:"foreignkey:SaleId;association_foreignkey:ID;association_autoupdate:false;assoaciation_autocreate:false"`
	ProductID int64    `json:"-"`
	Product   *Product `json:"product,omitempty" gorm:"foreignkey:ProductId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Quantity  int      `json:"quantity"`
	Increase  float64  `json:"increase"`
	Discount  float64  `json:"discount"`
}
