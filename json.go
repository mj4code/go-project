package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Error occured with code 5xx :", msg)
	}

	type Error struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, Error{
		Error: msg,
	})

}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal the payload %v", payload)
		w.WriteHeader(500)
		return
	}
	//We're responding with JSON
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
