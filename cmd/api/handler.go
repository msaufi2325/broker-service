package main

import (
	"net/http"
)

// Broker is the handler function for the broker endpoint
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Create a payload with the response data
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker endpoint",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
