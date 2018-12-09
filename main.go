package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// HomeHandler is exported
func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "Home")
}

// PostsIndexHandler is exported
func PostsIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "posts index")
}

// PostsCreateHandler is exported
func PostsCreateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "posts create ")
}

// PostShowHandler is exported
func PostShowHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(w, "showing post", id)
}

// PostUpdateHandler is exported
func PostUpdateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "post update")
}

// PostDeleteHandler is exported
func PostDeleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "post delete")
}

// PostEditHandler is exported
func PostEditHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, "post edit")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
<<<<<<< HEAD
	// for index page
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.ListenAndServe(":"+port, nil)
=======

	r := httprouter.New()
	r.GET("/", HomeHandler)

	//Posts collection
	r.GET("/posts", PostsIndexHandler)
	r.POST("/posts", PostsCreateHandler)

	// Posts singular
	r.GET("/posts/:id", PostShowHandler)
	r.PUT("/posts/:id", PostUpdateHandler)
	r.DELETE("/posts/:id", PostDeleteHandler)
	r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":"+port, r)
>>>>>>> URL-Routing
}
