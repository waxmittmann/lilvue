package articles

import (
	// "fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// package main

// import (
//   "encoding/json"
//   "net/http"
// )

// type Profile struct {
//   Name    string
//   Hobbies []string
// }

// func main() {
//   http.HandleFunc("/", foo)
//   http.ListenAndServe(":3000", nil)
// }

// func foo(w http.ResponseWriter, r *http.Request) {
//   profile := Profile{"Alex", []string{"snowboarding", "programming"}}

//   js, err := json.Marshal(profile)
//   if err != nil {
//     http.Error(w, err.Error(), http.StatusInternalServerError)
//     return
//   }

//   w.Header().Set("Content-Type", "application/json")
//   w.Write(js)
// }

type Profile struct {
	Name    string
	Hobbies []string
  }
  
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    // w.WriteHeader(http.StatusOK)
	// // fmt.Fprintf(w, "Category: %v\n", vars["category"])
	// fmt.Fprintf(w, "{ 'category': '%v' }", vars["category"])

	profile := Profile{"Alex", []string{ vars["category"], "programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
