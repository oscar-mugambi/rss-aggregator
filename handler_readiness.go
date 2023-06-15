package main

import "net/http"

// this is the method signature you have to use if you want GO define a HTTP handler in the way GO expects
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
