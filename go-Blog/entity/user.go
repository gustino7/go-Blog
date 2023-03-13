package entity

import (
	"tugas4/utils"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Nama     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" type:"varchar(100)"`
	Role     string `json:"role"`
	Blog     []Blog `json:"blog,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = utils.PasswordHash(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password != "" && tx.Statement.Changed("Password") {
		var err error
		u.Password, err = utils.PasswordHash(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}
