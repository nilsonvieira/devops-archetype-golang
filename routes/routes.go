package routes

import (
	"go-api/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	return router
}
