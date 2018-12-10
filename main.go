package main

import (
	"fmt"
	"net/http"
)

// HelloWorld is exported
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":3000", nil)
}
