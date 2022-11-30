package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func TestIDHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-test-id", "test")
		//_ = json.NewEncoder(w).Encode([]string{"Test"})
		h(w, r)

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
			w.Header().Add("x-inner-id-", "inner")
		}
	}
}

func InnerIDHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-inner-id-", "inner")
		h(w, r)
		fmt.Println()
	}
}

type ResponseWriterWrapper struct {
	http.ResponseWriter
	body       bytes.Buffer
	statusCode int
}
