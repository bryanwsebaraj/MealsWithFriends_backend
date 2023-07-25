package models

import (
	"errors"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

type TimePreference struct {
	UserID         uint32    `gorm:"primary_key;not null" json:"user_id"`
	Date           time.Time `gorm:"primary_key" json:"date"`
	LunchSlot      uint32    `gorm:"size:255" json:"lunchslot"`
	DinnerSlot     uint32    `gorm:"size:255" json:"dinnerslot"`
	LunchResponse  bool      `gorm:"size:255" json:"lunchres"`
	DinnerResponse bool      `gorm:"size:100" json:"dinnerres"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func cleanDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}

func (timePref *TimePreference) SaveTimePreference(db *gorm.DB, uid uint32) (*TimePreference, error) {
	var err error
	timePref.UserID = uid
	timePref.Date = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	timePref.LunchSlot = 0
	timePref.DinnerSlot = 0
	timePref.LunchResponse = false
	timePref.DinnerResponse = false
	timePref.CreatedAt = time.Now()
	timePref.UpdatedAt = time.Now()
	err = db.Debug().Create(&timePref).Error
	if err != nil {
		return &TimePreference{}, err
	}

	return timePref, nil
}

func (timePref *TimePreference) ValidateTimePref() error {
	timePref.UpdatedAt = time.Now()
	if reflect.TypeOf(timePref.LunchSlot).Kind() != reflect.Uint32 {
		return errors.New("Lunch Slot not uint32")
	}
	if reflect.TypeOf(timePref.DinnerSlot).Kind() != reflect.Uint32 {
		return errors.New("Lunch Slot not uint32")
	}
	return nil
}

func (timePref *TimePreference) FindAllTimePrefences(db *gorm.DB) (*[]TimePreference, error) {
	var err error
	timePreferences := []TimePreference{}
	err = db.Debug().Model(&TimePreference{}).Limit(1000).Find(&timePreferences).Error
	if err != nil {
		return &[]TimePreference{}, err
	}
	return &timePreferences, err
}

func (timePref *TimePreference) FindTimePrefsByUser(db *gorm.DB, uid uint32) (*[]TimePreference, error) {
	var err error
	timePrefs := []TimePreference{}
	err = db.Debug().Model(TimePreference{}).Where("user_id = ?", uid).Find(&timePrefs).Error
	if err != nil {
		return &[]TimePreference{}, err
	}
	return &timePrefs, err
}

func (timePref *TimePreference) FindTimePrefsByDate(db *gorm.DB, date time.Time) (*[]TimePreference, error) {
	var err error
	timePrefs := []TimePreference{}
	err = db.Debug().Model(TimePreference{}).Where("date = ?", cleanDate(date)).Find(&timePrefs).Error
	if err != nil {
		return &[]TimePreference{}, err
	}
	return &timePrefs, err
}

func (timePref *TimePreference) FindTimePrefByUserDate(db *gorm.DB, uid uint32, date time.Time) (*TimePreference, error) {
	var err error
	err = db.Debug().Model(TimePreference{}).Where("user_id = ? AND date = ?", uid, cleanDate(date)).Take(&timePref).Error
	if err != nil {
		return &TimePreference{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &TimePreference{}, errors.New("Time Preference Not Found")
	}
	return timePref, err
}

func (timePref *TimePreference) UpdateTimePref(db *gorm.DB, uid uint32, date time.Time) (*TimePreference, error) {

	db = db.Debug().Model(&TimePreference{}).Where("user_id = ? AND date = ?", uid, cleanDate(date)).Take(&TimePreference{}).UpdateColumns(
		map[string]interface{}{
			"lunch_slot":      timePref.LunchSlot,
			"dinner_slot":     timePref.DinnerSlot,
			"lunch_response":  timePref.LunchResponse,
			"dinner_response": timePref.DinnerResponse,
			"updated_at":      time.Now(),
		},
	)
	if db.Error != nil {
		return &TimePreference{}, db.Error
	}
	err := db.Debug().Model(&TimePreference{}).Where("user_id = ? AND date = ?", uid, cleanDate(date)).Take(&timePref).Error
	if err != nil {
		return &TimePreference{}, err
	}
	return timePref, nil
}

func (timePref *TimePreference) DeleteATimePref(db *gorm.DB, uid uint32, date time.Time) (int64, error) {

	db = db.Debug().Model(&TimePreference{}).Where("user_id = ? AND date = ?", uid, cleanDate(date)).Take(&TimePreference{}).Delete(&TimePreference{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
