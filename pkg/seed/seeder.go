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

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
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
}
