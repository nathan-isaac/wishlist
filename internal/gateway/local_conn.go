//go:build local

package gateway

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func NewConnection() (*sql.DB, error) {
	db, err := sql.Open("libsql", os.Getenv("DATABASE_URL"))

	if err != nil {
		return nil, fmt.Errorf("failed to open db %s", err)
	}

	return db, nil
}
