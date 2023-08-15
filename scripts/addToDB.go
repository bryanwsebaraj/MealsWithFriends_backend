package scripts

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func AddOnTime(db *gorm.DB) {
	for {
		now := time.Now()

		if now.Hour() == 4 && now.Minute() == 0 {
			fmt.Println("Adding new preferences for users")
			AddTimePrefs(db)
		}

		if now.Hour() == 4 && now.Minute() == 30 {
			fmt.Println("Adding new meal match entries in DB")
			AddMealMatches(db)
			DeactivateYesterdayMeals(db)
		}

		// Wait before checking again
		time.Sleep(1 * time.Minute)
	}
}
