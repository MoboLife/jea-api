package models

type Group struct {
	Model
	Name		string		`json:"name"`
	Permission	int64		`json:"permission"`
}