package scripts

import (
	//"log"
	"fmt"
	"time"

	//"github.com/bryanwsebaraj/mealswithfriends/pkg/models"

	"github.com/jinzhu/gorm"
)

func AddOnTime(db *gorm.DB) {
	for {
		now := time.Now()

		if now.Hour() == now.Hour() && now.Minute() == now.Minute() {
			fmt.Println("Adding new preferences for users")
			AddTimePrefs(db)
		}

		if now.Hour() == 11 && now.Minute() == 55 {
			fmt.Println("Adding new meal match entries in DB")
			AddMealMatches(db)
		}

		// Wait for a short duration before checking again
		time.Sleep(60 * time.Minute)
	}
}
