package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/alive", alive).Methods("GET")
	log.Printf("Starting server on %s", ":5000")
	err := http.ListenAndServe(":5000", r)
	log.Fatal(err)
}
