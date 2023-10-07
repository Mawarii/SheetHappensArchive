package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"id"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Characters []Character
}
