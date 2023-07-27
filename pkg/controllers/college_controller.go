package controllers

import (
	"net/http"
	"strconv"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/models"
	"github.com/bryanwsebaraj/mealswithfriends/pkg/responses"

	"github.com/gorilla/mux"
)

func (server *Server) GetColleges(w http.ResponseWriter, r *http.Request) {
	college := models.College{}
	colleges, err := college.FindAllColleges(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, colleges)
}

func (server *Server) GetCollegesByUni(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	unid, err := strconv.ParseUint(vars["university"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	college := models.College{}
	collegesGotten, err := college.FindCollegesByUni(server.DB, uint32(unid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, collegesGotten)
}

func (server *Server) GetCollege(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	college := models.College{}
	collegeGotten, err := college.FindCollegeByID(server.DB, uint32(cid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, collegeGotten)
}
