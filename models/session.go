package models

import "time"

type SessionAccess struct {
	Id			int64		`json:"id" gorm:"primary_key"`
	IpAddress	string		`json:"ipAddress"`
	SessionId	int64		`json:"-"`
	AccessAt	time.Time	`json:"accessAt"`
}

type SessionType string

var (
	MobileSession	SessionType = "MOBILE"
	WebSession		SessionType = "WEB"
)

type Session struct {
	Id				int64				`json:"id" gorm:"primary_key"`
	User			*User				`json:"user,omitempty" gorm:"foreignkey:UserId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	UserId			int64				`json:"-"`
	Access			[]*SessionAccess	`json:"access" gorm:"foreignkey:SessionId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	Token			string				`json:"-"`
	Model			string				`json:"model"`
	DeviceId		string				`json:"deviceId"`
	Platform		string				`json:"platform"`
	Version			string				`json:"version"`
	Type			SessionType			`json:"type"`
}
