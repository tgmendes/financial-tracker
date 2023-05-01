package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tgmendes/financial-tracker/pkg/dao"
	"golang.org/x/exp/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))

	connStr := os.Getenv("DATABASE_URL")

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

	_, err = q.CreateSymbol(ctx, dao.CreateSymbolParams{
		ID:       "VWRL.L",
		Type:     "ETF",
		Name:     pgtype.Text{String: "Vanguard FTSE All-World UCITS ETF", Valid: true},
		Exchange: pgtype.Text{String: "LSE", Valid: true},
	})
	if err != nil {
		logger.Error("unable to create symbol", "error", err)
		os.Exit(1)
	}

}
