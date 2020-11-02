package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response represents data returned to the client
type Response struct {
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"errorMessage"`
}

// JSONResponse converts Response object to Json and return to the client
func JSONResponse(response interface{}, w http.ResponseWriter, log *log.Logger) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
