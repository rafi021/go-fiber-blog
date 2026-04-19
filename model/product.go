package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `gorm:"size:255;not null" validate:"required,min=4,max=200" json:"name"`
	Description string  `gorm:"size:255;" json:"description"`
	Price       float64 `gorm:"not null" validate:"required,gt=0" json:"price"`
}
