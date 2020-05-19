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
			&models.Company{},
			&models.Group{},
			&models.Product{},
			&models.Sale{},
			&models.User{},
		},
		Private:       false,
		Tenancy:       true,
	},
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

