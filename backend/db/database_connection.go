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

func create_db() *sql.DB {
	print(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Ping() {
	db := create_db()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Ping Success")
}

func Query() { 
	db := create_db()

	var (
		id int
		name string
		email string
	)
	rows, err := db.Query("SELECT UserId, Name, Email FROM Users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
