package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	err := godotenv.Load("env.local")
	if err != nil {
		panic(err)
	}

	connStr := os.Getenv("POSTGRES_URL")

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	queryCreateHypertable := `SELECT create_hypertable('stock_data', 'time');`
	_, err = dbpool.Exec(ctx, queryCreateHypertable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create the `stock_data` hypertable: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully created hypertable `stock_data`")
}
