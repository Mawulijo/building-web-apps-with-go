package main

import (
	"html/template"
	"net/http"
	"path"
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
	http.HandleFunc("/", showBooks)
	http.ListenAndServe(":8080", nil)
}
