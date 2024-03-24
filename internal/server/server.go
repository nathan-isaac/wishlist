package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"wishlist/internal/domain"
	"wishlist/internal/gateway"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type AdminUser struct {
	Username string
	Password string
}

type Server struct {
	port    int
	host    string
	db      *sql.DB
	ctx     context.Context
	queries *gateway.Queries
	domain  *domain.App
	admin   AdminUser
}

type Wishlist struct {
	Id          string
	Name        string
	Description string
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE_URL")

	db, err := sql.Open("sqlite3", database)

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
		domain: &domain.App{
			Queries: queries,
			Ctx:     ctx,
		},
		admin: AdminUser{
			Username: os.Getenv("ADMIN_USERNAME"),
			Password: os.Getenv("ADMIN_PASSWORD"),
		},
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
