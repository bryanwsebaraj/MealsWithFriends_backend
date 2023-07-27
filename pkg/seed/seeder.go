package seed

import (
	"fmt"
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
		//UniversityID: 1,
	},
	{
		FirstName:  "George",
		LastName:   "Washington",
		Email:      "georgew@freedom.usa",
		Password:   "usa",
		Gender:     "Male",
		GradeLevel: "senior",
		CollegeID:  2,
		//UniversityID: 1,
	},
	{
		FirstName:  "John",
		LastName:   "Adams",
		Email:      "ja@gmail.com",
		Password:   "usa",
		Gender:     "Male",
		GradeLevel: "junior",
		CollegeID:  2,
		//UniversityID: 1,
	},
}

var meals = []models.Meal{
	{
		MealType: "lunch",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 0, 0, 0, 0, time.UTC),
		Location: "Timothy Dwight",
		TimeSlot: 1,
		IsActive: true,
	},
	{
		MealType: "dinner",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day()), 0, 0, 0, 0, time.UTC),
		Location: "Silliman",
		TimeSlot: 2,
		IsActive: true,
	},
	{
		MealType: "lunch",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day()), 0, 0, 0, 0, time.UTC),
		Location: "Silliman",
		TimeSlot: 2,
		IsActive: true,
	},
	{
		MealType: "dinner",
		Date:     time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day()), 0, 0, 0, 0, time.UTC),
		Location: "TD",
		TimeSlot: 2,
		IsActive: true,
	},
}

var user_meals = []models.User{}

var timePreferences = []models.TimePreference{
	{
		UserID:         2,
		Date:           time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 0, 0, 0, 0, time.UTC),
		LunchSlot:      1,
		DinnerSlot:     2,
		LunchResponse:  false,
		DinnerResponse: true,
	},
	{
		UserID:         1,
		Date:           time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 0, 0, 0, 0, time.UTC),
		LunchSlot:      2,
		DinnerSlot:     2,
		LunchResponse:  true,
		DinnerResponse: true,
	},
	{
		UserID:         3,
		Date:           time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 0, 0, 0, 0, time.UTC),
		LunchSlot:      2,
		DinnerSlot:     2,
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

	for i, _ := range user_meals {
		errMeal = db.Debug().Model(&models.Meal{}).Create(&meals[i]).Error
		if errMeal != nil {
			log.Fatalf("cannot seed users table: %v", errMeal)
		}

	}

	// move to matching algorithm. how to create new entries in junction table

	user := models.User{}
	userGotten1, err := user.FindUserByID(db, 1)
	user2 := models.User{}
	userGotten2, err := user2.FindUserByID(db, 2)
	user3 := models.User{}
	userGotten3, err := user3.FindUserByID(db, 3)

	meal := models.Meal{}
	mealGotten, err := meal.FindMealByID(db, 1)

	db.Unscoped().Model(&userGotten1).Association("Meals").Clear()
	db.Unscoped().Model(&userGotten2).Association("Meals").Clear()
	db.Unscoped().Model(&userGotten3).Association("Meals").Clear()

	db.Model(&mealGotten).Association("Users").Delete(users[2])
	fmt.Println("here")

	db.Model(&userGotten1).Association("Meals").Append([]models.Meal{meals[0]}) // meal 1, user 1
	db.Model(&userGotten1).Association("Meals").Append([]models.Meal{meals[2]}) // meal 3, user 1
	db.Model(&userGotten1).Association("Meals").Append([]models.Meal{meals[3]}) // meal 4, user 1
	db.Model(&userGotten2).Association("Meals").Append([]models.Meal{meals[3]}) // meal 4, user 2
	db.Model(&userGotten2).Association("Meals").Append([]models.Meal{meals[2]}) // meal 3, user 2
	db.Model(&userGotten3).Association("Meals").Append([]models.Meal{meals[2]}) // meal 3, user 3

	db.Model(&mealGotten).Association("Users").Append([]models.User{users[2]}) // meal 1, user 3
	db.Model(&mealGotten).Association("Users").Append([]models.User{users[1]}) // meal 1, user 2

	//userGotten1.UpdateAUser(db, 1)
	fmt.Println(userGotten1)
	//db.Save(userGotten1)
	//fmt.Println(userGotten1.Meals)

	//listmeals := db.Model(&userGotten1).Association("Meals").Find(&meals)
	//fmt.Println(listmeals)
	//fmt.Println(db.Model(&userGotten2).Association("Meals").Find(&meals))

	//exDate := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	mealcount := db.Model(&userGotten1).Association("Meals").Count()
	usercount := db.Model(&mealGotten).Association("Users").Count()
	fmt.Println(mealcount, "", usercount)
	/*
		var users10 []models.User
		err = db.Model(&models.User{}).Preload("Meals").Find(&users10).Error
		fmt.Println(users10)
	*/
	fmt.Println("space")
	fmt.Println(mealGotten.FindUsersByMealID(db, 1))
	fmt.Println(mealGotten.FindTimePrefsByMealID(db, 1))
}
