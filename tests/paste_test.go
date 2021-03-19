package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/112RG/Curator/testing"
)

func TestMainRoute(t *testing.T) {
	router := routes.Build()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
}
