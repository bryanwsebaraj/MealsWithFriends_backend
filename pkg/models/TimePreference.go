package models

import (
	//"errors"
	//"html"
	//"log"
	//"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type TimePreference struct {
	UserID User `gorm:"primary_key;not null;autoIncrement:false" json:"id"`
	//Date   time.Time

	Day   int        `gorm:"primary_key;not null" json:"day"`
	Month time.Month `gorm:"primary_key;not null" json:"month"`
	Year  int        `gorm:"primary_key;not null" json:"year"`

	LunchSlot      uint32    `gorm:"size:255" json:"lunchslot"`
	DinnerSlot     uint32    `gorm:"size:255" json:"dinnerslot"`
	LunchResponse  bool      `gorm:"size:255" json:"lunchres"`
	DinnerResponse bool      `gorm:"size:100" json:"dinnerres"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (timePref *TimePreference) Prepare(u *User) {
	//timePref.Date = (time.Now()).Format("2006/01/02")
	timePref.Year, timePref.Month, timePref.Day = time.Now().Date()
	timePref.LunchResponse = false
	timePref.DinnerResponse = false
	timePref.CreatedAt = time.Now()
	timePref.UpdatedAt = time.Now()
}

func (timePref *TimePreference) ValidateUpdate() error {

	return nil
}

func (timePref *TimePreference) CreateTimePreference(db *gorm.DB, u *User) (*TimePreference, error) {
	var err error
	err = db.Debug().Create(&timePref).Error
	if err != nil {
		return &TimePreference{}, err
	}
	return timePref, nil
}

func (timePref *TimePreference) UpdateTimePref() error {

	return nil
}
