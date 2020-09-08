package routes

import (
	"github.com/estebanMe/edcomments/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//SetUserRouter route to user register
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}