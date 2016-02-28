package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyRouterAndHandler(t *testing.T) {
	svr := httptest.NewServer(newHelloServer())
	defer svr.Close()

	res, err := http.Get(svr.URL + "/hello/world")

	if err != nil {
		t.Fatal("Problem calling hello server", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Error("Expected a 200 but got", res.StatusCode)
	}

	// etc...
}
