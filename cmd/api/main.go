package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
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

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      api.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelDebug),
	}

	logger.Info("starting server", "address", server.Addr, "environment", config.env)
	err := server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func (svc *service) healthzHandler(w http.ResponseWriter, r *http.Request) {
	payload := capsule{
		"status":      "available",
		"environment": svc.config.env,
		"version":     version,
	}

	err := SendJSON(w, http.StatusOK, nil, payload)
	if err != nil {
		svc.Send500Error(w, r, err)
	}
}
