package models

// Sale model
type Sale struct {
	Model
	Purchaser   *Client        `json:"purchaser,omitempty" gorm:"foreignkey:PurchaserID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	PurchaserID int32          `json:"-"`
	Status      int            `json:"status"`
	Products    []*SaleProduct `json:"products,omitempty" gorm:"foreignkey:SaleID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	CompanyID   int64          `json:"-"`
	Company     *Company       `json:"company,omitempty" gorm:"foreignkey:CompanyID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	SellerID    int64          `json:"-"`
	Seller      *Employer      `json:"seller,omitempty" gorm:"foreignkey:SellerID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Total       float32        `json:"total"`
	Discount    float32        `json:"discount"`
	Increase    float32        `json:"increase"`
	PaymentType int            `json:"paymentType"`
}

func (s *Sale) GetFilters() Filters {
	return []ModelFilter{
		CompanyFilter,
		CreatedFilter,
		Filter("purchaser", "purchaser_id", false, Integer),
		Filter("seller", "seller_id", false, Integer),
		Filter("type", "payment_type", false, Integer),
		Filter("status", "status", false, Integer),
	}
}

// SaleProduct model
type SaleProduct struct {
	Model
	SaleID    int64    `json:"-"`
	Sale      *Sale    `json:"sale,omitempty" gorm:"foreignkey:SaleID;association_foreignkey:ID;association_autoupdate:false;assoaciation_autocreate:false"`
	ProductID int64    `json:"-"`
	Product   *Product `json:"product,omitempty" gorm:"foreignkey:ProductID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	Quantity  int      `json:"quantity"`
	Increase  float64  `json:"increase"`
	Discount  float64  `json:"discount"`
}
