package main

import (
	"database/sql"

	"github.com/rcarvalho-pb/go-authentication-service/data"
)

const webPort = "8080"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
