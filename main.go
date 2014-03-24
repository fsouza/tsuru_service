package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func main() {
	m := pat.New()
	m.Post("/resources", HttpHandler(AddInstance))
	http.Handle("/", m)
	http.ListenAndServe(":8080", nil)
}
