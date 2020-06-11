package models

type Client struct {
	*Model
	Name		string		`json:"name"`
	Type		uint8		`json:"type"`
}
