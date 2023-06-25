package impl

import (
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	uPond "aquafarm/internal/usecase/pond"
)

type usecase struct {
	rDbFarm rDbFarm.Repository
	rDbPond rDbPond.Repository
}

func New(rDbFarm rDbFarm.Repository, rDbPond rDbPond.Repository) uPond.Usecase {
	return &usecase{
		rDbFarm: rDbFarm,
		rDbPond: rDbPond,
	}
}
