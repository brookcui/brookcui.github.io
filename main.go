package main

import (
	models "github.com/brookcui.github.io/models"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseGlob("./templates/*"))

func main() {
	r := newRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", makeHandler(indexHandler)).Methods("GET")
	r.HandleFunc("/about", makeHandler(aboutHandler)).Methods("GET")

	staticFileDirectory := http.Dir("./static/")
	staticFileHandler :=
		http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	return r
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(about)?$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		if m[0] == "/" {
			fn(w, r, "index")
			return
		}
		fn(w, r, m[1])
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, title, nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request, title string) {
	renderTemplate(w, title, nil)
}
