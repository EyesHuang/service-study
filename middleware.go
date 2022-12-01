package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func TestIDHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := newResponseWriterWrapper(w)

		w.Header().Add("x-test-id", "test")
		rw.Header().Add("Trailer", "x-test-id-v2")

		h(w, r)

		rw.Header().Set("x-test-id-v2", "hello")
		_ = json.NewEncoder(&rw).Encode([]string{"Test"})
	}
}

func InnerIDHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-inner-id", "inner")
		h(w, r)
		fmt.Println()
	}
}

type ResponseWriterWrapper struct {
	w          *http.ResponseWriter
	body       *bytes.Buffer
	statusCode *int
}

func (rw ResponseWriterWrapper) Header() http.Header {
	return (*rw.w).Header()
}

func (rw ResponseWriterWrapper) WriteHeader(statusCode int) {
	(*rw.w).WriteHeader(statusCode)
}

func (rw ResponseWriterWrapper) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return (*rw.w).Write(b)
}

func newResponseWriterWrapper(w http.ResponseWriter) ResponseWriterWrapper {
	var body bytes.Buffer
	var statusCode int

	return ResponseWriterWrapper{
		w:          &w,
		body:       &body,
		statusCode: &statusCode,
	}
}
