-- name: CreateStockData :copyfrom
INSERT INTO stock_data (time, symbol_id, open, high, low, close, adjusted_close, volume, dividend_amount,
                        split_coefficient)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: BatchCreateStockData :batchexec
INSERT INTO stock_data (time, symbol_id, open, high, low, close, adjusted_close, volume, dividend_amount,
                        split_coefficient)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT do nothing ;