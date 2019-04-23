package main

import (
	/* External imports*/
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
	// "html/template"
	"github.com/gorilla/mux"

	/* My imports*/
	"github.com/waxmittmann/lilvue_backend/articles"
)

// Dumb trick to make it shut up about unused import. 
var _ = fmt.Printf

func spaHandler(w http.ResponseWriter, r *http.Request) {
// func SpaHandler(w http.ResponseWriter, r *http.Request) {
    // vars := mux.Vars(r)
    // w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Category: %v\n", vars["category"])
	
	// http.ServeFile(w, r, "./frontend/dist")
	http.ServeFile(w, r, "./frontend/dist/index.html")
}

// func SpaHandler(w http.ResponseWriter, r *http.Request) {

func main() {
	r := mux.NewRouter()

	// Serve Single Page App
	
	// Run / as file server
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// http.ServeFile(w, r, "./frontend/dist")
	// http.HandleFunc("/", spaHandler)

	// r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", articles.ArticlesCategoryHandler)
	// r.HandleFunc("/spa/", SpaHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
