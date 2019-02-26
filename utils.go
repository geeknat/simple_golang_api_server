package main

import (
	"net/http"
	"encoding/json"
	"os"
)

// response struct holds response payload
type response struct {
	Success int         `json:"success"`
	Message interface{} `json:"message"`
}

// respondWithJSON writes server to the client in JSON
func respondWithJSON(w http.ResponseWriter, httpStatus int, successCode int, payload interface{}) {

	formattedResponse := response{successCode, payload}

	response, _ := json.Marshal(formattedResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,Content-Disposition")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.WriteHeader(httpStatus)
	w.Write(response)
}

// getEnvironmentVariable fetches the value for specified key
// in .env file
func getEnvironmentVariable(key string) string {
	return os.Getenv(key)
}
