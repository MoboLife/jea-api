package models

import "time"

type CarRental struct {
	Model
	Client			*Client		`json:"client,omitempty" gorm:"foreignkey:ClientID;association_foreignkey:ID;association_autocreate:false;association_autoupdate:false"`
	ClientID		int64		`json:"-"`
	VehiclePlate	string		`json:"plate"`
	Days			int			`json:"days"`
	DayValue		float64		`json:"dayValue"`
	Discount		float64		`json:"discount"`
	OutDate			*time.Time	`json:"outDate"`
	Km				float32		`json:"km"`
	Fuel			float32		`json:"fuel"`
	PaymentType		int			`json:"paymentType"`
	Total			float64		`json:"total"`
	Observations	string		`json:"observations"`
}

func (c *CarRental) GetFilters() Filters {
	return Filters{
		ClientFilter,
		CreatedFilter,
		Filter("paymentType", "payment_type", false, Integer),
	}
}