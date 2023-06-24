package impl

import (
	rRedis "aquafarm/common/redisw"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	redisConn redis.Client
}

func New(redisConn redis.Client) rRedis.Repository {
	return &repository{
		redisConn: redisConn,
	}
}
