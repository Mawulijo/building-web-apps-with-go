package main

import (
	"net/http"
	"os"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// GenerateMarkdown is exported
func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.Run([]byte(r.FormValue("body")))
	w.Write(markdown)
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// For home
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.ListenAndServe(":"+port, nil)
}
