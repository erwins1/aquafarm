package impl

import (
	"aquafarm/common/redisw"
	rRedisLog "aquafarm/internal/repository/redis/log"
)

type repository struct {
	redisConnection redisw.Repository
}

func New(redisConnection redisw.Repository) rRedisLog.Repository {
	return &repository{
		redisConnection: redisConnection,
	}
}
