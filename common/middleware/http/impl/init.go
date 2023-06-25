package impl

import (
	mHttp "aquafarm/common/middleware/http"
	"aquafarm/common/redisw"
)

type middleware struct {
	redisConnection redisw.Repository
}

func New(redisConnection redisw.Repository) mHttp.Middleware {
	return &middleware{
		redisConnection: redisConnection,
	}
}
