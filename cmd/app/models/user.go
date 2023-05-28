package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Auth0ID       string
	Relationships []UserBusinessRelationship `gorm:"foreignKey:UserID"`
	Settings GeneralSettings `gorm:"foreignKey:UserID"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

type BusinessAccount struct {
	gorm.Model
	Name          string
	Relationships []UserBusinessRelationship `gorm:"foreignKey:BusinessAccountID"`
	AccountSettings AccountSettings `gorm:"foreignKey:BusinessAccountID"`
	Posts []Post `gorm:"foreignKey:BusinessAccountID"`
}

type UserBusinessRelationship struct {
	gorm.Model
	UserID            uint
	User              User
	BusinessAccountID uint
	BusinessAccount   BusinessAccount
	Role              string
}