package main

import "net/http"

// Handle any error (not going into a valid route)
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
