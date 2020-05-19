package controller

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"jea-api/environment"
)

type EnvironmentController interface {
	Create(eid string) error
	Delete(eid string) error
	Exists(eid string) bool
}

type EnvironmentControllerContext struct {
	DB	*gorm.DB
}

func (e *EnvironmentControllerContext) Create(eid string) error {
	if e.Exists(eid) {
		return errors.New("this schema already exists")
	}
	database := environment.UseEnvironment(eid, e.DB)
	err := database.Exec(fmt.Sprintf("CREATE SCHEMA %s;", eid)).Error
	if err != nil {
		return err
	}
	err = database.AutoMigrate(environment.GetStructure(environment.ERP).Models...).Error
	if err != nil{
		return err
	}
	return nil
}

func (e *EnvironmentControllerContext) Delete(eid string) error {
	return e.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec(fmt.Sprintf("drop schema %s cascade;", eid)).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (e *EnvironmentControllerContext) Exists(eid string) bool {
	rows, err := e.DB.Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name = ?;", eid).Rows()
	if err != nil {
		return false
	}
	return rows.Next()
}

func NewEnvironmentController(db *gorm.DB) EnvironmentController {
	return &EnvironmentControllerContext{DB: db}
}