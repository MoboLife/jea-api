package models

type Filters []ModelFilter

func Filter(query string, field string, multiple bool, filterType ModelFilterType) ModelFilter {
	return ModelFilter{
		Query:    query,
		Field:    field,
		Multiple: multiple,
		Type:     filterType,
	}
}

var CompanyFilter = ModelFilter{
	Query:    "company",
	Field:    "company_id",
	Multiple: false,
	Type:     Integer,
}

var ClientFilter = ModelFilter{
	Query:    "client",
	Field:    "client_id",
	Multiple: false,
	Type:     Integer,
}

var DescriptionFilter = ModelFilter{
	Query:    "search",
	Field:    "description",
	Multiple: false,
	Type:     String,
}

var ValidationFilter = ModelFilter{
	Query:    "validation",
	Field:    "validation_date",
	Multiple: true,
	Type:     Date,
}

var CreatedFilter = ModelFilter{
	Query:    "createdAt",
	Field:    "created_at",
	Multiple: true,
	Type:     Date,
}