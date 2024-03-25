package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s *service) routes() http.Handler {
	mux := chi.NewRouter()

	// TODO: configure router middleware
	mux.Use(s.RecoverPanic)

	// Use custom NotFound and MethodNotAllowed handlers
	mux.NotFound(s.NotFoundHandler)
	mux.MethodNotAllowed(s.MethodNotAllowedHandler)

	// Routing
	mux.Get("/v1/healthz", s.healthzHandler)

	mux.Get("/v1/nfts/{nftid}", s.ShowNFTHandler())
	mux.Post("/v1/nfts", s.CreateNFTHandler())

	return mux
}
