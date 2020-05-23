package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := newRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}