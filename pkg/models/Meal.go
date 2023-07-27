package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Meal struct {
	MealID    uint32    `gorm:"primary_key;auto_increment=true" json:"meal_id"`
	MealType  string    `gorm:"size:255;not null" json:"mealtype"`
	Date      time.Time `gorm:"not null" json:"date"`
	Location  string    `gorm:"not null" json:"location"`
	TimeSlot  uint32    `gorm:"not null" json:"timeslot"`
	Users     []User    `gorm:"many2many:user_meals" json:"users"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (m *Meal) SaveMeal(db *gorm.DB, mealType string, date time.Time, timeslot uint32) (*Meal, error) {
	m.MealID = 0
	m.MealType = mealType
	m.Date = cleanDate(date)
	m.TimeSlot = timeslot
	m.IsActive = true
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	err := db.Debug().Create(&m).Error
	if err != nil {
		return &Meal{}, err
	}
	return m, nil
}

func ValidateMealType(mealType string) (string, error) {
	if mealType == "lunch" || mealType == "dinner" {
		return mealType, nil
	} else {
		return mealType, errors.New("Invalid string")
	}
}

func (m *Meal) DeactivateMeal(db *gorm.DB) {
	m.IsActive = false
}

func (m *Meal) FindMealsByDate(db *gorm.DB, date time.Time) (*[]Meal, error) {
	var err error
	meals := []Meal{}
	err = db.Debug().Model(Meal{}).Where("date = ?", cleanDate(date)).Find(&meals).Error
	if err != nil {
		return &[]Meal{}, err
	}
	return &meals, err
}

func (m *Meal) FindMealsByUserID(db *gorm.DB, uid uint32) (*[]Meal, error) {
	user := User{}
	userGotten, err := user.FindUserByID(db, uid)
	if err != nil {
		return &[]Meal{}, errors.New("User Not Found")
	}

	err = db.Model(&User{}).Preload("Meals").Take(&userGotten).Error
	if err != nil {
		return &[]Meal{}, errors.New("Meals Not Found")
	}
	meals := userGotten.Meals
	return &meals, err
}

func (m *Meal) FindMealByUserIDDate(db *gorm.DB, uid uint32, date time.Time, meal string) (*Meal, error) {
	user := User{}
	userGotten, err := user.FindUserByID(db, uid)
	if err != nil {
		return &Meal{}, errors.New("User Not Found")
	}

	err = db.Model(&User{}).Preload("Meals").Take(&userGotten).Error
	if err != nil {
		return &Meal{}, errors.New("Meals Not Found")
	}

	meals := userGotten.Meals
	for i := range meals {
		if cleanDate(date).Equal(meals[i].Date) && meals[i].MealType == meal {
			return &meals[i], err
		}
	}
	return m, err

}

func (m *Meal) FindUsersByMealID(db *gorm.DB, mid uint32) (*[]User, error) {
	meal := Meal{}
	mealGotten, err := meal.FindMealByID(db, mid)
	if err != nil {
		return &[]User{}, errors.New("Meal Not Found")
	}

	err = db.Model(&Meal{}).Preload("Users").Take(&mealGotten).Error
	if err != nil {
		return &[]User{}, errors.New("Users Not Found")
	}

	users := mealGotten.Users
	return &users, err

}

func (m *Meal) FindTimePrefsByMealID(db *gorm.DB, mid uint32) (*[]TimePreference, error) {
	users, err := m.FindUsersByMealID(db, mid)
	userList := *users
	timePrefs := []TimePreference{}
	meal, err := m.FindMealByID(db, mid)

	for i := range *users {
		timePref := TimePreference{}
		timePref1, err := timePref.FindTimePrefByUserDate(db, userList[i].ID, meal.Date)
		if err != nil {
			return &[]TimePreference{}, errors.New("Time Preference Not Found")
		}
		timePrefs = append(timePrefs, *timePref1)
	}

	return &timePrefs, err

}

func (m *Meal) FindMealByID(db *gorm.DB, mid uint32) (*Meal, error) {
	var err error
	err = db.Debug().Model(Meal{}).Where("meal_id = ?", mid).Take(&m).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Meal{}, errors.New("Meal Not Found")
	}
	if err != nil {
		return &Meal{}, err
	}

	return m, err
}
