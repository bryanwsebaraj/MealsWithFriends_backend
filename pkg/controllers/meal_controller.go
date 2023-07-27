package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"
	"github.com/bryanwsebaraj/mealswithfriends/pkg/responses"
	"github.com/gorilla/mux"
)

func (server *Server) GetMealByUserDateMeal(w http.ResponseWriter, r *http.Request) {
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

	mealType, err := models.ValidateMealType(string(vars["mealtype"]))
	meal := models.Meal{}
	mealsGotten, err := meal.FindMealByUserIDDate(server.DB, uint32(uid), date, mealType)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, mealsGotten)
}

func (server *Server) GetMealsForUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	meal := models.Meal{}
	mealsGotten, err := meal.FindMealsByUserID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, mealsGotten)
}

func (server *Server) GetUsersByMeal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mid, err := strconv.ParseUint(vars["mid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	meal := models.Meal{}
	usersGotten, err := meal.FindUsersByMealID(server.DB, uint32(mid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, usersGotten)
}

func (server *Server) GetMealByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mid, err := strconv.ParseUint(vars["mid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	meal := models.Meal{}
	mealGotten, err := meal.FindMealByID(server.DB, uint32(mid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, mealGotten)
}

func (server *Server) GetTimePrefByMeal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mid, err := strconv.ParseUint(vars["mid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	meal := models.Meal{}
	timePrefGotten, err := meal.FindTimePrefsByMealID(server.DB, uint32(mid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, timePrefGotten)
}
