package models

import "time"

type Address struct {
	Street       	string		`json:"street"`
	Number       	string		`json:"number"`
	Complement   	string		`json:"complement"`
	Neighborhood 	string		`json:"neighborhood"`
	PostalCode	 	int			`json:"postalCode"`
	UF				string		`json:"uf"`
	City			string		`json:"city"`
	Reference		string		`json:"reference"`
}

// Client model
type Client struct {
	*Model
	Name 		string 		`json:"name"`
	CPF			string		`json:"cpf"`
	RG			string		`json:"rg"`
	Birthday	time.Time	`json:"birthday"`
	CNPJ		string		`json:"cnpj"`
	Type 		uint8  		`json:"type"`
	Address
	Phone1		string		`json:"phone1"`
	Phone2		string		`json:"phone2"`
	Celular1	string		`json:"celular1"`
	Celular2	string		`json:"celular2"`
	Email1		string		`json:"email1"`
	Email2		string		`json:"email2"`
	FatherName	string		`json:"fatherName"`
	MotherName	string		`json:"motherName"`
	CivilState	int			`json:"civilState"`
	Schooling	int			`json:"schooling"`
	Observation	string		`json:"observation"`

}

func (c *Client) GetFilters() Filters {
	return Filters{
		CreatedFilter,
		Filter("type", "type", false, Integer),
	}
}