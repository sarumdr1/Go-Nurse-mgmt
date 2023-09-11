package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required min=6"`
}
