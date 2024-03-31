package model

import "gorm.io/gorm"

type User_Table struct {
	gorm.Model
	ID       uint32 `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
