package main

import (
	"github.com/brookcui.github.io/routes"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	r := newRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newRouter() *mux.Router {
	tmpl = template.Must(template.ParseGlob("./templates/*"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler :=
		http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", routes.GetIndexPageData())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
