package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Key         string `gorm:"unique;not null"`
	Description string
}
