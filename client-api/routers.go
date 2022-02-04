package main

import (
	"Bleenco/client-api/utils"
	"github.com/gorilla/mux"
)

func registerRoutes(router *mux.Router) {
	state := new(utils.ParserState)
	router.HandleFunc("/parse", handleParser(state)).Methods("GET")
	router.HandleFunc("/select/{page}", handleSelect).Methods("GET")
}
