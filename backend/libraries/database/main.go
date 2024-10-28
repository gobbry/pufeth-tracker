package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

const schemaPath = "./libraries/database/schema.sql"

func setupDatabase(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Read schema file
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file: %w", err)
	}

	// Execute schema
	if _, err := db.ExecContext(ctx, string(schema)); err != nil {
		return nil, fmt.Errorf("error creating schema: %w", err)
	}

	return db, nil
}

func main() {
	ctx := context.Background()

	db, err := setupDatabase(ctx)
	if err != nil {
		slog.Error("failed to setup database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("database setup completed successfully")

	// Rest of your application code...
}
