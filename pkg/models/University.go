package models

import (
	//"errors"
	//"html"
	//"log"
	//"strings"
	"time"
	//"github.com/jinzhu/gorm"
)

type University struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UniversityName string    `gorm:"size:255;not null;unique" json:"uniname"`
	AthleticConf   string    `gorm:"size:255" json:"athconf"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
