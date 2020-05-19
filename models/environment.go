package models

type Environment struct {
	Model
	EID					string		`json:"eid" gorm:"column:eid"`
	Client				*Client		`json:"client,omitempty" gorm:"foreignkey:PurchaserId;association_foreignkey:ID"`
	ClientId			int64		`json:"-"`
	StructureType		string		`json:"structure_type"`
	Created				bool		`json:"created"`
}
