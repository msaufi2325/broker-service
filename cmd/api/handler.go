package main

import (
	"encoding/json"
	"net/http"
)

// jsonResponse represents the structure of the JSON response
type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Broker is the handler function for the broker endpoint
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Create a payload with the response data
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker endpoint",
	}

	// Convert the payload to JSON format
	out, _ := json.MarshalIndent(payload, "", "\t")

	// Set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)

	// Write the JSON response to the response writer
	w.Write(out)
}
