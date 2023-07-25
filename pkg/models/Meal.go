package models

import (
	//"errors"
	//"html"
	//"log"
	//"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Meal struct {
	MealID    uint32    `gorm:"primary_key;auto_increment" json:"mealid"`
	MealType  string    `gorm:"size:255;not null" json:"mealtype"`
	Date      time.Time `gorm:"default:CURRENT_TIME" json:"day"`
	TimeSlot  uint32    `gorm:"default:CURRENT_TIME" json:"time"`
	Users     []User    `gorm:"many2many:user_matches" json:"user_matches"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (m *Meal) SaveMeal(db *gorm.DB, mealType string, date time.Time, timeslot uint32) (*Meal, error) {
	m.MealID = 0
	m.MealType = mealType
	m.Date = date
	m.TimeSlot = timeslot
	//m.Users =
	m.IsActive = true
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	err := db.Debug().Create(&m).Error
	if err != nil {
		return &Meal{}, err
	}

	return m, nil
}

func (m *Meal) FindMealsByDate() {}

func (m *Meal) FindMealsByUser() {}

func (m *Meal) FindMealsByUserDate() {}
