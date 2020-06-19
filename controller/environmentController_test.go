package controller

import (
	"jea-api/environment"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var connectionString = "host=localhost port=5432 user=postgres password=super dbname=jea sslmode=disable"

func TestEnvironmentControllerExists(t *testing.T) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		t.Error("Error in connect with database. Error:", err.Error())
	}
	var environmentController = NewEnvironmentController(db)
	if !environmentController.Exists("vitor") {
		t.Error("Failed in verify if schema 'vitor' exists.")
	}
}

func TestEnvironmentControllerCreate(t *testing.T) {
	gorm.DefaultTableNameHandler = environment.TableNameHandler
	db, err := gorm.Open("postgres", connectionString)
	db.LogMode(true)
	if err != nil {
		t.Error("Error in connect with database. Error: ", err.Error())
	}
	var environmentController = NewEnvironmentController(db)
	err = environmentController.Create("teste")
	if err != nil {
		t.Error("Error in create teste schema. Error:", err.Error())
	}
}

func TestEnvironmentControllerDelete(t *testing.T) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		t.Error("Error in connect with database. Error:", err.Error())
	}
	var environmentController = NewEnvironmentController(db)
	err = environmentController.Delete("teste")
	if err != nil {
		t.Error("Error in delete teste schema. Error:", err.Error())
	}
}
