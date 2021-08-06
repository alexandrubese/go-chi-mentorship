package main

import (
	"alexapp.pck.com/handler"
	"alexapp.pck.com/initscripts"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	c := chi.NewRouter()

	initscripts.InitMiddlewares(c)
	initscripts.InitAppRoutes(c, &handler.Handler{DB: initscripts.InitDbConnection()})

	log.Println("Launched backend server on port" + initscripts.AppPort)
	err := http.ListenAndServe(initscripts.AppPort, c)
	if err != nil {
		log.Fatal("Server initialisation error : " + err.Error())
	}
}
