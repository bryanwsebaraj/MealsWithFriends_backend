package scripts

import (
	//"log"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"

	"github.com/jinzhu/gorm"
)

func AddMealMatches(db *gorm.DB) error {
	// for university in universities
	university := models.University{}
	universitiesPointer, err := university.FindAllUniversities(db)
	if err != nil {
		return errors.New("Universities Not Found")
	}
	universities := *universitiesPointer
	for i, _ := range universities {
		// for all colleges in university, make list of all users in university
		college := models.College{}
		collegesForUniPointer, err := college.FindCollegesByUni(db, universities[i].ID)
		if err != nil {
			return errors.New("Colleges For University Not Found")
		}

		collegesForUni := *collegesForUniPointer
		if len(collegesForUni) == 0 {
			fmt.Println("uni skipped bc no colleges in uni")
			continue
		}
		var collegeIDs []uint32

		for j, _ := range collegesForUni {
			collegeIDs = append(collegeIDs, collegesForUni[j].ID)
		}
		// get all users that go to those colleges
		users := []models.User{}
		db.Model(&models.User{}).Where("college_id BETWEEN ? AND ?", collegeIDs[0], collegeIDs[len(collegeIDs)-1]).Find(&users)

		// find time prefs for those users, and categorize users into lists for lunch
		timePrefs := []models.TimePreference{}
		date := time.Now()
		now := time.Date(date.Year(), date.Month(), date.Day()-1, 20, 0, 0, 0, time.Local)

		for k, _ := range users {
			timePref := models.TimePreference{}
			timePrefGottenPointer, err := timePref.FindTimePrefByUserDate(db, users[k].ID, now)
			if err != nil {
				fmt.Println("err")
				continue
			}
			timePrefGotten := *timePrefGottenPointer
			timePrefs = append(timePrefs, timePrefGotten)
		}

		// 4 lunch and 3 dinner slots
		numLunchSlots := 4
		numDinnerSlots := 3
		lunch := make([][]uint32, numLunchSlots) // array of userIDs
		for l, _ := range timePrefs {
			lunch[timePrefs[l].LunchSlot-1] = append(lunch[timePrefs[l].LunchSlot-1], timePrefs[l].UserID)
		}

		dinner := make([][]uint32, numDinnerSlots) // array of userIDs
		for m, _ := range timePrefs {
			dinner[timePrefs[m].DinnerSlot-1] = append(dinner[timePrefs[m].DinnerSlot-1], timePrefs[m].UserID)
		}

		// randomly shuffle users for lunch
		rand.Seed(time.Now().UnixNano())
		for n := 0; n < numLunchSlots; n++ {
			rand.Shuffle(len(lunch[n]), func(i, j int) { lunch[n][i], lunch[n][j] = lunch[n][j], lunch[n][i] })
		}

		// randomly shuffle users for dinner
		rand.Seed(time.Now().UnixNano())
		for o := 0; o < numDinnerSlots; o++ {
			rand.Shuffle(len(dinner[o]), func(i, j int) { dinner[o][i], dinner[o][j] = dinner[o][j], dinner[o][i] })
		}

		// create meals with random location (implement meal location table in future) and randomly pair users based on categories
		for p := 0; p < numLunchSlots; p++ {
			for len(lunch[p]) > 1 {
				if len(lunch[p])%2 != 0 {
					meal := models.Meal{}
					mealSaved, err := meal.SaveMeal(db, "lunch", time.Now(), "commons", uint32(p+1))
					if err != nil {
						fmt.Println("meal not saved")
						continue
					}

					user := models.User{}
					userGotten1, err := user.FindUserByID(db, lunch[p][len(lunch[p])-1])
					user2 := models.User{}
					userGotten2, err := user2.FindUserByID(db, lunch[p][len(lunch[p])-2])
					user3 := models.User{}
					userGotten3, err := user3.FindUserByID(db, lunch[p][len(lunch[p])-3])

					lunch[p] = lunch[p][:len(lunch[p])-3]

					db.Model(&userGotten1).Association("Meals").Append(&mealSaved)
					db.Model(&userGotten2).Association("Meals").Append(&mealSaved)
					db.Model(&userGotten3).Association("Meals").Append(&mealSaved)
				} else {
					meal := models.Meal{}
					mealSaved, err := meal.SaveMeal(db, "lunch", time.Now(), "commons", uint32(p+1))
					if err != nil {
						fmt.Println("meal not saved")
						continue
					}

					user := models.User{}
					userGotten1, err := user.FindUserByID(db, lunch[p][len(lunch[p])-1])
					user2 := models.User{}
					userGotten2, err := user2.FindUserByID(db, lunch[p][len(lunch[p])-2])

					lunch[p] = lunch[p][:len(lunch[p])-2]

					db.Model(&userGotten1).Association("Meals").Append(mealSaved)
					db.Model(&userGotten2).Association("Meals").Append(mealSaved)
				}
			}

		}

		// dinner
		for p := 0; p < numDinnerSlots; p++ {
			for len(dinner[p]) > 1 {
				if len(dinner[p])%2 != 0 {
					meal := models.Meal{}
					mealSaved, err := meal.SaveMeal(db, "dinner", time.Now(), "Silliman", uint32(p+1))
					if err != nil {
						fmt.Println("meal not saved")
						continue
					}

					user := models.User{}
					userGotten1, err := user.FindUserByID(db, dinner[p][len(dinner[p])-1])
					user2 := models.User{}
					userGotten2, err := user2.FindUserByID(db, dinner[p][len(dinner[p])-2])
					user3 := models.User{}
					userGotten3, err := user3.FindUserByID(db, dinner[p][len(dinner[p])-3])

					dinner[p] = dinner[p][:len(dinner[p])-3]

					db.Model(&userGotten1).Association("Meals").Append(&mealSaved)
					db.Model(&userGotten2).Association("Meals").Append(&mealSaved)
					db.Model(&userGotten3).Association("Meals").Append(&mealSaved)
				} else {
					meal := models.Meal{}
					mealSaved, err := meal.SaveMeal(db, "dinner", time.Now(), "Silliman", uint32(p+1))
					if err != nil {
						fmt.Println("meal not saved")
						continue
					}

					user := models.User{}
					userGotten1, err := user.FindUserByID(db, dinner[p][len(dinner[p])-1])
					user2 := models.User{}
					userGotten2, err := user2.FindUserByID(db, dinner[p][len(dinner[p])-2])

					dinner[p] = dinner[p][:len(dinner[p])-2]

					db.Model(&userGotten1).Association("Meals").Append(mealSaved)
					db.Model(&userGotten2).Association("Meals").Append(mealSaved)
				}
			}

		}
	}
	return nil

}

func DeactivateYesterdayMeals(db *gorm.DB) error {
	meal := models.Meal{}
	yesterdayTime := time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() - 1), 20, 0, 0, 0, time.Local)
	mealsPointer, err := meal.FindMealsByDate(db, yesterdayTime)
	if err != nil {
		return err
	}
	meals := *mealsPointer
	for i, _ := range meals {
		meals[i].DeactivateMeal(db)
		meals[i].UpdateMeal(db, meals[i].MealID)
	}
	return nil
}
