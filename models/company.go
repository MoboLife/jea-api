package models

type Company struct {
	Model
	Name	string		`json:"name"`
	CNPJ	string		`json:"cnpj"`
	City	string		`json:"city"`
}
