package main

import (
	"jea-api/api"
	"jea-api/common"
	"jea-api/database"
	"jea-api/environment"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	var engine = gin.Default()
	var environmentManager = LoadEnvironment()
	db, err := database.NewDatabase(database.ConnectionInfo{
		Host:     environmentManager.DatabaseHost,
		Port:     environmentManager.DatabasePort,
		User:     environmentManager.DatabaseUser,
		Database: environmentManager.DatabaseDatabase,
		Password: environmentManager.DatabasePassword,
		URL:	  environmentManager.DatabaseURL,
		Driver:   "postgres",
	}, true)
	if err != nil {
		logrus.WithField("error", err.Error()).Panic("Error in connect with database")
		return
	}
	engine.Use(common.CORS())
	engine.Use(database.UseDatabase(db))
	environment.SetupDatabase(db)
	api.NewAPI(engine)
	err = engine.Run()
	if err != nil {
		logrus.Error("Error in start Server. Error: ", err.Error())
	}
}
