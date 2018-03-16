package main

import (
	"fmt"
	"net/http"
)

type HomePageHandler struct{}

func (h *HomePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.Handle("/", &HomePageHandler{})

	http.ListenAndServe(":3000", nil)
}
