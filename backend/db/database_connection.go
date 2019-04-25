package db

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

// Todo: from config
var user = "postgres"
var password = "test"
var host = "127.0.0.1:5432"
var dbName = "lilvue"
var connStr = "postgres://" + user + ":" + password + "@" + host +"/" + dbName + "?sslmode=disable"

func Ping() {
	print(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Ping Success")
}
