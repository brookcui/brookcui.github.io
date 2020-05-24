package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var homeTmpl *template.Template

func main() {
	r := newRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	homeTmpl = template.Must(template.ParseFiles("./templates/index.html"))

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler :=
		http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := homeTmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
