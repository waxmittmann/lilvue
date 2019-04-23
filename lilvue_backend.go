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
	// "articles_handler"
	// "github.com/mwittmann/lilvue_backend/articles"
	// "mwittmann/lilvue_backend/articles"
	// "lilvue_backend/articles"
	// "articles"
	"github.com/waxmittmann/lilvue_backend/articles"
)

// Dumb trick to make it shut up about unused import. 
var _ = fmt.Printf

// func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
//     vars := mux.Vars(r)
//     w.WriteHeader(http.StatusOK)
//     fmt.Fprintf(w, "Category: %v\n", vars["category"])
// }

func main() {
	r := mux.NewRouter()

	
	// r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", articles.ArticlesCategoryHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
