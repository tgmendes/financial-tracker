package handler

import (
	"context"
	"errors"
	"fmt"
	alphavantage2 "github.com/tgmendes/financial-tracker/pkg/alpha-vantage"
	dao2 "github.com/tgmendes/financial-tracker/pkg/dao"
	"github.com/tgmendes/financial-tracker/pkg/marshal"
	"golang.org/x/exp/slog"
	"golang.org/x/time/rate"
	"time"
)

const (
	// we can make 5 requests per minute - which means a request every 12 seconds
	rpsBurst   = 5
	rpsLimit   = 60 / rpsBurst * time.Second
	dailyLimit = 500
)

type StockFetcher struct {
	avClient     *alphavantage2.Client
	q            *dao2.Queries
	logger       *slog.Logger
	burstLimiter *rate.Limiter
	dailyLimiter *rate.Limiter
}

func NewStockFetcher(avApiKey string, q *dao2.Queries, log *slog.Logger) *StockFetcher {
	client := alphavantage2.NewClient(alphavantage2.BASE_URL, avApiKey)

	return &StockFetcher{
		avClient:     client,
		q:            q,
		logger:       log,
		burstLimiter: rate.NewLimiter(rate.Every(rpsLimit), rpsBurst),
		dailyLimiter: rate.NewLimiter(rate.Every(24*time.Hour), dailyLimit),
	}
}

func (s *StockFetcher) AllTimeData(ctx context.Context) error {
	symbols, err := s.q.ListSymbolIDs(ctx)
	if err != nil {
		return err
	}

	rateErrLimit := 5
	rateErrCount := 0
	maxIter := 5 * len(symbols)
	currIter := 0
	for len(symbols) > 0 {
		currIter += 1
		if currIter > maxIter {
			return fmt.Errorf("unable to fetch stock data within %d iterations", maxIter)
		}

		if rateErrCount > rateErrLimit {
			return fmt.Errorf("rate limit error count above threshold (%d)", rateErrLimit)
		}

		if err := s.burstLimiter.Wait(ctx); err != nil {
			return fmt.Errorf("error waiting for limiter: %w", err)
		}

		if !s.dailyLimiter.Allow() {
			return fmt.Errorf("exceeded daily rate limit: %w", err)
		}

		sym := symbols[0]
		symbols = symbols[1:]

		ts, err := s.avClient.DailyAdjusted(sym, "compact")
		if errors.Is(err, alphavantage2.ErrTooManyRequests) {
			symbols = append(symbols, sym)
			rateErrCount += 1
			continue
		}
		if err != nil {
			return err
		}

		err = s.batchStoreTSData(ctx, ts)
		if err != nil {
			fmt.Printf("unbale to store data for %s: %s\n", sym, err)
			continue
		}
	}

	return nil
}

func (s *StockFetcher) SeedAllTimeData(ctx context.Context) error {
	symbols, err := s.q.ListSymbolIDs(ctx)
	if err != nil {
		return err
	}

	rateErrLimit := 5
	rateErrCount := 0
	maxIter := 5 * len(symbols)
	currIter := 0
	for len(symbols) > 0 {
		currIter += 1
		if currIter > maxIter {
			return fmt.Errorf("unable to fetch stock data within %d iterations", maxIter)
		}

		if rateErrCount > rateErrLimit {
			return fmt.Errorf("rate limit error count above threshold (%d)", rateErrLimit)
		}

		if err := s.burstLimiter.Wait(ctx); err != nil {
			return fmt.Errorf("error waiting for limiter: %w", err)
		}

		if !s.dailyLimiter.Allow() {
			return fmt.Errorf("exceeded daily rate limit: %w", err)
		}

		sym := symbols[0]
		symbols = symbols[1:]

		ts, err := s.avClient.DailyAdjusted(sym, "compact")
		if errors.Is(err, alphavantage2.ErrTooManyRequests) {
			symbols = append(symbols, sym)
			rateErrCount += 1
			continue
		}
		if err != nil {
			return err
		}

		err = s.batchStoreTSData(ctx, ts)
		if err != nil {
			fmt.Printf("unbale to store data for %s: %s\n", sym, err)
			continue
		}
	}

	return nil
}

func (s *StockFetcher) storeTSData(ctx context.Context, ts *alphavantage2.TimeSeriesResponse) error {
	sdParams := make([]dao2.CreateStockDataParams, len(ts.TimeSeries))
	idx := 0
	for timestamp, row := range ts.TimeSeries {
		data, err := marshal.AVtoDAOTimeseries(row, timestamp, ts.Metadata.Symbol)
		if err != nil {
			return fmt.Errorf("unable to marshal timeseries: %w", err)
		}
		sdParams[idx] = data
		idx += 1
	}

	_, err := s.q.CreateStockData(ctx, sdParams)
	if err != nil {
		return fmt.Errorf("unable to create stock data: %w", err)
	}

	return nil
}

func (s *StockFetcher) batchStoreTSData(ctx context.Context, ts *alphavantage2.TimeSeriesResponse) error {
	sdParams := make([]dao2.BatchCreateStockDataParams, len(ts.TimeSeries))
	idx := 0
	for timestamp, row := range ts.TimeSeries {
		data, err := marshal.AVtoBatchDAOTimeseries(row, timestamp, ts.Metadata.Symbol)
		if err != nil {
			return fmt.Errorf("unable to marshal timeseries: %w", err)
		}
		sdParams[idx] = data
		idx += 1
	}

	res := s.q.BatchCreateStockData(ctx, sdParams)
	res.Exec(func(i int, err error) {
		if err != nil {
			fmt.Printf("unable to batch insert: %s\n", err)
			return
		}
		fmt.Printf("batch result size: %d\n", i)
	})
	return nil
}
