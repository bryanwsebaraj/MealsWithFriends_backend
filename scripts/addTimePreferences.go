package scripts

import (
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

	// for user in users, create new time preference using userID
	timePreference := models.TimePreference{}
	for _, individUser := range *users {
		_, err := timePreference.SaveTimePreference(db, individUser.GetID())
		if err != nil {
			return errors.New("Time Preference Not Created")
		}
	}
	return nil

}
