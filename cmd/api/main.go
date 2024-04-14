package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"wishlist/internal/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		slog.Warn("Error loading .env file", "error", err)
	}

	s := server.NewServer()

	err = s.ListenAndServe()

	if err != nil {
		slog.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
