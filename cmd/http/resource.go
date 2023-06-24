package main

import (
	"aquafarm/common/datastore"
	"aquafarm/common/redisw"

	rediswImpl "aquafarm/common/redisw/impl"

	"github.com/jmoiron/sqlx"
)

type resourceHTTP struct {
	Redis  redisw.Repository
	DbConn *sqlx.DB
}

func initResourceHTTP() *resourceHTTP {
	var (
		resources resourceHTTP
	)
	// init redis
	redisConn := datastore.InitRedisConnection()
	// put redis connection to redis wrapper
	resources.Redis = rediswImpl.New(*redisConn)
	resources.DbConn = datastore.InitPostgres()

	return &resources
}
