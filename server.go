package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Pinger interface {
	PingContext(ctx context.Context) error
}

type server struct {
	router *chi.Mux
}

func (s *server) routes(p ...Pinger) {
	s.router.Get("/tests", TestIDHandler(s.HandlerGetAll()))
}

func (s *server) HandlerGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]string{"Success"})
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func NewServer(p ...Pinger) *server {
	s := server{
		router: chi.NewRouter(),
	}
	s.routes(p...)
	return &s
}
