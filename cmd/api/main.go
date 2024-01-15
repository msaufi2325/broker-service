package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPortNumber = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPortNumber)

	// Define the HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPortNumber),
		Handler: app.routes(),
	}

	// Start the HTTP server
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
