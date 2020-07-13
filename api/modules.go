package api

import "jea-api/models"

// Modules list with all API and routers
type Modules struct {
	Company 	models.Company 	`router:"/companies"`
	Client  	models.Client  	`router:"/clients"`
	User    	models.User    	`router:"/users"`
	Group   	models.Group   	`router:"/groups"`
}
