package model

import (
	"time"
)

// user
type User struct {
	UID                uint           `gorm:"primaryKey;column:uid"`
	Cert               string         `gorm:"column:cert;size:100;not null"`
	Password           string         `gorm:"column:password;size:100;not null"`
	RegisteredAt       time.Time      `gorm:"column:registered_datetime;autoCreateTime"`
	LastLoginAt        *time.Time     `gorm:"column:last_login_datetime"`

	Detail             UserDetail     `gorm:"foreignKey:UserUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "u_user"
}

type UserDetail struct {
	UserUID   uint       `gorm:"primaryKey;column:user_uid"`
	Email     string     `gorm:"size:100;not null"`
	Name      string     `gorm:"size:50"`
	Birthday  *time.Time
	UserType  string     `gorm:"type:enum('admin','user','creator');default:'user'"`

	User      *User      `gorm:"foreignKey:UserUID;references:UID"`
}

func (UserDetail) TableName() string {
	return "u_user_detail"
}


