package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	dungeonEndpoints "github.com/Aorioli/procedural/endpoints/dungeon"
	mazeEndpoints "github.com/Aorioli/procedural/endpoints/maze"
	musicEndpoints "github.com/Aorioli/procedural/endpoints/music"
	mazeService "github.com/Aorioli/procedural/services/maze"
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

	musicRouter := router.PathPrefix("/music").Subrouter()
	for _, route := range musicEndpoints.HTTP(root) {
		log.Println("/music" + route.Path)
		musicRouter.Handle(route.Path, route.Handler).Methods(route.Method)
	}

	log.Fatalln(http.ListenAndServe(
		":"+port,
		handlers.LoggingHandler(
			os.Stdout,
			router,
		),
	))
}
