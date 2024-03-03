package sqlc

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

const connStr = "user=root password=1234 dbname=market sslmode=disable port=5435"
const Dbdriver = "postgres"

func OpenPsgtreConnection() (*sql.DB, error){
	db, err := sql.Open(Dbdriver, connStr)
	if err != nil {
		log.Printf("Failed to connect to database %v\n", err)
		os.Exit(1)
	}
	return db, nil
}

