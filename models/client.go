package models

import "github.com/jinzhu/gorm"

type Client struct {
	*Model
	Name		string		`json:"name"`
	Type		uint8		`json:"type"`
}

func (c *Client) Setup(db *gorm.DB) {

}
