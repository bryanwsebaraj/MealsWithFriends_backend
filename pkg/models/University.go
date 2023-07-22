package models

import (
	"errors"
	//"html"
	//"log"
	//"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type University struct {
	ID             uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UniversityName string    `gorm:"size:255;not null;unique" json:"uniname"`
	AthleticConf   string    `gorm:"size:255" json:"athconf"`
	Colleges       []College `json:"college_list"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (uni *University) FindAllUniversities(db *gorm.DB) (*[]University, error) {
	var err error
	universities := []University{}
	err = db.Debug().Model(&University{}).Limit(100).Find(&universities).Error
	if err != nil {
		return &[]University{}, err
	}
	return &universities, err
}

func (uni *University) FindUniversityByID(db *gorm.DB, unid uint32) (*University, error) {
	var err error
	err = db.Debug().Model(University{}).Where("id = ?", unid).Take(&uni).Error
	if err != nil {
		return &University{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &University{}, errors.New("University Not Found")
	}
	return uni, err
}
