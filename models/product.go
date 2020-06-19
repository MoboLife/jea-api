package models

// Product model
type Product struct {
	Model
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}
