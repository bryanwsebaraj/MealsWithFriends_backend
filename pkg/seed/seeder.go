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
		Email:      "bryanwsebaraj@gmail.com",
		Password:   "pass",
		Gender:     "Male",
		GradeLevel: "sophomore",
		College:    colleges[0],
	},
	models.User{
		FirstName:  "George",
		LastName:   "Washington",
		Email:      "georgew@freedom.usa",
		Password:   "usa",
		Gender:     "Male",
		GradeLevel: "senior",
		College:    colleges[1],
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

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}

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
}
