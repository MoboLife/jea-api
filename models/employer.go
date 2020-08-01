package models

// Employer model
type Employer struct {
	Model
	Name 		string 		`json:"name"`
	CPF			string		`json:"cpf"`
	RG			string		`json:"rg"`
	Type		int			`json:"type"`
	Workload	int			`json:"workload"`
	Address
	Sales		[]*Sale		`json:"sales,omitempty" gorm:"foreignkey:SellerID;association_foreignkey:ID;association_autocreate:false;association_autoupdate:false"`
	CompanyID	int64		`json:"-"`
	Company		*Company	`json:"company,omitempty" gorm:"foreignkey:CompanyID;association_foreignkey:ID;association_autocreate:false;association_autoupdate:false"`
}

func (e *Employer) GetFilters() Filters {
	return Filters {
		CreatedFilter,
		Filter("name", "name", false, String),
	}
}
