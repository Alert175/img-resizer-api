package models

import (
	"gorm.io/gorm"
)

type ImageLog struct {
	gorm.Model
	Url    string `gorm:"not null"`
	Status string `gorm:"not null"`
}
