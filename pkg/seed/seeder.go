package seed

import (
	"log"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		FirstName:  "Bryan",
		LastName:   "SebaRaj",
		Email:      "donaldjtrump@gmail.com",
		Password:   "pass",
		Gender:     "Male",
		GradeLevel: "soph",
		//University: "Yale",
		//College:    "TD",
	},
	models.User{
		FirstName:  "Bryan",
		LastName:   "SebaRaj",
		Email:      "joebiden@gmail.com",
		Password:   "not a password",
		Gender:     "Male",
		GradeLevel: "soph",
		//University: "Yale",
		//College:    "Morse",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

	}
}
