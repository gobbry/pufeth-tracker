package front

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"github.com/gobbry/puffering/libraries/comms"
	"github.com/gobbry/puffering/libraries/database/db"
	"github.com/redis/go-redis/v9"
)

const (
	YEAR_IN_ETH_BLOCKS = 2_628_000
)

type RateManagerBuffer struct {
	BlockNumber    int64
	BlockTimestamp int64
	ConversionRate string
	TotalAsset     string
	TotalSupply    string
}

type RateManager struct {
	buffer   *RingBuffer
	rd       *redis.Client
	database *sql.DB
}

func NewRateManager(rd *redis.Client, database *sql.DB) (*RateManager, error) {
	rm := &RateManager{
		buffer:   NewRingBuffer(YEAR_IN_ETH_BLOCKS),
		rd:       rd,
		database: database,
	}

	err := rm.loadHistorical(context.Background())
	if err != nil {
		return nil, err
	}

	return rm, nil
}

func (rm *RateManager) Start(ctx context.Context) error {
	//TODO: use error channel
	go func() {
		if err := rm.subUpdates(ctx); err != nil {
			slog.Error("unable to subscribe: %v", "err", err)
		}
	}()

	return nil
}

func (rm *RateManager) GetRates() []RateManagerBuffer {
	return rm.buffer.GetAll()
}

func (rm *RateManager) loadHistorical(ctx context.Context) error {
	querier := db.New(rm.database)
	chains, err := querier.GetChains(ctx)
	if err != nil {
		return err
	}

	if len(chains) == 0 {
		return nil
	}

	for _, chain := range chains {
		rates, err := querier.GetHistoricalPufferRates(ctx, db.GetHistoricalPufferRatesParams{
			Chain: chain,
			Limit: int32(rm.buffer.capacity), // should use cursor if time series too large
		})
		if err != nil {
			return err
		}

		for i := len(rates) - 1; i >= 0; i-- {
			rm.buffer.Push(RateManagerBuffer{
				BlockNumber:    rates[i].BlockNumber,
				BlockTimestamp: rates[i].BlockTimestampMs,
				ConversionRate: rates[i].ConversionRate,
				TotalAsset:     rates[i].TotalAsset,
				TotalSupply:    rates[i].TotalSupply,
			})
		}
	}

	return nil
}

func (rm *RateManager) subUpdates(ctx context.Context) error {
	topic, err := comms.PufEthConversionRateTopic("")
	if err != nil {
		return err
	}

	sub := rm.rd.PSubscribe(ctx, topic)
	defer sub.Close()

	ch := sub.Channel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-ch:
			if !ok {
				return fmt.Errorf("unable to subscribe: %w", err)
			}

			event := new(comms.PufferCRateEvent)
			err := json.Unmarshal([]byte(msg.Payload), event)
			if err != nil {
				slog.Error("unable to marshal message", "err", err, "payload", msg.Payload)
				continue
			}

			slog.Info("received new conversion event", "block", event.BlockNumber, "rate", event.ConversionRate)
			rm.buffer.Push(RateManagerBuffer{
				BlockNumber:    event.BlockNumber,
				BlockTimestamp: event.BlockTimestampMs,
				ConversionRate: event.ConversionRate,
				TotalAsset:     event.TotalAsset,
				TotalSupply:    event.TotalSupply,
			})
		}
	}
}

/* ///////////////////////////////// CLAUDE GENERATED CODE /////////////////////////////////// */
type RingBuffer struct {
	data     []RateManagerBuffer
	head     int // Points to the newest item
	count    int // Number of items in buffer
	capacity int
	mu       sync.RWMutex
}

func NewRingBuffer(capacity int) *RingBuffer {
	return &RingBuffer{
		data:     make([]RateManagerBuffer, capacity),
		capacity: capacity,
		head:     -1, // Start with -1 to indicate empty buffer
	}
}

func (rb *RingBuffer) Push(item RateManagerBuffer) {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	// Move head forward (with wrap-around)
	rb.head = (rb.head + 1) % rb.capacity

	// Insert new item at head
	rb.data[rb.head] = item

	// Update count if we haven't filled the buffer yet
	if rb.count < rb.capacity {
		rb.count++
	}
}

func (rb *RingBuffer) GetAll() []RateManagerBuffer {
	rb.mu.RLock()
	defer rb.mu.RUnlock()

	if rb.count == 0 {
		return []RateManagerBuffer{}
	}

	result := make([]RateManagerBuffer, rb.count)

	// Copy items in descending order (newest to oldest)
	for i := 0; i < rb.count; i++ {
		// Calculate source index with wrap-around
		idx := (rb.head - i + rb.capacity) % rb.capacity
		result[i] = rb.data[idx]
	}

	return result
}

// Optional: Add method to get capacity utilization
func (rb *RingBuffer) Utilization() float64 {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return float64(rb.count) / float64(rb.capacity)
}

// Optional: Add method to get current size
func (rb *RingBuffer) Size() int {
	rb.mu.RLock()
	defer rb.mu.RUnlock()
	return rb.count
}
