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

// func serveStatic(router *mux.Router, staticDirectory string) {
// 	staticPaths := map[string]string{
// 		"styles":           staticDirectory + "/css/",
// 		// "bower_components": staticDirectory + "/bower_components/",
// 		// "images":           staticDirectory + "/images/",
// 		"scripts":          staticDirectory + "/js/",
// 	}
// 	for pathName, pathValue := range staticPaths {
// 		pathPrefix := "/" + pathName + "/"
// 		router.PathPrefix(pathPrefix).Handler(http.StripPrefix(pathPrefix,
// 			http.FileServer(http.Dir(pathValue))))
// 	}

// 	// staticDirectory := "/frontend/dist/"
// 	ServeStatic(router, staticDirectory)
// }
	// router := mux.NewRouter()

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/", articles.ArticlesCategoryHandler)


	// Serve Single Page App
	
	// Run / as file server
	// fs := http.FileServer(http.Dir("./frontend/dist"))
	// http.Handle("/", fs)

	// serveStatic(r, "/frontend/dist")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/dist")))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/dist"))))

	// http.ServeFile(w, r, "./frontend/dist")
	// http.HandleFunc("/", spaHandler)

	// r.HandleFunc("/products/{key}", ProductHandler)


	// r.HandleFunc("/products/{key}", ProductHandler)
	// r.HandleFunc("/spa/", SpaHandler)
	// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8080",
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Fatal(srv.ListenAndServe())
}
