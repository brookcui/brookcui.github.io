package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("./templates/*"))

func main() {
	r := newRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler :=
		http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}
