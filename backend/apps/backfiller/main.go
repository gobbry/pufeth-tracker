package main

import (
	"context"
	"database/sql"
	"log/slog"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobbry/puffering/libraries/comms"
	"github.com/gobbry/puffering/libraries/database/service"
)

const (
	MAINNET_RPC = "https://eth-pokt.nodies.app"
	MAINNET_WS  = "wss://ethereum-rpc.publicnode.com"
	DB_URL      = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	REDIS_URL   = "localhost:6379"
)

func main() {
	database, err := comms.NewTimescaleDB()
	if err != nil {
		slog.Error("unable to db connection", "err", err)
		return
	}
	defer database.Close()

	rd, err := comms.NewRedis()
	if err != nil {
		slog.Error("unable to redis connection", "err", err)
		return
	}
	defer rd.Close()

	c, err := ethclient.DialContext(context.Background(), MAINNET_RPC)
	if err != nil {
		slog.Error("unable to eth client", "err", err)
		return
	}
	defer c.Close()

	tx, err := database.BeginTx(context.Background(), &sql.TxOptions{ReadOnly: true})
	if err != nil {
		slog.Error("failed to begin transaction", "error", err)
		return
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(context.Background(), `
        DECLARE puffer_cursor CURSOR FOR
        SELECT block_number, block_timestamp_ms FROM puffer_conversion_rates
        WHERE chain = $1
        ORDER BY block_number DESC;
    `, "ethereum")
	if err != nil {
		slog.Error("failed to declare cursor", "error", err)
		return
	}

	jobs := make(chan int64, 1_000)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for blockNumber := range jobs {
				b, err := c.BlockByNumber(context.Background(), big.NewInt(blockNumber))
				if err != nil {
					slog.Error("unable to get block by number", "blockNumber", blockNumber, "err", err)
					continue
				}
				err = service.UpdateConversionRate(c, database, rd, blockNumber, int64(b.Time())*1_000, context.Background())
				if err != nil {
					slog.Error("unable to update conversion rate", "blockNumber", blockNumber, "err", err)
				}
			}
		}()
	}

	prevBlockNumber := int64(-1)
	for {
		rows, err := tx.QueryContext(context.Background(), `FETCH FORWARD 100 FROM puffer_cursor`)
		if err != nil {
			slog.Error("failed to fetch rows", "error", err)
			break
		}

		hasRows := false
		for rows.Next() {
			hasRows = true
			var blockNumber, blockTimestampMs int64
			if err := rows.Scan(&blockNumber, &blockTimestampMs); err != nil {
				slog.Error("failed to scan row", "error", err)
				continue
			}

			if prevBlockNumber > 0 && prevBlockNumber-blockNumber > 1 {
				slog.Info("curr block", "blockNum", blockNumber, "blockTimestampMs", blockTimestampMs)
				for missingBlock := blockNumber + 1; missingBlock < prevBlockNumber; missingBlock++ {
					slog.Warn("missing block detected", "blockNumber", missingBlock)
					jobs <- missingBlock
				}
			}

			prevBlockNumber = blockNumber
		}
		rows.Close()

		if !hasRows {
			break
		}
	}
	tx.Rollback()

	wg.Wait()
	close(jobs)
}
