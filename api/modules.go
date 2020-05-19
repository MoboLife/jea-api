package api

import "jea-api/models"

type Modules struct {
	Company		models.Company		`router:"/companies"`
	Product		models.Product		`router:"/products"`
	Client		models.Client		`router:"/clients"`
	User		models.User			`router:"/users"`
	Group		models.Group		`router:"/groups"`
}
