package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	Name        string    `gorm:"not null"`
	PictureUrl  string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	CategoryID  *uint     `gorm:"default:null"`
	Category    *Category `gorm:"foreignKey:CategoryID"`
}
