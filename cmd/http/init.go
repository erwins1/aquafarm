package main

import (
	"github.com/julienschmidt/httprouter"

	// repository
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbFarmImpl "aquafarm/internal/repository/db/farm/impl"
	rDbPond "aquafarm/internal/repository/db/pond"
	rDbPondImpl "aquafarm/internal/repository/db/pond/impl"

	// usecase
	uFarm "aquafarm/internal/usecase/farm"
	uFarmImpl "aquafarm/internal/usecase/farm/impl"
	uPond "aquafarm/internal/usecase/pond"
	uPondImpl "aquafarm/internal/usecase/pond/impl"

	// handler
	hFarmImpl "aquafarm/internal/handler/farm/impl"
	hPondImpl "aquafarm/internal/handler/pond/impl"
)

type repositoryHTTP struct {
	rDbFarm rDbFarm.Repository
	rDbPond rDbPond.Repository
}

type usecaseHTTP struct {
	uFarm uFarm.Usecase
	uPond uPond.Usecase
}

func InitApp(httpRouter *httprouter.Router, resouces *resourceHTTP) {
	repositories := InitRepository(resouces)
	usecases := InitUsecase(repositories)
	InitHandler(httpRouter, usecases)
}

func InitRepository(resouces *resourceHTTP) *repositoryHTTP {
	var (
		repositories repositoryHTTP
	)

	repositories.rDbFarm = rDbFarmImpl.New(resouces.Db)
	repositories.rDbPond = rDbPondImpl.New(resouces.Db)

	return &repositories
}

func InitUsecase(repositories *repositoryHTTP) *usecaseHTTP {
	var (
		usecases usecaseHTTP
	)

	usecases.uFarm = uFarmImpl.New(repositories.rDbFarm, repositories.rDbPond)
	usecases.uPond = uPondImpl.New(repositories.rDbFarm, repositories.rDbPond)

	return &usecases
}

func InitHandler(httpRouter *httprouter.Router, usecases *usecaseHTTP) {
	hFarmImpl.New(httpRouter, usecases.uFarm)
	hPondImpl.New(httpRouter, usecases.uPond)
}
