package model

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Race   string `json:"race"`
	UserID uint   `gorm:"foreignKey:UserID" json:"user_id"`
}
