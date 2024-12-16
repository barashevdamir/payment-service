package main

import (
	"golang.org/x/exp/slog"
	"os"
	"payment-service/internal/config"
	"payment-service/internal/lib/logger/handlers/slogpretty"
	"payment-service/internal/lib/logger/sl"
	"payment-service/internal/storage/postgresql"
	"payment-service/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting payment-service",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	switch cfg.StorageType {
	case "sqlite":
		_, err := sqlite.New(cfg.StoragePath)
		if err != nil {
			log.Error("failed to init storage", sl.Err(err))
			os.Exit(1)
		}
	case "postgres":
		_, err := postgresql.New(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Dbname, cfg.Postgres.Sslmode)
		if err != nil {
			log.Error("failed to init storage", sl.Err(err))
			os.Exit(1)
		}
	default:
		log.Error("unknown storage type: " + cfg.StorageType)
		os.Exit(1)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
