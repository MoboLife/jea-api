package main

import (
	"github.com/VitorEmanoel/menv"
	"log"
)

type Environment struct {
	DatabaseHost		string		`name:"DATABASE_HOST"`
	DatabasePort		string		`name:"DATABASE_PORT"`
	DatabaseUser		string		`name:"DATABASE_USER"`
	DatabasePassword	string		`name:"DATABASE_PASSWORD"`
	DatabaseDatabase	string		`name:"DATABASE_DATABASE"`
}

func LoadEnvironment() Environment {
	var environment = Environment{}
	err := menv.LoadEnvironment(&environment)
	if err != nil {
		log.Fatalln("Error in load environment variables. Error: ", err.Error())
	}
	return environment
}
