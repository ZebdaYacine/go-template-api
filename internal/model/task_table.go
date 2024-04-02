package model

import (
	"time"

	"gorm.io/gorm"
)

type Task_Table struct {
	gorm.Model
	ID      uint32    `json:"id" gorm:"primaryKey"`
	Code    string    `json:"code" gorm:"Unique"`
	Cost    float64   `json:"coste" gorm:"Unique"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	Status  string    `json:"status"`
	UserID  uint32    `json:"userID" gorm:"references:UserID constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
