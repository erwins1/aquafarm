package datastore

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func InitRedisConnection() *redis.Client {
	host := os.Getenv("REDIS_HOST")        // get redis host from env file
	rdb := redis.NewClient(&redis.Options{ // init redis connection
		Addr: host,
	})
	// check connection status
	rStatus := rdb.Ping(context.Background())
	if rStatus.Err() != nil {
		log.Fatal("fail connect to redis, err: ", rStatus.Err())
	}
	// return redis connection
	return rdb
}

func CloseRedisConnection(redis *redis.Client) {
	// close redis connection
	err := redis.Close()
	if err != nil {
		log.Println("fail to close redis")
	}
}
