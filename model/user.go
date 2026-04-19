package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" validate:"required,min=4,max=200" json:"name"`
	Email    string `gorm:"size:255;not null;unique" validate:"required,email" json:"email"`
	Password string `gorm:"size:255;not null" validate:"required,min=4,max=255" json:"-"`
	Phone    string `gorm:"size:15;not null" validate:"required,min=11,max=15" json:"phone"`
}
