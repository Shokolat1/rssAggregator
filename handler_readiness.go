package main

import "net/http"

// Function signature to use in order to define an HTTP handler (first using a response writer, and then the request itself)
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
