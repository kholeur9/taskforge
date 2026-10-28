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

func main() {
	if err := run(); err != nil {
		slog.Error("arrêt avec erreur", "error", err)
		os.Exit(1)
	}
}

func run() error {
	port := "8080"
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	serverErr := make(chan error, 1)
	go func() {
		slog.Info("Serveur démarré", "port", port)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <-err
		}
	}()

	var runErr error
	select {
	case runErr = <-serverErr:
		return runErr
	case <-ctx.Done():
		slog.Info("signal d'arrêt reçu, arrêt gracieux...")
	}

	if err := shutdownSrv(srv, 30*time.Second); err != nil {
		return err
	}

	slog.Info("arrêt propre terminé")
	return runErr
}

func shutdownSrv(srv *http.Server, timeShutdown time.Duration) error {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), timeShutdown)
	defer cancel()
	var shutdownErr error
	if err := srv.Shutdown(shutdownCtx); err != nil {
		if errClose := srv.Close(); errClose != nil {
			shutdownErr = errClose
		}
		shutdownErr = err
	}
	return fmt.Errorf("arrêt gracieux: %w", shutdownErr)
}