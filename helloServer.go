package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newHelloServer() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", helloHandler)
	return r
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name, exists := mux.Vars(r)["name"]

	if !exists {
		name = "world"
	}

	w.Write([]byte("Hello, " + name))
}
