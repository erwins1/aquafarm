package main

import (
	"github.com/julienschmidt/httprouter"

	// repository
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbFarmImpl "aquafarm/internal/repository/db/farm/impl"

	// usecase
	uFarm "aquafarm/internal/usecase/farm"
	uFarmImpl "aquafarm/internal/usecase/farm/impl"

	// handler
	hFarmImpl "aquafarm/internal/handler/farm/impl"
)

type repositoryHTTP struct {
	rDbFarm rDbFarm.Repository
}

type usecaseHTTP struct {
	uFarm uFarm.Usecase
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

	return &repositories
}

func InitUsecase(repositories *repositoryHTTP) *usecaseHTTP {
	var (
		usecases usecaseHTTP
	)

	usecases.uFarm = uFarmImpl.New(repositories.rDbFarm)

	return &usecases
}

func InitHandler(httpRouter *httprouter.Router, usecases *usecaseHTTP) {
	hFarmImpl.New(httpRouter, usecases.uFarm)
}
