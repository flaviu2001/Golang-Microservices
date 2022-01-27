package main

import "github.com/gorilla/mux"

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/parse", handleParser).Methods("GET")
	router.HandleFunc("/select/{page}", handleSelect).Methods("GET")
}
