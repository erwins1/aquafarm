package impl

import (
	rRedisLog "aquafarm/internal/repository/redis/log"
	uLog "aquafarm/internal/usecase/log"
)

type usecase struct {
	rRedisLog rRedisLog.Repository
}

func New(rRedisLog rRedisLog.Repository) uLog.Usecase {
	return &usecase{
		rRedisLog: rRedisLog,
	}
}
