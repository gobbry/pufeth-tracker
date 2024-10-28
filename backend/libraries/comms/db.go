package comms

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewTimescaleDB() (*sql.DB, error) {
	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		return nil, fmt.Errorf("database env not set")
	}

	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
