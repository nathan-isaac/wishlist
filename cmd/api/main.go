package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"wishlist/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()

	if err != nil {
		logger.Warn("Error loading .env file", "error", err)
	}

	s, err := server.NewServer()

	if err != nil {
		logger.Error("Error initializing server", "error", err)
	}

	err = s.ListenAndServe()

	if err != nil {
		logger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
