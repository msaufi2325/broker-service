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

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPortNumber),
		Handler: app.routes(),
	}

	// start http server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
