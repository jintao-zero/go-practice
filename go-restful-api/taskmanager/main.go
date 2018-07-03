package main

import (
	"github.com/urfave/negroni"
	"go-practice/go-restful-api/taskmanager/common"
	"go-practice/go-restful-api/taskmanager/routers"
	"log"
	"net/http"
)

// Entry point of the program
func main() {
	// Calls startup logic
	common.StartUp()

	// Get the mux router object
	router := routers.InitRoutes()

	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
