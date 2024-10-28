-- name: InsertPufferRate :one
INSERT INTO
    puffer_conversion_rates (
        chain,
        block_number,
        block_timestamp_ms,
        total_asset,
        total_supply,
        conversion_rate
    )
VALUES
    ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetLatestPufferRate :one
SELECT *
FROM puffer_conversion_rates
WHERE chain = $1
ORDER BY block_number DESC
LIMIT 1;

-- name: GetChains :many
SELECT DISTINCT chain
FROM puffer_conversion_rates;


-- name: GetHistoricalPufferRates :many
SELECT *
FROM puffer_conversion_rates
WHERE chain = $1
ORDER BY block_number DESC
LIMIT $2;