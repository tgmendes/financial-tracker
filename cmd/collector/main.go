package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tgmendes/financial-tracker/pkg/dao"
	"github.com/tgmendes/financial-tracker/pkg/handler"
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
		logger.Error("DB URL not set")
		os.Exit(1)
	}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		logger.Error("unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	q := dao.New(dbpool)

	h := handler.NewStockFetcher(apiKey, q, logger)
	err = h.AllTimeData(ctx)
	if err != nil {
		logger.Error("unable to collect data", "error", err)
		os.Exit(1)
	}
}
