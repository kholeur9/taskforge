package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	if err := run(); err != nil {
		slog.Error("Arrêt de l'application", "error", err)
		os.Exit(1)
	}
}

func run() error {
	config := Config{
		Port:            getEnv("PORT", "8080"),
		ShutdownTimeout: 15 * time.Second,
	}
	mux := http.NewServeMux()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: mux,
	}
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("Serveur démarré", "port", config.Port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			// On fait remonter l'erreur dans le channel pout le transmettre à la goroutine principale.
			serverErr <- err
		}
	}()
	var runErr error
	select {
	case runErr = <-serverErr:
		return fmt.Errorf("Serveur: %w", runErr)
	case <-ctx.Done():
		slog.Info("Signal d'arrêt reçu, arrêt gracieux...")
	}
	shutdownCtx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		var shutdownErr, closeErr error
		if errClose := server.Close(); errClose != nil {
			closeErr = fmt.Errorf("Arrêt immédiat du serveur: %w", errClose)
		}
		shutdownErr = fmt.Errorf("Arrêt gracieux: %w", err)
		return errors.Join(shutdownErr, closeErr)
	}
	slog.Info("Arrêt propre terminé")
	return nil
}
