package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string `json:"token" form:"token"`
}
