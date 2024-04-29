package main

import (
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"log/slog"
	"os"
	"wishlist/internal/gateway"
	"wishlist/internal/server"
	"wishlist/schema"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := godotenv.Load()

	if err != nil {
		logger.Warn("Error loading .env file", "error", err)
	}

	db, err := gateway.NewConnection()

	if err != nil {
		panic(err)
	}

	goose.SetBaseFS(schema.EmbedMigrations)

	if err := goose.SetDialect("sqlite"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "."); err != nil {
		panic(err)
	}

	s, err := server.NewServer(db)

	if err != nil {
		logger.Error("Error initializing server", "error", err)
	}

	err = s.ListenAndServe()

	if err != nil {
		logger.Error("Error starting server", "error", err)
		os.Exit(1)
	}
}
