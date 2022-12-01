package main

import (
	"log"
	"net/http/httptest"
	"testing"
)

func Test_Trailer(t *testing.T) {
	srv := NewServer()

	r := httptest.NewRequest("GET", "/tests", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)

	rsp := w.Result()
	defer rsp.Body.Close()

	log.Printf("%+v", rsp.Trailer)
}
