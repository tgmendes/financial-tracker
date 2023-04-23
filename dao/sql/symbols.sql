-- name: CreateSymbol :one
INSERT INTO symbols (id, type, name, exchange)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListSymbolIDs :many
SELECT id FROM symbols
ORDER BY id;