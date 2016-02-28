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
	name := mux.Vars(r)["name"]
	w.Write([]byte("Hello, " + name))
}
