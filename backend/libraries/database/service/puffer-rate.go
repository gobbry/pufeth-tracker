package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gobbry/puffering/libraries/comms"
	"github.com/gobbry/puffering/libraries/database/db"
	pufferclient "github.com/gobbry/puffering/libraries/puffer-client"
	"github.com/redis/go-redis/v9"
)

func UpdateConversionRate(c *ethclient.Client, database *sql.DB, rd *redis.Client, blockNumber int64, blockTimestampMs int64, ctx context.Context) error {
	querier := db.New(database)
	rate, err := pufferclient.GetPufEthConversionRate(ctx, c, big.NewInt(blockNumber))
	if err != nil {
		return fmt.Errorf("unable to fetch conversion rate: %w", err)
	}

	params := &db.InsertPufferRateParams{
		Chain:            "ethereum",
		BlockNumber:      blockNumber,
		BlockTimestampMs: blockTimestampMs,
		TotalAsset:       rate.TotalAssets.String(),
		TotalSupply:      rate.TotalSupply.String(),
		ConversionRate:   rate.ConversionRate.String(),
	}
	_, err = querier.InsertPufferRate(ctx, *params)
	if err != nil {
		return fmt.Errorf("unable to store conversion rate: %w", err)
	}

	slog.Info("stored rate", "rate", rate.ConversionRate.String(), "totalAsset", rate.TotalAssets.String(), "totalSupply", rate.TotalSupply.String())

	rateEvent := comms.PufferCRateEvent{
		InsertPufferRateParams: *params,
	}
	msg, err := json.Marshal(rateEvent)
	if err != nil {
		return fmt.Errorf("unable to marshal rate event: %w", err)
	}

	topic, err := comms.PufEthConversionRateTopic("ethereum")
	if err != nil {
		return err
	}

	err = rd.Publish(ctx, topic, msg).Err()
	if err != nil {
		return fmt.Errorf("unable to pubsub conversion rate: %w", err)
	}

	return nil
}
