package models

// Group model
type Group struct {
	Model
	Name       string `json:"name"`
	Permission int64  `json:"permission"`
}

func (g* Group) GetFilters() Filters {
	return Filters{
		CreatedFilter,
		Filter("name", "name", false, String),
		Filter("permission", "permission", false, Integer),
	}
}
