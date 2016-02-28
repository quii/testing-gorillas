package main

import (
	"io/ioutil"
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
	svr := httptest.NewServer(newHelloServer())
	defer svr.Close()

	res, err := http.Get(svr.URL + "/hello/chris")

	if err != nil {
		t.Fatal("Problem calling hello server", err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if string(greeting) != "Hello, chris" {
		t.Error("Expected hello Chris but got ", greeting)
	}
}
