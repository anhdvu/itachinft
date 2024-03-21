package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

const version = "1.0.0"

type config struct {
	env  string
	port int
}

type service struct {
	logger *slog.Logger
	config config
}

func main() {
	var config config

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	api := &service{
		logger: logger,
		config: config,
	}

	mux := chi.NewRouter()
	mux.Get("/v1/healthz", api.healthzHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info("starting server", "address", server.Addr, "environment", config.env)

	err := server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func (svc *service) healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", svc.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
