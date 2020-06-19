package models

import "time"

// SessionAccess model
type SessionAccess struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	IPAddress string    `json:"ipAddress"`
	SessionID int64     `json:"-"`
	AccessAt  time.Time `json:"accessAt"`
}

// SessionType is type of session
type SessionType string

var (
	// MobileSession type Mobile of Session
	MobileSession SessionType = "MOBILE"
	// WebSession type Web of Session
	WebSession SessionType = "WEB"
)

// Session model
type Session struct {
	ID       int64            `json:"id" gorm:"primary_key"`
	User     *User            `json:"user,omitempty" gorm:"foreignkey:UserId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:false"`
	UserID   int64            `json:"-"`
	Access   []*SessionAccess `json:"access" gorm:"foreignkey:SessionId;association_foreignkey:ID;association_autoupdate:false;association_autocreate:true"`
	Token    string           `json:"-"`
	Model    string           `json:"model"`
	DeviceID string           `json:"deviceId"`
	Platform string           `json:"platform"`
	Version  string           `json:"version"`
	Type     SessionType      `json:"type"`
}
