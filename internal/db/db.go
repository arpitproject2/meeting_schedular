package db

import (
	"database/sql"
	//"github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "user=postgres dbname=meetingdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}
