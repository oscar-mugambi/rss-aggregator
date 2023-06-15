package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// a little helper function that helps send JSON responses
// responseWriter is exposed by the std lib. Its what http handlers in GO use
// code is the status code
// interface is something we can marshall into a JSON structure
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//marshall the payload into a JSON obj
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		w.WriteHeader(500)
		return
	}
	// add a header to the response to say we are responding with JSON data
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error", msg)
	}

	type errResponse struct {
		// this is a json reflect tag that specifies how we want the json.Marshal func
		// to convert the struct to a json object
		// we are saying we have an error field  and the key to the string is error
		// {"error":""}
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{Error: msg})
}
