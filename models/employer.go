package models

// Employer model
type Employer struct {
	Model
	Name 		string 		`json:"name"`
	CPF			string		`json:"cpf"`
	RG			string		`json:"rg"`
	Type		int			`json:"type"`
	Address
}

func (e *Employer) GetFilters() []ModelFilter {
	return Filters {
		CreatedFilter,
		Filter("name", "name", false, String),
	}
}
