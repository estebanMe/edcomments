package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/estebanMe/edcomments/migration"
	"github.com/estebanMe/edcomments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzo la migración...")
		migration.Migrate()
		log.Println("Finalizo la migración...")
	}
    //Begin routes
	router := routes.InitRoutes()



	//Begin the middlewares
	n := negroni.Classic()
	n.UseHandler(router)


	server := &http.Server{
		Addr: "8080",
		Handler: n,
	}

	log.Println("Iniciando el servidor en http://localhost:8080")
	log.Println(server.ListenAndServe())
	log.Println("Finalizo la ejecución programa")

}
