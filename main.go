package main

import (
	"encoding/json"
	"net/http"
)

// Book is exported
type Book struct {
	Title string `json:"title"`
	Autor string `json:"autthor"`
}

func showBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{
		"Bulding Web Apps With Go",
		"Jeremy Saenz",
	}
	books, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(books)
}

func main() {
	http.HandleFunc("/", showBooks)
	http.ListenAndServe(":8080", nil)
}
