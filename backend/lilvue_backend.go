package main

import (
	/* External imports*/
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"time"

	/* My imports*/
	"github.com/waxmittmann/lilvue_backend/articles"
	"github.com/waxmittmann/lilvue_backend/db"
)

// Dumb trick to make it shut up about unused import. 
var _ = fmt.Printf

func main() {
	setupDb()
	setupHttp()
}

// Set up and test db connection
func setupDb() {
	db.Ping()
	db.Query()
}

// Set up routing and start server
func setupHttp() {
	// Check DB connection
	r := mux.NewRouter()

	// Serve API endpoints
	r.HandleFunc("/articles/{category}/", articles.ArticlesCategoryHandler)

	// Serve Single Page App
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../frontend/dist")))

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8080",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}
