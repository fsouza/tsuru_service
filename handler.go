package main

import (
	"net/http"
)

type HttpHandler func(http.ResponseWriter, *http.Request) error

func (fn HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fn(w, r)
}
