package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welcome, your unique visitor number is", uuid.New().String())
}
func alive(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Service is alive")
}
