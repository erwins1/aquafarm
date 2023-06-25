package main

import (
	"aquafarm/common/datastore"
	"aquafarm/common/dbw"
	"aquafarm/common/redisw"

	dbwImpl "aquafarm/common/dbw/impl"
	rediswImpl "aquafarm/common/redisw/impl"
)

type resourceHTTP struct {
	Redis redisw.Repository
	Db    dbw.Repository
}

func initResourceHTTP() *resourceHTTP {
	var (
		resources resourceHTTP
	)
	// init redis
	redisConn := datastore.InitRedisConnection()
	// put redis connection to redis wrapper
	resources.Redis = rediswImpl.New(*redisConn)

	// init db
	DbConn := datastore.InitPostgres()
	resources.Db = dbwImpl.New(DbConn)

	return &resources
}
