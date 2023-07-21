package models

import (
	//"errors"
	//"html"
	//"log"
	//"strings"
	"time"
	//"github.com/jinzhu/gorm"
)

type College struct {
	ID         uint32     `gorm:"primary_key;auto_increment" json:"id"` //  should this be a primary key or should college and uni be primary key?
	College    string     `gorm:"size:255;not null" json:"college"`
	University University `gorm:"size:255;not null" json:"uniname"`
	City       string     `gorm:"size:255" json:"city"`
	State      string     `gorm:"size:100" json:"state"`
	CreatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
