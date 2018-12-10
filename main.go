package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

// MyMiddleware is exported
func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there...")
	if r.URL.Query().Get("password") == "secret123" {
		next(w, r)
	} else {
		http.Error(w, "Not Authorized", 401)
	}
	log.Println("Logging on the way back...")
}

func main() {
	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)
	n.Run(":8080")
}
