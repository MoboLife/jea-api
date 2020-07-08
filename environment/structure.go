package environment

import (
	"jea-api/models"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// StructureConfig configuration of database structure
type StructureConfig struct {
	StructureType StructureType
	Models        []interface{}
	Private       bool
	Tenancy       bool
}

// StructureType type of structure
type StructureType string

var (
	// ERP Structure type ERP
	ERP StructureType = "ERP"
	// Manager Struct type Manager
	Manager StructureType = "MANAGER"
)

var structures = map[StructureType]StructureConfig{
	Manager: {
		StructureType: Manager,
		Models: []interface{}{
			&models.Group{},
			&models.User{},
			&models.Session{},
			&models.SessionAccess{},
			&models.Client{},
			&models.Environment{},
		},
		Private: true,
		Tenancy: false,
	},
	ERP: {
		StructureType: ERP,
		Models: []interface{}{
			&models.Client{},
			&models.Employer{},
			&models.Company{},
			&models.SaleProduct{},
			&models.Group{},
			&models.ProductGroup{},
			&models.Product{},
			&models.ProductStock{},
			&models.ProductStockTransfer{},
			&models.PurchaseProduct{},
			&models.Purchase{},
			&models.Sale{},
			&models.User{},
			&models.Session{},
			&models.SessionAccess{},
			&models.AccountPayable{},
			&models.AccountReceivable{},
			&models.CarRental{},
			&models.CargoMap{},
		},
		Private: false,
		Tenancy: true,
	},
}

// MigrateTables method for create non exists table
func MigrateTables(db *gorm.DB, tables ...interface{}) error {
	for _, table := range tables {
		if !db.HasTable(table) {
			err := db.Create(table).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// GetStructure get structure from type
func GetStructure(structureType StructureType) StructureConfig {
	return structures[structureType]
}

// SetupDatabase setup database structure
func SetupDatabase(db *gorm.DB) {
	err := db.AutoMigrate(GetStructure(Manager).Models...).Error
	if err != nil {
		logrus.Error("Error in Migrate Models. Error: ", err.Error())
	}
}
