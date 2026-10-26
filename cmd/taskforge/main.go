package main

import (
	"context"
	"errors"
	//"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"log/slog"
	"net/http"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Ajouter un channel pour éviter que le serveur ne se ferme avec main
	serverErr := make(chan error, 1)
	go func() {
		slog.Info("Server started", "port", port)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		log.Println("Serveur: %w", err)
	case <-ctx.Done():
		slog.Info("Signal d'arrêt reçu, arrêt gracieux...")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Arrêt gracieux: %w", err)
	}
	slog.Info("arrêt propre terminé")
}
