package models

import "gorm.io/gorm"

type UserAuth struct {
	gorm.Model
	Email    string `gorm:"not null"`
	Password string `gorm:"not null"`
}
