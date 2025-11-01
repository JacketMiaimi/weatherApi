package main

import (
	l "log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go.mod/internal/config"
	"go.mod/internal/http-server/get"
	"go.mod/internal/redis"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.LoadConfig("PATH_CONFIG")

	// logs
	log := setupLogger(cfg.Env)

	log = log.With(slog.String("env", cfg.Env))

	log.Info("starting weather-api")
	log.Info("initializing server", slog.String("address", cfg.Adress))

	// redis 
	if err := redis.InitRedis(); err != nil {
		l.Fatal("redis is not connected", err)
	}
	

	// router & middleware
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/get/{city}", get.New(log))

	log.Info("starting server", slog.String("address", cfg.Adress))

	srv := &http.Server{
		Addr:         cfg.Adress,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.TimeOut,
		WriteTimeout: cfg.HTTPServer.TimeOut,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Info("server started")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
