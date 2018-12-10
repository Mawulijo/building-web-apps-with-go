package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/unrolled/render"
)

// Book is exported
type Book struct {
	Title  string `json:"title"`
	Author string `json:"autthor"`
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{
		"Bulding Web Apps With Go",
		"Jeremy Saenz",
	}
	filePath := path.Join("templates", "index.html")
	templ, err := template.ParseFiles(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := templ.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	r := render.New(render.Options{})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome, visit sub pages now."))
	})
	mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
		r.Data(w, http.StatusOK, []byte("Some binary data here."))
	})

	mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
	})
	mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
		name := "Joshua Agbeku"
		// Assumes you have a template in ./templates called "example.tmpl"
		// $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
		r.HTML(w, http.StatusOK, "example", name)
	})
	http.ListenAndServe(":8080", mux)
}
