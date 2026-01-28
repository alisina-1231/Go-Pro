package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respodWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Respond: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respodWithJSON(w, code, errResponse{
		Error: msg,
	})

}

func respodWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json respose: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

