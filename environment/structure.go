package environment

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"jea-api/models"
)

type StructureConfig struct {
	StructureType	StructureType
	Models			[]interface{}
	Private			bool
	Tenancy			bool
}

type StructureType string

var (
	ERP StructureType = "ERP"
	Manager StructureType = "MANAGER"
)

var structures = map[StructureType]StructureConfig{
	Manager: {
		StructureType: Manager,
		Models:        []interface{}{
			&models.Group{},
			&models.User{},
			&models.Session{},
			&models.SessionAccess{},
			&models.Client{},
			&models.Environment{},
		},
		Private:       true,
		Tenancy:       false,
	},
	ERP: {
		StructureType: ERP,
		Models:        []interface{}{
			&models.Client{},
			&models.Employer{},
			&models.Company{},
			&models.SaleProduct{},
			&models.Group{},
			&models.Product{},
			&models.Sale{},
			&models.User{},
			&models.Session{},
			&models.SessionAccess{},
			&models.AccountPayable{},
			&models.AccountReceivable{},
		},
		Private:       false,
		Tenancy:       true,
	},
}

func MigrateTables(db *gorm.DB, tables ...interface{}) error {
	for _, table := range tables{
		if !db.HasTable(table) {
			err := db.Create(table).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetStructure(structureType StructureType) StructureConfig {
	return structures[structureType]
}

func SetupDatabase(db *gorm.DB) {
	err := db.AutoMigrate(GetStructure(Manager).Models...).Error
	if err != nil {
		logrus.Error("Error in Migrate Models. Error: ", err.Error())
	}
}

