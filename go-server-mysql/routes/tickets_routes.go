package routes

import (
	"main/controllers"

	"github.com/gorilla/mux"
)

func SetTicketsRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/tickets/api").Subrouter()
	subRoute.HandleFunc("/lista", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/guardar", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")
}
