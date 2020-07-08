package models

import "time"

type CargoMap struct {
	Model
	OutDate			*time.Time		`json:"outDate"`
	Driver			*Employer		`json:"driver,omitempty" gorm:"foreignkey:DriverID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	DriverID		int64			`json:"-"`
	Purchase		*Purchase		`json:"purchase,omitempty" gorm:"foreignkey:PurchaseID;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	PurchaseID		int64			`json:"-"`
	VehiclePlate	string			`json:"plate"`
	Observations	string			`json:"observations"`
}

func (c *CargoMap) GetFilters() Filters {
	return Filters {
		CreatedFilter,
		Filter("purchase", "purchase_id", false, Integer),
		Filter("driver", "driver_id", false, Integer),
	}
}
