package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Responds with error msg instead of JSON
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX  error:", msg)
	}

	// Pass the error struct into the "respondWithJSON" function so it can be converted into a JSON object, where the field name will be called "error"
	type errResp struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResp{Error: msg})
}

// Takes in a response writer, a status code, and what we will work with (as JSON info)
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Try to convert JSON data into byte for security reasons
	data, err := json.Marshal(payload)
	// If error, return failed response
	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	// If no error, create and send response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
