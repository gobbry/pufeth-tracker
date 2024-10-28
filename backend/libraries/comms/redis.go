package comms

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func NewRedis() (*redis.Client, error) {
	godotenv.Load()

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	db, err := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 0)

	if err != nil {
		return nil, err
	}
	if host == "" || port == "" {
		return nil, fmt.Errorf("REDIS_HOST, REDIS_PORT or REDIS_DB is not set")
	}

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   int(db),
	})

	return client, nil
}
