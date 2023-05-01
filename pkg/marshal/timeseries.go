package marshal

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
	"github.com/tgmendes/financial-tracker/pkg/alpha-vantage"
	dao2 "github.com/tgmendes/financial-tracker/pkg/dao"
	"time"
)

func AVtoDAOTimeseries(data alphavantage.TimeSeries, timestamp, symbolID string) (dao2.CreateStockDataParams, error) {
	tsTime, err := time.Parse("2006-01-02", timestamp)
	if err != nil {
		return dao2.CreateStockDataParams{}, fmt.Errorf("unable to parse time: %w", err)
	}

	open := decimal.NewFromFloat(data.Open)
	high := decimal.NewFromFloat(data.High)
	low := decimal.NewFromFloat(data.Low)
	closeVal := decimal.NewFromFloat(data.Close)
	adjustedClose := decimal.NewFromFloat(data.AdjustedClose)
	dividendAmt := decimal.NewFromFloat(data.DividendAmount)

	return dao2.CreateStockDataParams{
		Time: pgtype.Timestamptz{
			Time:  tsTime,
			Valid: true,
		},
		SymbolID: symbolID,
		Open: pgtype.Numeric{
			Int:   open.Coefficient(),
			Exp:   open.Exponent(),
			Valid: true,
		},
		High: pgtype.Numeric{
			Int:   high.Coefficient(),
			Exp:   high.Exponent(),
			Valid: true,
		},
		Low: pgtype.Numeric{
			Int:   low.Coefficient(),
			Exp:   low.Exponent(),
			Valid: true,
		},
		Close: pgtype.Numeric{
			Int:   closeVal.Coefficient(),
			Exp:   closeVal.Exponent(),
			Valid: true,
		},
		AdjustedClose: pgtype.Numeric{
			Int:   adjustedClose.Coefficient(),
			Exp:   adjustedClose.Exponent(),
			Valid: true,
		},
		Volume: pgtype.Int8{
			Int64: data.Volume,
			Valid: true,
		},
		DividendAmount: pgtype.Numeric{
			Int:   dividendAmt.Coefficient(),
			Exp:   dividendAmt.Exponent(),
			Valid: true,
		},
		SplitCoefficient: pgtype.Float8{
			Float64: data.SplitCoefficient,
			Valid:   true,
		},
	}, nil
}

func AVtoBatchDAOTimeseries(data alphavantage.TimeSeries, timestamp, symbolID string) (dao2.BatchCreateStockDataParams, error) {
	tsTime, err := time.Parse("2006-01-02", timestamp)
	if err != nil {
		return dao2.BatchCreateStockDataParams{}, fmt.Errorf("unable to parse time: %w", err)
	}

	open := decimal.NewFromFloat(data.Open)
	high := decimal.NewFromFloat(data.High)
	low := decimal.NewFromFloat(data.Low)
	closeVal := decimal.NewFromFloat(data.Close)
	adjustedClose := decimal.NewFromFloat(data.AdjustedClose)
	dividendAmt := decimal.NewFromFloat(data.DividendAmount)

	return dao2.BatchCreateStockDataParams{
		Time: pgtype.Timestamptz{
			Time:  tsTime,
			Valid: true,
		},
		SymbolID: symbolID,
		Open: pgtype.Numeric{
			Int:   open.Coefficient(),
			Exp:   open.Exponent(),
			Valid: true,
		},
		High: pgtype.Numeric{
			Int:   high.Coefficient(),
			Exp:   high.Exponent(),
			Valid: true,
		},
		Low: pgtype.Numeric{
			Int:   low.Coefficient(),
			Exp:   low.Exponent(),
			Valid: true,
		},
		Close: pgtype.Numeric{
			Int:   closeVal.Coefficient(),
			Exp:   closeVal.Exponent(),
			Valid: true,
		},
		AdjustedClose: pgtype.Numeric{
			Int:   adjustedClose.Coefficient(),
			Exp:   adjustedClose.Exponent(),
			Valid: true,
		},
		Volume: pgtype.Int8{
			Int64: data.Volume,
			Valid: true,
		},
		DividendAmount: pgtype.Numeric{
			Int:   dividendAmt.Coefficient(),
			Exp:   dividendAmt.Exponent(),
			Valid: true,
		},
		SplitCoefficient: pgtype.Float8{
			Float64: data.SplitCoefficient,
			Valid:   true,
		},
	}, nil
}
