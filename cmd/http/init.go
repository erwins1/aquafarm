package main

import (
	"github.com/julienschmidt/httprouter"

	// repository
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbFarmImpl "aquafarm/internal/repository/db/farm/impl"
	rDbPond "aquafarm/internal/repository/db/pond"
	rDbPondImpl "aquafarm/internal/repository/db/pond/impl"
	rRedisLog "aquafarm/internal/repository/redis/log"
	rRedisLogImpl "aquafarm/internal/repository/redis/log/impl"

	// usecase
	uFarm "aquafarm/internal/usecase/farm"
	uFarmImpl "aquafarm/internal/usecase/farm/impl"
	uLog "aquafarm/internal/usecase/log"
	uLogImpl "aquafarm/internal/usecase/log/impl"
	uPond "aquafarm/internal/usecase/pond"
	uPondImpl "aquafarm/internal/usecase/pond/impl"

	// handler
	hFarmImpl "aquafarm/internal/handler/farm/impl"
	hLogImpl "aquafarm/internal/handler/log/impl"
	hPondImpl "aquafarm/internal/handler/pond/impl"

	// middleware
	mHttp "aquafarm/common/middleware/http"
	mHttpImpl "aquafarm/common/middleware/http/impl"
)

type repositoryHTTP struct {
	rDbFarm   rDbFarm.Repository
	rDbPond   rDbPond.Repository
	rRedisLog rRedisLog.Repository
}

type usecaseHTTP struct {
	uFarm uFarm.Usecase
	uPond uPond.Usecase
	uLog  uLog.Usecase
}

type middlewareHttp struct {
	mHttp mHttp.Middleware
}

func InitApp(httpRouter *httprouter.Router, resouces *resourceHTTP) {
	repositories := InitRepository(resouces)
	usecases := InitUsecase(repositories)
	middlewares := InitMiddleware(resouces)
	InitHandler(httpRouter, usecases, middlewares)
}

func InitRepository(resouces *resourceHTTP) *repositoryHTTP {
	var (
		repositories repositoryHTTP
	)

	repositories.rDbFarm = rDbFarmImpl.New(resouces.Db)
	repositories.rDbPond = rDbPondImpl.New(resouces.Db)
	repositories.rRedisLog = rRedisLogImpl.New(resouces.Redis)

	return &repositories
}

func InitUsecase(repositories *repositoryHTTP) *usecaseHTTP {
	var (
		usecases usecaseHTTP
	)

	usecases.uFarm = uFarmImpl.New(repositories.rDbFarm, repositories.rDbPond)
	usecases.uPond = uPondImpl.New(repositories.rDbFarm, repositories.rDbPond)
	usecases.uLog = uLogImpl.New(repositories.rRedisLog)

	return &usecases
}

func InitMiddleware(resouces *resourceHTTP) *middlewareHttp {
	var (
		middlewares middlewareHttp
	)

	middlewares.mHttp = mHttpImpl.New(resouces.Redis)
	return &middlewares
}

func InitHandler(httpRouter *httprouter.Router, usecases *usecaseHTTP, middlewares *middlewareHttp) {
	hFarmImpl.New(httpRouter, usecases.uFarm, middlewares.mHttp)
	hPondImpl.New(httpRouter, usecases.uPond, middlewares.mHttp)
	hLogImpl.New(httpRouter, usecases.uLog, middlewares.mHttp)
}
