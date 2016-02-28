package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func IgnoreTestMyHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/chris", nil)
	res := httptest.NewRecorder()

	helloHandler(res, req)

	if res.Body.String() == "Hello, world" {
		t.Error("Fail! It should not use the default, it should see Chris!")
	}
}

func TestMyRouterAndHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hello/chris", nil)
	res := httptest.NewRecorder()
	newHelloServer().ServeHTTP(res, req)

	if res.Body.String() != "Hello, chris" {
		t.Error("Expected hello Chris but got ", res.Body.String())
	}
}
