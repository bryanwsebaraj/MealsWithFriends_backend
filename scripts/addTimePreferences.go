package scripts

import (
	//"log"
	//"fmt"
	//"time"
	"errors"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"

	"github.com/jinzhu/gorm"
)

func AddTimePrefs(db *gorm.DB) error {
	user := models.User{}
	users, err := user.FindAllUsers(db)
	if err != nil {
		return errors.New("Users Not Found")
	}
	//fmt.Print(users)
	// for user in user, take userid, create new time preference, and then add it to DB
	timePreference := models.TimePreference{}
	for _, individUser := range *users {
		// element is the element from someSlice for where we are
		_, err := timePreference.SaveTimePreference(db, individUser.GetID())
		if err != nil {
			return errors.New("Time Preference Not Created")
		}
		//fmt.Println(timePrefCreated)
	}
	return nil

}
