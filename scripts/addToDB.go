package scripts

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

func AddOnTime(db *gorm.DB) {
	for {
		now := time.Now()

		if now.Hour() == now.Hour() && now.Minute() == now.Minute() {
			fmt.Println("Adding new preferences for users")
			AddTimePrefs(db)
		}

		if now.Hour() == now.Hour() && now.Minute() == now.Minute() {
			fmt.Println("Adding new meal match entries in DB")
			AddMealMatches(db)
		}

		// Wait before checking again
		time.Sleep(60 * time.Minute)
	}
}
