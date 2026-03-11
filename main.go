package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Wahbi8/PM_Golang/apis"
	"github.com/Wahbi8/PM_Golang/logger"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /email/invoice", apis.SendEmailApi)
	mux.HandleFunc("GET /health", healthCheck)

	server := &http.Server{
		Addr:    ":1212",
		Handler: mux,
	}

	go func() {
		logger.Log.Info().Str("port", "1212").Msg("Starting HTTP server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("HTTP server failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	logger.Log.Info().Msg("Server exited")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
