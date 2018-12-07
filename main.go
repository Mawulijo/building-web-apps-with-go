package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

// GenerateMarkdown is exported
func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	w.Write(markdown)
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.ListenAndServe(":"+port, nil)
}
