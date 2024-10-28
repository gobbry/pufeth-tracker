package main

import (
	"context"
	"log/slog"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobbry/puffering/libraries/comms"
	"github.com/gobbry/puffering/libraries/database/service"
	_ "github.com/lib/pq"
)

const (
	MAINNET_RPC = "https://eth-pokt.nodies.app"
	MAINNET_WS  = "wss://ethereum-rpc.publicnode.com"
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

	c, err := ethclient.DialContext(context.Background(), MAINNET_WS)
	if err != nil {
		slog.Error("unable to eth client", "err", err)
		return
	}
	defer c.Close()

	headers := make(chan *types.Header)
	sub, _ := c.SubscribeNewHead(context.Background(), headers)
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			slog.Error("subscription error", "err", err)
			return

		case header := <-headers:
			slog.Info("new block", "block", header.Number)
			err := service.UpdateConversionRate(c, database, rd, header.Number.Int64(), int64(header.Time)*1_000, context.Background())
			if err != nil {
				slog.Error("unable to update conversion rate", "err", err)
			}
		}
	}
}
