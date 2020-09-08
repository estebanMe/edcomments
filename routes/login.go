package routes

import "github.com/gorilla/mux"

//SerLoginRouter router para login
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}

