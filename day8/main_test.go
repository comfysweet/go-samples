package main

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	r, err := http.Get("http://localhost/phones/Bill")
	if err != nil {
		t.Error(err)
	}
	if r.Status != "200 OK" {
		t.Error("TestGet failed")
	}
}
