package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/112RG/Curator/routes"
	_ "github.com/112RG/Curator/testing"
	"github.com/stretchr/testify/assert"
)

func TestMainRoute(t *testing.T) {
	router := routes.Build()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
}
