package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connect() *sql.DB {
	// Get database URL and auth token from environment variables
	dbUrl := os.Getenv("TURSO_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	url := "libsql://" + dbUrl + ".turso.io?authToken=" + authToken

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	//defer db.Close()
	println("Connect to database")
	return db
}
