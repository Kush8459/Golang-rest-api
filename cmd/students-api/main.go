package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kush8459/students-api/internal/config"
	"github.com/kush8459/students-api/internal/http/handlers/student"
	"github.com/kush8459/students-api/internal/storage/sqlite"
)

func main() {
	// load config

	cfg := config.MustLoad()

	// database setup

	storage, dbErr := sqlite.New(cfg)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shut down server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")
}
