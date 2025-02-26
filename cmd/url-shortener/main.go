package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sglite"

	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)


func main() {
	//TODO config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// init logger: slog
	log := setupLogger(cfg.Env)

	log.Info("Star url-shortered", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	// init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		// slog.StringValue(err.Error())
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	defer storage.Close()
	// init router: chi, "chi render"

	// init: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
