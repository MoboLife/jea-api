package controller

import (
	"errors"
	"fmt"
	"jea-api/environment"

	"github.com/jinzhu/gorm"
)

// EnvironmentController controller for environments
type EnvironmentController interface {
	Create(eid string) error
	Delete(eid string) error
	Update(eid string) error
	Exists(eid string) bool
}

// EnvironmentControllerContext context of Controller
type EnvironmentControllerContext struct {
	DB *gorm.DB
}

// Create create environment method
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
	if err != nil {
		return err
	}
	return nil
}

// Delete delete environment
func (e *EnvironmentControllerContext) Delete(eid string) error {
	return e.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec(fmt.Sprintf("drop schema %s cascade;", eid)).Error
		if err != nil {
			return err
		}
		return nil
	})
}

// Update update environment
func (e *EnvironmentControllerContext) Update(eid string) error {
	return e.DB.Transaction(func(tx *gorm.DB) error {
		if !e.Exists(eid) {
			return errors.New("this schema not exists")
		}
		database := environment.UseEnvironment(eid, e.DB)
		environment.MigrateTables(database, environment.GetStructure(environment.ERP).Models...)
		return nil
	})
}

// Exists check if environment exists
func (e *EnvironmentControllerContext) Exists(eid string) bool {
	rows, err := e.DB.Raw("SELECT schema_name FROM information_schema.schemata WHERE schema_name = ?;", eid).Rows()
	if err != nil {
		return false
	}
	return rows.Next()
}

// NewEnvironmentController create environment controller
func NewEnvironmentController(db *gorm.DB) EnvironmentController {
	return &EnvironmentControllerContext{DB: db}
}
