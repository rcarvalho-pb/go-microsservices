package data

import (
	"database/sql"
	"time"
)

type Models struct {
	User *User
}

var dbTimeout = 3 * time.Second

var db *sql.DB

func New(dbPool *sql.DB) *Models {
	db = dbPool

	return &Models{
		User: &User{},
	}
}
