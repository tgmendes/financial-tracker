package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	connStr := os.Getenv("DATABASE_URL")
	apiKey := os.Getenv("AV_API_KEY")

	if apiKey == "" {
		logger.Error("Alpha Vantage API key not found")
	}

	if connStr == "" {
		slog.Error("DB URL not set")
		os.Exit(1)
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		slog.Error("unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	cmd, err := dbpool.Exec(ctx, "SELECT 1")
	if err != nil {
		logger.Error("selecting", "error", err)
	}
	logger.Info("got from select", "msg", cmd.String())
	//
	//q := dao.New(dbpool)
	//
	//h := handler.NewStockFetcher(apiKey, q)
	//err = h.AllTimeData(ctx)
	//if err != nil {
	//	panic(err)
	//}
}
