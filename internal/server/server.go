package server

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	port int
	db   *sqlx.DB
}

type Wishlist struct {
	Id          string
	Name        string
	Description string
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db := sqlx.MustConnect("sqlite3", ":memory:")

	db.MustExec(`
CREATE TABLE wishlist (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT
);
CREATE TABLE wishlist_item (
    id TEXT PRIMARY KEY,
    wishlist_id TEXT,
    link TEXT NOT NULL,
    description TEXT,
    bought BOOLEAN DEFAULT FALSE,
    FOREIGN KEY(wishlist_id) REFERENCES wishlist(id)
);
CREATE TABLE purchase (
    id TEXT PRIMARY KEY,
    wishlist_item_id TEXT,
    bought_at DATETIME NOT NULL,
    FOREIGN KEY(wishlist_item_id) REFERENCES wishlist_item(id)
);
`)
	wId, err := GenerateId()

	if err != nil {
		panic(err)
	}

	db.MustExec(`
INSERT INTO wishlist (id, name, description)
VALUES (?, ?, ?)
`, wId, "Name", "Description")

	wishlists := []Wishlist{}

	err = db.Select(&wishlists, "SELECT id, name, description FROM wishlist")

	if err != nil {
		panic(err)
	}

	fmt.Println(wishlists)

	NewServer := &Server{
		port: port,
		db:   db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
