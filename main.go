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
	r.HandleFunc("/", HelloHandler)
	r.HandleFunc("/hello", HelloHandler).Methods("GET")
	return r
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world!")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}