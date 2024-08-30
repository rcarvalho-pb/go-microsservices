package main

import (
	"fmt"
	"log"
	"net/http"
)

var webPort = "80"

type Config struct{}

func main() {

	app := Config{}

	log.Printf("Starting server on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
