package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	
)

var db *sql.DB



func TestMain(m *testing.M) {
	
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to connect to database %v\n", err)
		os.Exit(1)
	}

	defer db.Close()

	code := m.Run()

	os.Exit(code)

}

