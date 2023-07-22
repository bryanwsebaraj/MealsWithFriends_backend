package controllers

import (
	"net/http"
	"strconv"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"
	"github.com/bryanwsebaraj/mealswithfriends/pkg/responses"

	"github.com/gorilla/mux"
)

func (server *Server) GetUniversities(w http.ResponseWriter, r *http.Request) {

	university := models.University{}

	universities, err := university.FindAllUniversities(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, universities)
}

func (server *Server) GetUniversity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	unid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	university := models.University{}
	universityGotten, err := university.FindUniversityByID(server.DB, uint32(unid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, universityGotten)
}
