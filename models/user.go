package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Name			string		`json:"name"`
	Username		string		`json:"username"`
	Permissions		int64		`json:"permissions"`
	Password		string		`json:"password,omitempty" gorm:"-"`
	Hash			string		`json:"-"`
	Groups			[]*Group	`json:"groups,omitempty" gorm:"many2many:user_groups"`
}

func (u *User) BeforeSave() (err error){
	if u.Password == "" {
		return errors.New("password is empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Hash = string(hash)
	return nil
}
