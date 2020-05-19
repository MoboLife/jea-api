package models

import "github.com/jinzhu/gorm"

type Product struct {
	Model
	Description		string		`json:"description"`
	Price			float32		`json:"price"`
}

func (p *Product) Setup(db *gorm.DB) {

}