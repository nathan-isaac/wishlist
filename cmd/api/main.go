package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"whishlist/internal/server"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := server.NewServer()

	err = s.ListenAndServe()

	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
