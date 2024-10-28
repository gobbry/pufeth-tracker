// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
)

const getChains = `-- name: GetChains :many
SELECT DISTINCT chain
FROM puffer_conversion_rates
`

func (q *Queries) GetChains(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getChains)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var chain string
		if err := rows.Scan(&chain); err != nil {
			return nil, err
		}
		items = append(items, chain)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getHistoricalPufferRates = `-- name: GetHistoricalPufferRates :many
SELECT chain, block_number, block_timestamp_ms, total_asset, total_supply, conversion_rate, created_at
FROM puffer_conversion_rates
WHERE chain = $1
ORDER BY block_number DESC
LIMIT $2
`

type GetHistoricalPufferRatesParams struct {
	Chain string `json:"chain"`
	Limit int32  `json:"limit"`
}

func (q *Queries) GetHistoricalPufferRates(ctx context.Context, arg GetHistoricalPufferRatesParams) ([]PufferConversionRate, error) {
	rows, err := q.db.QueryContext(ctx, getHistoricalPufferRates, arg.Chain, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []PufferConversionRate{}
	for rows.Next() {
		var i PufferConversionRate
		if err := rows.Scan(
			&i.Chain,
			&i.BlockNumber,
			&i.BlockTimestampMs,
			&i.TotalAsset,
			&i.TotalSupply,
			&i.ConversionRate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLatestPufferRate = `-- name: GetLatestPufferRate :one
SELECT chain, block_number, block_timestamp_ms, total_asset, total_supply, conversion_rate, created_at
FROM puffer_conversion_rates
WHERE chain = $1
ORDER BY block_number DESC
LIMIT 1
`

func (q *Queries) GetLatestPufferRate(ctx context.Context, chain string) (PufferConversionRate, error) {
	row := q.db.QueryRowContext(ctx, getLatestPufferRate, chain)
	var i PufferConversionRate
	err := row.Scan(
		&i.Chain,
		&i.BlockNumber,
		&i.BlockTimestampMs,
		&i.TotalAsset,
		&i.TotalSupply,
		&i.ConversionRate,
		&i.CreatedAt,
	)
	return i, err
}

const insertPufferRate = `-- name: InsertPufferRate :one
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
    ($1, $2, $3, $4, $5, $6) RETURNING chain, block_number, block_timestamp_ms, total_asset, total_supply, conversion_rate, created_at
`

type InsertPufferRateParams struct {
	Chain            string `json:"chain"`
	BlockNumber      int64  `json:"block_number"`
	BlockTimestampMs int64  `json:"block_timestamp_ms"`
	TotalAsset       string `json:"total_asset"`
	TotalSupply      string `json:"total_supply"`
	ConversionRate   string `json:"conversion_rate"`
}

func (q *Queries) InsertPufferRate(ctx context.Context, arg InsertPufferRateParams) (PufferConversionRate, error) {
	row := q.db.QueryRowContext(ctx, insertPufferRate,
		arg.Chain,
		arg.BlockNumber,
		arg.BlockTimestampMs,
		arg.TotalAsset,
		arg.TotalSupply,
		arg.ConversionRate,
	)
	var i PufferConversionRate
	err := row.Scan(
		&i.Chain,
		&i.BlockNumber,
		&i.BlockTimestampMs,
		&i.TotalAsset,
		&i.TotalSupply,
		&i.ConversionRate,
		&i.CreatedAt,
	)
	return i, err
}