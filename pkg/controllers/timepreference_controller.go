package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/auth"
	"github.com/bryanwsebaraj/mealswithfriends/pkg/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetAllTimePref(w http.ResponseWriter, r *http.Request) {
	timePref := models.TimePreference{}
	timePrefs, err := timePref.FindAllTimePrefences(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, timePrefs)
}

func (server *Server) GetTimePrefsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	timePrefs := models.TimePreference{}
	timePrefsGotten, err := timePrefs.FindTimePrefsByUser(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, timePrefsGotten)
}

func (server *Server) GetTimePrefsByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date, err := time.Parse("2006-Jan-02", vars["date"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	timePrefs := models.TimePreference{}
	timePrefsGotten, err := timePrefs.FindTimePrefsByDate(server.DB, time.Time(date))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, timePrefsGotten)
}

func (server *Server) GetTimePrefByUserDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse("2006-Jan-02", vars["date"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	timePrefs := models.TimePreference{}
	timePrefsGotten, err := timePrefs.FindTimePrefByUserDate(server.DB, uint32(uid), time.Time(date))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, timePrefsGotten)
}

func (server *Server) UpdateTimePref(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	date, err := time.Parse("2006-Jan-02", vars["date"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	timePref := models.TimePreference{}
	err = json.Unmarshal(body, &timePref)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	err = timePref.ValidateTimePref()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	updatedTimePref, err := timePref.UpdateTimePref(server.DB, uint32(uid), time.Time(date))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedTimePref)

}
