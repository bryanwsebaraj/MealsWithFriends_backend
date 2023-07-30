package seed

import (
	"log"
	"time"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		FirstName:  "Bryan",
		LastName:   "SebaRaj",
		Email:      "bryanwsebaraj@gmail.com",
		Password:   "pass",
		Gender:     "Male",
		GradeLevel: "sophomore",
		CollegeID:  1,
	},
	{
		FirstName:  "George",
		LastName:   "Washington",
		Email:      "georgew@president.usa",
		Password:   "usa",
		Gender:     "Male",
		GradeLevel: "senior",
		CollegeID:  2,
	},
	{
		FirstName:  "Sam",
		LastName:   "Adams",
		Email:      "sam@foundingfather.beer",
		Password:   "usa",
		Gender:     "Male",
		GradeLevel: "junior",
		CollegeID:  2,
	},
}

// sample meals without users associated with them
var meals = []models.Meal{
	{
		MealType: "lunch",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		Location: "Timothy Dwight",
		TimeSlot: 1,
		IsActive: true,
	},
	{
		MealType: "dinner",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		Location: "Silliman",
		TimeSlot: 2,
		IsActive: true,
	},
	{
		MealType: "lunch",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		Location: "Silliman",
		TimeSlot: 2,
		IsActive: true,
	},
	{
		MealType: "dinner",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		Location: "TD",
		TimeSlot: 2,
		IsActive: true,
	},
}

// sameple timepreferences for yesterday (day before boot up)
var timePreferences = []models.TimePreference{
	{
		UserID:         2,
		Date:           time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		LunchSlot:      1,
		DinnerSlot:     2,
		LunchResponse:  false,
		DinnerResponse: true,
	},
	{
		UserID:         1,
		Date:           time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		LunchSlot:      1,
		DinnerSlot:     3,
		LunchResponse:  true,
		DinnerResponse: true,
	},
	{
		UserID:         3,
		Date:           time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local),
		LunchSlot:      2,
		DinnerSlot:     3,
		LunchResponse:  true,
		DinnerResponse: true,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	errUni := db.Debug().DropTableIfExists(&models.University{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", errUni)
	}

	errCollege := db.Debug().DropTableIfExists(&models.College{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", errCollege)
	}

	errTP := db.Debug().DropTableIfExists(&models.TimePreference{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", errTP)
	}

	errMeal := db.Debug().DropTableIfExists(&models.Meal{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", errMeal)
	}

	errUni = db.Debug().AutoMigrate(&models.University{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", errUni)
	}

	errCollege = db.Debug().AutoMigrate(&models.College{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", errCollege)
	}

	errTP = db.Debug().AutoMigrate(&models.TimePreference{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", errTP)
	}

	errMeal = db.Debug().AutoMigrate(&models.Meal{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", errMeal)
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	for i, _ := range universities {
		errUni = db.Debug().Model(&models.University{}).Create(&universities[i]).Error
		if errUni != nil {
			log.Fatalf("cannot seed users table: %v", errUni)
		}

	}

	for i, _ := range colleges {
		errCollege = db.Debug().Model(&models.College{}).Create(&colleges[i]).Error
		if errCollege != nil {
			log.Fatalf("cannot seed users table: %v", errCollege)
		}

	}

	for i, _ := range timePreferences {
		errTP = db.Debug().Model(&models.TimePreference{}).Create(&timePreferences[i]).Error
		if errTP != nil {
			log.Fatalf("cannot seed users table: %v", errTP)
		}

	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}

	for i, _ := range meals {
		errMeal = db.Debug().Model(&models.Meal{}).Create(&meals[i]).Error
		if errMeal != nil {
			log.Fatalf("cannot seed users table: %v", errMeal)
		}

	}

	for i, _ := range meals {
		meal := models.Meal{}
		mealGotten, err := meal.FindMealByID(db, uint32(i+1))
		if err != nil {
			log.Fatalf("cannot clear user associations: %v", err)
		}
		db.Unscoped().Model(&mealGotten).Association("Users").Clear()
	}
	/*
		var users10 []models.User
		err = db.Model(&models.User{}).Preload("Meals").Find(&users10).Error
		fmt.Println(users10)

		mealGotten := models.Meal{}
		fmt.Println(mealGotten.FindUsersByMealID(db, 1))
		fmt.Println(mealGotten.FindTimePrefsByMealID(db, 1))
	*/

}
