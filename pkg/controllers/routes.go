package controllers

import "github.com/bryanwsebaraj/mealswithfriends/pkg/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//College routes
	s.Router.HandleFunc("/colleges", middlewares.SetMiddlewareJSON(s.GetColleges)).Methods("GET")
	s.Router.HandleFunc("/colleges/{id}", middlewares.SetMiddlewareJSON(s.GetCollege)).Methods("GET")
	s.Router.HandleFunc("/collegesByUni/{university}", middlewares.SetMiddlewareJSON(s.GetCollegesByUni)).Methods("GET")

	//University routes
	s.Router.HandleFunc("/universities", middlewares.SetMiddlewareJSON(s.GetUniversities)).Methods("GET")
	s.Router.HandleFunc("/universities/{id}", middlewares.SetMiddlewareJSON(s.GetUniversity)).Methods("GET")

	//TimePreference routes
	s.Router.HandleFunc("/timepreferences", middlewares.SetMiddlewareJSON(s.GetAllTimePref)).Methods("GET")
	s.Router.HandleFunc("/timepreferences/{user_id}", middlewares.SetMiddlewareJSON(s.GetTimePrefsByUser)).Methods("GET")
	s.Router.HandleFunc("/timepreferencesByDate/{date}", middlewares.SetMiddlewareJSON(s.GetTimePrefsByDate)).Methods("GET")
	s.Router.HandleFunc("/timepreferences/{user_id}/{date}", middlewares.SetMiddlewareJSON(s.GetTimePrefByUserDate)).Methods("GET")
	s.Router.HandleFunc("/timepreferences/{user_id}/{date}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateTimePref))).Methods("PUT")

}
