package model

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Race   string `json:"race"`
	UserID int    `gorm:"foreignKey:UserID" json:"user_id"`
}
