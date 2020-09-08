package routes

import "github.com/gorilla/mux"

//InitRoutes Begin Routes and return router
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)

	return router
}