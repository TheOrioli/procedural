package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	mazeEndpoints "github.com/Aorioli/procedural/endpoints/maze"
	mazeService "github.com/Aorioli/procedural/services/maze"
	dungeonEndpoints "github.com/Aorioli/procedural/endpoints/dungeon"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}

	router := mux.NewRouter().StrictSlash(true)

	mazeSvc := mazeService.New()
	root := context.Background()

	mazeRouter := router.PathPrefix("/maze").Subrouter()
	for _, route := range mazeEndpoints.HTTP(mazeSvc, root) {
		log.Println("/maze" + route.Path)
		mazeRouter.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	dungeonRouter := router.PathPrefix("/dungeon").Subrouter()
	for _, route := range dungeonEndpoints.HTTP(root) {
		log.Println("/dungeon" + route.Path)
		dungeonRouter.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	log.Fatalln(http.ListenAndServe(
		":"+port,
		handlers.LoggingHandler(
			os.Stdout,
			router,
		),
	))
}
