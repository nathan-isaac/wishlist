package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"whishlist/internal/gateway"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	port    int
	host    string
	db      *sql.DB
	ctx     context.Context
	queries *gateway.Queries
}

type Wishlist struct {
	Id          string
	Name        string
	Description string
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")

	db, err := sql.Open("sqlite3", "./wishlist.db")

	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	queries := gateway.New(db)

	NewServer := &Server{
		port:    port,
		host:    host,
		db:      db,
		ctx:     ctx,
		queries: queries,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", NewServer.host, NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
