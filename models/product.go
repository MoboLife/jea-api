package models

// Product model
type Product struct {
	Model
	Description 	string  					`json:"description"`
	StockTransfers	[]*ProductStockTransfer		`json:"stockTransfers,omitempty" gorm:"foreignkey:ProductID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	Stock			[]*ProductStock				`json:"stock,omitempty" gorm:"foreignkey:ProductID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	Price       	float32 					`json:"price"`
	Profit			float32						`json:"profit"`
	ProductGroup	*ProductGroup				`json:"group,omitempty" gorm:"foreignkey:ProductGroupID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ProductGroupID	int64						`json:"-"`
}

type ProductStock struct {
	Model
	Company		*Company	`json:"company,omitempty" gorm:"foreignkey:CompanyID;association_foreignkey:ID;assoaction_autoupdate:false;association_autocreate:false"`
	CompanyID	int64		`json:"-"`
	ProductID	int64		`json:"-"`
	Quantity	float32		`json:"quantity"`
}

type ProductStockTransfer struct {
	Model
	From		*Company	`json:"from,omitempty" gorm:"foreignkey:CompanyID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	CompanyID	int64		`json:"-"`
	To			*Company	`json:"to,omitempty" gorm:"foreignkey:ToCompanyID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ToCompanyID	int64		`json:"-"`
	Product		*Product	`json:"product,omitempty" gorm:"foreignkey:ProductID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	ProductID	int64		`json:"-"`
	Quantity	int			`json:"quantity"`
}


type ProductGroup struct {
	Model
	Name	string 		`json:"name"`
}

func (p *Product) GetFilters() Filters {
	return Filters {
		CreatedFilter,
		DescriptionFilter,
		Filter("price", "price", false, Integer),
	}
}
