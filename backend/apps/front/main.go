package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	front "github.com/gobbry/puffering/apps/front/rate"
	"github.com/gobbry/puffering/libraries/comms"
)

func main() {
	database, err := comms.NewTimescaleDB()
	if err != nil {
		slog.Error("unable to db connection", "err", err)
		return
	}
	defer database.Close()

	rd, err := comms.NewRedis()
	if err != nil {
		slog.Error("unable to redis connection", "err", err)
		return
	}
	defer rd.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rm, err := front.NewRateManager(rd, database)
	if err != nil {
		slog.Error("unable to front manager", "err", err)
		return
	}
	err = rm.Start(ctx)
	if err != nil {
		slog.Error("unable to start front manager", "err", err)
		return
	}
	slog.Info("started front manager...")

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/rate-series", func(w http.ResponseWriter, r *http.Request) {
		rates := rm.GetRates()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"rates": rates,
			"count": len(rates),
		})
	})

	slog.Info("starting server...")
	http.ListenAndServe(":3001", r)
}
