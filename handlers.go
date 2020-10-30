package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Enter: home")
	defer log.Println("Exit: home")

	w.Header().Set("Content-Type", "application/json")

	message := Message{
		Number: uuid.New().String(),
	}
	messageJson, err := json.Marshal(message)
	if err != nil {
	}

	for i := 0; i < 3; i++ {
		req, err := http.NewRequest("POST", "http://storage:5010/store", bytes.NewBuffer(messageJson))
		if err != nil {
			log.Fatalln(err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Print("Unable to connect to storage")
			if i == 2 {
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode("Sorry our service is currently unavailable")
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				return
			}
			time.Sleep(1 * time.Second)
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		err = resp.Body.Close()
		if err != nil {
			log.Print("Error")
		}
		log.Print("Response Body:", string(body))
		i = 3
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Welcome, your unique visitor number is: " + message.Number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	return
}

func alive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Service is alive")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	return
}
