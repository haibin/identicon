package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestIdenticonHandler(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/monster/{name}", identiconHandler)

	req, _ := http.NewRequest("GET", "http://example.com/monster/haha", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
