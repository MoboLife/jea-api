package models

// Client model
type Client struct {
	*Model
	Name string `json:"name"`
	Type uint8  `json:"type"`
}
