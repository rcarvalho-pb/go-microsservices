package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/go-authentication-service/data"
)

const webPort = "8080"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting Authentication service")
	app := &Config{}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		if err != nil {
			log.Panic(err)
		}
	}
}
