package controllers

import (
	"net/http"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "API successfully interfaced")

}
