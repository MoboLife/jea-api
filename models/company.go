package models

// Company model
type Company struct {
	Model
	SocialName	string		`json:"socialName"`
	Name 		string 		`json:"name"`
	CNPJ 		string 		`json:"cnpj"`
	Phone		string		`json:"phone"`
	Celular		string		`json:"celular"`
	Email		string		`json:"email"`
	Address
}

func (c *Company) GetFilters() Filters {
	return Filters{
		CreatedFilter,
		Filter("name", "name", false, String),
		Filter("cnpj", "cnpj", false, String),
	}
}
