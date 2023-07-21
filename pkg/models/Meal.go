package models

import (
	//"errors"
	//"html"
	//"log"
	//"strings"
	"time"
	//"github.com/jinzhu/gorm"
)

type Meal struct {
	MealID    uint32    `gorm:"primary_key;auto_increment" json:"mealid"`
	MealType  string    `gorm:"size:255;not null" json:"mealtype"`
	Date      time.Time `gorm:"default:CURRENT_DATE" json:"date"`
	Time      time.Time `gorm:"default:CURRENT_TIME" json:"time"`
	Users     []User    `gorm:"many2many:user_matches" json:"user_matches"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
