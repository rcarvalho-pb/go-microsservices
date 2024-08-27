package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rcarvalho-pb/go-authentication-service/data"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"

type Config struct {
	DB     *sql.DB
	Models *data.Models
}

func main() {
	log.Println("Starting Authentication service")

	db := initDB()

	app := &Config{
		DB:     db,
		Models: data.New(db),
	}

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

func initDB() *sql.DB {
	db := connectToDB()
	if db == nil {
		log.Panic("Couldn't connect to DB")
	}

	return db
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	count := 0

	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("DB not ready yet...")
			count++
		} else {
			log.Println("Connected to DB")
			return conn
		}

		if count >= 10 {
			return nil
		}

		log.Println("Backing off 2 sec...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
