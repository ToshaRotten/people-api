package main

import (
	"log/slog"
	"net/http"
	"os"
	"people-api/internal/config"
	"people-api/internal/http-server/handlers/person/save"
	"people-api/internal/http-server/mw"
	"people-api/internal/storage/postgres"
	"people-api/utils/extended_slog"
)

// slog - logger
// postgres - storage
// chi - router

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig("./config/local.env")

	log := setupLogger(cfg.Env)

	log.Info("starting people-api service", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := postgres.New(cfg.Storage)
	if err != nil {
		log.Error("failed to init storage:", extended_slog.Error(err))
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/user/save", save.Save(log, storage))

	handler := mw.Logging(mux)

	srv := http.Server{
		Addr:         cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port,
		Handler:      handler,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	log.Error("server stopped: ", srv.ListenAndServe())
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		//log = slog.NewJSONHandler()
	case envProd:
		//log = slog.NewJSONHandler()
	}

	return log
}
