package model

type Character struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Race string `json:"race"`
}
