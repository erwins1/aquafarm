package impl

import (
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	uFarm "aquafarm/internal/usecase/farm"
)

type usecase struct {
	rDbFarm rDbFarm.Repository
	rDbPond rDbPond.Repository
}

func New(rDbFarm rDbFarm.Repository, rDbPond rDbPond.Repository) uFarm.Usecase {
	return &usecase{
		rDbFarm: rDbFarm,
		rDbPond: rDbPond,
	}
}
