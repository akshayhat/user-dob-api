package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	connStr := "postgres://postgres:postgres2309@localhost:5433/user_dob_db?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	fmt.Println("Database connected successfully")
	return db
}
