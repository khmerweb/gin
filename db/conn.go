package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connect() *sql.DB {
	// Get database URL and auth token from environment variables
	dbUrl := os.Getenv("TURSO_DATABASE_URL_GO")
	authToken := os.Getenv("TURSO_AUTH_TOKEN_GO")
	if dbUrl == "" || authToken == "" {
		log.Fatal("TURSO_DATABASE_URL and TURSO_AUTH_TOKEN environment variables must be set.")
	}

	db, err := sql.Open("libsql", dbUrl+"?authToken="+authToken)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db: %s", err)
		os.Exit(1)
	} else {
		println("Connect to database")
	}

	return db
}
