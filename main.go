package main

import (
	"github.com/gin-gonic/gin"
	"jea-api/api"
	"jea-api/common"
	"jea-api/database"
	"jea-api/environment"
	"log"
)

func main() {
	var engine = gin.Default()
	var environmentManager = LoadEnvironment()
	var db = database.NewDatabase(database.ConnectionInfo{
		Host:     environmentManager.DatabaseHost,
		Port:     environmentManager.DatabasePort,
		User:     environmentManager.DatabaseUser,
		Database: environmentManager.DatabaseDatabase,
		Password: environmentManager.DatabasePassword,
	})
	engine.Use(common.CORS())
	engine.Use(database.UseDatabase(db))
	environment.SetupDatabase(db)
	api.NewAPI(engine)
	err := engine.Run()
	if err != nil {
		log.Println("Error in start Server. Error: ", err.Error())
	}
}
