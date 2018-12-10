package main

import (
	"net/http"

	"github.com/unrolled/render"
)

type MyController struct {
	AppController
	*render.Render
}

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to a controller.
type Action func(w http.ResponseWriter, r *http.Request) error

// This is our Base Controller
type AppController struct{}

// The action function helps with error handling in a controller
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":8080", c.Action(c.Location))
}
